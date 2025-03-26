package config

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // Importa el driver de PostgreSQL
)

var DB *sql.DB

// ConectarDB establece la conexión a la base de datos
func ConectarDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://admin:admin123@localhost:5432/marketplace_db")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// Verifica la conexión
	if err := DB.Ping(); err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	log.Println("Conexión exitosa a la base de datos")
}
