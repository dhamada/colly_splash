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
	Method                  string
	EndPoint                string
	SplashUrl               string
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

func Method(method string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.Method = method
	}
}

func EndPoint(endPoint string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		c.EndPoint = endPoint
	}
}

func SplashUrl(splashUrl string) func(*SplashCollector) {
	return func(c *SplashCollector) {
		url := splashUrl

		if strings.LastIndex(splashUrl, "/") == -1 {
			url += "/"
		}

		c.SplashUrl = url
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
	c.Method = "GET"
	c.EndPoint = "render.html"
	c.SplashUrl = "http://127.0.0.1:8050/"
	c.SplashHeaders = []string{}
	c.DontProcessResponse = false
	c.DontSendHeaders = false
	c.MagicResponse = true
	c.SessionID = "default"
	c.HttpStatusFromErrorCode = true
	c.Collector = colly.Collector{}
}

func (c *SplashCollector) Visit(url string) error {
	requestUrl := fmt.Sprintf("%s%s", c.SplashUrl, c.EndPoint)

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	fmt.Println(c.Args)

	requestBody, err := json.Marshal(&c.Args)
	fmt.Printf("Request Body: %q \n", requestBody)
	fmt.Printf("Error: %s\n", err)

	if err != nil {
		return err
	}

	c.Collector.Request("POST", requestUrl, bytes.NewReader(requestBody), nil, header)

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
