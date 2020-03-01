package database

import (
	"database/sql"
	"fmt"
	"log"
)

/* funcion para conectarse a la base de datos */
func GetConnectionDB() *sql.DB {
	dsn := "postgres://postgres:postgres@127.0.0.1:5432/go_api_rest?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error al conectar BD: %s", err))
	}
	return db
}