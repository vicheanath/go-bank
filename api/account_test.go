package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	mockdb "github.com/vicheanath/go-bank/db/mock/sqlc"
	db "github.com/vicheanath/go-bank/db/sqlc"
	"github.com/vicheanath/go-bank/util"
	"go.uber.org/mock/gomock"
)


func TestGetAccountAPI(t *testing.T){
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		CreateAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// 1. Start a test server
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	// 2. Create a new HTTP request
	url := fmt.Sprintf("/accounts/%d", account.ID)
	request := httptest.NewRequest(http.MethodGet, url, nil)

	// 3. Call the handler function

	server.router.ServeHTTP(recorder, request)
	require.NoError(t, nil)

	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       int64(util.RandomInt(1, 1000)),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}