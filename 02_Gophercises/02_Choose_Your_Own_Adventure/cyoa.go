package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func StoryFromJSON(r io.Reader) (Story, error) {
	decoded := json.NewDecoder(r)
	var story Story
	if err := decoded.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

var defaultHandlerTemplate string = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>

    <body>
        <h1>{{ .Title }}</h1>
        {{range .Paragraph}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
                <li>
                    <a href="/{{.Chapter}}">{{.Text}}</a>
                </li>
            {{end}}
        </ul>
    </body>
</html>
`

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, opts ...HandlerOption) handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
			panic(err)
		}
		return
	}
	http.Error(w, "Chapter Not Found", http.StatusNotFound)
}
