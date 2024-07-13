package raftcluster

import (
    "log"

    "github.com/spf13/viper"
)

type Config struct {
    NodeID   string   `mapstructure:"id"`
    BindAddr string   `mapstructure:"bind_addr"`
    DataDir  string   `mapstructure:"data_dir"`
    Peers    []string `mapstructure:"peers"`
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("config")

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.UnmarshalKey("node", &config); err != nil {
        return nil, err
    }

    return &config, nil
}

func MustLoadConfig() *Config {
    config, err := LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    return config
}
