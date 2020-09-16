package dlproxy

import (
	"fmt"
	"net/http"
)

type Playlist struct {
	Entries     []PlaylistEntry `json:"entries"`
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Uploader    string          `json:"uploader"`
	UploaderID  string          `json:"uploader_id"`
	UploaderURL string          `json:"uploader_url"`
}

type PlaylistEntry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Playlist fetches detailed information about the playlist.
func (api *API) Playlist(id string) (playlist *Playlist, err error) {
	url := fmt.Sprint(api.url, "/dl/playlist?id=", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return playlist, api.doRequest(req, playlist)
}
