package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

// deserializer is used to deserialize AdmissionReview requests
var deserializer = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()

// patchOperation represents a JSON patch operation
type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

// Handler function for the /add-marvel-label endpoint
func handleAddMarvelLabel(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling webhook request ...")

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	// Deserialize the request body into an AdmissionReview object
	var admissionReviewReq v1beta1.AdmissionReview
	if _, _, err := deserializer.Decode(body, nil, &admissionReviewReq); err != nil {
		http.Error(w, "could not deserialize request", http.StatusBadRequest)
		return
	}

	// Extract the raw Pod object from the AdmissionReview request
	raw := admissionReviewReq.Request.Object.Raw

	// Initialize a map to unmarshal the raw Pod object into
	var pod map[string]interface{}
	if err := json.Unmarshal(raw, &pod); err != nil {
		// Return an error if unmarshalling fails
		http.Error(w, "could not unmarshal raw object", http.StatusBadRequest)
		return
	}

	// Extract the metadata from the Pod object
	metadata, ok := pod["metadata"].(map[string]interface{})
	if !ok {
		// Return an error if metadata is not found
		http.Error(w, "could not get metadata", http.StatusBadRequest)
		return
	}

	// Check if the Pod has any owner references (i.e., is managed by a controller)
	ownerReferences, ok := metadata["ownerReferences"].([]interface{})

	// If the Pod has owner references, skip the webhook logic
	if ok && len(ownerReferences) > 0 {
		log.Println("Skipping Pod managed by a controller")
		// Construct the AdmissionReview response
		admissionReviewResponse := v1beta1.AdmissionReview{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "admission.k8s.io/v1",
				Kind:       "AdmissionReview",
			},
			Response: &v1beta1.AdmissionResponse{
				UID:     admissionReviewReq.Request.UID,
				Allowed: true, // Indicate that the request is allowed
			},
		}

		// Serialize the AdmissionReview response to JSON
		responseBytes, err := json.Marshal(admissionReviewResponse)
		if err != nil {
			http.Error(w, "could not serialize response", http.StatusInternalServerError)
			log.Println("could not serialize response:", err)
			return
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		log.Println("Webhook request handled successfully, skipping Pod managed by a controller")
		return
	}

	// Generate Marvel API URL dynamically
	ts := fmt.Sprintf("%v", time.Now().Unix())
	publicKey := "84f75d5854abb64040a580afa56dd9c0"
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")
	if privateKey == "" {
		http.Error(w, "missing private key", http.StatusInternalServerError)
		log.Println("MARVEL_PRIVATE_KEY environment variable not set")
		return
	}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(ts+privateKey+publicKey)))
	apiURL := constructMarvelAPIURL(ts, publicKey, hash)

	// Fetch a random Marvel name
	maxOffset := 800 // Maximum offset for the Marvel API
	marvelName, err := getRandomMarvelName(apiURL, maxOffset)
	if err != nil {
		http.Error(w, "could not fetch Marvel name", http.StatusInternalServerError)
		log.Println("could not fetch Marvel name:", err)
		return
	}

	// Sanitize the Marvel name to make it a valid Kubernetes label
	sanitizedMarvelName := sanitizeLabel(marvelName)

	// Construct the JSON patch operations
	patchOps := []patchOperation{
		{
			Op:    "add",
			Path:  "/metadata/labels",
			Value: map[string]string{"marvel": sanitizedMarvelName},
		},
	}

	// Serialize the JSON patch operations to JSON
	patchBytes, err := json.Marshal(patchOps)
	if err != nil {
		http.Error(w, "could not serialize patch operations", http.StatusInternalServerError)
		log.Println("could not serialize patch operations:", err)
		return
	}

	// Construct the AdmissionReview response
	admissionReviewResponse := v1beta1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Response: &v1beta1.AdmissionResponse{
			UID:     admissionReviewReq.Request.UID,
			Allowed: true,
			Patch:   patchBytes,
			PatchType: func() *v1beta1.PatchType {
				pt := v1beta1.PatchTypeJSONPatch
				return &pt
			}(),
		},
	}

	// Serialize the AdmissionReview response to JSON
	responseBytes, err := json.Marshal(admissionReviewResponse)
	if err != nil {
		http.Error(w, "could not serialize response", http.StatusInternalServerError)
		log.Println("could not serialize response:", err)
		return
	}

	log.Printf("Sending AdmissionReview response: %s\n", string(responseBytes))

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
	log.Println("Webhook request handled successfully")
}
