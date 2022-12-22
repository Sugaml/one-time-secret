package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/berrybytes/simplesecrets/util"
	"github.com/stretchr/testify/require"
)

func createRandomSecret(t *testing.T) Secret {
	arg := CreateSecretParams{
		Creator: util.RandomUser(),
		Content: util.RandomContent(), //randomly generated
	}
	secret, err := testQueries.CreateSecret(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, secret)
	require.Equal(t, arg.Content, secret.Content)
	require.NotZero(t, secret.ID)
	require.NotZero(t, secret.CreatedAt)
	return secret
}
func TestCreateSecret(t *testing.T) {
	createRandomSecret(t)
}

func TestGetSecret(t *testing.T) {
	secret1 := createRandomSecret(t)
	secret2, err := testQueries.GetSecret(context.Background(), secret1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, secret2)

	require.Equal(t, secret1.ID, secret2.ID)
	require.Equal(t, secret1.Content, secret2.Content)
	require.WithinDuration(t, secret1.CreatedAt, secret2.CreatedAt, time.Second)

}

func TestDeleteSecret(t *testing.T) {
	secret1 := createRandomSecret(t)
	err := testQueries.DeleteSecret(context.Background(), secret1.ID)
	require.NoError(t, err)
	secret2, err := testQueries.GetSecret(context.Background(), secret1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, secret2)
}

func TestListSecret(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSecret(t)
	}
	arg := ListSecretsParams{
		Limit:  5,
		Offset: 5,
	}
	secrets, err := testQueries.ListSecrets(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, secrets, 5)
	for _, secret := range secrets {
		require.NotEmpty(t, secret)
	}
}

func TestUpdateAccount(t *testing.T) {
	secret1 := createRandomSecret(t)

	arg := UpdateSecretParams{
		ID:     secret1.ID,
		Isview: true,
	}

	secret2, err := testQueries.UpdateSecret(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, secret2)

	require.Equal(t, secret1.ID, secret2.ID)
	require.Equal(t, secret1.Content, secret2.Content)
	require.WithinDuration(t, secret1.CreatedAt, secret2.CreatedAt, time.Second)
}
