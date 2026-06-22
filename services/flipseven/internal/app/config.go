package app

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Zadigo/flipseven/internal/models"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	models.AppConfigInterface
	Django struct {
		SecretKey string `yaml:"secret_key"`
	}
	Env      map[string]string `yaml:"env"`
	Emailing struct {
		TestProvider struct {
			FromEmail string `yaml:"from_email"`
			Password  string `yaml:"password"`
			Host      string `yaml:"host"`
			HostName  string `yaml:"host_name"`
			Port      uint16 `yaml:"port"`
		} `yaml:"test_provider"`
	}
	Tests struct {
		Auth struct {
			Url        string `yaml:"url"`
			AccessEnv  string `yaml:"access_env"`
			RefreshEnv string `yaml:"refresh_env"`
		} `yaml:"auth"`
	}
}

func (a *AppConfig) checkPrefix(value string) string {
	if strings.HasPrefix(value, "$") {
		trueValue := strings.TrimPrefix(value, "$")
		return os.Getenv(trueValue)
	}
	return value
}

func (a *AppConfig) setEnvironment() {
	for key, value := range a.Env {
		os.Setenv(key, value)
	}

	a.Emailing.TestProvider.Password = a.checkPrefix(a.Emailing.TestProvider.Password)
	a.Tests.Auth.AccessEnv = a.checkPrefix(a.Tests.Auth.AccessEnv)
	a.Tests.Auth.RefreshEnv = a.checkPrefix(a.Tests.Auth.RefreshEnv)
}

func (a *AppConfig) LoadConfig(app models.AppInterface) error {
	filePath := path.Join(app.GetBaseDir(), "settings.yaml")

	var fileName string
	if file, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("Config file does not exist: %s", filePath)
	} else {
		fileName = file.Name()
	}

	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		return fmt.Errorf("Invalid config file name: %s", fileName)
	}

	if ext := parts[len(parts)-1]; ext != "yaml" && ext != "yml" {
		return fmt.Errorf("Unsupported config file format: %s", ext)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(content, a); err != nil {
		return fmt.Errorf("Failed to unmarshal config file: %w", err)
	}
	a.setEnvironment()
	return nil
}
