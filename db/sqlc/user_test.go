package db

import (
	"context"
	"testing"
	"time"

	"github.com/russelshane/hyperbank/util"
	"github.com/stretchr/testify/require"
)

// Function to create a new random account and persist it into the database
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName: util.RandomOwner(),
		Email: util.RandomEmail(),
	}

	// Will attempt to create a new account based on random arguments and validate the inputs
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// Validates that the arguments provided matches the new values in the database
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	// Checks if the ID and Timestamp fields are generated by postgres
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

// Tests the creation of a new user
func TestCreateUser(t *testing.T){
	createRandomUser(t)
}

// Tests to create a new account and retrieve its data
func TestGetUser(t *testing.T){
	user1 := createRandomUser(t)	
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, 1*time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, 1*time.Second)
}