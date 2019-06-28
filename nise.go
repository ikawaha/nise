package main

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
)

func init() {
	_ = tokenizer.New()
}

func dropHiragana(s string) string {
	var ret bytes.Buffer
	for _, r := range s {
		if unicode.In(r, unicode.Hiragana) {
			continue
		}
		ret.WriteRune(r)
	}
	return ret.String()
}

func filterToken(prev, token tokenizer.Token) string {
	if token.Class == tokenizer.DUMMY {
		return ""
	}
	f := token.Features()
	if len(f) == 0 {
		return ""
	}
	pos := f[0]
	if pos == "助動詞" {
		if v, ok := filter助動詞(prev, token, f); ok {
			return v
		}
	}
	if pos == "名詞" {
		if v, ok := filter名詞(token, f); ok {
			return v
		}
	}
	if pos == "動詞" {
		if v, ok := filter動詞(prev, token, f); ok {
			return v
		}
	}
	if v, ok := filterDefault(prev, token, f); ok {
		return v
	}
	return dropHiragana(token.Surface)
}

func filter助動詞(prev, token tokenizer.Token, features []string) (string, bool) {
	if f := prev.Features(); len(f) > 5 {
		if f[4] == "サ変・スル" && token.Surface == "ます" {
			return "也", true
		}
	}
	if len(features) > 5 {
		if features[4] == "特殊・ナイ" && prev.Pos() == "動詞" {
			return "不", true
		}
	}
	switch token.Surface {
	case "する":
		return "也", true
	case "たい":
		return "希望", true
	case "た", "ます":
		return "了", true
	case "ず", "ん":
		return "不", true
	}
	return "", false
}

func filter動詞(prev, _ tokenizer.Token, features []string) (string, bool) {
	if f := prev.Features(); len(f) == 0 || len(f) > 2 && f[1] != "サ変接続" {
		if len(features) > 5 && features[4] == "サ変・スル" {
			return "実行", true
		}
	}
	return "", false
}

func filter名詞(token tokenizer.Token, _ []string) (string, bool) {
	switch token.Surface {
	case "私", "わたし":
		return "我", true
	case "あなた", "お前":
		return "你", true
	case "た":
		return "了", true
	}
	return "", false
}

func filterDefault(prev, token tokenizer.Token, features []string) (string, bool) {
	if len(features) > 2 && features[0] == "助詞" && features[1] == "連体化" {
		return "的", true
	}
	if token.Pos() == "形容詞" && token.Surface == "ない" {
		return "非", true
	}
	if f := token.Features(); len(f) > 2 && f[1] == "副助詞／並立助詞／終助詞" {
		return "如何?", true
	}
	if token.Surface == "?" || token.Surface == "？" {
		if f := prev.Features(); len(f) == 0 || len(f) > 2 && f[1] != "副助詞／並立助詞／終助詞" {
			return "如何?", true
		}
		return "", true
	}

	return "", false
}

func Filter(s string) string {
	t := tokenizer.New()
	tokens := t.Tokenize(s)
	ret := make([]string, 0, len(tokens))

	for i := 1; i < len(tokens); i++ {
		ret = append(ret, filterToken(tokens[i-1], tokens[i]))
	}
	return strings.Join(ret, "")
}
