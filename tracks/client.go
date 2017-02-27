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

func getClient() Client {
	return Client{Backend: musixmatch.GetBackend(), Key: musixmatch.Key}
}

// Get uses the backend client to retrieve a single track by ID.
func Get(params *musixmatch.TrackGetParams) (*musixmatch.Tracks, error) {
	return getClient().Get(params)
}

// Get hits track.get endpoint and returns a single track by ID.
func (c Client) Get(params *musixmatch.TrackGetParams) (*musixmatch.Tracks, error) {
	if params.ID == nil && params.MBID == nil {
		return nil, errors.New("TrackGetParams must have either ID or MBID set")
	}
	queryParams := musixmatch.SetTrackGetParams(params)

	v := &musixmatch.Return{}

	if err := c.Backend.Call("GET", "track.get", queryParams, v); err != nil {
		return nil, err
	}

	tracks := &musixmatch.Tracks{}
	err := json.Unmarshal(*v.Message.Body, &tracks)

	return tracks, err
}

// Search uses the backend client to search for specific tracks.
func Search(params *musixmatch.TrackListParams) (*musixmatch.TrackList, error) {
	return getClient().Search(params)
}

// Search hits track.search endpoint and returns a list of Tracks that fits the search params
func (c Client) Search(params *musixmatch.TrackListParams) (*musixmatch.TrackList, error) {
	v := &musixmatch.Return{}

	queryParams := musixmatch.SetTrackListParams(params)

	if err := c.Backend.Call("GET", "track.search", queryParams, v); err != nil {
		return nil, err
	}

	trackList := &musixmatch.TrackList{}
	return trackList, json.Unmarshal(*v.Message.Body, &trackList)
}

// GetLyric uses the backed client to retrieve a single track lyric info.
func GetLyric(params *musixmatch.TrackLyricsParams) (*musixmatch.LyricInfo, error) {
	return getClient().GetLyric(params)
}

// GetLyric hits track.lyrics.get endpoint and returns information about where to request the lyrics for a specific track.
func (c Client) GetLyric(params *musixmatch.TrackLyricsParams) (*musixmatch.LyricInfo, error) {
	queryParams := musixmatch.SetTrackLyricsParams(params)

	v := &musixmatch.Return{}
	method := "track.lyrics.get"

	if params.Artist != nil && params.TrackTitle != nil {
		method = "matcher.lyrics.get"
	}

	if err := c.Backend.Call("GET", method, queryParams, v); err != nil {
		return nil, err
	}

	lyric := &musixmatch.LyricInfo{}
	return lyric, json.Unmarshal(*v.Message.Body, &lyric)
}
