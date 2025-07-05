package client

import (
	"github.com/cheokman/sms-go/config"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// TwilioProvider sends SMS via Twilio API
type TwilioProvider struct {
	client *twilio.RestClient
	from   string
}

// NewTwilioProvider initializes Twilio client using config
func NewTwilioProvider(cfg *config.Config) *TwilioProvider {
	cli := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TwilioAccountSID,
		Password: cfg.TwilioAuthToken,
	})
	return &TwilioProvider{client: cli, from: cfg.TwilioFromNumber}
}

// SendMessage sends SMS messages one by one via Twilio
func (p *TwilioProvider) SendMessage(ref string, recipients []string, content, lang string) ([]SendResponse, error) {
	var results []SendResponse
	for _, to := range recipients {
		params := &openapi.CreateMessageParams{}
		params.SetTo(to)
		params.SetFrom(p.from)
		params.SetBody(content)
		resp, err := p.client.Api.CreateMessage(params)
		if err != nil {
			return nil, err
		}
		results = append(results, SendResponse{Recipient: to, JobID: *resp.Sid})
	}
	return results, nil
}
