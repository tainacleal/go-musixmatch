package tracks

import (
	"encoding/json"
	"errors"

	"github.com/tainacleal/go-musixmatch"
)

// Client sets the Backend that implements BackendService and the API Key
type Client struct {
	Backend musixmatch.BackendService
	Key     string
}

// Get returns a track by ID
func Get(params *musixmatch.TrackGetParams) (*musixmatch.Tracks, error) {
	return getClient().Get(params)
}

// Get returns a track by ID
func (c Client) Get(params *musixmatch.TrackGetParams) (*musixmatch.Tracks, error) {
	if params.ID == nil && params.MBID == nil {
		return nil, errors.New("TrackGetParams must have either ID or MBID set")
	}
	queryParams := musixmatch.SetTrackGetParams(params)

	v := musixmatch.NewReturn()

	if err := c.Backend.Call("GET", "track.get", queryParams, v); err != nil {
		return nil, err
	}

	tracks := &musixmatch.Tracks{}
	err := json.Unmarshal(*v.Message.Body, &tracks)

	return tracks, err
}

// List returns a list of Tracks
func List(params *musixmatch.TrackListParams) (*musixmatch.TrackList, error) {
	return getClient().List(params)
}

// List returns a list of Tracks
func (c Client) List(params *musixmatch.TrackListParams) (*musixmatch.TrackList, error) {
	v := musixmatch.NewReturn()

	queryParams := musixmatch.SetTrackListParams(params)

	if err := c.Backend.Call("GET", "track.search", queryParams, v); err != nil {
		return nil, err
	}

	trackList := &musixmatch.TrackList{}
	err := json.Unmarshal(*v.Message.Body, &trackList)

	return trackList, err
}

func getClient() Client {
	return Client{Backend: musixmatch.GetBackend(), Key: musixmatch.Key}
}
