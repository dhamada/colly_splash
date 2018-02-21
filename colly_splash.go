package colly_splash

import (
	"github.com/gocolly/colly"
	"net/http"
	"bytes"
	"fmt"
	"strings"
	"encoding/json"
)

type SplashCollector struct {
	Collector               colly.Collector
	EndPoint                string
	SplashURL               string
	SplashHeaders           []string
	DontProcessResponse     bool
	DontSendHeaders         bool
	MagicResponse           bool
	SessionID               string
	HttpStatusFromErrorCode bool
	Args                    *SplashArgs
}

func NewCollySplash(options ...func(*SplashCollector)) *SplashCollector {
	c := &SplashCollector{}
	c.Init()

	for _, f := range options {
		f(c)
	}

	return c
}

func Collector(collector colly.Collector) func(*SplashCollector) {
	return func(splashCollector *SplashCollector) {
		splashCollector.Collector = collector
	}
}

func EndPoint(endPoint string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.EndPoint = endPoint
	}
}

func SplashURL(splashURL string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		url := splashURL

		if strings.LastIndex(splashURL, "/") == -1 {
			url += "/"
		}

		c.SplashURL = url
	}
}

func SplashHeaders(headers []string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.SplashHeaders = headers
	}
}

func DontProcessResponse(dontProcessResponse bool) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.DontProcessResponse = dontProcessResponse
	}
}

func DontSendHeaders(dontSendHeaders bool) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.DontSendHeaders = dontSendHeaders
	}
}

func MagicResponse(magicResponse bool) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.MagicResponse = magicResponse
	}
}

func SessionID(sessionID string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.SessionID = sessionID
	}
}

func HttpStatusFromErrorCode(httpStatusFromErrorCode bool) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.HttpStatusFromErrorCode = httpStatusFromErrorCode
	}
}

func Args(args *SplashArgs) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.Args = args
	}
}

func (c *SplashCollector) Init() {
	c.EndPoint = "render.html"
	c.SplashURL = "http://127.0.0.1:8050/"
	c.SplashHeaders = []string{}
	c.DontProcessResponse = false
	c.DontSendHeaders = false
	c.MagicResponse = true
	c.SessionID = "default"
	c.HttpStatusFromErrorCode = true
	c.Collector = colly.Collector{}
}

func (c *SplashCollector) Visit(url string) error {
	var requestURL string
	if url == "" {
		requestURL = fmt.Sprintf("%s%s", c.SplashURL, c.EndPoint)
	} else {
		requestURL = fmt.Sprintf("%s%s?url=%s", c.SplashURL, c.EndPoint, url)
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json;")

	b, err := json.Marshal(&c.Args)

	if err != nil {
		return err
	}

	var f interface{}
	if err = json.Unmarshal(b, &f); err != nil {
		return err
	}

	// delete default values and empty values
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {

		case string:
			if v == "" {
				delete(m, k)
			}

		case int:
			if v == 0 {
				delete(m, k)
			}

		case float32:
			if v == float32(0) {
				delete(m, k)
			}

		case float64:
			if v == float64(0) {
				delete(m, k)
			}

		case []interface{}:
			if cap(vv) == 0 {
				delete(m, k)
			}

		default:
			continue
		}
	}

	b, err = json.Marshal(f)
	// send post method to splash
	c.Collector.Request("POST", requestURL, bytes.NewReader(b), nil, header)
	return nil
}

func (c *SplashCollector) OnRequest(f colly.RequestCallback) {
	c.Collector.OnRequest(f)
}

func (c *SplashCollector) OnResponse(f colly.ResponseCallback) {
	c.Collector.OnResponse(f)
}

func (c *SplashCollector) OnHTML(goquerySelector string, f colly.HTMLCallback) {
	c.Collector.OnHTML(goquerySelector, f)
}

func (c *SplashCollector) OnXML(xpathQuery string, f colly.XMLCallback) {
	c.Collector.OnXML(xpathQuery, f)
}

func (c *SplashCollector) OnHTMLDetach(goquerySelector string) {
	c.Collector.OnHTMLDetach(goquerySelector)
}

func (c *SplashCollector) OnXMLDetach(xpathQuery string) {
	c.Collector.OnXMLDetach(xpathQuery)
}

func (c *SplashCollector) OnError(f colly.ErrorCallback) {
	c.Collector.OnError(f)
}

func (c *SplashCollector) OnScraped(f colly.ScrapedCallback) {
	c.Collector.OnScraped(f)
}

func (c *SplashCollector) Wait() {
	c.Collector.Wait()
}
