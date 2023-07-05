package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input file containing subdomains
	inputFile := "all.txt"
	outputFile := "httpsonly.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var modifiedSubdomains []string

	// Process each line in the input file
	for scanner.Scan() {
		subdomain := scanner.Text()

		// Remove "http://" if present and replace with "https://"
		modified := strings.Replace(subdomain, "http://", "https://", 1)
		modifiedSubdomains = append(modifiedSubdomains, modified)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input file:", err)
		return
	}

	// Write modified subdomains to output file
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, subdomain := range modifiedSubdomains {
		_, err := writer.WriteString(subdomain + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing output file writer:", err)
		return
	}

	fmt.Println("Modified subdomains written to", outputFile)
}
