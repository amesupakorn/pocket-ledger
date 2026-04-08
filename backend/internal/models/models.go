package models

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedAt time.Time
}

type Wallet struct {
	ID        int64
	UserID    int64
	Name      string
	Balance   float64
	CreatedAt time.Time
}

type Category struct {
	ID        int64
	Name      string
	Type      string
	CreatedAt time.Time
}

type Transaction struct {
	ID         int64
	UserID     int64
	WalletID   int64
	CategoryID int64
	Amount     float64
	Type       string
	Note       string
	CreatedAt  time.Time
}
