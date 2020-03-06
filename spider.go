package gluaspider

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

// spider type
type spider struct {
	restyClient *resty.Client
}

// NewSpider NewSpider
func NewSpider() *spider {
	return &spider{
		restyClient: resty.New(),
	}
}

// Get Simple Get Url
func (s *spider) Get(l *lua.LState) int {
	resp, err := s.restyClient.R().Get(l.CheckString(1))

	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))

		return 2
	}

	return s.newDocumentFromString(l, resp.String())
}

// RestyClient Get RestyClient
func (s *spider) RestyClient(l *lua.LState) int {
	l.Push(luar.New(l, s.restyClient))

	return 1
}

// newDocumentFromString New Goquery Document From String
func (s *spider) newDocumentFromString(l *lua.LState, html string) int {
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
func (s *spider) NewDocumentFromString(l *lua.LState) int {
	return s.newDocumentFromString(l, l.CheckString(1))
}

// Loader Loader
func (s *spider) Loader(l *lua.LState) int {
	// register functions to the table
	mod := l.SetFuncs(l.NewTable(), map[string]lua.LGFunction{
		"RestyClient":           s.RestyClient,
		"NewDocumentFromString": s.NewDocumentFromString,
		"Get":                   s.Get,
	})

	// returns the module
	l.Push(mod)

	return 1
}
