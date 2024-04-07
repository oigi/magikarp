package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	VIPER  *viper.Viper
	CONFIG *Config
	LOG    *zap.Logger
)
