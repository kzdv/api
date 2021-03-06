package config

type Config struct {
	Server   ConfigServer   `yaml:"server"`
	Database ConfigDatabase `yaml:"database"`
	Redis    ConfigRedis    `yaml:"redis"`
	Session  ConfigSession  `yaml:"session"`
	OAuth    ConfigOAuth    `yaml:"oauth"`
}

type ConfigServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ConfigDatabase struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Automigrate bool   `yaml:"automigrate"`
}

type ConfigRedis struct {
	Password      string   `yaml:"password"`
	Database      int      `yaml:"database"`
	Address       string   `yaml:"address"`
	Sentinel      bool     `yaml:"sentinel"`
	MasterName    string   `yaml:"master_name"`
	SentinelAddrs []string `yaml:"sentinel_addrs"`
}

type ConfigSession struct {
	Cookie ConfigSessionCookie `yaml:"cookie"`
}

type ConfigSessionCookie struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
	Domain string `yaml:"domain"`
	Path   string `yaml:"path"`
	MaxAge int    `yaml:"max_age"`
}

type ConfigOAuth struct {
	BaseURL      string `yaml:"base_url"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	MyBaseURL    string `yaml:"my_base_url"`

	Endpoints ConfigOAuthEndpoints `yaml:"endpoints"`
}

type ConfigOAuthEndpoints struct {
	Authorize string `yaml:"authorize"`
	Token     string `yaml:"token"`
	UserInfo  string `yaml:"user"`
}
