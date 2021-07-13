package misc

import (
	"errors"
	"log"
	"os"
)

//CheckErr handles all error if not nil
func CheckErr(err error, isPanic bool, customError string) {
	if err != nil {
		LogError(err)
		LogError(errors.New(customError))
	}
}

//LogError logs all error to file
func LogError(err error) {
	f, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	logger := log.New(f, "Error: ", log.LstdFlags)
	if err != nil {
		logger.Println(err.Error())
	}
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

//set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
