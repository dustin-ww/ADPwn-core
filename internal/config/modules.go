package config

import (
	"ADPwn-core/pkg/model/adpwn"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const (
	moduleConfigPath  = "../../configs"
	moduleConfigName  = "modules"
	enumerationPrefix = "enumeration." // Changed: removed "modulelib." prefix
	attackPrefix      = "attack."      // Changed: removed "modulelib." prefix

	// Common module config keys
	attackIDKey        = ".attack_id"
	nameKey            = ".name"
	versionKey         = ".version"
	descriptionKey     = ".description"
	authorKey          = ".author"
	executionMetricKey = ".execution_metric"
	dependsOnKey       = ".depends_on"
	lootPathKey        = ".loot_path"
	inheritsKey        = ".inherits"
	optionsKey         = ".options"
)

// ModuleFromConfig loads a module configuration from the specified key
func ModuleFromConfig(key string) (*adpwn.Module, []*adpwn.ModuleDependency, error) {
	viper.SetConfigName(moduleConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(moduleConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("error while parsing module configuration: %w", err)
	}

	// Debug
	log.Printf("Loading ADpwn Module with key: %s\n", key)

	var module *adpwn.Module
	var inherits []*adpwn.ModuleDependency
	var prefix string
	// Check if the enumeration module exists

	switch {
	// enum module
	case viper.Get("enumeration."+key) != nil:
		prefix = enumerationPrefix + key
		module = buildEnumerationModule(prefix, key)

	//attack module
	case viper.Get("attack."+key) != nil:
		prefix = attackPrefix + key
		module = buildAttackModule(prefix, key)

	default:
		return nil, nil, fmt.Errorf("module with key %s not found", key)
	}

	inherits = buildDependencyEdges(prefix, key)
	module.Options = buildModuleOptions(prefix, key)
	return module, inherits, nil
}

// buildModuleOptions method to build specified module options from config yml
func buildModuleOptions(prefix, moduleKey string) []*adpwn.ModuleOption {
	optionsPath := prefix + optionsKey
	log.Printf("Loading module options from: %s\n", optionsPath)
	optionsMap := viper.GetStringMap(optionsPath)

	var keys []string
	for key := range optionsMap {
		keys = append(keys, key)
	}

	var moduleOptions []*adpwn.ModuleOption
	for _, optionKey := range keys {
		typeString := viper.GetString(optionsPath + "." + optionKey + ".type")
		typeObj, err := adpwn.ParseModuleOptionType(typeString)
		label := viper.GetString(optionsPath + "." + optionKey + ".label")
		placeholder := viper.GetString(optionsPath + "." + optionKey + ".placeholder")
		if err != nil {
			log.Printf("Error parsing type: %s", typeString)
			continue
		}
		required := viper.GetBool(optionsPath + optionKey + ".required")

		moduleOption := &adpwn.ModuleOption{
			ModuleKey:   moduleKey,
			Key:         optionKey,
			Type:        typeObj,
			Label:       label,
			Required:    required,
			Placeholder: placeholder,
		}
		moduleOptions = append(moduleOptions, moduleOption)
	}
	return moduleOptions
}

func buildDependencyEdges(prefix, actualModuleKey string) []*adpwn.ModuleDependency {
	log.Println("Starting to Build Module Dependencies for ADPwn Module with key: " + actualModuleKey)
	inheritModuleKeys := viper.GetStringSlice(prefix + inheritsKey)
	var dependencyEdges []*adpwn.ModuleDependency
	for _, previousModuleKey := range inheritModuleKeys {
		log.Printf("Building dependency edge with key: %s\n", previousModuleKey)
		dependencyEdges = append(dependencyEdges, &adpwn.ModuleDependency{PreviousModule: previousModuleKey, NextModule: actualModuleKey})
	}
	return dependencyEdges
}

// buildEnumerationModule creates an enumeration module from config
func buildEnumerationModule(prefix, key string) *adpwn.Module {
	return &adpwn.Module{
		AttackID:        viper.GetString(prefix + attackIDKey),
		ExecutionMetric: viper.GetString(prefix + executionMetricKey),
		Description:     viper.GetString(prefix + descriptionKey),
		Name:            viper.GetString(prefix + nameKey),
		Version:         viper.GetString(prefix + versionKey),
		Author:          viper.GetString(prefix + authorKey),
		ModuleType:      adpwn.EnumerationModule,
		LootPath:        viper.GetString(prefix + lootPathKey),
		Key:             key,
	}
}

// buildAttackModule creates an attack module from config
func buildAttackModule(prefix, key string) *adpwn.Module {
	return &adpwn.Module{
		AttackID:        viper.GetString(prefix + attackIDKey),
		ExecutionMetric: viper.GetString(prefix + executionMetricKey),
		Description:     viper.GetString(prefix + descriptionKey),
		Name:            viper.GetString(prefix + nameKey),
		Version:         viper.GetString(prefix + versionKey),
		Author:          viper.GetString(prefix + authorKey),
		ModuleType:      adpwn.AttackModule,
		LootPath:        viper.GetString(prefix + lootPathKey),
		Key:             key,
	}
}
