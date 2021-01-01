package config

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port string
}

type DatabaseConfigurations struct {
	Name     string
	User     string
	Password string
	Host     string
}
