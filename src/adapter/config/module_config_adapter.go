package config

import (
	"ADPwn/core/model/adpwn"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const (
	moduleConfigPath  = "../../config"
	moduleConfigName  = "modules"
	enumerationPrefix = "enumeration." // Changed: removed "modules." prefix
	attackPrefix      = "attack."      // Changed: removed "modules." prefix

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
func ModuleFromConfig(key string) (*adpwn.Module, []*adpwn.ModuleInheritanceEdge, error) {
	viper.SetConfigName(moduleConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(moduleConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("error while parsing module configuration: %w", err)
	}

	// Debug
	fmt.Printf("Looking for module with key: %s\n", key)
	fmt.Printf("Found keys: %v\n", viper.AllKeys())

	var module *adpwn.Module
	var inherits []*adpwn.ModuleInheritanceEdge
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

	log.Printf("Found option keys: %v", keys)

	var moduleOptions []*adpwn.ModuleOption
	for _, optionKey := range keys {
		typeString := viper.GetString(optionsPath + "." + optionKey + ".type")
		typeObj, err := adpwn.ParseModuleOptionType(typeString)
		if err != nil {
			log.Printf("Error parsing type: %s", typeString)
			continue
		}
		log.Printf(typeObj.String())
		required := viper.GetBool(optionsPath + optionKey + ".required")

		moduleOption := &adpwn.ModuleOption{
			ModuleKey: moduleKey,
			Key:       optionKey,
			Type:      typeObj,
			Required:  required,
		}
		moduleOptions = append(moduleOptions, moduleOption)
	}
	return moduleOptions
}

func buildDependencyEdges(prefix, actualModuleKey string) []*adpwn.ModuleInheritanceEdge {
	log.Println("Building dependency edges " + prefix + inheritsKey)
	inheritModuleKeys := viper.GetStringSlice(prefix + inheritsKey)
	log.Printf("Inherit module keys: %v\n", inheritModuleKeys)
	var dependencyEdges []*adpwn.ModuleInheritanceEdge
	for _, previousModuleKey := range inheritModuleKeys {
		log.Printf("Building dependency edge with key: %s\n", previousModuleKey)
		dependencyEdges = append(dependencyEdges, &adpwn.ModuleInheritanceEdge{PreviousModule: previousModuleKey, NextModule: actualModuleKey})
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
