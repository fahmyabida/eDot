package config

type JWTConfig struct {
	SecretKey string `envconfig:"SECRET_KEY" required:"true"`
}

// LoadForJWTConfig loads JWT configuration and returns it
func LoadForJWTConfig() (config *JWTConfig) {
	config = &JWTConfig{}

	mustLoad("JWT", config)

	return
}
