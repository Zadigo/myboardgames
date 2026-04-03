package backend

// ServerBackendConfig represents the configuration for a specific backend in the YAML file
type ServerBackendConfig struct {
	Url string `json:"url" yaml:"url"`
}

type DatabaeClientConfig struct {
	Client string `json:"client" yaml:"client"`
}

// ServerBackendsConfig represents the backend configuration in the YAML file
type ServerBackendsConfig struct {
	Database *DatabaeClientConfig `json:"database" yaml:"database"`
	Postgres *ServerBackendConfig `json:"postgres" yaml:"postgres"`
	Redis    *ServerBackendConfig `json:"redis" yaml:"redis"`
	RabbitMQ *ServerBackendConfig `json:"rabbitmq" yaml:"rabbitmq"`
}

type ServerBaseConfig struct {
	Backends *ServerBackendsConfig `json:"backends" yaml:"backends"`
}

// YamlConfig represents the structure of the configuration YAML file
type ServerConfig struct {
	Config ServerBaseConfig `json:"config" yaml:"config"`
}
