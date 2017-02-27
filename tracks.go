package musixmatch

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Tracks represents the main track type
type Tracks struct {
	Track `json:"track"`
}

// Track represents the fields of the track type
type Track struct {
	ID               int64    `json:"track_id"`
	MBID             string   `json:"track_mbid"`
	Name             string   `json:"track_name"`
	Rating           int      `json:"track_rating"`
	Instrumental     FlexBool `json:"instrumental"`
	Explicit         FlexBool `json:"explicit"`
	Favourites       int      `json:"num_favourite"`
	Genres           Genres   `json:"primary_genres"`
	HasLyrics        FlexBool `json:"has_lyrics"`
	HasSubtitles     FlexBool `json:"has_subtitles"`
	LyricsID         int64    `json:"lyrics_id"`
	SubtitleID       int64    `json:"subtitle_id"`
	AlbumID          int64    `json:"album_id"`
	AlbumName        string   `json:"album_name"`
	ArtistID         int64    `json:"artist_id"`
	ArtistName       string   `json:"artist_name"`
	ShareURL         string   `json:"track_share_url"`
	Restricted       FlexBool `json:"restricted"`
	FirstReleaseDate FlexTime `json:"first_release_date"`
	UpdatedTime      FlexTime `json:"updated_time"`
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

// TrackList represents a list of Track instances
type TrackList struct {
	Tracks []*Tracks `json:"track_list"`
}

// LyricInfo holds a Lyric instance
type LyricInfo struct {
	Lyric `json:"lyrics"`
}

// Lyric represents info returned from get.lyrics.track
type Lyric struct {
	ID                int64    `json:"lyrics_id"`
	Restricted        FlexBool `json:"restricted"`
	Instrumental      FlexBool `json:"instrumental"`
	Body              string   `json:"lyrics_body"`
	Language          string   `json:"lyrics_language"`
	ScriptTrackingURL string   `json:"script_tracking_url"`
	PixelTrackingURL  string   `json:"pixel_tracking_url"`
	HTMLTrackinURL    string   `json:"html_tracking_url"`
	Copyright         string   `json:"lyrics_copyright"`
	UpdatedAt         FlexTime `json:"updated_time"`
}

// TrackGetParams are the acceptable params for the get.track request
type TrackGetParams struct {
	ID   *int64
	MBID *string
}

// SetTrackGetParams ...
func SetTrackGetParams(tgp *TrackGetParams) string {
	var queryParams []string
	if tgp != nil {

		if tgp.ID != nil {
			queryParams = append(queryParams, fmt.Sprintf("track_id=%s", strconv.FormatInt(*tgp.ID, 10)))
		}
		if tgp.MBID != nil {
			queryParams = append(queryParams, fmt.Sprintf("track_mbid=%s", url.QueryEscape(*tgp.MBID)))
		}
	}

	return strings.Join(queryParams, "&")
}

// TrackLyricsParams are the acceptable params for the track.lyrics.get and matcher.track.get requests
type TrackLyricsParams struct {
	TrackTitle *string
	Artist     *string
	ID         *int64
	MBID       *string
}

// SetTrackLyricsParams ...
func SetTrackLyricsParams(tlp *TrackLyricsParams) string {
	var queryParams []string
	if tlp != nil {
		// If track title and artist is passed, use matcher.lyrics.get
		if tlp.TrackTitle != nil && tlp.Artist != nil {
			queryParams = append(queryParams,
				fmt.Sprintf("q_track=%s", url.QueryEscape(*tlp.TrackTitle)),
				fmt.Sprintf("q_artist=%s", url.QueryEscape(*tlp.Artist)),
			)
			return strings.Join(queryParams, "&")
		}
		if tlp.ID != nil {
			queryParams = append(queryParams, fmt.Sprintf("track_id=%s", strconv.FormatInt(*tlp.ID, 10)))
		}
		if tlp.MBID != nil {
			queryParams = append(queryParams, fmt.Sprintf("track_mbid=%s", url.QueryEscape(*tlp.MBID)))
		}
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
