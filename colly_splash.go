package colly_splash

import (
	"github.com/gocolly/colly"
)

type Collector struct {
	colly.Collector
	Method                  string
	EndPoint                string
	SplashUrl               string
	SplashHeaders           []string
	DontProcessResponse     bool
	DontSendHeaders         bool
	MagicResponse           bool
	SessionID               string
	HttpStatusFromErrorCode bool
	Meta                    []byte
	Args                    SplashArgs
}

var collector colly.Collector

func NewCollySplash(options ...func(*Collector)) *Collector {
	c := &Collector{}
	c.Init()

	for _, f := range options {
		f(c)
	}

	return c
}

func Method(method string) func(*Collector) {
	return func(c *Collector) {
		c.Method = method
	}
}

func EndPoint(endPoint string) func(*Collector) {
	return func(c *Collector) {
		c.EndPoint = endPoint
	}
}

func SplashUrl(url string) func(*Collector) {
	return func(c *Collector) {
		c.SplashUrl = url
	}
}

func SplashHeaders(headers []string) func(*Collector) {
	return func(c *Collector) {
		c.SplashHeaders = headers
	}
}

func DontProcessResponse(dontProcessResponse bool) func(*Collector) {
	return func(c *Collector) {
		c.DontProcessResponse = dontProcessResponse
	}
}

func DontSendHeaders(dontSendHeaders bool) func(*Collector) {
	return func(c *Collector) {
		c.DontSendHeaders = dontSendHeaders
	}
}

func MagicResponse(magicResponse bool) func(*Collector) {
	return func(c *Collector) {
		c.MagicResponse = magicResponse
	}
}

func SessionID(sessionID string) func(*Collector) {
	return func(c *Collector) {
		c.SessionID = sessionID
	}
}

func HttpStatusFromErrorCode(httpStatusFromErrorCode bool) func(*Collector) {
	return func(c *Collector) {
		c.HttpStatusFromErrorCode = httpStatusFromErrorCode
	}
}

func Meta(meta []byte) func(*Collector) {
	return func(c *Collector) {
		c.Meta = meta
	}
}

func Args(args *SplashArgs) func(*Collector) {
	return func(c *Collector) {
		c.Args = *args
	}
}

func (c *Collector) Init() {
	c.Method = "GET"
	c.EndPoint = "render.html"
	c.SplashUrl = "http://127.0.0.1:8050/"
	c.SplashHeaders = []string{}
	c.DontProcessResponse = false
	c.DontSendHeaders = false
	c.MagicResponse = true
	c.SessionID = "default"
	c.HttpStatusFromErrorCode = true
	c.Meta = []byte{}
}

func (c *Collector) Visit(url string, collyCollector *colly.Collector) {
	collector = *collyCollector
	requestUrl := ""

	collyCollector.Visit(requestUrl)
}

func (c *Collector) OnRequest(f colly.RequestCallback) {
	collector.OnRequest(f)
}

func (c *Collector) OnResponse(f colly.ResponseCallback) {
  collector.OnResponse(f)
}

func (c *Collector) OnHTML(goquerySelector string, f colly.HTMLCallback) {
  collector.OnHTML(goquerySelector, f)
}

func (c *Collector) OnXML(xpathQuery string, f colly.XMLCallback) {
  collector.OnXML(xpathQuery, f)
}

func (c *Collector) OnHTMLDetach(goquerySelector string) {
  collector.OnHTMLDetach(goquerySelector)
}

func (c *Collector) OnXMLDetach(xpathQuery string) {
  collector.OnXMLDetach(xpathQuery)
}

func (c *Collector) OnError(f colly.ErrorCallback) {
  collector.OnError(f)
}

func (c *Collector) OnScraped(f colly.ScrapedCallback) {
  collector.OnScraped(f)
}