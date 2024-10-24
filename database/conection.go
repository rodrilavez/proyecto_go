package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:proyectogo@tcp(127.0.0.1:3306)/sistema_escolar")
	if err != nil {
		log.Println("Error conectando a la base de datos:", err)
		return nil, err
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		log.Println("Error al verificar la conexión a la base de datos:", err)
		return nil, err
	}

	return db, nil
}
