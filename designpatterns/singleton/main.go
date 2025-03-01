package main

import (
	"fmt"
	"sync"
)

// Database is a singleton struct
type Database struct {
	ConnectionString string
}

// Declare a variable to hold the singleton instance
var instance *Database

// Mutex for thread-safe access
var once sync.Once

// GetInstance returns the singleton instance of the Database
func GetInstance() *Database {
	// Ensure that the instance is initialized only once
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		instance = &Database{ConnectionString: "localhost:5432"}
	})
	return instance
}
func main() {
	// First call to GetInstance creates the singleton
	db1 := GetInstance()
	fmt.Printf("DB1: %v\n", db1.ConnectionString)

	// Subsequent calls return the same instance
	db2 := GetInstance()
	fmt.Printf("DB2: %v\n", db2.ConnectionString)

	// Verify if db1 and db2 point to the same instance
	if db1 == db2 {
		fmt.Println("Both instances are the same.")
		fmt.Printf("db1 : %p\n", db1)
		fmt.Printf("db2 : %p\n", db2)
	}
}

/*
output:
Creating single instance now.
DB1: localhost:5432
DB2: localhost:5432
Both instances are the same.
db1 : 0x14000010040
db2 : 0x14000010040
*/
