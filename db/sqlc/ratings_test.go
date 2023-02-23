package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/jbattistella/movieratings/utils"
)

func createRandomRating(t *testing.T) Rating {

	user := createRandomUser(t)
	movie := createRandomMovie(t)

	userid := sql.NullInt32{Int32: int32(user.ID), Valid: true}
	movieid := sql.NullInt32{Int32: int32(movie.ID), Valid: true}

	arg := CreateRatingParams{
		Score:   int32(utils.RandomRating()),
		MovieID: movieid,
		UserID:  userid,
	}

	ratings, err := testQueries.CreateRating(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating rating: %s", err)
	}

	if ratings.Score == 0 {
		t.Errorf("got %v, want %v", ratings.Score, arg.Score)
	}

	return ratings

}

func TestCreateRating(t *testing.T) {
	createRandomRating(t)
}

func TestGetRatings(t *testing.T) {
	ratings1 := createRandomRating(t)
	ratings2, err := testQueries.GetRating(context.Background(), ratings1.ID)
	if err != nil {
		log.Println(err)
	}

	if ratings1.Score != ratings2.Score {
		t.Errorf("wanted %v, got %v", ratings1.Score, ratings2.Score)
	}
}

func TestGetMovieRatings(t *testing.T) {
	ratings1 := createRandomRating(t)
	ratings2, err := testQueries.GetMovieRatings(context.Background(), ratings1.MovieID)
	if err != nil {
		t.Errorf("error getting ratings: %s", err)
	}

	for _, rate := range ratings2 {

		fmt.Println(rate.Score)

		if rate.ID != rate.ID {
			t.Errorf("wanted %v, got %v", rate.ID, rate.ID)
		}
		if rate.Score != rate.Score {
			t.Errorf("wanted %v, got %v", rate.Score, rate.Score)
		}
		if rate.MovieID != rate.MovieID {
			t.Errorf("wanted %v, got %v", rate.MovieID, rate.MovieID)
		}
		if rate.UserID != rate.UserID {
			t.Errorf("wanted %v, got %v", rate.UserID, rate.UserID)
		}
	}
}
func TestGetUserRatings(t *testing.T) {
	ratings1 := createRandomRating(t)
	ratings2, err := testQueries.GetUserRatings(context.Background(), ratings1.UserID)
	if err != nil {
		t.Errorf("error getting ratings: %s", err)
	}

	for _, rate := range ratings2 {

		if rate.Score != rate.Score {
			t.Errorf("wanted %v, got %v", rate.Score, rate.Score)
		}
	}
}

func TestDeleteRating(t *testing.T) {
	rate1 := createRandomRating(t)
	err := testQueries.DeleteRating(context.Background(), rate1.ID)
	if err != nil {
		t.Errorf("error occured while deleting: %v", err)
	}

	rate2, err := testQueries.GetRating(context.Background(), rate1.ID)
	if err == nil {
		t.Errorf("not empty: %v. got: %v", err, rate2.ID)
	}
}

func TestUpdateRatingParams(t *testing.T) {
	oldRating := createRandomRating(t)

	newScore := int32(utils.RandomRating())

	args := UpdateRatingParams{
		Score: newScore,
		ID:    oldRating.ID,
	}

	updatedRating, err := testQueries.UpdateRating(context.Background(), args)
	if err != nil {
		t.Errorf("error updating rating: %s", err)
	}

	if newScore != updatedRating.Score {
		t.Errorf("wanted %v, got %v", newScore, updatedRating.Score)
	}
}
