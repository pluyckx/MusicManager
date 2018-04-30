package main

import (
	"fmt"

	"github.com/pluyckx/go-discogs"
)

func main() {
	client, err := discogs.NewClient(&discogs.Options{
		UserAgent: "MusicManager-pluyckx/0.0.1",
		Token:     "kVpbucdfFZKlSgfEuYcoHfITcnofsSHoePQouYAn"})

	if err != nil {
		panic(err)
	}

	resp, err := client.Search.Search(discogs.SearchRequest{Artist: "Alpha Twins", Title: "Unleashed"})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Pages: %d\nItems per page: %d\n", resp.Pagination.Pages, resp.Pagination.PerPage)

	for p := 0; p < resp.Pagination.Pages; p++ {
		for _, result := range resp.Results {
			release, err := client.Database.Release(result.ID)

			if err != nil {
				panic(err)
			}
			fmt.Println(release.ID)

			for _, track := range release.Tracklist {
				var artists []discogs.ArtistSource

				if len(track.Artists) > 0 {
					artists = track.Artists
				} else {
					artists = release.Artists
				}

				for _, artist := range artists {
					fmt.Printf("%s %s ", artist.Name, artist.Join)
				}

				fmt.Printf("- %s\n", track.Title)
			}

			fmt.Println()
		}
	}
}
