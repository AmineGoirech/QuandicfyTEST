package main

import (
	"database/sql"
	"fmt"
	"time"
)

// ExportCustomerRevenue exporte les données de chiffre d'affaires des clients dans une table 'test_export_DATE'.
func ExportCustomerRevenue(db *sql.DB, customerRevenue map[int]float64) error {
	// Récupérer la date actuelle
	date := time.Now().Format("20060102")

	// Vérifier si la table existe déjà
	tableExists := false
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'your_database' AND table_name = ?)", "test_export_"+date).Scan(&tableExists)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'existence de la table: %v", err)
	}

	// Si la table n'existe pas, la créer
	if !tableExists {
		_, err := db.Exec("CREATE TABLE test_export_" + date + " (CustomerID BIGINT, Email VARCHAR(255), Revenue DECIMAL(10, 2))")
		if err != nil {
			return fmt.Errorf("erreur lors de la création de la table: %v", err)
		}
	}

	// Préparer la requête d'insertion
	stmt, err := db.Prepare("INSERT INTO test_export_" + date + " (CustomerID, Email, Revenue) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE Revenue = VALUES(Revenue)")
	if err != nil {
		return fmt.Errorf("erreur lors de la préparation de la requête d'insertion: %v", err)
	}
	defer stmt.Close()

	// Récupérer les adresses e-mail à partir de la table CustomerData
	rows, err := db.Query("SELECT c.CustomerID, cd.ChannelValue FROM Customer c JOIN CustomerData cd ON c.CustomerID = cd.CustomerID WHERE cd.ChannelTypeID = 1")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération des e-mails à partir de la table CustomerData: %v", err)
	}
	defer rows.Close()

	// Créer une carte pour stocker les e-mails des clients
	customerEmails := make(map[int]string)

	// Parcourir les résultats et remplir la carte des e-mails
	for rows.Next() {
		var customerID int
		var email string
		err := rows.Scan(&customerID, &email)
		if err != nil {
			return fmt.Errorf("erreur lors de la lecture des résultats de la requête: %v", err)
		}
		customerEmails[customerID] = email
	}

	// Insérer les données de chiffre d'affaires des clients dans la table
	for customerID, revenue := range customerRevenue {
		email := customerEmails[customerID]
		_, err := stmt.Exec(customerID, email, revenue)
		if err != nil {
			return fmt.Errorf("erreur lors de l'insertion des données dans la table: %v", err)
		}
	}

	fmt.Println("Exportation des données de chiffre d'affaires des clients réussie!")
	return nil
}
