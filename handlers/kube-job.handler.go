package handlers

import (
	"blogs/kube-jobs/job"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct to hold the parsed JSON body
type bodyStruct struct {
    Concurrency int `json:"concurrency"`
}

func HandleKubeJobPost(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }

   /*
		At this point you can have custom logic
		Maybe a file upload service,
		maybe some other service. For this example,
		we will just take a concurrency value from the body
   */

	var parsedBody bodyStruct
	err = json.Unmarshal(body, &parsedBody)
	if err!= nil{
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = job.SpawnBasicJob(parsedBody.Concurrency)
	if err!=nil{
		http.Error(w, "failed creating jobs", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
    // Send a response
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "POST data received successfully!")
}