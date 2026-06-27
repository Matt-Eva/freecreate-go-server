package config

import (
	"fmt"
	"freecreate/logger"
	"os"
	"path/filepath"
)

type PgCoreQueries struct {
	getUserByEmail string
}

func (q PgCoreQueries) GetUserByEmail () string{
	return q.getUserByEmail
}

func ConfigPgCoreQueries ()(PgCoreQueries, error){
	q := PgCoreQueries{}
	getUserByEmail, getUserEmailErr := os.ReadFile(filepath.Join("./db/pg_core/query_files", "get_user_by_email.sql"))
	if getUserEmailErr != nil {
		logger.Log(getUserEmailErr)
		return q, getUserEmailErr
	}
	q.getUserByEmail = string(getUserByEmail)

	msg := fmt.Sprintf("Pg core queries loaded. Query 1: %s", q.GetUserByEmail())
	fmt.Println(msg)

	return q, nil
}