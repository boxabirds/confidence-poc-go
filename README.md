# Experiment: Log Probability Analysis in Text Prediction

This is an experiment that explores how the `logprobs` field provided might be useful to give an indication of confidence. `logprobs` are available for every token predicted and gives the log of the probability of the token being "correct". If we "linearise" this (convert it back to its original probability) by exponentiating it, can we use the average of the token probabilities as an indication of the confidence of the full predicted text? As you can see with the provided samples, it depends on the model! 🤣

## Sample Input / Output

`bash go run main.go --file prompts.txt`

Notice how GPT-4 seems to give decent logprobs but GPT 3.5 is… less so. 

➜  confidence-poc-go git:(main) ✗ go run main.go --file prompts.txt --model gpt-4-0125-preview

### Using model: gpt-4-0125-preview
Prompt 'What is the capital of France?' 
* Predicted response: The capital of France is Paris.
* Average probability: 0.999833

Prompt 'What is the capital of London?'
* Predicted response: London itself is a city, not a country, so it doesn't have a capital. London is the capital city of England and the United Kingdom.
* Average probability: 0.921332

Prompt 'What is the capital of Cracklebackenstan?': 
* Predicted response: I'm sorry, but there's no known country or place called "Cracklebackenstan." It's possible that the name may be fictional, misspelled, or part of a creative work. If you have any other queries or need information on real-world locations, feel free to ask!
* Average probability: 0.757107


### Using model: gpt-3.5-turbo-0125
Prompt 'What is the capital of France?': 
* Predicted response: The capital of France is Paris.
* Average probability: 0.994263

Prompt 'What is the capital of London?': 
* Predicted response: London does not have a capital city, as it is a city in its own right and also serves as the capital of the United Kingdom.
* Average probability: 0.762048

Prompt 'What is the capital of Cracklebackenstan?': 
* Predicted response: The capital of Cracklebackenstan is Crackleton.
* Average probability: 0.932297
🤦‍♂️




## Usage

-  `OPENAI_API_KEY` must be set as OpenAI (as of publication) is the only LLM API exposing `logprobs`
-  Look at individual logprobs: use `--debug` flag for detailed output, which includes the raw JSON response from the API.
-  `--model` can be used to override the model which defaults to GPT 3.5. 




