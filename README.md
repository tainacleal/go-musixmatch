# go-musixmatch
Go client for the [Musixmatch API](https://developer.musixmatch.com/)

This is primarily a practice project, strongly based on Stripe's go client.

Finished:
- Track endpoints

TODO:
- Chart endpoints
- Artist endpoints
- Album endpoints
- Unit tests

Usage Example:

```go
func main() {
  // setup musixmatch api key
	musixmatch.Key = config.Key
  
  // get track by ID
	track, err := tracks.Get(&musixmatch.TrackGetParams{
		ID: musixmatch.Int64(15445186),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(track.ArtistName)

  // search for tracks
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

  // get specific lyrics
	lyric, err := tracks.GetLyric(&musixmatch.TrackLyricsParams{
		ID: musixmatch.Int64(15445186),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(lyric.Body)
}
```
