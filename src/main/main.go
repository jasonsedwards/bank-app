package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Balance struct {
	EffectiveBalance struct {
		MinorUnits int `json:"minorUnits"`
	} `json:"effectiveBalance"`
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}

	return client
}

func getBalance(client *http.Client, method string) float64 {
	url := "https://api.starlingbank.com/api/v2/accounts/" + os.Getenv("account_id") + "/balance"
	bearer := "Bearer " + os.Getenv("token")

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Error while making request", err)
	}

	req.Header.Add("Authorization", bearer)

	res, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	// Close the connection to reuse it
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error while closing the connection", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var response Balance

	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0
	}

	bal := float64(response.EffectiveBalance.MinorUnits) / 100

	return bal
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	// c should be re-used for further calls
	c := httpClient()
	response := getBalance(c, http.MethodGet)

	msg := "Balance: £" + strconv.FormatFloat(response, 'f', 2, 64)

	log.Println(msg)
}
