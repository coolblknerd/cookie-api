package config

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	Name     string
	User     string
	Password string
}
