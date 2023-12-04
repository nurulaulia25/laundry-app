package main

import (
	"database/sql"
	"enigma-laundry/entity"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Tingkatkanvalue25"
	dbname   = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	// view
	/*customers := viewCustomers()
	  for _, customer := range customers {
	      entryDateStr := customer.EntryDate.Format("2006-01-02")
	      outDateStr := customer.OutDate.Format("2006-01-02")
	      fmt.Printf("%d %s %s %s %s %.2f\n", customer.Id, customer.Name, customer.Phone, entryDateStr, outDateStr, customer.Bill)
	  }*/
	// Update customers
	/*customer := entity.Customers{
		Id:        67890,
		Name:      "Cyntia",
		Phone:     "0897375863",
		EntryDate: time.Date(2022, 8, 28, 0, 0, 0, 0, time.Local),
		OutDate:   time.Date(2022, 8, 30, 0, 0, 0, 0, time.Local),
		Bill:      21000,
	}
	addUpdateCustomers(customer)*/

	//deletecustomers
	/*err := deleteCustomers("67890")
	  if err != nil {
	      fmt.Println("Error:", err)
	  }*/

	//select table master services
	/*services := viewServices()
	  for _, service := range services {
	      fmt.Printf("%d %s %.2f\n", service.Id, service.Name, service.Price)
	  }*/

	//update service
	/*service := entity.Services{
	      Id: 4,
	      Name: "Cuci saja",
	      Price: 5000,
	  }
	  addUpdateServices(service)*/

	//delete service
	/*err := deleteService("4")
	  if err != nil {
	      fmt.Println("Error:", err)
	  }*/

	// select
	/*transactions := viewTransactions()
	for _, transaction := range transactions {
		entryDate := transaction.DateEntry.Format("2006-01-02")
		fmt.Printf("%d %d %d %s %s %.2f %.2f\n", transaction.CustomerId, transaction.ServiceId, transaction.Quantity, transaction.Unit, entryDate, transaction.Price, transaction.TotalPrice)
	}*/
	// select total price customer_id
	customerID := 12345
	totalAmount, err := sumTotalPriceByCustomerId(customerID)
	if err != nil {
		fmt.Println("Error fetching total amount:", err)
	} else {
		fmt.Printf("Total amount for customer %d: %.2f\n", customerID, totalAmount)
	}
}

func viewCustomers() []entity.Customers {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT *  FROM customers"
	rows, err := db.Query(sqlStatement)
	handleError(err, "Query Data")
	defer rows.Close()

	var customers []entity.Customers

	for rows.Next() {
		customer := entity.Customers{}
		var entryDate, outDate time.Time
		var bill sql.NullFloat64

		err = rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &entryDate, &outDate, &bill)
		handleError(err, "Scan Data")

		customer.EntryDate = entryDate
		customer.OutDate = outDate

		if bill.Valid {
			customer.Bill = bill.Float64
		} else {
			customer.Bill = 0
		}

		customers = append(customers, customer)
		err = rows.Err()
		handleError(err, "Rows Error")
	}
	return customers
}

func addUpdateCustomers(customer entity.Customers) {
	db := connectDb()
	defer db.Close()

	var err error
	sqlStatement := "INSERT INTO customers(customer_id, customer_name, customer_phone, entry_date, out_date, customer_bill) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT(customer_id) DO UPDATE SET customer_name = $2, customer_phone = $3, entry_date = $4, out_date = $5, customer_bill = $6"

	entryDateStr := customer.EntryDate.Format("2006-01-02")
	outDateStr := customer.OutDate.Format("2006-01-02")

	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone, entryDateStr, outDateStr, customer.Bill)
	handleError(err, "Insert or Update Data")
	rowCount, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Error getting rows affected: %s\n", err)
		return
	}

	if rowCount > 0 {
		fmt.Println("Update successful!")
	} else {
		fmt.Println("No rows were updated. Possibly no matching row found.")
	}
}

