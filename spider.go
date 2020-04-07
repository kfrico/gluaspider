package gluaspider

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

// Spider type
type Spider struct {
	restyClient *resty.Client
}

// NewSpider NewSpider
func NewSpider() *Spider {
	return &Spider{
		restyClient: resty.New(),
	}
}

// GetRestyClient Get RestyClient
func (s *Spider) GetRestyClient() *resty.Client {
	return s.restyClient
}

func (s *Spider) SetProxy(url string) {
	s.restyClient.SetProxy(url)
}

// Get Simple Get Url
func (s *Spider) Get(l *lua.LState) int {
	resp, err := s.restyClient.R().Get(l.CheckString(1))

	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))

		return 2
	}

	return s.newDocumentFromString(l, resp.String())
}

// RestyClient Get RestyClient
func (s *Spider) RestyClient(l *lua.LState) int {
	l.Push(luar.New(l, s.restyClient))

	return 1
}

// newDocumentFromString New Goquery Document From String
func (s *Spider) newDocumentFromString(l *lua.LState, html string) int {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))

		return 2
	}

	l.Push(luar.New(l, doc))
	l.Push(lua.LNil)

	return 2
}

// NewDocumentFromString New Goquery Document From String
func (s *Spider) NewDocumentFromString(l *lua.LState) int {
	return s.newDocumentFromString(l, l.CheckString(1))
}

// Regexp
func (s *Spider) Regexp(l *lua.LState) int {
	reg, err := regexp.Compile(l.CheckString(1))

	if err != nil {
		l.Push(lua.LBool(false))

		return 1
	}

	l.Push(lua.LBool(reg.MatchString(l.CheckString(2))))

	return 1
}

// Loader Loader
func (s *Spider) Loader(l *lua.LState) int {
	// register functions to the table
	mod := l.SetFuncs(l.NewTable(), map[string]lua.LGFunction{
		"RestyClient":           s.RestyClient,
		"NewDocumentFromString": s.NewDocumentFromString,
		"Get":                   s.Get,
		"Regexp":                s.Regexp,
	})

	// returns the module
	l.Push(mod)

	return 1
}
