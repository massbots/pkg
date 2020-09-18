package dlproxy

import (
	"fmt"
	"net/http"
)

type Video struct {
	ID           string                 `json:"id"`
	Artist       string                 `json:"artist"`
	Track        string                 `json:"track"`
	Title        string                 `json:"title"`
	Author       string                 `json:"uploader"`
	UploaderID   string                 `json:"uploader_id"`
	UploaderURL  string                 `json:"uploader_url"`
	IsLive       bool                   `json:"is_live"`
	Categories   []string               `json:"categories"`
	ChannelID    string                 `json:"channel_id"`
	ChannelURL   string                 `json:"channel_url"`
	Description  string                 `json:"description"`
	Duration     float32                `json:"duration"`
	Fps          float32                `json:"fps"`
	Width        int                    `json:"width"`
	Height       int                    `json:"height"`
	LikeCount    int                    `json:"like_count"`
	ViewCount    int                    `json:"view_count"`
	DislikeCount int                    `json:"dislike_count"`
	UploadDate   string                 `json:"upload_date"`
	Tags         []string               `json:"tags"`
	Thumbnail    string                 `json:"thumbnail"`
	WebpageURL   string                 `json:"webpage_url"`
	Formats      map[string]VideoFormat `json:"formats"`
}

type VideoFormat struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Size int    `json:"size"`
}

// Video fetches detailed information about the video.
func (api *API) Video(id string) (*Video, error) {
	url := fmt.Sprint(api.url, "/dl/video?id=", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var video Video
	return &video, api.doRequest(req, &video)
}
