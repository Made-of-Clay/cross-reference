package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

/*
This script is a practice with Go to handle reading/writing
and formatting data that I'll use elsewhere.
*/

func main() {
	fmt.Println("hello world")
	writeBooks(formatBooks(readCsv()))
}

func readCsv() [][]string {
	file, err := os.Open("./cross_references.csv")
	if err != nil {
		log.Fatal("Unable to read input file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file")
	}

	return records
}

func formatBooks(data [][]string) string {
	books := make([]string, 0, 66)
	for _, row := range data[1:] {
		fromRef := strings.Trim(row[0], " ")
		book := strings.Split(fromRef, ".")[0]
		if !slices.Contains(books, book) {
			books = append(books, book)
		}
	}
	fmt.Println(books)
	dat, err := json.Marshal(books)
	if err != nil {
		log.Fatal("JSON formatting failed")
	}
	return string(dat)
}

func writeBooks(books string) {
	file, err := os.Create("books-go.json")
	if err != nil {
		log.Fatal("Could not create new JSON file")
	}
	defer file.Close()

	_, err = file.WriteString(books)
	if err != nil {
		panic(err)
	}
}
