package store

import "github.com/jmoiron/sqlx"

func CreateSchema(db *sqlx.DB) error {
	var characters = `
Drop table if exists characters;
CREATE TABLE characters (
	id int,
	name text,
	birthday text,
	img text,
	status text,
	nickname text,
	portrayed text
);
`
	var quotes = `
Drop table if exists quotes;
CREATE TABLE quotes (
	id int,
	quote text,
	author text,
	series text
);`
	var episodes = `
Drop table if exists episodes;
CREATE TABLE episodes (
	id int,
	title text,
	season int,
	episode int,
	air_date text,
	series text
);
`
	tx := db.MustBegin()
	db.MustExec(characters)
	db.MustExec(quotes)
	db.MustExec(episodes)
	return tx.Commit()
}
