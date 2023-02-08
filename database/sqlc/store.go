package database

import "database/sql"

type Store interface {
	Querier
}

type SQLStore struct {
	//Using a queries struct like this is called composition. It is said to be a better
	//decision than inheritance
	//This line exactly is the composition, adding this pointer gives the Store struct the behaviour of the Queries struct. Make sesne pa
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
