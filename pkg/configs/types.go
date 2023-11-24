package configs

type Config struct {
	Server             *ServerConfig `json:"server" validate:"required"`
	AttributesMaxLimit uint          `json:"attributesMaxLimit" validate:"required"`
	TraitsMaxLimit     uint          `json:"traitsMaxLimit" validate:"required"`
}

type ServerConfig struct {
	Host              string `json:"host" validate:"required"`
	Port              uint   `json:"port" validate:"required"`
	ServerReadTimeout uint   `json:"serverReadTimeout" validate:"required"`
}
