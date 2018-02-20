package egn

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/iamGreedy/Inspire/variable"
	sq "github.com/Masterminds/squirrel"
)


type Database struct {
	connector *sql.DB
}
type Serializer interface {
	Save() []byte
	Load(data []byte)
}
func SQLite3() *Database {
	db, err := sql.Open("sqlite", variable.Resource("db"))
	if err != nil {
		panic(err)
	}
	return &Database{
		connector:db,
	}
}

func (s *Database ) Store(key string, value Serializer) error {
	tx, err := s.connector.Begin()
	if err != nil {
		return err
	}

	tx.Commit()
}
func (s *Database ) Load(key string) (value Serializer, err error)  {

}