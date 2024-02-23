package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Fonction pour établir la connexion à la base de données
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:aminegoirech@tcp(localhost:3306)/QuanticfyDB")
	if err != nil {
		return nil, err
	}

	// Vérifier la connexion à la base de données
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Connexion à la base de données réussie!")
	fmt.Println("*****************************************************")
	return db, nil
}
