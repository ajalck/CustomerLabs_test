package handler

import (
	"CustomerLabsTest/model"
	"CustomerLabsTest/worker"
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

func HandleJSONRequest(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	} else if body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	inputData := make(chan map[string]interface{})
	outputData := make(chan *model.OuputData)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		go worker.Worker(inputData, outputData)
		inputData <- body
		close(inputData)

		outputbytes, err := json.Marshal(<-outputData)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		webhookURL := "https://webhook.site/48f11e84-bdb1-4d0a-a545-ffca7d28c2f0"
		resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(outputbytes))
		if err != nil {
			w.WriteHeader(resp.StatusCode)
			w.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()

		// Check the response status
		if resp.StatusCode == http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			w.Write([]byte("Webhook sent successfully"))
			return
		} else {
			w.WriteHeader(resp.StatusCode)
			w.Write([]byte("Webhook request failed with status: " + resp.Status))
			return
		}
	}()
	wg.Wait()
}
