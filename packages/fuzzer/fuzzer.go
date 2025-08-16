package fuzzer

import (
	"net/http"
	"strings"

	"github.com/cmd-e/my-fuzzer/packages/logger"
)

type Fuzzer struct {
	Words []string
	URL   string
}

func (f Fuzzer) Fuzz() {
	if !strings.Contains(f.URL, "FUZZ") {
		logger.ErrorLog.Println("URL does not contain 'FUZZ' keyword for fuzzing.")
		return
	}
	logger.InfoLog.Println("Starting fuzzing with", len(f.Words), "words against URL:", f.URL)
	client := getHttpClient()
	for i, word := range f.Words {
		if i%1000 == 0 {
			logger.InfoLog.Println("Fuzzing word", i+1, "of", len(f.Words), ":", word)
		}
		fuzzUrl := strings.Replace(f.URL, "FUZZ", word, 1)
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

func getHttpClient() *http.Client {
	return &http.Client{}
}
