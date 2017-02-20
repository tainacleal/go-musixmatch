package musixmatch

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Tracks represents the main track type
type Tracks struct {
	Track `json:"track"`
}

// Track represents the fields of the track type
type Track struct {
	ID               int64     `json:"track_id"`
	MBID             string    `json:"track_mbid"`
	Rating           int       `json:"track_rating"`
	Instrumental     FlexBool  `json:"instrumental"`
	Explicit         FlexBool  `json:"explicit"`
	Favourites       int       `json:"num_favourite"`
	Genres           Genres    `json:"primary_genres"`
	HasLyrics        FlexBool  `json:"has_lyrics"`
	HasSubtitles     FlexBool  `json:"has_subtitles"`
	LyricsID         int64     `json:"lyrics_id"`
	SubtitleID       int64     `json:"subtitle_id"`
	AlbumID          int64     `json:"album_id"`
	AlbumName        string    `json:"album_name"`
	ArtistID         int64     `json:"artist_id"`
	ArtistName       string    `json:"artist_name"`
	ShareURL         string    `json:"track_share_url"`
	Restricted       FlexBool  `json:"restricted"`
	FirstReleaseDate time.Time `json:"first_release_date"`
	UpdatedTime      time.Time `json:"updated_time"`
}

// Genres represents the list of primary music genres
type Genres struct {
	GenreList []*MusicGenre `json:"music_genre_list"`
}

// MusicGenre represents the music genre type
type MusicGenre struct {
	Genre GenreFields `json:"music_genre"`
}

// GenreFields represents the fields of the music genre type
type GenreFields struct {
	ID           int64  `json:"music_genre_id"`
	ParentID     int64  `json:"music_genre_parent_id"`
	Name         string `json:"music_genre_name"`
	NameExtended string `json:"music_genre_name_extended"`
	Vanity       string `json:"music_genre_vanity"`
}

// TrackList ...
type TrackList struct {
	Tracks []*Tracks `json:"track_list"`
}

// TrackGetParams are the acceptable params for the get.track request
type TrackGetParams struct {
	ID   *int64
	MBID *string
}

// SetTrackGetParams ...
func SetTrackGetParams(tgp *TrackGetParams) string {
	var queryParams []string
	if tgp.ID != nil {
		queryParams = append(queryParams, fmt.Sprintf("track_id=%s", strconv.FormatInt(*tgp.ID, 10)))
	}
	if tgp.MBID != nil {
		queryParams = append(queryParams, fmt.Sprintf("track_mbid=%s", url.QueryEscape(*tgp.MBID)))
	}
	return strings.Join(queryParams, "&")
}

// TrackListParams are the param options acceptable for requests returning a track list
type TrackListParams struct {
	TrackTitle           *string
	Artist               *string
	Lyrics               *string
	FilterArtistID       *int64
	FilterMusicGenreID   *int64
	FilterLyricsLang     *string
	FilterLyricsOnly     *bool
	FilterByArtistRating *string
	FilterByTrackRating  *string
	Page                 *int
	PageSize             *int
}

// SetTrackListParams ...
func SetTrackListParams(tlp *TrackListParams) string {
	var queryParams []string
	if tlp != nil {
		if tlp.TrackTitle != nil {
			queryParams = append(queryParams, fmt.Sprintf("q_track=%s", url.QueryEscape(*tlp.TrackTitle)))
		}
		if tlp.Artist != nil {
			queryParams = append(queryParams, fmt.Sprintf("q_artist=%s", url.QueryEscape(*tlp.Artist)))
		}
		if tlp.Lyrics != nil {
			queryParams = append(queryParams, fmt.Sprintf("q_lyrics=%s", url.QueryEscape(*tlp.Lyrics)))
		}
		if tlp.FilterArtistID != nil {
			queryParams = append(queryParams, fmt.Sprintf("f_artist_id=%s", strconv.FormatInt(*tlp.FilterArtistID, 10)))
		}
		if tlp.FilterMusicGenreID != nil {
			queryParams = append(queryParams, fmt.Sprintf("f_music_genre_id=%s", strconv.FormatInt(*tlp.FilterMusicGenreID, 10)))
		}
		if tlp.FilterLyricsLang != nil {
			queryParams = append(queryParams, fmt.Sprintf("f_lyrics_lang=%s", url.QueryEscape(*tlp.FilterLyricsLang)))
		}
		if tlp.FilterLyricsOnly != nil && *tlp.FilterLyricsOnly {
			queryParams = append(queryParams, "f_has_lyrics=1")
		}
		if tlp.FilterByArtistRating != nil {
			queryParams = append(queryParams, fmt.Sprintf("s_artist_rating=%s", url.QueryEscape(*tlp.FilterByArtistRating)))
		}
		if tlp.FilterByTrackRating != nil {
			queryParams = append(queryParams, fmt.Sprintf("s_track_rating=%s", url.QueryEscape(*tlp.FilterByTrackRating)))
		}
		if tlp.Page != nil {
			queryParams = append(queryParams, fmt.Sprintf("page=%s", strconv.Itoa(*tlp.Page)))
		}
		if tlp.PageSize != nil {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%s", strconv.Itoa(*tlp.PageSize)))
		}
	}

	return strings.Join(queryParams, "&")
}
