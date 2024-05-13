# Experiment: Log Probability Analysis in Text Prediction

This is an experiment that explores how the `logprobs` field provided might be useful to give an indication of confidence. `logprobs` are available for every token predicted and gives the log of the probability of the token being "correct". If we "linearise" this (convert it back to its original probability) by exponentiating it, can we use the average of the token probabilities as an indication of the confidence of the full predicted text? As you can see with the provided samples, probably not! 🤣

## Sample Input / Output

`bash go run main.go --file prompts.txt`


- Prompt 'What is the capital of France?': Predicted response: The capital of France is Paris.
  - Average probability: 0.993785
- Prompt 'What is the capital of London?': Predicted response: There is no capital of London as London itself is the capital city of England.
  - Average probability: 0.779275
- Prompt 'What is the capital of Cracklebackenstan?': Predicted response: The capital of Cracklebackenstan is Crackleton.
  - Average probability: 0.918923

Ok then! 🤦‍♂️

## Usage

-  `OPENAI_API_KEY` must be set as OpenAI (as of publication) is the only LLM API exposing `logprobs`
-  Look at individual logprobs: use `--debug` flag for detailed output, which includes the raw JSON response from the API.




