package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Scripts struct {
    actor string
    detail string
}

func main() {
    database, err := sql.Open("sqlite3", "./database.sqlite")

    if err != nil {
        panic(err)
    }

    rows, _ := database.Query("SELECT actor, detail FROM scripts")

    var actor,details string

    for rows.Next() {
        rows.Scan(&actor, &details)
    }

}