package api

import (
	"fmt"

	"github.com/Maow-Nam/read-api/config"
	"gopkg.in/mgo.v2"
)

// MongoDB holds metadata about session database and collections name.
type (
	MongoDB struct {
		Conn *mgo.Session
		BCol *mgo.Collection
	}
)

// NewMongoDB creates a new macOddsTeamDB backed by a given Mongo server.
func NewMongoDB() (*MongoDB, error) {
	s := config.Spec()
	conn, err := mgo.Dial(s.DBHost)

	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	return &MongoDB{
		Conn: conn,
		BCol: conn.DB(s.DBName).C(s.DBBookCol),
	}, nil
}

// Close closes the database.
func (db *MongoDB) Close() {
	db.Conn.Close()
}
