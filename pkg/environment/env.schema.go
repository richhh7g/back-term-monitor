package environment

var envSchema EnvSchema

type EnvSchema struct {
	APP_ENV         string `validate:"required"`
	MONGO_DB        string `validate:"required"`
	MONGO_URL       string `validate:"required"`
	MONGO_HOST      string `validate:"required"`
	MONGO_PORT      string `validate:"required"`
	MONGO_USER      string `validate:"required"`
	MONGO_PASS      string `validate:"required"`
	MAIL_NAME       string `validate:"required"`
	MAIL_SENDER     string `validate:"required"`
	RESEND_BASE_URL string `validate:"required,url"`
	RESEND_API_KEY  string `validate:"required"`
}
