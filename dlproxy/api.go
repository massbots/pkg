package dlproxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// API allows you to interact with dlproxy instance.
// See New().
type API struct {
	url string

	// PollTick is used for polling functions as ticker duration.
	PollTick time.Duration
}

// New creates a new API dlproxy instance with the given URL.
func New(url string) *API {
	return &API{
		url:      url,
		PollTick: time.Second,
	}
}

func (api *API) doRawRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		m := make(map[string]string)
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}
		return nil, Error(m["error"])
	}

	return data, nil
}

func (api *API) doRequest(req *http.Request, v interface{}) error {
	data, err := api.doRawRequest(req)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
