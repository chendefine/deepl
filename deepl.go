package deepl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	LanguageChinese = "ZH"
	LanguageEnglish = "EN"
)

const (
	apiHost                  = "https://api-free.deepl.com"
	apiAuthScheme            = "DeepL-Auth-Key"
	apiEndpointTranslateText = "/v2/translate"
)

type DeeplConfig struct {
	AuthKey string `json:"auth_key"`
}

type DeeplClient struct {
	client *resty.Client
	config *DeeplConfig
}

type TranslateTextReq struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang"`
	TargetLang string   `json:"target_lang"`
}

type TranslateTextRsp struct {
	Message      string `json:"message"`
	Translations []struct {
		SourceLang string `json:"detected_source_language"`
		Text       string `json:"text"`
	} `json:"translations"`
}

func NewDeeplConfig(key string) *DeeplConfig {
	return &DeeplConfig{AuthKey: key}
}

func NewDeeplClient(cfg *DeeplConfig) *DeeplClient {
	client := resty.New().SetBaseURL(apiHost).SetAuthScheme(apiAuthScheme).SetAuthToken(cfg.AuthKey)
	return &DeeplClient{client: client, config: cfg}
}

func (dc *DeeplClient) TranslateText(ctx context.Context, req *TranslateTextReq) (*TranslateTextRsp, error) {
	data, err := dc.client.R().SetContext(ctx).SetBody(req).Post(apiEndpointTranslateText)
	if err != nil {
		return nil, err
	} else if data.RawResponse.ContentLength == 0 && data.StatusCode() != http.StatusOK {
		return nil, errors.New(data.Status())
	}
	raw, rsp := data.Body(), new(TranslateTextRsp)
	if err = json.Unmarshal(raw, rsp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v ; raw data: %s", err, string(raw))
	} else if rsp.Message != "" {
		return nil, errors.New(rsp.Message)
	}
	return rsp, nil
}
