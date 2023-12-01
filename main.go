package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cinemast/mortal-gombat/db"
	_ "github.com/mattn/go-sqlite3"
)

//go:generate docker run --rm -v .:/src -w /src sqlc/sqlc generate

func connect() *db.Db {
	c, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	return db.NewDb(c)
}

func main() {
	ctx := context.Background()

	database := connect()

	a, err := database.CreateAuthor(ctx, "user@example.org")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", a)

	err = db.Tx(ctx, database, func(ctx context.Context, tx *db.Db) error {
		for i := 0; i < 30; i++ {
			e, err := tx.CreateEntry(ctx, db.CreateEntryParams{
				Title: fmt.Sprintf("Title: %d", i),
				Body: sql.NullString{
					String: "body ....",
					Valid:  true,
				},
				AuthorID: a.ID,
			})
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", e)

			for j := 0; j < 5; j++ {
				comment, err := tx.CreateComment(ctx, db.CreateCommentParams{
					Body:     "some comment",
					AuthorID: a.ID,
					EntryID:  e.ID,
				})
				if err != nil {
					return err
				}
				fmt.Printf("%+v\n", comment)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

}
