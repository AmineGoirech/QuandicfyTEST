package main

import (
	"fmt"
	"sort"
	"time"
)

// CalculateCustomerRevenue calcule le chiffre d'affaires total de chaque client depuis le 01/04/2020.
func CalculateCustomerRevenue(customers []Customer, customerEvents []CustomerEventData, prices []Price) map[int]float64 {
	// Initialise une carte pour stocker le chiffre d'affaires de chaque client
	customerRevenue := make(map[int]float64)

	// Parcourt tous les clients
	for _, customer := range customers {
		// Initialise le chiffre d'affaires du client à 0
		customerRevenue[customer.CustomerID] = 0

		// Parcourt tous les événements client associés à ce client
		for _, event := range customerEvents {
			// Vérifie si l'événement est postérieur au 01/04/2020 et si c'est un achat
			eventDate, err := time.Parse("2006-01-02", event.EventDate[:10])
			if err != nil {
				fmt.Println("Erreur lors de la conversion de la date:", err)
				continue
			}
			if event.CustomerID == customer.CustomerID && eventDate.After(time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC)) && event.EventTypeID == 6 {
				// Recherche le prix du contenu associé à cet événement
				for _, price := range prices {
					if event.ContentID == price.ContentID {
						// Calcule le chiffre d'affaires et l'ajoute au total du client
						revenue := float64(event.Quantity) * price.Price
						customerRevenue[customer.CustomerID] += revenue
						break
					}
				}
			}
		}
	}

	return customerRevenue
}

// Fonction pour trier les clients par chiffre d'affaires
func SortCustomersByRevenue(customerRevenue map[int]float64) []struct {
	CustomerID int
	Revenue    float64
} {
	var sortedCustomers []struct {
		CustomerID int
		Revenue    float64
	}
	for k, v := range customerRevenue {
		sortedCustomers = append(sortedCustomers, struct {
			CustomerID int
			Revenue    float64
		}{k, v})
	}
	sort.Slice(sortedCustomers, func(i, j int) bool {
		return sortedCustomers[i].Revenue > sortedCustomers[j].Revenue
	})
	return sortedCustomers
}

// Fonction pour sélectionner les clients du premier quantile
func SelectTopCustomers(customerRevenue map[int]float64, quantile float64) map[int]float64 {
	sortedCustomers := SortCustomersByRevenue(customerRevenue)
	numCustomers := len(sortedCustomers)
	numTopCustomers := int(float64(numCustomers) * quantile)
	topCustomers := make(map[int]float64)
	for i := 0; i < numTopCustomers; i++ {
		customerID := sortedCustomers[i].CustomerID
		topCustomers[customerID] = customerRevenue[customerID]
	}
	return topCustomers
}

// Structure pour stocker les statistiques de chiffre d'affaires par quantile
type QuantileStats struct {
	NumCustomers int
	MaxRevenue   float64
	MinRevenue   float64
	Quantile     float64
}

// Fonction pour calculer les statistiques de chiffre d'affaires par quantile
func CalculateRevenueQuantiles(customerRevenue map[int]float64, numQuantiles int) map[int]QuantileStats {
	quantileStats := make(map[int]QuantileStats)
	sortedCustomers := SortCustomersByRevenue(customerRevenue)
	numCustomers := len(sortedCustomers)
	quantileSize := numCustomers / numQuantiles
	for i := 0; i < numQuantiles; i++ {
		start := i * quantileSize
		end := (i + 1) * quantileSize
		if i == numQuantiles-1 {
			end = numCustomers
		}
		quantileCustomers := sortedCustomers[start:end]
		maxRevenue := quantileCustomers[0].Revenue
		minRevenue := quantileCustomers[len(quantileCustomers)-1].Revenue
		quantileStats[i+1] = QuantileStats{
			NumCustomers: len(quantileCustomers),
			MaxRevenue:   maxRevenue,
			MinRevenue:   minRevenue,
			Quantile:     float64(i+1) / float64(numQuantiles),
		}
	}
	return quantileStats
}

//Cette fonction compte le nombre de clients dont le chiffre d'affaires est inférieur ou égal au quantile spécifié.
func CountCustomersInQuantile(customerRevenue map[int]float64, quantile float64) int {
	numCustomers := 0
	for _, revenue := range customerRevenue {
		if revenue <= quantile {
			numCustomers++
		}
	}
	return numCustomers
}
