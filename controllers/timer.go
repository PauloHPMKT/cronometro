package timer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimerRequest struct {
	Minutes int `json:"minutes"`
}

func HandleTimer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var timerRequest TimerRequest
	err := json.NewDecoder(r.Body).Decode(&timerRequest)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	go startCountdown(timerRequest.Minutes)
}

func startCountdown(minutes int) {
	seconds := minutes * 60
	fmt.Printf("Starting countdown for %d minutes \n", minutes)

	for seconds > 0 {
		min := seconds / 60
		sec := seconds % 60
		fmt.Printf("\rTime remaining: %02d:%02d", min, sec)

		time.Sleep(1 * time.Second)

		seconds--
	}

	fmt.Println("\rTime is up!")
}
