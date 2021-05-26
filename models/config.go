package models

import (
	"os"

	"github.com/pramodshenkar/golang_todo_api/db"
)

var server = os.Getenv("DATABASE")
var databaseName = os.Getenv("DATABASE_NAME")
var dbConnect = db.NewConnection(server, databaseName)
