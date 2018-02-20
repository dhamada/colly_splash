package colly_splash

import "fmt"

type SplashArgs struct {
	// render.args
	BaseUrl               string
	Timeout               float32
	ResourceTimeout       float32
	Wait                  float32
	Proxy                 string
	JS                    string
	JSSource              string
	Filters               []string
	AllowedDomains        []string
	AllowedContentTypes   []string
	ForbiddenContentTypes []string
	Viewport              string
	Images                int
	Headers               []byte
	Body                  string
	SaveArgs              []byte
	LoadArgs              []byte
	args5Media            int
	// render.png: above + following
	Width       int
	Height      int
	RenderAll   int
	ScaleMethod string
	// render.jpeg: above + following
	Quality int
	// render.har
	ResponseBody int
	// render.json
	Html    int
	Png     int
	Jpeg    int
	Iframes int
	Script  int
	Console int
	History int
	Har     int
}

type viewport struct {
	Width  int
	Height int
}

func (v *viewport) String() string {
	return fmt.Sprintf("%dx%d", v.Width, v.Height)
}

func NewArgs(options ...func(args *SplashArgs)) *SplashArgs {
	args := &SplashArgs{}
	args.Init()

	for _, f := range options {
		f(args)
	}

	return args
}

func BaseUrl(baseUrl string) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.BaseUrl = baseUrl
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
		args.Viewport = viewport{width, height}.String()
	}
}

func Images(images int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.Images = images
	}
}

func Headers(headers []byte) func(*SplashArgs)  {
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

func args5Media(args5Media int) func(*SplashArgs) {
	return func(args *SplashArgs) {
		args.args5Media = args5Media
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
	args.BaseUrl = ""
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
	args.Viewport = viewport{1024, 768}.String()
	args.Images = 1
	args.Headers = []byte{}
	args.Body = ""
	args.SaveArgs = []byte{}
	args.LoadArgs = []byte{}
	args.args5Media = 0
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
