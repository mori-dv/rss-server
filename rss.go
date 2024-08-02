package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer response.Body.Close()
	dat, er := io.ReadAll(response.Body)
	if er != nil {
		return RSSFeed{}, er
	}
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssFeed, nil
}
