// client/client.go
package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"sms-go/config"
	"sms-go/models"
)

// Client holds API config + HTTP client
type Client struct {
	cfg     *config.Config
	httpCli *http.Client
}

// New creates a new SMS API client
func New(cfg *config.Config) *Client {
	return &Client{
		cfg:     cfg,
		httpCli: &http.Client{},
	}
}

// SendMessage sends SMS to one or more recipients
func (c *Client) SendMessage(ref string, recipients []string, content, lang string) (*models.MsgSendRet, error) {
	// build request XML
	jds := models.JDS{
		Account: models.Account{
			Acid:    c.cfg.Account,
			LoginID: c.cfg.LoginID,
			Passwd:  c.cfg.Password,
			MsgSend: &models.MsgSend{
				Ref:        ref,
				Recipients: recipients,
				Content:    content,
				Language:   lang,
			},
		},
	}
	buf, err := xml.MarshalIndent(jds, "", "  ")
	if err != nil {
		return nil, err
	}

	// wrap with XML header & DOCTYPE if your API requires it
	payload := append([]byte(xml.Header), buf...)

	req, err := http.NewRequest("POST", c.cfg.BaseURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml;charset=UTF-8")

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// parse response
	respBody, _ := ioutil.ReadAll(resp.Body)
	var wrapper struct {
		XMLName    xml.Name          `xml:"jds"`
		MsgSendRet models.MsgSendRet `xml:"msg_send_ret"`
	}
	if err := xml.Unmarshal(respBody, &wrapper); err != nil {
		return nil, err
	}
	return &wrapper.MsgSendRet, nil
}
