package helper

import (
	"encoding/json"
	"math"
	"net/http"
	"time"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func GetJson(url string, target interface{}) error {

	var myClient = &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&target)
}
