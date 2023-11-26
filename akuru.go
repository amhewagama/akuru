package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"bufio"
	
	_ "github.com/mattn/go-sqlite3"
)

const dbName = "data.db" // Change this to your desired database name

var db *sql.DB

func main() {
	dbPath := "./" + dbName

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		// Read word from user input
		fmt.Print("Enter a word (or 'exit' to quit): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		word := strings.TrimSpace(scanner.Text())

		// Check if the user wants to exit
		if strings.ToLower(word) == "exit" {
			fmt.Println("Exiting the program.")
			break
		}

		// Validate that a word was entered
		if word == "" {
			fmt.Println("Please enter a valid word.")
			continue
		}

		results, err := searchWord(word)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Search Results:")
		for _, result := range results {
			var final1 = DecodeString(result)
			SplitAndPrint(final1)
		}
	}
}
func DecodeString(input string) string {
	var result string
	for _, c := range input {
		decodedChar := int(c) - 3
		result += string(decodedChar)
	}
	return result
}

func SplitAndPrint(input string) {
	parts := strings.Split(input, "|")
	for _, part := range parts {
		fmt.Println(part)
	}
}

func searchWord(word string) ([]string, error) {
	rows, err := db.Query("SELECT value FROM dict WHERE key=?", word)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var result string
		if err := rows.Scan(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
