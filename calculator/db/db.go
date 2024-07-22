package db

import (
	"calculator/models"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "calculation"
  )


func InitDB() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        return nil, err
    }

	fmt.Println("Successfully connected!")

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS calculation_results (
        ID SERIAL PRIMARY KEY,
        FirstValue FLOAT NOT NULL,
        SecondValue FLOAT NOT NULL,
        result FLOAT NOT NULL
    );
    `

    _, err = db.Exec(createTableSQL)

	if err != nil {
        panic(err)
    }

	if err = db.Ping(); err != nil {
        return nil, err
    }

    fmt.Println("Table created successfully!")

    return db, nil
}

func AddCalculation(user models.Calculation,dbConn *sql.DB) (*models.Calculation, error) {
	var userDb models.Calculation

	qStr := `INSERT INTO calculation_results (id,firstvalue,secondvalue, result) VALUES ($1,$2,$3, $4) RETURNING *`

	if dbConn == nil {
        fmt.Println("1111111database connection is not initialized")
    }

	err := dbConn.QueryRow(qStr, "2", user.FirstValue, user.SecondValue, user.Result).Scan(
		&userDb.ID,
		&userDb.FirstValue,
		&userDb.SecondValue,
		&userDb.Result,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Error IS ::", err)

	return &userDb, nil
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
