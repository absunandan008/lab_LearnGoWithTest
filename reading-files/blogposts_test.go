package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/absunandan008/lab_LearnGoWithTest/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d but wanted %d posts", len(posts), len(fs))
	}
}
