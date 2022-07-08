package api

import (
	db "github/islmaghany/bank/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transfereRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (s *Server) createTransfer(ctx *gin.Context) {
	var t transfereRequest
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// args := db.TransferTxParams{
	// 	FromAccountID: ,
	// }

	ts, err := s.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: t.FromAccountID,
		ToAccountID:   t.ToAccountID,
		Amount:        t.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, ts)

}
