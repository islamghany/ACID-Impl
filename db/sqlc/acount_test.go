package db

import (
	"context"
	"database/sql"
	"github/islmaghany/bank/util"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(ctx, account1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	args := UpdateAccountParams{
		Balance: util.RandomMoney(),
		ID:      account1.ID,
	}

	account2, err := testQueries.UpdateAccount(ctx, args)

	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(ctx, account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(ctx, account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}
