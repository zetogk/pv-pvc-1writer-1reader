package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		writeTimeToFile("/config/time.json")
		time.Sleep(10 * time.Second)
	}
}

func writeTimeToFile(filename string) {
	currentTime := time.Now()
	data := map[string]interface{}{
		"time": currentTime.Format(time.RFC3339),
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Time written to", filename)
}
