package environment

var envSchema EnvSchema

type EnvSchema struct {
	APP_ENV string `validate:"required"`
}
