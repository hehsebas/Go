package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Importar el driver de MySQL
)

var DB *sql.DB

// InitMySQL se encarga de inicializar la conexión a MySQL
func InitMySQL() {
	var err error
	// Cambia los valores de usuario, contraseña y nombre de la base de datos
	DB, err = sql.Open("mysql", "root:admin123@@tcp(localhost:3306)/movilGO")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión a MySQL exitosa")
}
