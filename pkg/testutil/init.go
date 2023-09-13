package testutil

import "github.com/davidtrse/go-clean-arch/config"

func InitConfigForIntegrationTest() {
	var err error
	config.AppConfig, err = config.ViperLoadConfig("../../config/config-local-itest.yml")
	if err != nil {
		panic(err)
	}
}
