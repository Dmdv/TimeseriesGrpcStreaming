package common

import (
	"github.com/dmdv/timeseriesgrpcstreaming/models"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// InitConfig sets up Viper and loads configuration from `config/<ENVIRONMENT>.yml`
// ENVIRONMENT is an ENV variable. Should be in {development, production}. Default: `development`
func InitConfig() {
	env := "development"
	if e := os.Getenv("ENVIRONMENT"); e != "" {
		env = e
	}

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic().Err(err).Msg("fatal error config file")
	}
	viper.Set("env", env)

	err = ReadDefaultApplicationSettings()
	if err != nil {
		log.Fatal().Err(err).Msg("App settings config error")
	}
}

var DefaultApplicationSettings models.ApplicationSettings

func ReadDefaultApplicationSettings() error {
	err := viper.UnmarshalKey("web.default_application_settings", &DefaultApplicationSettings)
	if err != nil {
		return errors.Wrap(err, "unable to read default application settings")
	}
	err = DefaultApplicationSettings.Validate()
	if err != nil {
		return errors.Wrap(err, "invalid default application settings")
	}
	return nil
}