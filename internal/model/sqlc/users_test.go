package db

import (
	"context"
	"testing"
	"time"

	"github.com/berrybytes/simplesecrets/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomUser(), //randomly generated
		Hashpassword: hashPassword,
		Email:        util.RandomEmail(),
		FullName:     util.RandomName(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Hashpassword, user.Hashpassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}
