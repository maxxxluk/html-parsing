package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// URL of the HTML page to parse
	url := "https://example.com"

	// Perform HTTP request to fetch the HTML content
	response, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Parse and extract information from the HTML
	info := extractInfo(response)

	// Create the file name with current date and time
	fileName := fmt.Sprintf("info_%s.txt", time.Now().Format("20060102_150405"))

	// Write the information to a text file
	if err := writeToFile(fileName, info); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Information collected and saved to", fileName)
}

// Extracts information from the HTML document
func extractInfo(doc *goquery.Document) string {
	// TODO: Implement the logic to extract information from the HTML document
	// Example: extracting title and paragraph text
	title := doc.Find("title").Text()
	paragraph := doc.Find("p").Text()

	// Concatenate the extracted information
	info := strings.Join([]string{title, paragraph}, "\n")
	return info
}

// Writes content to a file
func writeToFile(fileName string, content string) error {
	// Write the content to a temporary file
	tmpFile := fileName + ".tmp"
	err := ioutil.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		return err
	}

	// Rename the temporary file to the final name
	err = os.Rename(tmpFile, fileName)
	if err != nil {
		return err
	}

	return nil
}
