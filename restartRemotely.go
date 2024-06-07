package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const (
	remoteURL     = "https://checkitservices.xyz:3000/should_restart"
	checkInterval = 60 * time.Second
	scriptPath    = "./restart_push_notification.sh"
)

type RestartResponse struct {
	ShouldRestart bool `json:"should_restart"`
}

func main() {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for range ticker.C {
		shouldRestart, err := checkShouldRestart()
		if err != nil {
			log.Printf("Error checking restart status: %v", err)
			continue
		}

		if shouldRestart {
			log.Println("Restarting pushNotification...")
			if err := restartPushNotification(); err != nil {
				log.Printf("Error restarting pushNotification: %v", err)
			} else {
				log.Println("pushNotification restarted successfully.")
			}
		}
	}
}

func checkShouldRestart() (bool, error) {
	resp, err := http.Get(remoteURL)
	if err != nil {
		return false, fmt.Errorf("failed to GET %s: %v", remoteURL, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response body: %v", err)
	}

	var restartResp RestartResponse
	if err := json.Unmarshal(body, &restartResp); err != nil {
		return false, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return restartResp.ShouldRestart, nil
}

func restartPushNotification() error {
	cmd := exec.Command(scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute script: %v\n%s", err, output)
	}
	log.Printf("Script output: %s", output)
	return nil
}
