package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type ImageResponse struct {
	ImageURLs []string `json:"image_urls"`
}

func main() {
	http.HandleFunc("/generate-lgtm", handleGenerateLGTM)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handleGenerateLGTM(w http.ResponseWriter, r *http.Request) {
	imageURLs, err := generateLGTMImages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ImageResponse{
		ImageURLs: imageURLs,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func generateLGTMImages() ([]string, error) {
	endpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	apiKey := os.Getenv("AZURE_OPENAI_API_KEY")

	cred := azcore.NewKeyCredential(apiKey)

	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	req := azopenai.ImageGenerationRequest{
		Prompt: "LGTM image for GitHub pull requests",
		N:      5,                         // 5つの画像候補を生成
		Size:   azopenai.ImageSize512x512, // 修正後の定数名
	}

	resp, err := client.GenerateImage(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error generating images: %v", err)
	}

	var imageUrls []string
	for _, data := range resp.Data {
		imageUrls = append(imageUrls, data.URL)
	}

	return imageUrls, nil
}
