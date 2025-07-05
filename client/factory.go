package client

import (
	"fmt"

	"sms-go/config"
)

// NewProvider returns the SMSProvider based on config.Platform
func NewProvider(cfg *config.Config) (SMSProvider, error) {
	switch cfg.Platform {
	case "3mo":
		return NewThreeMOProvider(cfg), nil
	case "twilio":
		return NewTwilioProvider(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported SMS platform: %s", cfg.Platform)
	}
}
