// ADPwn/core/service/adpwn_module_service.go
package service

import (
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
	loadedModule := plugin.GetAll()
	var apiModules []*model.ADPwnModule

	for _, module := range loadedModule {
		uid := strings.ToLower(strings.ReplaceAll(module.GetName(), " ", "_")) + "_" + module.GetVersion()
		apiModules = append(apiModules, &model.ADPwnModule{
			UID:         uid,
			AttackID:    module.GetName(),
			Metric:      module.GetExecutionMetric(),
			Description: module.GetDescription(),
			Name:        module.GetName(),
			Version:     module.GetVersion(),
			Author:      module.GetAuthor(),
		})
	}
	return apiModules
}
