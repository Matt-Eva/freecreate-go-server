package config

import (
	"fmt"
	"freecreate/internal/lib/logger"
	"os"
	"path/filepath"
)

type PgCoreQueries struct {
	getUserByEmail string
	createUser     string
	getUserInfo    string
	deleteUser     string
	updateUserInfo string
	createCreator  string
	getMyCreators  string
	getMyCreator   string
	createWriting  string
}

func (q PgCoreQueries) GetUserByEmail() string {
	return q.getUserByEmail
}

func (q PgCoreQueries) CreateUser() string {
	return q.createUser
}

func (q PgCoreQueries) GetUserInfo() string {
	return q.getUserInfo
}

func (q PgCoreQueries) UpdateUserInfo() string {
	return q.updateUserInfo
}

func (q PgCoreQueries) DeleteUser() string {
	return q.deleteUser
}

func (q PgCoreQueries) CreateCreator() string {
	return q.createCreator
}

func (q PgCoreQueries) GetMyCreators() string {
	return q.getMyCreators
}

func (q PgCoreQueries) GetMyCreator() string {
	return q.getMyCreator
}

func (q PgCoreQueries) CreateWriting() string {
	return q.createWriting
}

func ConfigPgCoreQueries() (PgCoreQueries, error) {
	q := PgCoreQueries{}

	getUserByEmail, getUserEmailErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "get_user_by_email.sql"))
	if getUserEmailErr != nil {
		logger.Log(getUserEmailErr)
		return q, getUserEmailErr
	}
	q.getUserByEmail = string(getUserByEmail)

	createUser, createUserErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "create_user.sql"))
	if createUserErr != nil {
		logger.Log(createUserErr)
		return q, createUserErr
	}
	q.createUser = string(createUser)

	getUserInfo, getUserInfoErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "get_user_info.sql"))
	if getUserInfoErr != nil {
		logger.Log(getUserInfoErr)
		return q, getUserInfoErr
	}
	q.getUserInfo = string(getUserInfo)

	updateUserInfo, updateUserInfoErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "update_user_info.sql"))
	if updateUserInfoErr != nil {
		logger.Log(updateUserInfoErr)
		return q, updateUserInfoErr
	}
	q.updateUserInfo = string(updateUserInfo)

	deleteUser, deleteUserErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "delete_user.sql"))
	if deleteUserErr != nil {
		logger.Log(deleteUserErr)
		return q, deleteUserErr
	}
	q.deleteUser = string(deleteUser)

	createCreator, createCreatorErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "create_creator.sql"))
	if createCreatorErr != nil {
		logger.Log(createCreatorErr)
		return q, createCreatorErr
	}
	q.createCreator = string(createCreator)

	getMyCreators, getMyCreatorsErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "get_my_creators.sql"))
	if getMyCreatorsErr != nil {
		logger.Log(getMyCreatorsErr)
		return q, getMyCreatorsErr
	}
	q.getMyCreators = string(getMyCreators)

	getMyCreator, getMyCreatorErr := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "get_my_creator.sql"))
	if getMyCreatorErr != nil {
		logger.Log(getMyCreatorsErr)
		return q, getMyCreatorsErr
	}
	q.getMyCreator = string(getMyCreator)

	createWriting, createWritingError := os.ReadFile(filepath.Join("./internal/db/pg_core/query_files", "create_writing.sql"))
	if createWritingError != nil {
		logger.Log(createWritingError)
		return q, createWritingError
	}
	q.createWriting = string(createWriting)

	msg := "Pg core queries loaded."
	fmt.Println(msg)

	return q, nil
}
