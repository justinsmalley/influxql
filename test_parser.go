package main

import (
	"fmt"
	"log"
	
	"github.com/influxdata/influxql"
)

func main() {
	// Test DROP FIELD statement
	stmt1, err := influxql.ParseStatement("DROP FIELD temperature FROM sensors")
	if err != nil {
		log.Fatalf("Error parsing DROP FIELD: %v", err)
	}
	
	if dropStmt, ok := stmt1.(*influxql.DropFieldStatement); ok {
		fmt.Printf("DROP FIELD parsed successfully:\n")
		fmt.Printf("  Field: %s\n", dropStmt.Name)
		fmt.Printf("  Measurement: %s\n", dropStmt.Measurement)
		fmt.Printf("  Database: %s\n", dropStmt.Database)
		fmt.Printf("  String: %s\n", dropStmt.String())
	} else {
		log.Fatalf("Expected DropFieldStatement, got %T", stmt1)
	}
	
	// Test RENAME FIELD statement
	stmt2, err := influxql.ParseStatement("ALTER MEASUREMENT sensors RENAME FIELD temperature TO air_temperature")
	if err != nil {
		log.Fatalf("Error parsing RENAME FIELD: %v", err)
	}
	
	if renameStmt, ok := stmt2.(*influxql.RenameFieldStatement); ok {
		fmt.Printf("\nRENAME FIELD parsed successfully:\n")
		fmt.Printf("  Old Name: %s\n", renameStmt.OldName)
		fmt.Printf("  New Name: %s\n", renameStmt.NewName)
		fmt.Printf("  Measurement: %s\n", renameStmt.Measurement)
		fmt.Printf("  Database: %s\n", renameStmt.Database)
		fmt.Printf("  String: %s\n", renameStmt.String())
	} else {
		log.Fatalf("Expected RenameFieldStatement, got %T", stmt2)
	}
	
	fmt.Println("\nAll tests passed!")
}
