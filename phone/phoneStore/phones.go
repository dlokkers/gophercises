package phones

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "woReni9u"
	dbname   = "phone_numbers"
)

var db *sql.DB

type PhoneNumber struct {
	Id int
	Number string
}

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable dbname=%s",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS phoneNumbers (id SERIAL, phoneNumber VARCHAR NOT NULL)`)
	if err != nil {
		panic(err)
	}
}

func InsertPhone(phoneNumber string) (int, error) {
	var id int

	err := db.QueryRow(`INSERT INTO phoneNumbers(phoneNumber) VALUES($1) RETURNING id`, phoneNumber).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetAllNumbers() ([]PhoneNumber, error) {
	var numbers []PhoneNumber

	rows, err := db.Query(`SELECT id, phoneNumber FROM phoneNumbers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var p PhoneNumber
		if err := rows.Scan(&p.Id, &p.Number); err != nil {
			return nil, err
		} else {
			numbers = append(numbers, p)
		}
	}
	return numbers, nil
}
