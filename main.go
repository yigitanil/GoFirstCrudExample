package main

import "GoFirstCrudExample/config"

func main() {

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDb.Close()

	config.CreateServer(db)
}
