// In service/attack_vector_service.go
package service

import (
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/interfaces"
	"ADPwn-core/pkg/model/adpwn"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

// ModuleExecutor represents something that can execute an ADPwn module
type ModuleExecutor interface {
	ExecuteModule(key string, params *adpwn.Parameter) error
}

// RunAttackVector runs an attack vector starting with the target module
func RunAttackVector(ctx context.Context, targetModuleKey string, params *adpwn.Parameter, executor interfaces.ModuleExecutor) error {
	runID := uuid.New().String()
	logger := sse.GetLogger(runID)

	log.Println("Starting Execution with runID: " + runID)

	if params == nil {
		params = &adpwn.Parameter{}
	}
	params.RunID = runID

	logger.Event("run_start", map[string]interface{}{
		"runId":     runID,
		"timestamp": time.Now().Unix(),
		"module":    targetModuleKey,
	})

	moduleLogger := logger.ForModule(targetModuleKey)
	moduleLogger.Info("Starting module execution")

	// Log the start of the attack vector
	logger.Info(fmt.Sprintf("Starting attack vector: %s", targetModuleKey))

	moduleService, err := NewADPwnModuleService(nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create module service: %v", err))
		return fmt.Errorf("failed to create module service: %v", err)
	}

	moduleDependencies, err := moduleService.GetAttackVectorByKey(ctx, targetModuleKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get attack vector: %v", err))
		return fmt.Errorf("failed to get attack vector: %v", err)
	}

	// Execute each module in the attack vector
	for _, module := range moduleDependencies {
		logger.Info(fmt.Sprintf("Executing module: %s", module.Name), map[string]string{
			"moduleKey": module.Key,
		})

		err := executor.ExecuteModule(module.Key, params)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to execute module: %s", module.Name), map[string]interface{}{
				"moduleKey": module.Key,
				"error":     err.Error(),
			})

			// Send run_error event
			logger.Event("run_error", map[string]interface{}{
				"runId":     runID,
				"timestamp": time.Now().Unix(),
				"module":    module.Key,
				"error":     err.Error(),
			})

			return fmt.Errorf("failed to execute module %s: %v", module.Key, err)
		}

		logger.Info(fmt.Sprintf("Successfully executed module: %s", module.Name))
	}

	// Send run_complete event
	logger.Event("run_complete", map[string]interface{}{
		"runId":     runID,
		"timestamp": time.Now().Unix(),
		"module":    targetModuleKey,
	})

	return nil
}
