package my

import (
	"code.google.com/p/go.net/html"
	"io"
	"net/http"
	"strings"
)

func GetTitle(s string) (string, error) {
	resp, err := http.Get(s)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				return "", io.EOF
			}
			//More error handling code
			return "", z.Err()
		}

		t := z.Token()
		if tt == html.StartTagToken && t.Data == "title" {
			tt = z.Next()
			t = z.Token()
			return strings.TrimSpace(html.UnescapeString(t.Data)), nil
		}
	}
}
