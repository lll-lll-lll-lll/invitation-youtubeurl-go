package config

import (
	"errors"
	"fmt"
	"net/url"
)

type YoutubeURL string

const YoutubeHOST = "www.youtube.com"

// Validate urlが正しいか検査
func (y YoutubeURL) Validate() error {
	u, err := url.ParseRequestURI(string(y))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to parse URL: %s. error: %s", string(y), err.Error()))
	}
	if u.Hostname() != YoutubeHOST {
		return errors.New(fmt.Sprintf("your host is %s, correct host name is %s.", u.Hostname(), YoutubeHOST))
	}
	return nil
}

// ToStruct urlを構造体に落としこむメソッド
func (y YoutubeURL) ToStruct() (*url.URL, error) {
	u, err := url.ParseRequestURI(string(y))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to parse URL: %s. error: %s", string(y), err.Error()))
	}
	return u, nil
}

// ToYouTubeURL urlをYoutubeURL型に変換するメソッド
func ToYouTubeURL(url string) YoutubeURL {
	u := YoutubeURL(url)
	return u
}
