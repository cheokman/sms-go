package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	BaseURL          string `envconfig:"MLP_BASE_URL" default:"https://mlpsmsapi.three.com.mo/v1/externalApi/message"`
	Account          string `envconfig:"MLP_ACCOUNT"   required:"true"`
	LoginID          string `envconfig:"MLP_LOGIN_ID"  required:"true"`
	Password         string `envconfig:"MLP_PASSWORD" required:"true"`
	Platform         string `envconfig:"SMS_PLATFORM" default:"3mo"`
	TwilioAccountSID string `envconfig:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken  string `envconfig:"TWILIO_AUTH_TOKEN"`
	TwilioFromNumber string `envconfig:"TWILIO_FROM_NUMBER"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
