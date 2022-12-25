package expanse

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func UpdateExpenseHandler(c echo.Context) error {
	db := setDB()
	rowID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	e := Expenses{}
	err = c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	stmt, err := db.Prepare("UPDATE expenses SET title = $1, amount = $2, note = $3, tags = $4 WHERE id = $5")
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	if _, err := stmt.Exec(e.Title, e.Amount, e.Note, pq.Array(e.Tags), rowID); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "Can't update data"})
	}

	return c.JSON(http.StatusOK, e)
}
