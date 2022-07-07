package api

import (
	"fmt"
	mockdb "github/islmaghany/bank/db/mock"
	db "github/islmaghany/bank/db/sqlc"
	"github/islmaghany/bank/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	account := createAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1)

	server := NewServer(store)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/accounts/%d", account.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)
	//require.

}
func createAccount() *db.Account {
	return &db.Account{
		ID:       util.RandomInt(1, 100),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
