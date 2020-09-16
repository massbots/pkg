package dlproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Download struct {
	VideoID   string        `json:"video_id"`
	FormatID  string        `json:"format_id"`
	ChatID    int64         `json:"chat_id"`
	Caption   string        `json:"caption"`
	Silent    bool          `json:"silent"`
	ParseMode string        `json:"parse_mode,omitempty"`
	Progress  *ProgressData `json:"progress,omitempty"`
}

// ProgressData is a raw templated data which will be used for the
// progress and done messages.
type ProgressData struct {
	MessageID         int    `json:"message_id,omitempty"`
	Template          string `json:"template"`
	BarDone           string `json:"bar_done"`
	BarLeft           string `json:"bar_left"`
	StatusDownloading string `json:"status:download"`
	StatusUploading   string `json:"status:upload"`
}

type DownloadResult struct {
	OK        bool    `json:"ok"`
	Error     Error   `json:"error"`
	Server    string  `json:"server"`
	Proxy     string  `json:"proxy"`
	AverSpeed float64 `json:"aver_speed"`
	FileID    string  `json:"file_id"`
}

// Download tells dlproxy to immediately start the downloading.
// Returns a unique ID, which can be used for result polling.
func (api *API) Download(d Download) (string, error) {
	url := fmt.Sprint(api.url, "/dl/download?video_id=", d.VideoID)

	body, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	data, err := api.doRawRequest(req)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// PollDownload starts a polling ticker and returns the result
// as soon as such comes. Uses PollTick field as a ticker duration.
func (api *API) PollDownload(id string) (r *DownloadResult, err error) {
	t := time.NewTicker(api.PollTick)
	for range t.C {
		r, err = api.pollDownload(id)
		if err != nil {
			return
		}
		if r.Error != "" {
			err = r.Error
			return
		}
		if r.OK {
			t.Stop()
			break
		}
	}
	return
}

func (api *API) pollDownload(id string) (result *DownloadResult, err error) {
	url := fmt.Sprint(api.url, "/dl/download?id=", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return result, api.doRequest(req, result)
}
