package book

import (
	"embed"
	"html/template"
	"strings"

	"github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/extensions/shortcode"
	"gopkg.in/yaml.v3"
)

//go:embed templates
var templates embed.FS

//go:embed public
var public embed.FS

type Book struct {
	Title  string
	Author string
	Image  string
}

func init() {
	def, _ := yaml.Marshal(Book{})
	shortcode.RegisterShortCode("book", shortcode.ShortCode{
		Render:  bookSC,
		Default: strings.TrimSpace(string(def)),
	})
	xlog.RegisterTemplate(templates, "templates")
	xlog.RegisterStaticDir(public)
	xlog.RegisterWidget(xlog.HEAD_WIDGET, 0, style)
}

func style(xlog.Page) template.HTML {
	return `<link rel="stylesheet" href="/public/book_style.css">`
}

func bookSC(in xlog.Markdown) template.HTML {
	var b Book

	if err := yaml.Unmarshal([]byte(in), &b); err != nil {
		return template.HTML(err.Error())
	}

	output := xlog.Partial("book", xlog.Locals{"book": b})

	return template.HTML(output)
}
