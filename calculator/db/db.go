package db

import (
	"calculator/models"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

// var DB *sql.DB

var dbConn *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "calculation"
  )

func InitDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Error")
	} 

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS calculation_results (
        ID SERIAL PRIMARY KEY,
        FirstValue FLOAT NOT NULL,
        SecondValue FLOAT NOT NULL,
        result FLOAT NOT NULL
    );
    `

    _, err = db.Exec(createTableSQL)

	dbConn = db
}

func AddCalculation(user models.Calculation) (*models.Calculation) {
	var userDb models.Calculation

	qStr := `INSERT INTO calculation_results (id,firstvalue,secondvalue, result) VALUES ($1,$2,$3, $4) RETURNING *`

	if dbConn == nil {
        fmt.Println("No database connection!")
    }
	
	err := dbConn.QueryRow(qStr, "2", user.FirstValue, user.SecondValue, user.Result).Scan(
		&userDb.ID,
		&userDb.FirstValue,
		&userDb.SecondValue,
		&userDb.Result,
	)
	
	return &userDb
}

// func (c *DB) UpdateCalculation(id string, user models.User) (*models.User, error) {
// 	var userDb models.User

// 	qStr := `UPDATE users SET Firstvalue = $1, Secondvalue = $2, result = $3 WHERE id = $4 RETURNING *`
// 	err := c.Pg.QueryRowx(qStr, user.FirstName, user.LastName, result, id).StructScan(&userDb)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &userDb, nil
// }

// func (c *DB) GetCalculationByID(id string) (*models.User, error) {
// 	var userDb models.User

// 	qStr := `SELECT * FROM users WHERE id = $1`
// 	err := c.Pg.QueryRowx(qStr, id).StructScan(&userDb)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &userDb, nil
// }

// func (c *DB) DeleteCalculationByID(id string) error {

// 	qStr := `DELETE FROM users WHERE id = $1`
// 	_, err := c.Pg.Exec(qStr, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
