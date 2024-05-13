package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main [prompt]",
	Short: "Process prompts using OpenAI",
	Args:  cobra.MaximumNArgs(1),
	Run:   run,
}

var fileFlag string
var debugFlag bool
var modelFlag string

func init() {
	rootCmd.PersistentFlags().StringVarP(&fileFlag, "file", "f", "", "Specify a file to read prompts from")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable detailed debug output")
	rootCmd.PersistentFlags().StringVarP(&modelFlag, "model", "m", "gpt-3.5-turbo-0125", "Specify the model used by OpenAI")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	fmt.Printf("Using model: %s\n", modelFlag)

	if fileFlag != "" {
		processFile(ctx, client, fileFlag)
	} else if len(args) > 0 {
		processPrompt(ctx, client, args[0])
	} else {
		log.Fatal("No prompt provided. Use the command with a prompt or --file flag.")
	}
}

func processFile(ctx context.Context, client *openai.Client, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processPrompt(ctx, client, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from file: %v", err)
	}
}

func processPrompt(ctx context.Context, client *openai.Client, prompt string) {
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: modelFlag,
		Messages: []openai.ChatCompletionMessage{
			{Role: "user", Content: prompt},
		},
		LogProbs: true,
	})
	if err != nil {
		log.Fatalf("Error generating chat completion: %v", err)
		return
	}
	prettyJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	if len(resp.Choices) > 0 && resp.Choices[0].Message.Content != "" {
		fmt.Printf("Prompt '%s': Predicted response: %s\n", prompt, resp.Choices[0].Message.Content)
		// Assuming log probabilities are still needed and available in the response
		// This part of the code will need to be adjusted based on actual response structure
		// Check if log probabilities are included and calculate the average log probability
		if len(resp.Choices) > 0 && resp.Choices[0].LogProbs != nil {
			var probs []float64 // Slice to store probabilities

			// Collecting log probabilities and converting them to probabilities
			for _, logProb := range resp.Choices[0].LogProbs.Content {
				prob := math.Exp(logProb.LogProb) // Convert log probability to probability
				probs = append(probs, prob)
			}

			// Calculate the average probability
			var totalProb float64
			for _, prob := range probs {
				totalProb += prob
			}
			avgProb := totalProb / float64(len(probs))

			fmt.Printf("Average probability: %f\n", avgProb)
		} else {
			fmt.Println("No log probabilities available.")
		}

	} else {
		fmt.Println("No response received.")
	}

	if debugFlag {
		fmt.Println(string(prettyJSON))
	}
}
