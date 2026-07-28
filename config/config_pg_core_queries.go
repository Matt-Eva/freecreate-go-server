package config

import (
	"fmt"
	"freecreate/lib/logger"
	"os"
	"path/filepath"
)

type PgCoreQueries struct {
	getUserByEmail string
	createUser     string
}

func (q PgCoreQueries) GetUserByEmail() string {
	return q.getUserByEmail
}

func (q PgCoreQueries) CreateUser() string {
	return q.createUser
}

func ConfigPgCoreQueries() (PgCoreQueries, error) {
	q := PgCoreQueries{}
	getUserByEmail, getUserEmailErr := os.ReadFile(filepath.Join("./db/pg_core/query_files", "get_user_by_email.sql"))
	if getUserEmailErr != nil {
		logger.Log(getUserEmailErr)
		return q, getUserEmailErr
	}
	q.getUserByEmail = string(getUserByEmail)

	createUser, createUserErr := os.ReadFile(filepath.Join("./db/pg_core/query_files", "create_user.sql"))
	if createUserErr != nil {
		logger.Log(createUserErr)
		return q, createUserErr
	}
	q.createUser = string(createUser)

	msg := fmt.Sprintf("Pg core queries loaded. Query 1: %s, Query2: %s", q.GetUserByEmail(), q.CreateUser())
	fmt.Println(msg)

	return q, nil
}
