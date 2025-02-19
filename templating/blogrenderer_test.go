package blogrenderer_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	blogrenderer "github.com/natac13/learn-go-with-tests/templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	pr, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := pr.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		// want := `<h1>hello world</h1><p>This is a description</p>
		// Tags: <ul><li>go</li><li>tdd</li></ul>`
		fmt.Printf("got: %s\n", buf.String())
		// if got != want {
		// 	t.Errorf("got '%s' want '%s'", got, want)
		// }
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := pr.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>
`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pr.Render(io.Discard, aPost)
	}
}
