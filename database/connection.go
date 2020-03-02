package database

import (
	"database/sql"
	"fmt"
	"log"
)

/* funcion para conectarse a la base de datos */
func GetConnectionDB() *sql.DB {
	// dsn := "postgres://postgres:postgres@127.0.0.1:5432/go_api_rest?sslmode=disable"
	dsn := "postgres://vzmemnlocgrknl:ff0e6b953b864cbb9e23f46dfe47ef2e86ca389fae4a3cddbe30ba91725ed058@ec2-52-202-185-87.compute-1.amazonaws.com:5432/ddk2ombgpk8bdq"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error al conectar BD: %s", err))
	}
	return db
}