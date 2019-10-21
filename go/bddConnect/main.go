package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	// Terminar conexión al terminar función
	defer db.Close()

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	// Listo, aquí ya podemos usar a db!
	fmt.Printf("Conectado correctamente")

	getAnswer(db)
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "default"
	pass := "test"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "promofarma_api"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Answer struct {
	id          int
	question_id int
	text        string
	is_correct  bool
	letter      string
}

func getAnswer(db *sql.DB) {

	db.Ping()
	fmt.Println(Answer{3, 20, "hola", true, "A"})
}
