package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Data struct {
	Wind   float64 `json:"wind"`
	Water  float64 `json:"water"`
	Status string  `json:"status"`
}

func fetchData() {
	for {
		resp, err := http.Get("http://localhost:8080/update") // Example values
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Printf("Received Data:\nWater: %.2f meters\nWind: %.2f meters/sec\nStatus: %s\n\n", data.Water, data.Wind, data.Status)

		time.Sleep(15 * time.Second)
	}
}

func main() {
	fmt.Println("Client is running...")
	fetchData()
}
