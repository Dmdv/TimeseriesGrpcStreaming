package models

import (
	validator "gopkg.in/go-playground/validator.v8"
)

type ApplicationSettings struct {
	// "MMDDYYYY" | "DDMMYYYY"
	DateFormat string `json:"dateFormat" asv:"eq=MMDDYYYY|eq=DDMMYYYY"`

	// "12h" | "24h"
	TimeFormat string `json:"timeFormat" asv:"eq=12h|eq=24h"`

	// Only "en" at this moment, after ISO 639-1
	Language string `json:"language" asv:"eq=en"`
}

var appSettingsValidator = validator.New(&validator.Config{TagName: "asv"})

func (as *ApplicationSettings) Validate() error {
	return appSettingsValidator.Struct(as)
}