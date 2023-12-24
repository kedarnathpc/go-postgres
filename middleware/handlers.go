package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/kedarnathpc/go-postgres/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// response is a struct for encoding JSON responses.
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// createConnection establishes a connection to the PostgreSQL database.
func createConnection() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to PostgreSQL database using the provided URL
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	// Ping the database to ensure a successful connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging DB:", err)
	}

	log.Println("Successfully connected to Postgres...")
	return db
}

// CreateStock handles the creation of a new stock record.
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	// Decode the request body into the Stock struct
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the request body: %v", err)
	}

	// Insert the stock into the database
	insertID := insertStock(stock)

	// Prepare a response and encode it as JSON
	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

// GetStock retrieves a stock by its ID.
func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert the string into an integer: %v", err)
	}

	// Retrieve the stock from the database
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatal("Unable to get stock: %v", err)
	}

	// Encode the stock as JSON and send the response
	json.NewEncoder(w).Encode(stock)
}

// GetAllStock retrieves all stocks from the database.
func GetAllStock(w http.ResponseWriter, r *http.Request) {
	// Retrieve all stocks from the database
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatal("Unable to get all stocks: %v", err)
	}

	// Encode the stocks as JSON and send the response
	json.NewEncoder(w).Encode(stocks)
}

// UpdateStock updates a stock record by its ID.
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert the string to an integer: %v", err)
	}

	var stock models.Stock

	// Decode the request body into the Stock struct
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the request body: %v", err)
	}

	// Update the stock in the database
	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected: %v", updatedRows)

	// Prepare a response and encode it as JSON
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteStock deletes a stock record by its ID.
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert string to integer: %v", err)
	}

	// Delete the stock from the database
	deletedRows := deleteStock(int64(id))

	// Prepare a response and encode it as JSON
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/records affected: %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// insertStock inserts a new stock into the database and returns the inserted ID.
func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlquery := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64

	// Execute the query and scan the inserted ID
	err := db.QueryRow(sqlquery, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatal("Unable to execute query: %v", err)
	}

	fmt.Printf("Inserted a single record. ID: %v", id)
	return id
}

// getStock retrieves a stock by its ID from the database.
func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stock models.Stock
	sqlquery := `SELECT * FROM  stocks WHERE stockid=$1`

	// Execute the query and scan the result into the Stock struct
	row := db.QueryRow(sqlquery, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return stock, err
}

// getAllStocks retrieves all stocks from the database.
func getAllStocks() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stocks []models.Stock
	sqlquery := `SELECT * FROM stocks`
	row, err := db.Query(sqlquery)
	if err != nil {
		log.Fatal("Unable to execute the query: %v", err)
	}

	defer row.Close()

	// Iterate over the result set and append each stock to the slice
	for row.Next() {
		var stock models.Stock
		err = row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatal("Unable to scan the row: %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

// updateStock updates a stock in the database by its ID and returns the number of affected rows.
func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()

	sqlquery := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := db.Exec(sqlquery, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatal("Unable to execute the query: %v", err)
	}

	// Get the number of affected rows
	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/records affected: %v", rowAffected)
	return rowAffected
}

// deleteStock deletes a stock from the database by its ID and returns the number of affected rows.
func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlquery := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlquery, id)
	if err != nil {
		log.Fatal("Unable to execute the query: %v", err)
	}

	// Get the number of affected rows
	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/records affected: %v", rowAffected)
	return rowAffected
}
