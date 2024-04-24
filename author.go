package main

import (
	"context"
	"hainguyen/authorrepo"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://your_username:your_password@localhost:5432/golang")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := authorrepo.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Debug().Msgf("Authors: %v", authors)
	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, authorrepo.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Debug().Msgf("Inserted author: %v", insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}
	log.Debug().Msgf("Fetched author: %v", fetchedAuthor)
	return nil
}

func queryAuthor() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msg("Error running the program")
	}
}
