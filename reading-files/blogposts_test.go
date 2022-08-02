package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/absunandan008/lab_LearnGoWithTest/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	/*
		t.Run("This is test 1, ignore it and that why it is commented", func(t *testing.T) {
			fs := fstest.MapFS{
				"hello world.md":  {Data: []byte("Title: Post 1")},
				"hello-world2.md": {Data: []byte("Title: Post 2")},
			}
			posts, err := blogposts.NewPostsFromFS(fs)

			if err != nil {
				t.Fatal(err)
			}

			if len(posts) != len(fs) {
				t.Errorf("got %d but wanted %d posts", len(posts), len(fs))
			}
		})
	*/
	t.Run("This is test 2", func(t *testing.T) {

		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}
		// rest of test code cut for brevity
		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})
}
func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v but want %+v", got, want)
	}

}
