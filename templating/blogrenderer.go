package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type PostViewModel struct {
	Title          string
	SanitisedTitle string
	Description    string
	Body           string
	Tags           []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	if err := r.templ.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}

	return nil
}

// type postViewModel struct {
// 	Post
// 	HTMLBody template.HTML
// }

// func newPostVM(p Post, r *PostRenderer) postViewModel {
// 	vm := postViewModel{Post: p}
// 	vm.HTMLBody = template.HTML(markdown)
// }
