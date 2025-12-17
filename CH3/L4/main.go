package main

import (
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedUrl.Host, nil
}
