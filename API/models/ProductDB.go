package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbuser = "root"
const dbpass = "VladBerkut17"
const dbname = "deliveries"

var element_id = 0

func GetProviders() []Provider {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:5917)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM provider")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	providers := []Provider{}

	for results.Next() {

		var prov Provider

		err = results.Scan(&prov.ProviderId, &prov.ProviderName, &prov.INN, &prov.ContactDetails, &prov.RF)

		if err != nil {
			panic(err.Error())
		}

		providers = append(providers, prov)

		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)
	}

	return providers

}

func GetProvider(provider_name string) []Provider {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:5917)/"+dbname)

	providers := []Provider{}

	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM deliveries.Provider WHERE provider_name = ?", provider_name)
	if err != nil {

		fmt.Println("Err", err.Error())
		return nil
	}

	for results.Next() {
		var prov Provider
		err = results.Scan(&prov.ProviderId, &prov.ProviderName, &prov.INN, &prov.ContactDetails, &prov.RF)

		if err != nil {
			panic(err.Error())
		}
		providers = append(providers, prov)
	}

	return providers

}

func AddProvider(product Provider) {
	element_id++
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:5917)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO Provider (Provider_name,INN,contact_details,RF) VALUES (?,?,?,?)",
		product.ProviderName, product.INN, product.ContactDetails, product.RF)

	/*
		// Or use fmt.Sprintf to concatenate SQL statement if prepared statement isn't worth here

		sqlstm :=
			fmt.Sprintf("INSERT INTO product (code,name,qty,last_updated)"+
				" VALUES ('%s','%s',%d, now())",
				product.Code, product.Name, product.Qty)

		insert, err := db.Query(sqlstm)
	*/

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func AddProduct(product Product) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:5917)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO Offers (element_id,element_name,quantity,price,amount,deadline) VALUES (?,?,?,?,?,?)",
		element_id, product.ElementName, product.Quantity, product.Price, product.Amount, product.Deadline)

	/*
		// Or use fmt.Sprintf to concatenate SQL statement if prepared statement isn't worth here

		sqlstm :=
			fmt.Sprintf("INSERT INTO product (code,name,qty,last_updated)"+
				" VALUES ('%s','%s',%d, now())",
				product.Code, product.Name, product.Qty)

		insert, err := db.Query(sqlstm)
	*/

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func GetProducts() []Product {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:5917)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM offers")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	product := []Product{}

	for results.Next() {

		var prod Product

		err = results.Scan(&prod.ProviderId, &prod.ElementId, &prod.ElementName, &prod.Quantity, &prod.Price, &prod.Amount, &prod.Deadline)

		if err != nil {
			panic(err.Error())
		}

		product = append(product, prod)

		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)
	}

	return product

}
