package main

import (
	"bufio"
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/cmd-e/my-fuzzer/packages/logger"
)

var wordListPath string
var URL string

func init() {
	logger.InitLogger()
	flag.StringVar(&wordListPath, "w", "", "Absolute path to the wordlist file")
	flag.StringVar(&URL, "u", "", "URL to fuzz")
}

func main() {
	flag.Parse()
	if wordListPath == "" {
		logger.ErrorLog.Println("No wordlist path provided. Use -w to specify the path.")
		return
	}
	if URL == "" {
		logger.ErrorLog.Println("No URL provided. Use -u to specify the URL. Include FUZZ keyword in the URL for fuzzing.")
		return
	}
	words := readWordlist(wordListPath)
	fuzz(words)
}

// This function reads a wordlist from a specified path and logs the process.
func readWordlist(path string) []string {
	var words []string
	logger.InfoLog.Println("Reading wordlist from:", path)
	wordList, err := os.Open(path)
	if os.IsNotExist(err) {
		logger.ErrorLog.Println("Wordlist file does not exist:", path)
		return nil
	}
	defer wordList.Close()

	scanner := bufio.NewScanner(wordList)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	logger.InfoLog.Println("Successfully read", len(words), "words from the wordlist.")

	return words
}

func fuzz(words []string) {
	if !strings.Contains(URL, "FUZZ") {
		logger.ErrorLog.Println("URL does not contain 'FUZZ' keyword for fuzzing.")
		return
	}
	logger.InfoLog.Println("Starting fuzzing with", len(words), "words against URL:", URL)
	client := &http.Client{}
	for i, word := range words {
		if i%1000 == 0 {
			logger.InfoLog.Println("Fuzzing word", i+1, "of", len(words), ":", word)
		}
		fuzzUrl := strings.Replace(URL, "FUZZ", word, -1)
		request, err := http.NewRequest("GET", fuzzUrl, nil)
		if err != nil {
			logger.ErrorLog.Println("Error creating HTTP request:", err)
			continue
		}
		response, err := client.Do(request)
		if err != nil {
			logger.ErrorLog.Println("Error making HTTP request:", err)
			continue
		}
		if response.StatusCode == http.StatusOK {
			logger.InfoLog.Println("Found valid response for word:", word, "Status Code:", response.StatusCode)
		}
	}
}
