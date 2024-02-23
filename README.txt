** FIRST STEPS TO CREATE THE DB USING DOCKER *** 

1- docker volume create QuanticfyVolume // CREATION DU VOLUME POUR LA PERSIS DES DONNEES 


2- docker network create quanticfy_network


2- docker run --name QuanticfyContainerSql \
-e MYSQL_ROOT_PASSWORD=aminegoirech \
-e MYSQL_USER=aminegoirech \
-e MYSQL_DATABASE=QuanticfyDB \
-p 3306:3306 \
-v QuanticfyVolume:/var/lib/mysql \
--network quanticfy_network \
-d mysql:latest


3- docker run --name AdminerContainer -p 8080:8080 --link QuanticfyContainerSql:db --network quanticfy_network -d adminer  


4- GRANT ALL PRIVILEGES ON QuanticfyDB.* TO 'aminegoirech'@'localhost';


5- docker exec -i QuanticfyContainerSql mysql -uroot -paminegoirech QuanticfyDB < /home/gouerch/Desktop/TESTGOSQL/QuandicfyTEST/script.sql // creation des tables 

6- docker exec -i QuanticfyContainerSql mysql -uroot -paminegoirech QuanticfyDB < /home/gouerch/Desktop/TESTGOSQL/QuandicfyTEST/insertion_donnes.sql //Remplissage des tables





Objectif:

- Identifier les clients ayant généré le plus de chiffre d'affaires (CA) et les classer par quantiles.
- Exporter les meilleurs clients (premier quantile) dans une nouvelle table MySQL.
- Analyser la répartition des clients par CA.



























/*
	events, err := LoadEvents(db) 
	if err != nil {log.Fatal("Erreur lors du chargement des événements:", err)}

	contents, err := LoadContents(db)
	if err != nil { log.Fatal("Erreur lors du chargement des contenus:", err)}

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
*/