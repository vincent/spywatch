package services

import (
	"context"
	"os"

	"github.com/mrk21/go-diff-fmt/difffmt"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func DiffDescription(oldText string, newText string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldText, newText, false)

	lineDiffs := difffmt.MakeLineDiffsFromDMP(diffs)
	hunks := difffmt.MakeHunks(lineDiffs, 3)
	unifiedFmt := difffmt.NewUnifiedFormat(difffmt.UnifiedFormatOption{
		ColorMode: difffmt.ColorNone,
	})

	before := difffmt.NewDiffTarget("before")
	after := difffmt.NewDiffTarget("after")
	unifiedFmt.Sprint(before, after, hunks)

	// return dmp.DiffPrettyText(diffs)
	return unifiedFmt.Sprint(before, after, hunks)
}

var prompt = "Here is a diff of a content that might have changed on a web page. Summarize the updated content with the minimal text. Ignore numbers which look like page numbers, likes count or similar.\n\n"
var AI_QUESTION_MAX_LENGTH = 16384

func DiffNarrative(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("no OPENAI_API_KEY defined")
	}

	client := openai.NewClient(apiKey)

	text = prompt + text
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
					Content: prompt + text,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
