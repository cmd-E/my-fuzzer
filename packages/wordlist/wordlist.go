package wordlist

import (
	"bufio"
	"os"

	"github.com/cmd-e/my-fuzzer/packages/logger"
)

type Wordlist struct {
	Words []string
}

func (w *Wordlist) ReadWordlist(path string) {
	var words []string
	logger.InfoLog.Println("Reading wordlist from:", path)
	wordList, err := os.Open(path)
	if err != nil {
		logger.ErrorLog.Println("Error opening wordlist file:", err)
		return
	}
	if os.IsNotExist(err) {
		logger.ErrorLog.Println("Wordlist file does not exist:", path)
		return
	}
	defer wordList.Close()

	scanner := bufio.NewScanner(wordList)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	logger.InfoLog.Println("Successfully read", len(words), "words from the wordlist.")
	w.Words = words
}
