package main

import (
	"fmt"

	"github.com/tainacleal/go-musixmatch"
	"github.com/tainacleal/go-musixmatch/config"
	"github.com/tainacleal/go-musixmatch/tracks"
)

func main() {
	musixmatch.Key = config.Key
	track, err := tracks.Get(&musixmatch.TrackGetParams{
		ID: musixmatch.Int64(15445186),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(track)

	trackList, err := tracks.List(&musixmatch.TrackListParams{
		Artist: musixmatch.String("lady gaga"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(trackList)
}