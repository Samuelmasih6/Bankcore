package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SQLStore provides access to all queries and transactions.
//
// db      -> connection pool to PostgreSQL
// Queries -> sqlc-generated query methods
type SQLStore struct {
	db *pgxpool.Pool
	*Queries
}

// NewStore creates a new SQLStore.
//
// The Queries field is initialized with the connection pool,
// allowing all generated sqlc methods to be called directly.
func NewStore(db *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function inside a database transaction.
//
// Approach:
//
// 1. Begin a transaction.
// 2. Create a Queries object bound to that transaction.
// 3. Execute all queries using the transaction.
// 4. If any query fails -> rollback.
// 5. Otherwise -> commit.
//
// This guarantees atomicity:
// either all operations succeed or none are saved.
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Start transaction
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}

	// Ensure queries run inside this transaction
	q := New(tx)

	// Execute business logic
	err = fn(q)
	if err != nil {
		// Undo all changes on failure
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error : %v, rb error : %v", err, rbErr)
		}
		return err
	}

	// Persist all changes
	return tx.Commit(ctx)
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"Transfer"`
	FromAccount Account  `json:"from_account_id"`
	ToAccount   Account  `json:"to_account_id"`
	Amount      int64    `json:"amount"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {

	account1, err = q.AddAccountBalance(
		ctx,
		AddAccountBalanceParams{
			ID:     accountID1,
			Amount: amount1,
		},
	)
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(
		ctx,
		AddAccountBalanceParams{
			ID:     accountID2,
			Amount: amount2,
		},
	)

	return
}

func (store *SQLStore) TransferTx(
	ctx context.Context,
	arg TransferTxParams,
) (TransferTxResult, error) {

	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		// Create transfer record
		result.Transfer, err = q.CreateTransfer(
			ctx,
			CreateTransferParams{
				FromAccountID: arg.FromAccountID,
				ToAccountID:   arg.ToAccountID,
				Amount:        arg.Amount,
			},
		)
		if err != nil {
			return err
		}

		// Create debit entry
		result.FromEntry, err = q.CreateEntry(
			ctx,
			CreateEntryParams{
				AccountID: arg.FromAccountID,
				Amount:    -arg.Amount,
			},
		)
		if err != nil {
			return err
		}

		// Create credit entry
		result.ToEntry, err = q.CreateEntry(
			ctx,
			CreateEntryParams{
				AccountID: arg.ToAccountID,
				Amount:    arg.Amount,
			},
		)
		if err != nil {
			return err
		}

		// Update balances
		if arg.FromAccountID < arg.ToAccountID {

			result.FromAccount,
				result.ToAccount,
				err = addMoney(
				ctx,
				q,
				arg.FromAccountID,
				-arg.Amount,
				arg.ToAccountID,
				arg.Amount,
			)

		} else {

			result.ToAccount,
				result.FromAccount,
				err = addMoney(
				ctx,
				q,
				arg.ToAccountID,
				arg.Amount,
				arg.FromAccountID,
				-arg.Amount,
			)

		}

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
