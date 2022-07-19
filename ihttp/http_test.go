package ihttp

import (
	"testing"
)

func TestHttpPost(t *testing.T) {
	requestMap := map[string]interface{}{
		"category": "treading",
		"period":   "day",
		"lang":     "go",
		"offset":   0,
		"limit":    2,
	}

	options := Options{
		URL:         "https://e.juejrewin.cn/resoewurces/github",
		RequestBody: requestMap,
		Encode:      EncodeForm,
		Retry:       6,
	}
	result, err := POST(&options)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result.ResponseBody)
}

func TestHttpGet(t *testing.T) {
	requestMap := map[string]interface{}{
		"code": "utf-8",
		"q":    "ps5",
	}
	options := Options{
		URL:         "https://sugg12est.tao321bao.com/sug",
		RequestBody: requestMap,
		Retry:       2,
	}
	result, err := GET(&options)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result.ResponseBody)
}
