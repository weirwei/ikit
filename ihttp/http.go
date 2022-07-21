package ihttp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/weirwei/ikit/ilog"
	"github.com/weirwei/ikit/iutil"
)

const (
	// EncodeJson 请求数据类型为json
	EncodeJson = "_json"

	// EncodeForm 请求数据类型为form
	EncodeForm = "_form"
)

// Options http request options
//
// URL request url
// RequestBody 请求体
// Encode default form
// Headers headers
// Cookies cookies
type Options struct {
	URL         string            `json:"url"`
	RequestBody interface{}       `json:"requestBody"`
	ContentType string            `json:"contentType"`
	Encode      string            `json:"encode"`
	Headers     map[string]string `json:"headers"`
	Cookies     map[string]string `json:"cookies"`

	Retry       int         `json:"retry"`   // 失败重试次数，0不重试
	RetryPolicy RetryPolicy `json:"-"`       // 失败监控，返回值为true表示失败
	Timeout     int         `json:"timeout"` // 超时时间，单位ms，默认3000ms
}

// RetryPolicy retry 策略
type RetryPolicy func(resp *http.Response, err error) bool

var defaultRetryPolicy RetryPolicy = func(resp *http.Response, err error) bool {
	return err != nil || resp == nil || resp.StatusCode >= http.StatusInternalServerError || resp.StatusCode == 0
}

func (o *Options) GetRetryPolicy() RetryPolicy {
	if o.RetryPolicy == nil {
		o.RetryPolicy = defaultRetryPolicy
	}
	return o.RetryPolicy
}

func (o *Options) GetTimeout() time.Duration {
	if o.Timeout <= 0 {
		o.Timeout = 3000
	}
	return time.Duration(o.Timeout) * time.Millisecond
}

// Result http request result
type Result struct {
	HttpCode     int
	ResponseBody string
}

// POST http post request
func POST(opt *Options) (*Result, error) {
	data, err := opt.getData()
	if err != nil {
		return nil, err
	}
	var (
		request  *http.Request
		response *http.Response
	)
	retry := opt.Retry
	for {
		request, err = http.NewRequest(httpPost, opt.URL, strings.NewReader(data))
		if err != nil {
			return nil, err
		}
		opt.makeRequest(request)
		client := http.Client{
			Timeout: opt.GetTimeout(),
		}
		response, err = client.Do(request)
		if err != nil {
			ilog.Errorf("POST err:%v,opt:%s", err, iutil.ToJson(opt))
		}
		if !opt.GetRetryPolicy()(response, err) || retry <= 0 {
			break
		}
		retry--
	}
	if err != nil {
		return nil, err
	}
	var res *Result
	if response != nil {
		res, err = responseToResult(response)
		if err != nil {
			return nil, err
		}
	}
	ilog.Infof("POST opt:%s,res:%s", iutil.ToJson(opt), iutil.ToJson(res))
	return res, nil
}

// GET http get request
func GET(opt *Options) (*Result, error) {
	data, err := opt.getUrlData()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s?%s", opt.URL, data)
	request, err := http.NewRequest(httpGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response *http.Response
	retry := opt.Retry
	for {
		client := http.Client{
			Timeout: opt.GetTimeout(),
		}
		response, err = client.Do(request)
		if err != nil {
			ilog.Errorf("GET err:%v,opt:%s", err, iutil.ToJson(opt))
		}
		if !opt.GetRetryPolicy()(response, err) || retry <= 0 {
			break
		}
		retry--
	}
	if err != nil {
		return nil, err
	}
	var res *Result
	if response != nil {
		res, err = responseToResult(response)
		if err != nil {
			return nil, err
		}
	}
	ilog.Infof("GET opt:%s,res:%s", iutil.ToJson(opt), iutil.ToJson(res))
	return res, nil
}

func (o *Options) getData() (string, error) {
	var data string
	var err error
	switch o.Encode {
	case EncodeJson:
		data, err = jsoniter.MarshalToString(o.RequestBody)
		if err != nil {
			return "", err
		}
	case EncodeForm:
		fallthrough
	default:
		data, err = o.getUrlData()
		if err != nil {
			return "", err
		}
	}

	return data, nil
}

func (o *Options) getUrlData() (data string, err error) {
	value := &url.Values{}
	if formData, ok := o.RequestBody.(map[string]string); ok {
		for k, v := range formData {
			value.Set(k, v)
		}
	} else if formData, ok := o.RequestBody.(map[string]interface{}); ok {
		for k, v := range formData {
			switch v := v.(type) {
			case string:
				value.Set(k, v)
			default:
				vStr, err := jsoniter.MarshalToString(v)
				if err != nil {
					return data, err
				}
				value.Set(k, vStr)
			}
		}
	} else {
		return data, errors.New("get requestBody error")
	}
	data = value.Encode()
	return data, nil
}

func (o *Options) makeRequest(req *http.Request) {
	for key, val := range o.Headers {
		req.Header.Set(key, val)
	}
	o.getContentType()
	req.Header.Set("Content-Type", o.ContentType)
	for key, val := range o.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: val,
		})
	}
}

func (o *Options) getContentType() {
	if len(o.ContentType) != 0 {
		return
	}
	switch o.Encode {
	case EncodeJson:
		o.ContentType = contentTypeJson
	case EncodeForm:
		fallthrough
	default:
		o.ContentType = contentTypeForm
	}
}

func responseToResult(response *http.Response) (*Result, error) {
	var res Result
	if response != nil {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		res.HttpCode = response.StatusCode
		res.ResponseBody = string(body)
		_ = response.Body.Close()
	}
	return &res, nil
}
