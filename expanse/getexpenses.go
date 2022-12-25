package expanse

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetIdExpensesHandler(c echo.Context) error {
	db := setDB()
	id := c.Param("id")
	e := Expenses{}
	stmt, err := db.Prepare("SELECT * FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query id statment:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&e.Id, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "id not found"})
	case nil:
		return c.JSON(http.StatusOK, e)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan id:" + err.Error()})
	}
}

func GetAllExpensesHandler(c echo.Context) error {
	db := setDB()
	var expenses = []Expenses{}
	stmt, err := db.Prepare("SELECT * FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusNotFound, Err{Message: "data not found"})
	}
	for rows.Next() {
		e := Expenses{}
		err := rows.Scan(&e.Id, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			return c.JSON(http.StatusNotFound, Err{Message: "data not found"})
		}
		expenses = append(expenses, e)
	}
	return c.JSON(http.StatusOK, expenses)
}
