package expanse

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetAllExpensesHandler(c echo.Context) error {
	db := setDB()
	expenses, err := GetAllExpenses(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}

func GetAllExpenses(db *sql.DB) ([]Expenses, error) {
	var expenses = []Expenses{}
	stmt, err := db.Prepare("SELECT * FROM expenses")
	if err != nil {
		return expenses, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return expenses, err
	}
	for rows.Next() {
		e := Expenses{}
		err := rows.Scan(&e.Id, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, e)
	}
	return expenses, nil
}
