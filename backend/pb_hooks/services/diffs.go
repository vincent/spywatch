package services

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/akedrou/textdiff"
	openai "github.com/sashabaranov/go-openai"
)

var ADD_DIFF, _ = regexp.Compile(`^\+(.*)$`)
var REM_DIFF, _ = regexp.Compile(`^\-(.*)$`)
var EMPTY_DIFF, _ = regexp.Compile(`^[\-\+]\s*$`)
var NO_NEWLINE = `\ No newline at end of file`

func DiffDescription(oldText string, newText string) string {
	diff := textdiff.Unified("before", "after", oldText, newText)

	// ignore header
	diffLines := strings.Split(diff, "\n")
	diffLines = diffLines[2:]

	lines := make([]string, 0)
	for _, line := range diffLines {
		if len(line) > 1 && line != NO_NEWLINE && !EMPTY_DIFF.MatchString(line) {
			formatted := ADD_DIFF.ReplaceAllString(line, "ADDED: $1")
			formatted = REM_DIFF.ReplaceAllString(formatted, "REMOVED: $1")
			lines = append(lines, formatted)
		}
	}
	return strings.Join(lines, "\n")
}

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
