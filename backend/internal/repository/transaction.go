package repository

import (
	"context"
	"time"

	"github.com/amesupakorn/pocket-ledger/internal/database"
	"github.com/amesupakorn/pocket-ledger/internal/dto"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

// insert transaction
func (r *TransactionRepository) CreateTx(
	ctx context.Context,
	tx pgx.Tx,
	userID, walletID, categoryID int64,
	amount float64,
	tType string,
	note string,
	createdAt time.Time,
) (int64, error) {

	var id int64

	err := tx.QueryRow(ctx, `
		INSERT INTO transactions
		(user_id, wallet_id, category_id, amount, type, note, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
	`, userID, walletID, categoryID, amount, tType, note, createdAt).Scan(&id)

	return id, err
}

// update wallet balance
func (r *TransactionRepository) UpdateWalletBalance(
	ctx context.Context,
	tx pgx.Tx,
	walletID int64,
	amount float64,
) error {

	_, err := tx.Exec(ctx, `
		UPDATE wallets
		SET balance = balance + $1
		WHERE id = $2
	`, amount, walletID)

	return err
}

func (r *TransactionRepository) GetAll(ctx context.Context) ([]dto.TransactionResponse, error) {
	rows, err := database.DB.Query(ctx, `
		SELECT id, amount, type, note, created_at
		FROM transactions
		ORDER BY created_at DESC`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.TransactionResponse

	for rows.Next() {
		var t dto.TransactionResponse

		err := rows.Scan(
			&t.ID,
			&t.Amount,
			&t.Type,
			&t.Note,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (r *TransactionRepository) GetByID(
	ctx context.Context,
	tx pgx.Tx,
	id int64,
) (dto.TransactionResponse, error) {

	var t dto.TransactionResponse

	err := tx.QueryRow(ctx, `
		SELECT id, amount, type, wallet_id
		FROM transactions
		WHERE id = $1
	`, id).Scan(
		&t.ID,
		&t.Amount,
		&t.Type,
		&t.WalletID,
	)

	return t, err
}

func (r *TransactionRepository) Delete(
	ctx context.Context,
	tx pgx.Tx,
	id int64,
) error {
	_, err := tx.Exec(ctx, `
	DELETE FROM transactions WHERE id = $1
	`, id)

	return err
}

func (r *TransactionRepository) Update(
	ctx context.Context,
	tx pgx.Tx,
	id int64,
	amount float64,
	note string,
) error {

	_, err := tx.Exec(ctx, `
		UPDATE transactions
		SET amount = $1, note $2
		WHERE id = $3
	`, amount, note, id)

	return err
}
