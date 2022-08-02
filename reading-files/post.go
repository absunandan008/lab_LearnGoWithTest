package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	//postData, err := io.ReadAll(postFile)
	//scanner.Scan()
	//titleLine := scanner.Text()

	//scanner.Scan()
	//descriptionLine := scanner.Text()

	//return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tagSlice := strings.Split(readMetaLine(tagSeparator), ", ")
	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: tagSlice, Body: body}, nil

}

func readBody(scanner *bufio.Scanner) string {
	//skipp "---"
	scanner.Scan()
	//now read next lines from body
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