func deleteCustomers(customerID string) error {
	db := connectDb()
	defer db.Close()

	var err error
	sqlStatement := "DELETE FROM customers WHERE customer_id = $1"
	result, err := db.Exec(sqlStatement, customerID)
	if err != nil {
		handleError(err, "Delete Data")
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		handleError(err, "Error getting rows affected")
		return err
	}

	if rowCount > 0 {
		fmt.Println("Delete successful!")
		return nil
	} else {
		fmt.Println("No rows were deleted. Possibly no matching row found.")
		return errors.New("No rows were deleted")
	}
}

func viewServices() []entity.Services {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM services"
	rows, err := db.Query(sqlStatement)
	handleError(err, "Query Data")
	defer rows.Close()

	var services []entity.Services

	for rows.Next() {
		service := entity.Services{}

		err = rows.Scan(&service.Id, &service.Name, &service.Price)
		handleError(err, "Scan Data")

		services = append(services, service)
		err = rows.Err()
		handleError(err, "Rows Error")
	}
	return services
}

func addUpdateServices(service entity.Services) {
	db := connectDb()
	defer db.Close()

	var err error
	sqlStatement := "INSERT INTO services(service_id, service_name, price) VALUES ($1, $2, $3) ON CONFLICT(service_id) DO UPDATE SET service_name =$2, price = $3"

	result, err := db.Exec(sqlStatement, service.Id, service.Name, service.Price)
	handleError(err, "Insert or Update Data")
	rowCount, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Eror getting rows affected: %s\n", err)
		return
	}
	if rowCount > 0 {
		fmt.Println("Update successful!")
	} else {
		fmt.Println("No rows were updated. Possibly no matching row found.")
	}

}

func deleteService(serviceId string) error {
	db := connectDb()
	defer db.Close()

	var err error
	sqlStatement := "DELETE FROM services WHERE service_id = $1"

	result, err := db.Exec(sqlStatement, serviceId)

	if err != nil {
		handleError(err, "Error getting rows affected")
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		handleError(err, "Error getting rows affected")
		return err
	}
	if rowCount > 0 {
		fmt.Println("Delete successful!")
		return nil
	} else {
		fmt.Println("No rows were deleted. Possibly no matching row found.")
		return errors.New("No rows were deleted")
	}
}

func viewTransactions() []entity.Transactions {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM transactions"
	rows, err := db.Query(sqlStatement)
	handleError(err, "Query Data")
	defer rows.Close()

	var transactions []entity.Transactions

	for rows.Next() {
		transaction := entity.Transactions{}
		var dateEntry time.Time
		var price, totalPrice sql.NullFloat64

		err = rows.Scan(
			&transaction.TransactionId,
			&transaction.CustomerId,
			&transaction.ServiceId,
			&transaction.Quantity,
			&transaction.Unit,
			&dateEntry,
			&price,
			&totalPrice,
		)
		handleError(err, "Scan Data")

		transaction.DateEntry = dateEntry

		if price.Valid && totalPrice.Valid {
			transaction.Price = price.Float64
			transaction.TotalPrice = totalPrice.Float64
		} else {
			transaction.Price = 0
			transaction.TotalPrice = 0
		}

		transactions = append(transactions, transaction)
	}

	err = rows.Err()
	handleError(err, "Rows Error")

	return transactions
}

func sumTotalPriceByCustomerId(customerId int) (float64, error) {
	db := connectDb()
	defer db.Close()

	query := "SELECT SUM(total_price) FROM transactions WHERE customer_id = $1"

	var totalAmount sql.NullFloat64
	err := db.QueryRow(query, customerId).Scan(&totalAmount)

	if err != nil {
		return 0, err
	}

	return totalAmount.Float64, nil
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	handleError(err, "Open Database")
	err = db.Ping()
	handleError(err, "Ping Database")
	fmt.Println("Success Connected")
	return db
}

func handleError(err error, action string) {
	if err != nil {
		fmt.Printf("Error %s: %s\n", action, err)
		panic(err)
	}
}
