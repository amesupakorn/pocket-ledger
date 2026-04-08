package dto

import "time"

type CreateTransactionRequest struct {
	WalletID   int64   `json:"wallet_id" binding:"required"`
	CategoryID int64   `json:"category_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	Type       string  `json:"type" binding:"required"`
	Note       string  `json:"note"`
}

type TransactionResponse struct {
	ID        int64     `json:"id"`
	WalletID  int64     `json:"wallet_id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
