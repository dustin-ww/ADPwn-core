// ADPwn/core/service/adpwn_module_service.go
package service

import (
	"ADPwn/core/interfaces"
	"ADPwn/core/internal/db"
	"ADPwn/core/model"
	"ADPwn/core/plugin"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"strings"
)

type ADPwnModuleService struct {
	db *dgo.Dgraph
}

func NewADPwnModuleService() (*ADPwnModuleService, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	return &ADPwnModuleService{
		db: db,
	}, nil
}

func (s *ADPwnModuleService) GetAll() []*model.ADPwnModule {
	enumerationModules := convertToModel(plugin.GetAllEnumerations(), false)
	attackModules := convertToModel(plugin.GetAllAttacks(), true)
	allModules := append(enumerationModules, attackModules...)
	return allModules
}

func (*ADPwnModuleService) Run(uid string) error {
	allModules := plugin.GetAll()
	for _, module := range allModules {
		moduleUID := strings.ToLower(strings.ReplaceAll(module.GetName(), " ", "_")) + "_" + module.GetVersion()
		if uid == moduleUID {
			//TODO
		}
	}
	return nil
}

func convertToModel(loadedModules []interfaces.ADPwnModule, isAttack bool) []*model.ADPwnModule {
	var apiModules []*model.ADPwnModule
	for _, module := range loadedModules {
		uid := strings.ToLower(strings.ReplaceAll(module.GetName(), " ", "_")) + "_" + module.GetVersion()
		modul := &model.ADPwnModule{
			UID:         uid,
			AttackID:    module.GetName(),
			Metric:      module.GetExecutionMetric(),
			Description: module.GetDescription(),
			Name:        module.GetName(),
			Version:     module.GetVersion(),
			Author:      module.GetAuthor(),
			IsAttack:    isAttack,
		}
		apiModules = append(apiModules, modul)
	}
	return apiModules
}
