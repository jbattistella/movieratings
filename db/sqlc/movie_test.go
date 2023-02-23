package db

import (
	"context"
	"log"
	"testing"

	"github.com/jbattistella/movieratings/utils"
)

func createRandomMovie(t *testing.T) Movie {

	arg := CreateMovieParams{
		Title:       utils.RandomString(6),
		Overview:    utils.RandomString(32),
		ReleaseDate: utils.RandomString(8),
		PosterUrl:   utils.RandomString(10),
	}

	movie, err := testQueries.CreateMovie(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating movie: %v", err)
	}
	if arg.Title != movie.Title {
		t.Errorf("error creating movie: got %v, want %v", movie.Title, arg.Title)
	}
	if arg.Overview != movie.Overview {
		t.Errorf("error creating movie: got %v, want %v", movie.Overview, arg.Overview)
	}
	if arg.ReleaseDate != movie.ReleaseDate {
		t.Errorf("error creating movie: got %v, want %v", movie.ReleaseDate, arg.ReleaseDate)
	}
	if arg.PosterUrl != movie.PosterUrl {
		t.Errorf("error creating movie: got %v, want %v", movie.PosterUrl, arg.PosterUrl)
	}
	return movie
}

func TestCreateMovie(t *testing.T) {
	createRandomMovie(t)
}

func TestGetMovie(t *testing.T) {
	movie1 := createRandomMovie(t)
	movie2, err := testQueries.GetMovies(context.Background(), movie1.ID)
	if err != nil {
		log.Println(err)
	}

	if movie1.Title != movie2.Title {
		t.Errorf("error creating getting username: got %v, want %v", movie1.Title, movie2.Title)
	}
	if movie1.Overview != movie2.Overview {
		t.Errorf("error creating getting password: got %v, want %v", movie1.Overview, movie2.Overview)
	}
	if movie1.ReleaseDate != movie2.ReleaseDate {
		t.Errorf("error creating getting ID: got %v, want %v", movie1.ReleaseDate, movie2.ReleaseDate)
	}
	if movie1.PosterUrl != movie2.PosterUrl {
		t.Errorf("error creating getting ID: got %v, want %v", movie1.PosterUrl, movie2.PosterUrl)
	}
}

func TestListMovies(t *testing.T) {

	args := ListMoviesParams{
		Limit:  5,
		Offset: 0,
	}

	movieList, err := testQueries.ListMovies(context.Background(), args)
	if err != nil {
		t.Errorf("error listing users: got %v", err)
	}

	for _, movie := range movieList {
		if movie.Title == "" {
			t.Errorf("wanted movie title got %v", movie.Title)
		}
		if movie.Overview == "" {
			t.Errorf("wanted movie overview got %v", movie.Overview)
		}
		if movie.ReleaseDate == "" {
			t.Errorf("wanted release date got %v", movie.ReleaseDate)
		}
		if movie.PosterUrl == "" {
			t.Errorf("wanted poster url got %v", movie.PosterUrl)
		}

	}
}

func TestDeleteMovie(t *testing.T) {
	movie1 := createRandomMovie(t)
	err := testQueries.DeleteMovies(context.Background(), movie1.ID)
	if err != nil {
		t.Errorf("error deleting user: got %v", err)
	}
	movie2, err := testQueries.GetMovies(context.Background(), movie1.ID)
	if err == nil {
		t.Errorf("error deleting user: got %v", movie2)
	}

}
