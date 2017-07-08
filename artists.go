package musixmatch

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Artist struct {
	ID            int64       `json:"artist_id"`
	MBID          string      `json:"artist_mbid"`
	Name          string      `json:"artist_name"`
	Comment       string      `json:"artist_comment"`
	Country       string      `json:"artist_country"`
	Rating        int         `json:"artist_rating"`
	PrimaryGenres Genres      `json:"primary_genres"`
	ShareURL      string      `json:"artist_share_url"`
	Credits       ArtistsList `json:"artist_credits"`
	Restricted    FlexBool    `json:"restricted"`
	Managed       FlexBool    `json:"managed"`
}

type Artists struct {
	Artist `json:"artist"`
}

type ArtistsList struct {
	Artists []*Artists `json:"artist_list"`
}

type ArtistGetParams struct {
	ID   *int64
	MBID *string
}

func SetArtistGetParams(agp *ArtistGetParams) string {
	var queryParams []string
	if agp != nil {

		if agp.ID != nil {
			queryParams = append(queryParams, fmt.Sprintf("artist_id=%s", strconv.FormatInt(*agp.ID, 10)))
		}
		if agp.MBID != nil {
			queryParams = append(queryParams, fmt.Sprintf("artist_mbid=%s", url.QueryEscape(*agp.MBID)))
		}
	}

	return strings.Join(queryParams, "&")
}

type ArtistSearchParams struct {
	Artist   *string
	ID       *int64
	MBID     *string
	Page     *int
	PageSize *int
}

func SetArtistSearchParams(asp *ArtistSearchParams) string {
	var queryParams []string

	if asp != nil {
		if asp.Artist != nil {
			queryParams = append(queryParams, fmt.Sprintf("q_artist=%s", url.QueryEscape(*asp.Artist)))
		}
		if asp.ID != nil {
			queryParams = append(queryParams, fmt.Sprintf("f_artist_id=%s", strconv.FormatInt(*asp.ID, 10)))
		}
		if asp.MBID != nil {
			queryParams = append(queryParams, fmt.Sprintf("f_artist_mbid=%s", url.QueryEscape(*asp.MBID)))
		}
		if asp.Page != nil {
			queryParams = append(queryParams, fmt.Sprintf("page=%s", strconv.Itoa(*asp.Page)))
		}
		if asp.PageSize != nil {
			queryParams = append(queryParams, fmt.Sprintf("page_se=%s", strconv.Itoa(*asp.PageSize)))
		}
	}
	return strings.Join(queryParams, "&")
}
