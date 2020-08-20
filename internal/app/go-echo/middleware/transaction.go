package middleware

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	// TxKey is key for tx
	TxKey = "Tx"
)

// TransactionHandler add middleware
func TransactionHandler(db *sqlx.DB) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx, _ := db.Beginx()

			c.Set(TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			logrus.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}
