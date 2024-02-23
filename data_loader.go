package main

import (
	"database/sql"
)

// Customer représente une entrée dans la table Customer de la base de données.
type Customer struct {
	CustomerID       int
	ClientCustomerID int
	InsertDate       string
}

// Event représente une entrée dans la table CustomerEvent de la base de données.
type Event struct {
	EventID       int
	ClientEventID int
	InsertDate    string
}

// Content représente une entrée dans la table Content de la base de données.
type Content struct {
	ContentID       int
	ClientContentID int
	InsertDate      string
}

// Price représente une entrée dans la table ContentPrice de la base de données.
type Price struct {
	ContentPriceID int
	ContentID      int
	Price          float64
	Currency       string
	InsertDate     string
}

// CustomerEventData représente une entrée dans la table CustomerEventData de la base de données.
type CustomerEventData struct {
	EventDataID int
	EventID     int
	ContentID   int
	CustomerID  int
	EventTypeID int
	EventDate   string
	Quantity    int
	InsertDate  string
}

// LoadCustomers charge les données des clients à partir de la base de données.
func LoadCustomers(db *sql.DB) ([]Customer, error) {
	var customers []Customer

	rows, err := db.Query("SELECT CustomerID, ClientCustomerID, InsertDate FROM Customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.CustomerID, &customer.ClientCustomerID, &customer.InsertDate); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// LoadEvents charge les données des événements clients à partir de la base de données.
func LoadEvents(db *sql.DB) ([]Event, error) {
	var events []Event

	rows, err := db.Query("SELECT EventID, ClientEventID, InsertDate FROM CustomerEvent")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.EventID, &event.ClientEventID, &event.InsertDate); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// LoadContents charge les données des contenus à partir de la base de données.
func LoadContents(db *sql.DB) ([]Content, error) {
	var contents []Content

	rows, err := db.Query("SELECT ContentID, ClientContentID, InsertDate FROM Content")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var content Content
		if err := rows.Scan(&content.ContentID, &content.ClientContentID, &content.InsertDate); err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	return contents, nil
}

// LoadPrices charge les données des prix de contenu à partir de la base de données.
func LoadPrices(db *sql.DB) ([]Price, error) {
	var prices []Price

	rows, err := db.Query("SELECT ContentPriceID, ContentID, Price, Currency, InsertDate FROM ContentPrice")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var price Price
		if err := rows.Scan(&price.ContentPriceID, &price.ContentID, &price.Price, &price.Currency, &price.InsertDate); err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}

// LoadCustomerEventData charge les données des événements clients à partir de la base de données.
func LoadCustomerEventData(db *sql.DB) ([]CustomerEventData, error) {
	var customerEvents []CustomerEventData

	rows, err := db.Query("SELECT EventDataID, EventID, ContentID, CustomerID, EventTypeID, EventDate, Quantity, InsertDate FROM CustomerEventData")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event CustomerEventData
		if err := rows.Scan(&event.EventDataID, &event.EventID, &event.ContentID, &event.CustomerID, &event.EventTypeID, &event.EventDate, &event.Quantity, &event.InsertDate); err != nil {
			return nil, err
		}
		customerEvents = append(customerEvents, event)
	}

	return customerEvents, nil
}
