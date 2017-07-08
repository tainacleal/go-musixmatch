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
	track, err := tracks.GetByID(15445186))
	if err != nil {
		// handle err
	}

  // search for tracks
	trackList, err := tracks.Search(&musixmatch.TrackListParams{
		Artist: musixmatch.String("lady gaga"),
	})
	if err != nil {
		// handle err
	}

	for _, t := range trackList.Tracks {
		fmt.Println(t.Name)
	}

  // get specific lyrics
	lyric, err := tracks.GetLyric(&musixmatch.TrackLyricsParams{
		Artist: musixmatch.String("adele"),
    TrackTitle: musixmatch.String("hello"),
	})
	if err != nil {
		// handle err
	}

  // get lyrics by ID
  lyric, err := tracks.GetLyricByID(15445186)
  if err!=nil{
    // handle err
  }
}
```
