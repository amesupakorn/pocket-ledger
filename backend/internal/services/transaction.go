package services

import (
	"context"
	"errors"

	"github.com/amesupakorn/pocket-ledger/internal/database"
	"github.com/amesupakorn/pocket-ledger/internal/dto"
	"github.com/amesupakorn/pocket-ledger/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		repo: repository.NewTransactionRepository(),
	}
}

func (s *TransactionService) CreateTransaction(
	ctx context.Context,
	userID int64,
	req dto.CreateTransactionRequest,
) (int64, error) {

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	// validate type
	if req.Type != "income" && req.Type != "expense" {
		return 0, errors.New("invalid type")
	}

	// insert transaction
	id, err := s.repo.CreateTx(
		ctx, tx,
		userID,
		req.WalletID,
		req.CategoryID,
		req.Amount,
		req.Type,
		req.Note,
	)
	if err != nil {
		return 0, err
	}

	// calculate balance change
	amount := req.Amount
	if req.Type == "expense" {
		amount = -amount
	}

	// update wallet
	if err := s.repo.UpdateWalletBalance(ctx, tx, req.WalletID, amount); err != nil {
		return 0, err
	}

	// commit
	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *TransactionService) GetTransactions(ctx context.Context) ([]dto.TransactionResponse, error) {
	return s.repo.GetAll(ctx)
}

func (s *TransactionService) DeleteTransaction(
	ctx context.Context,
	id int64,
) error {
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	t, err := s.repo.GetByID(ctx, tx, id)
	if err != nil {
		return err
	}

	var refund float64
	if t.Type == "expense" {
		refund = t.Amount
	} else {
		refund = -t.Amount
	}

	if err := s.repo.UpdateWalletBalance(ctx, tx, t.WalletID, refund); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) UpdateTransactions(
	ctx context.Context,
	id int64,
	amount float64,
	note string,
) error {
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	t, err := s.repo.GetByID(ctx, tx, id)

	if err != nil {
		return err
	}

	var diff float64
	if t.Type == "expense" {
		diff = amount - t.Amount
	} else {
		diff = t.Amount - amount
	}

	if err := s.repo.UpdateWalletBalance(ctx, tx, t.WalletID, diff); err != nil {
		return err
	}

	if err := s.repo.Update(ctx, tx, id, amount, note); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
