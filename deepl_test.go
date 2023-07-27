package deepl

import (
	"context"
	"fmt"
	"testing"
)

func TestTranslateText(t *testing.T) {
	deeplClient := NewDeeplClient(NewDeeplConfig(""))
	ctx := context.Background()
	req := &TranslateTextReq{Text: []string{"Hello!"}, SourceLang: LanguageEnglish, TargetLang: LanguageChinese}
	rsp, err := deeplClient.TranslateText(ctx, req)
	fmt.Println(rsp, err)
}
