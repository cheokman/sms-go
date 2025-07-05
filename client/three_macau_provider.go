package client

import (
	"github.com/cheokman/sms-go/config"
)

// ThreeMOProvider wraps the existing 3MO client.Client
type ThreeMOProvider struct {
	client *Client
}

// NewThreeMOProvider creates a provider using 3MO credentials
func NewThreeMOProvider(cfg *config.Config) *ThreeMOProvider {
	return &ThreeMOProvider{client: New(cfg)}
}

// SendMessage sends SMS via 3MO platform
func (p *ThreeMOProvider) SendMessage(ref string, recipients []string, content, lang string) ([]SendResponse, error) {
	resp, err := p.client.SendMessage(ref, recipients, content, lang)
	if err != nil {
		return nil, err
	}
	// map 3MO response to our SendResponse
	outs := make([]SendResponse, len(resp.Msgs))
	for i, m := range resp.Msgs {
		outs[i] = SendResponse{Recipient: m.Recipient, JobID: m.JobID}
	}
	return outs, nil
}
