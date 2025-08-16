package main

import (
	"flag"

	"github.com/cmd-e/my-fuzzer/packages/fuzzer"
	"github.com/cmd-e/my-fuzzer/packages/logger"
	"github.com/cmd-e/my-fuzzer/packages/wordlist"
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
	wl := wordlist.Wordlist{}
	wl.ReadWordlist(wordListPath)
	fuzz(wl.Words)
}

func fuzz(words []string) {
	f := fuzzer.Fuzzer{
		Words: words,
		URL:   URL,
	}
	f.Fuzz()
}
