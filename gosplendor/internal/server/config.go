package server

import (
	"context"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
}

type ServerConfig struct {
	RootDir    string
	Port       string
	YamlConfig *YamlConfig
}

func (s *ServerConfig) LoadYamlConfig(ctx context.Context) error {
	filePath := s.RootDir + "/config.yaml"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&s.YamlConfig)
	if err != nil {
		return err
	}
	log.Printf("🟢 Using YAML configuration from %s", filePath)
	return nil
}

func LoadConfig(rootDir string) *ServerConfig {
	return &ServerConfig{
		RootDir:    rootDir,
		Port:       "8000",
		YamlConfig: &YamlConfig{},
	}
}
