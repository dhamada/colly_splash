package colly_splash

import "fmt"

type SplashArgs struct {
	// render.args
	URL                   string   `json:"url"`
	BaseURl               string   `json:"base_url"`
	Timeout               float32  `json:"timeout"`
	ResourceTimeout       float32  `json:"resource_timeout"`
	Wait                  float32  `json:"wait"`
	Proxy                 string   `json:"proxy"`
	JS                    string   `json:"js"`
	JSSource              string   `json:"js_source"`
	Filters               []string `json:"filters"`
	AllowedDomains        []string `json:"allowed_domains"`
	AllowedContentTypes   []string `json:"allowed_content_types"`
	ForbiddenContentTypes []string `json:"forbidden_content_types"`
	Viewport              string   `json:"viewport"`
	Images                int      `json:"images"`
	Headers               []byte   `json:"headers"`
	Body                  string   `json:"body"`
	SaveArgs              []byte   `json:"save_args"`
	LoadArgs              []byte   `json:"load_args"`
	Html5Media            int      `json:"html5_media"`
	// render.png: above + following
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	RenderAll   int    `json:"render_all"`
	ScaleMethod string `json:"scale_method"`
	// render.jpeg: above + following
	Quality int `json:"quality"`
	// render.har
	ResponseBody int `json:"response_body"`
	// render.json
	Html    int `json:"html"`
	Png     int `json:"png"`
	Jpeg    int `json:"jpeg"`
	Iframes int `json:"iframes"`
	Script  int `json:"script"`
	Console int `json:"console"`
	History int `json:"history"`
	Har     int `json:"har"`
}

type viewport struct {
	Width  int
	Height int
}

func (v *viewport) String() string {
	return fmt.Sprintf("%dx%d", v.Width, v.Height)
}

func NewSsplashArgs(options ...func(args *SplashArgs)) *SplashArgs {
	args := &SplashArgs{}
	args.Init()

	for _, f := range options {
		f(args)
	}

	return args
}

func URL(url string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.URL = url
	}
}

func BaseURL(baseUrl string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.BaseURl = baseUrl
	}
}

func Timeout(timeout float32) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Timeout = timeout
	}
}

func ResourceTimeout(resourceTimeout float32) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.ResourceTimeout = resourceTimeout
	}
}

func Wait(wait float32) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Wait = wait
	}
}

func Proxy(proxy string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Proxy = proxy
	}
}

func JS(js string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.JS = js
	}
}

func JSSource(jsSource string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.JSSource = jsSource
	}
}

func Filters(filters []string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Filters = filters
	}
}

func AllowedDomains(allowedDomains []string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.AllowedDomains = allowedDomains
	}
}

func AllowedContentTypes(allowedContentTypes []string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.AllowedContentTypes = allowedContentTypes
	}
}

func ForbiddenContentTypes(forbiddenContentTypes []string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.ForbiddenContentTypes = forbiddenContentTypes
	}
}

func Viewport(width int, height int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		v := &viewport{1024, 768}
		args.Viewport = v.String()
	}
}

func Images(images int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Images = images
	}
}

func Headers(headers []byte) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Headers = headers
	}
}

func Body(body string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Body = body
	}
}

func SaveArgs(saveArgs []byte) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.SaveArgs = saveArgs
	}
}

func LoadArgs(loadArgs []byte) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.LoadArgs = loadArgs
	}
}

func Html5Media(html5Media int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Html5Media = html5Media
	}
}

func Width(width int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Width = width
	}
}

func Height(height int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Height = height
	}
}

func RenderAll(renderAll int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.RenderAll = renderAll
	}
}

func ScaleMethod(scaleMethod string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.ScaleMethod = scaleMethod
	}
}

func Quality(quality int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Quality = quality
	}
}

func ResponseBody(responseBody int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.ResponseBody = responseBody
	}
}

func Html(html int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Html = html
	}
}

func Png(png int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Png = png
	}
}

func Jpeg(jpeg int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Jpeg = jpeg
	}
}

func Iframes(iframes int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Iframes = iframes
	}
}

func Script(script int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Script = script
	}
}

func Console(console int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Console = console
	}
}

func History(history int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.History = history
	}
}

func Har(har int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Har = har
	}
}

func (args *SplashArgs) Init() {
	args.BaseURl = ""
	args.Timeout = 30
	args.ResourceTimeout = 10
	args.Wait = 0
	args.Proxy = ""
	args.JS = ""
	args.JSSource = ""
	args.Filters = []string{}
	args.AllowedDomains = []string{}
	args.AllowedContentTypes = []string{}
	args.ForbiddenContentTypes = []string{}

	v := &viewport{1024, 768}
	args.Viewport = v.String()

	args.Images = 1
	args.Headers = []byte{}
	args.Body = ""
	args.SaveArgs = []byte{}
	args.LoadArgs = []byte{}
	args.Html5Media = 0
	args.Width = -1
	args.Height = -1
	args.RenderAll = 0
	args.ScaleMethod = "raster"
	args.Quality = 75
	args.ResponseBody = 0
	args.Html = 0
	args.Png = 0
	args.Jpeg = 0
	args.Iframes = 0
	args.Script = 0
	args.Console = 0
	args.History = 0
	args.Har = 0
}
