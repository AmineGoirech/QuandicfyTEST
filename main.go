package main

import (
	"fmt"
	"log"
)

func main() {

	// Établir la connexion à la base de données
	db, err := connectDB()
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données:", err)
	}
	defer db.Close()

	// Chargez les données à partir de la base de données en utilisant les fonctions de chargement
	customers, err := LoadCustomers(db)
	if err != nil {
		log.Fatal("Erreur lors du chargement des clients:", err)
	}

	prices, err := LoadPrices(db)
	if err != nil {
		log.Fatal("Erreur lors du chargement des prix:", err)
	}

	customerEvents, err := LoadCustomerEventData(db)
	if err != nil {
		log.Fatal("Erreur lors du chargement des événements clients (CustomerEventData):", err)
	}

	// Calculez le chiffre d'affaires de chaque client
	customerRevenue := CalculateCustomerRevenue(customers, customerEvents, prices)

	// Affichez le chiffre d'affaires des 10 premiers clients (au hasard)
	count := 0
	fmt.Println("le chiffre d'affaires des 10 premiers clients (au hasard)")
	for customerID, revenue := range customerRevenue {
		if count >= 10 {
			break
		}
		fmt.Printf("CustomerID: %d, Chiffre d'affaires: %.2f\n", customerID, revenue)
		count++
	}

	fmt.Println("*****************************************************")

	// Trier la carte du chiffre d'affaires des clients dans l'ordre décroissant
	sortedCustomers := SortCustomersByRevenue(customerRevenue)

	// Imprimer les 10 premiers clients
	fmt.Println("les 10 premiers clients triés")
	for i := 0; i < 10 && i < len(sortedCustomers); i++ {
		fmt.Printf("ClientID: %d, Chiffre d'affaires: %.2f\n", sortedCustomers[i].CustomerID, sortedCustomers[i].Revenue)
	}

	fmt.Println("*****************************************************")

	// Calculer les statistiques de chiffre d'affaires par quantile
	numQuantiles := 4 // Vous pouvez ajuster le nombre de quantiles selon les besoins
	quantileStats := CalculateRevenueQuantiles(customerRevenue, numQuantiles)

	// Imprimer les statistiques de chiffre d'affaires par quantile
	for q, stats := range quantileStats {
		fmt.Printf("Quantile %d: Nombre de clients = %d, CA max = %.2f, CA min = %.2f\n", q, stats.NumCustomers, stats.MaxRevenue, stats.MinRevenue)
	}

	fmt.Println("*****************************************************")

	quantile := 0.75 // Par exemple, pour le 3eme quantile
	topCustomers := SelectTopCustomers(customerRevenue, quantile)
	fmt.Printf("Nombre de clients dans le 3eme quantile: %d\n", len(topCustomers))

	fmt.Println("*****************************************************")

	premierquantile := 0.025
	numCustomersQuantile0 := CountCustomersInQuantile(customerRevenue, premierquantile)
	fmt.Printf("Nombre de clients dans le premier quantile: %d\n", numCustomersQuantile0)

	// Exporter les données de chiffre d'affaires des clients
	err = ExportCustomerRevenue(db, customerRevenue)
	if err != nil {
		log.Fatal("Erreur lors de l'exportation des données de chiffre d'affaires des clients:", err)
	}
}
