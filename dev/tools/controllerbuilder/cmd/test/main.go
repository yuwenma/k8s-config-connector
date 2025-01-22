package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/vertexai/genai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
)

// getWeather is a placeholder for a real function that would fetch weather data.
func getWeather(location string) (string, error) {
	// In a real implementation, you would call a weather API here.
	fmt.Printf("Calling getWeather for location: %s\n", location)
	switch location {
	case "London":
		return "Sunny, 25°C", nil
	case "Paris":
		return "Cloudy, 20°C", nil
	default:
		return "", fmt.Errorf("unknown location: %s", location)
	}
}

func getFunctionDeclaration() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "getWeather",
		Description: "Get the current weather for a given location.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"location": {
					Type:        genai.TypeString,
					Description: "The city and state, e.g. San Francisco, CA",
				},
			},
			Required: []string{"location"},
		},
	}
}

func getTool() *genai.Tool {
	functionDeclaration := getFunctionDeclaration()
	return &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{functionDeclaration},
	}
}

func main() {
	ctx := context.Background()
	client, err := llm.BuildGeminiClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	tools := []*genai.Tool{getTool()}
	model := client.GenerativeModel("gemini-pro")
	model.Tools = tools
	cs := model.StartChat()

	// Initial history (optional)
	cs.History = []*genai.Content{
		genai.NewUserContent(genai.Text("What's the weather like?")),
		genai.NewUserContent(genai.Text("Where are you located?")),
	}

	send := func(msg string) *genai.GenerateContentResponse {
		fmt.Println("== User: ==")
		fmt.Println(msg)

		res, err := cs.SendMessage(ctx, genai.Text(msg))
		if err != nil {
			log.Fatal(err)
		}
		return res
	}

	res := send("What's the weather like in London?")
	printResponse(ctx, cs, res)

}
func sendFunc(ctx context.Context, cs *genai.ChatSession, location string, weather string) *genai.GenerateContentResponse {
	res, err := cs.SendMessage(ctx, genai.FunctionResponse{
		Name: "getWeather",
		Response: map[string]interface{}{
			"location": location,
			"weather":  weather,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func printResponse(ctx context.Context, cs *genai.ChatSession, resp *genai.GenerateContentResponse) {
	fmt.Println("== Model: ==")
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
				if call, ok := part.(genai.FunctionCall); ok {
					// Handle the function call
					if call.Name == "getWeather" {
						location, ok := call.Args["location"].(string)
						if !ok {
							log.Fatalf("invalid type for location: %v", call.Args["location"])
						}

						weather, err := getWeather(location)
						if err != nil {
							log.Fatalf("failed to get weather: %v", err)
						}
						res := sendFunc(ctx, cs, location, weather)
						printResponse(ctx, cs, res)
					}
				}
			}
		}
	}
	fmt.Println("---")
}
