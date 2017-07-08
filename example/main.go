package main

import (
	"fmt"

	"github.com/tainacleal/go-musixmatch"
	"github.com/tainacleal/go-musixmatch/config"
	"github.com/tainacleal/go-musixmatch/tracks"
)

func main() {
	musixmatch.Key = config.Key
	track, err := tracks.GetByID(15445186)
	if err != nil {
		panic(err)
	}

	fmt.Println(track.ArtistName)
	fmt.Println()

	trackList, err := tracks.Search(&musixmatch.TrackListParams{
		Artist: musixmatch.String("lady gaga"),
	})
	if err != nil {
		panic(err)
	}

	for _, t := range trackList.Tracks {
		fmt.Println(t.Name)
	}
	fmt.Println()

	lyric, err := tracks.GetLyricByID(15445186)
	if err != nil {
		panic(err)
	}

	fmt.Println(lyric.Body)
}
