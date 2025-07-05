// models/xml_models.go
package models

import "encoding/xml"

// jds is the root wrapper
type JDS struct {
	XMLName xml.Name `xml:"jds"`
	Account Account  `xml:"account"`
}

// Account holds credentials and one operation node
type Account struct {
	Acid    string `xml:"acid,attr"`
	LoginID string `xml:"loginid,attr"`
	Passwd  string `xml:"passwd,attr"`
	// one of ChangePwd, MsgSend, MsgRecv, Reqst, etc.
	MsgSend *MsgSend `xml:"msg_send,omitempty"`
}

// MsgSend request
type MsgSend struct {
	Ref        string   `xml:"ref,omitempty"`
	Recipients []string `xml:"recipient"`
	Content    string   `xml:"content"`
	Language   string   `xml:"language"`
	// optional fields omitted for brevity
}

// MsgSendRet response
type MsgSendRet struct {
	XMLName xml.Name   `xml:"msg_send_ret"`
	Msgs    []SendResp `xml:"msg"`
}

type SendResp struct {
	Recipient string `xml:"recipient"`
	JobID     string `xml:"jobid"`
}
