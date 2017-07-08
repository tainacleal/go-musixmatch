package artists

import (
	"encoding/json"
	"errors"

	"github.com/tainacleal/go-musixmatch"
)

type Client struct {
	Backend musixmatch.BackendService
	Key     string
}

func getClient() Client {
	return Client{Backend: musixmatch.GetBackend(), Key: musixmatch.Key}
}

func GetByID(id int64) (*musixmatch.Artists, error) {
	return getClient().Get(&musixmatch.ArtistGetParams{
		ID: musixmatch.Int64(id),
	})
}

func GetByMBID(mbid string) (*musixmatch.Artists, error) {
	return getClient().Get(&musixmatch.ArtistGetParams{
		MBID: musixmatch.String(mbid),
	})
}

func (c Client) Get(params *musixmatch.ArtistGetParams) (*musixmatch.Artists, error) {
	if params.ID == nil && params.MBID == nil {
		return nil, errors.New("ArtistGetParams must have either ID or MBID set")
	}
	queryParams := musixmatch.SetArtistGetParams(params)

	v := &musixmatch.Return{}

	if err := c.Backend.Call("GET", "artist.get", queryParams, v); err != nil {
		return nil, err
	}

	artists := &musixmatch.Artists{}
	err := json.Unmarshal(*v.Message.Body, &artists)
	return artists, err
}

func Search(params *musixmatch.ArtistSearchParams) (*musixmatch.ArtistsList, error) {
	return getClient().Search(params)
}

func (c Client) Search(params *musixmatch.ArtistSearchParams) (*musixmatch.ArtistsList, error) {
	v := &musixmatch.Return{}

	queryParams := musixmatch.SetArtistSearchParams(params)

	if err := c.Backend.Call("GET", "artist.search", queryParams, v); err != nil {
		return nil, err
	}

	artistList := &musixmatch.ArtistsList{}
	return artistList, json.Unmarshal(*v.Message.Body, &artistList)
}
