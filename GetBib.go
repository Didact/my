package my

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetBib(isbn, format string) string {
	if format == "" {
		format = "mla"
	}
	resp, err := http.Get("http://www.ottobib.com/isbn/" + isbn + "/" + format)
	Check(err)
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	Check(err)
	str := string(buf)
	if strings.Index(str, "No Results for") != -1 {
		return ""
	}
	div := "<div class=\"nine columns\">"
	start := strings.Index(str, div) + len(div) + 7
	end := strings.Index(str[start:], "      ")
	return str[start : start+end-1]
}
