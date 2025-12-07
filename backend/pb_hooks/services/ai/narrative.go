package ai

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

var AI_DIFF_NARRATIVE_PROMPT = `
Here is a diff of a content that might have changed on a web page.
Summarize the updated content with the minimal text.
Ignore content which look like page numbers, likes count, referrer codes or similar.
`
var AI_QUESTION_MAX_LENGTH = 16384

func DiffNarrative(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("no OPENAI_API_KEY defined")
	}

	client := openai.NewClient(apiKey)

	text = AI_DIFF_NARRATIVE_PROMPT + text
	if len(text) > AI_QUESTION_MAX_LENGTH {
		text = text[:AI_QUESTION_MAX_LENGTH]
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: AI_DIFF_NARRATIVE_PROMPT + text,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
