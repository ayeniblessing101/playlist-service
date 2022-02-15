package spotify

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

func GetTracks(t float64) (*spotify.SearchResult, error) {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	fmt.Println(t)
	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	switch {
	case t > 30:
		result, err := client.Search(ctx, "party", 8, spotify.Limit(5))
		if err != nil {
			return nil, err
		}
		return result, nil
	case t >= 15 && t <= 30:
		result, err := client.Search(ctx, "pop", 8, spotify.Limit(5))
		if err != nil {
			return nil, err
		}
		return result, nil
	case t >= 10 && t <= 14:
		result, err := client.Search(ctx, "rock", 8, spotify.Limit(5))
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result, err := client.Search(ctx, "classical", 8, spotify.Limit(5))
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}
