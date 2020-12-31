package configuration

type Configuration struct {
	Port         int    `config:"PROXY_PORT"`
	TargetRawUrl string `config:"TARGET_URL"`
}
