package embeddlement

import (
  "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func EmbedHtmlAll(anUrl, embedlyKey string) (string, error) {
	html, err := EmbedHtmlImage(anUrl)
	if err == nil {
		return html, nil
	}
	html, err = EmbedHtmlEmbedly(anUrl, embedlyKey)
	if err == nil {
		return html, nil
	}
	return "", err
}

func EmbedHtmlImage(anUrl string) (html string, err error) {
	resp, err := http.Head(anUrl)
	if err != nil {
		return "", fmt.Errorf("could not request URL")
	}
	defer resp.Body.Close()
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "image/") {
		return "", fmt.Errorf("not an image")
	}
	return fmt.Sprintf("<a href=\"%s\"><img src=\"%s\" /></a>", resp.Request.URL.String(), resp.Request.URL.String()), nil
}

func EmbedHtmlEmbedly(anUrl, embedlyKey string) (html string, err error) {
	queryString := url.Values{"url": {anUrl}, "key": {embedlyKey}, "format": {"json"}}.Encode()
	resp, err := http.Get("http://api.embed.ly/1/oembed?" + queryString)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var info struct {
		Type             string
		Thumbnail_url    string
		Thumbnail_width  int
		Thumbnail_height int
	}
	if err = json.Unmarshal(body, &info); err != nil {
		return "", err
	}
	if info.Type != "photo" && info.Type != "video" {
		return "", fmt.Errorf("not a reasonably thumbnailed URL")
	}
	if info.Thumbnail_url == "" {
		return "", fmt.Errorf("no thumbnail for URL")
	}
	return fmt.Sprintf("<a href=\"%s\"><img src=\"%s\" width=\"%d\" height=\"%d\" /></a>", anUrl, info.Thumbnail_url, info.Thumbnail_width, info.Thumbnail_height), nil
}
