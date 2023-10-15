package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"k8s.io/api/admission/v1beta1"
)

// Structure for JSON patch operation
type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

// Webhook handler function
func handleAddMarvelLabel(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling webhook request ...")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	// Deserialize the AdmissionReview request
	var admissionReviewReq v1beta1.AdmissionReview
	if _, _, err := deserializer.Decode(body, nil, &admissionReviewReq); err != nil {
		http.Error(w, "could not deserialize request", http.StatusBadRequest)
		return
	}

	// Fetch a random Marvel character name
	marvelName, err := getRandomMarvelName()
	if err != nil {
		http.Error(w, "could not fetch Marvel name", http.StatusInternalServerError)
		return
	}

	// Create JSON patch operation to add Marvel label
	patchOps := []patchOperation{
		{
			Op:    "add",
			Path:  "/metadata/labels",
			Value: map[string]string{"marvel": marvelName},
		},
	}

	// Construct the AdmissionReview response
	admissionReviewResponse := v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			UID:     admissionReviewReq.Request.UID,
			Allowed: true,
			Patch:   patchOps,
		},
	}

	// Serialize the AdmissionReview response to JSON
	responseBytes, err := json.Marshal(admissionReviewResponse)
	if err != nil {
		http.Error(w, "could not serialize response", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
	log.Println("Webhook request handled successfully")
}
