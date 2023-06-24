package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/edmundcheng221/banking/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := createAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.createAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	secondAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, secondAccount)

	require.Equal(t, account.ID, secondAccount.ID)
	require.Equal(t, account.Owner, secondAccount.Owner)
	require.Equal(t, account.Balance, secondAccount.Balance)
	require.Equal(t, account.Currency, secondAccount.Currency)

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	updateParams := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), updateParams)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.NotEqual(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
}

func TestDelete(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	secondAccount, err2 := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, secondAccount)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	params := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err)

	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
