package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/absunandan008/lab_LearnGoWithTest/blogrenderer"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("it converts single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<h1>hello world</h1>`

		if got != want {
			t.Errorf("got '%s' but want '%s' ", got, want)
		}

	})
}
