package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/jbattistella/movieratings/utils"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Username: utils.RandomString(6),
		Password: utils.RandomString(32),
		Email:    utils.RandomString(36),
	}

	testUser, err := testQueries.CreateUser(context.Background(), args)
	if err != nil {
		fmt.Printf("error creating user: %v", err)
	}
	if testUser.Username != args.Username {
		t.Errorf("error creating getting username: got %v, want %v", testUser.Username, args.Username)
	}
	if testUser.Password != args.Password {
		t.Errorf("error creating getting password: got %v, want %v", testUser.Password, args.Password)
	}
	if testUser.Email != args.Email {
		t.Errorf("error creating getting ID: got %v, want %v", testUser.Email, args.Email)
	}

	return testUser

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	if err != nil {
		log.Println(err)
	}

	if user1.Username != user2.Username {
		t.Errorf("error creating getting username: got %v, want %v", user1.Username, user2.Username)
	}
	if user1.Password != user2.Password {
		t.Errorf("error creating getting password: got %v, want %v", user1.Password, user2.Password)
	}
	if user1.Email != user2.Email {
		t.Errorf("error creating getting ID: got %v, want %v", user1.Email, user2.Email)
	}
}

func TestListUser(t *testing.T) {

	args := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}

	userList, err := testQueries.ListUsers(context.Background(), args)
	if err != nil {
		t.Errorf("error listing users: got %v", err)
	}

	for _, user := range userList {
		if user.Username == "" {
			t.Errorf("wanted username got %v", user.Username)
		}
		if user.Password == "" {
			t.Errorf("wanted username got %v", user.Password)
		}
		if user.Email == "" {
			t.Errorf("wanted username got %v", user.Email)
		}

	}
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	if err != nil {
		t.Errorf("error deleting user: got %v", err)
	}
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	if err == nil {
		t.Errorf("error deleting user: got %v", user2)
	}

}
