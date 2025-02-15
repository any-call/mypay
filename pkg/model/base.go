package model

import "errors"

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

const (
	NULL     = ""
	SUCCESS  = "SUCCESS"
	FAIL     = "FAIL"
	OK       = "OK"
	DebugOff = 0
	DebugOn  = 1
	Version  = "1.5.86"
)

var (
	MissWechatInitParamErr = errors.New("missing wechat init parameter")
	MissAlipayInitParamErr = errors.New("missing alipay init parameter")
	MissPayPalInitParamErr = errors.New("missing paypal init parameter")
	MissParamErr           = errors.New("missing required parameter")
	MarshalErr             = errors.New("marshal error")
	UnmarshalErr           = errors.New("unmarshal error")
	SignatureErr           = errors.New("signature error")
	VerifySignatureErr     = errors.New("verify signature error")
	CertNotMatchErr        = errors.New("cert not match error")
	GetSignDataErr         = errors.New("get signature data error")
)
