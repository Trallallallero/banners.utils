package utils

import (
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

func Sanitize(s string) bool {
	return strings.Contains(s, ";")
}

func Contains(ints []int, i int) bool {
	for index := 0; index < len(ints); index++ {
		if ints[index] == i {
			return true
		}
	}
	return false
}

func ExecuteQuery(query string, connection sqlx.DB) {
	queryResponse, queryError := connection.Exec(query)
	if queryError != nil {
		log.Fatal(queryError)
	} else {
		_, rowsAffectedError := queryResponse.RowsAffected()
		if rowsAffectedError != nil {
			log.Fatal(rowsAffectedError)
		}
	}
}
