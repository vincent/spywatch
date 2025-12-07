package services

import (
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var IGNORED_DIFF_LINE = regexp.MustCompile(`(email-protection)`)
var ADD_DIFF = regexp.MustCompile(`^[>\+](.*)$`)
var REM_DIFF = regexp.MustCompile(`^[<\-](.*)$`)
var EMPTY_DIFF = regexp.MustCompile(`^[\-\+]\s*$`)
var NO_NEWLINE = `\ No newline at end of file`

func DiffDescription(oldText string, newText string) string {
	// diff := textdiff.Unified("before", "after", oldText, newText)
	diff := diffCommand(oldText, newText)
	diffLines := strings.Split(diff, "\n")

	// ignore header
	// diffLines = diffLines[2:]

	lines := make([]string, 0)
	for _, line := range diffLines {
		if len(line) > 1 && line != NO_NEWLINE && !EMPTY_DIFF.MatchString(line) && !ignoredDiffLine(line) {
			formatted := ADD_DIFF.ReplaceAllString(line, "ADDED: $1")
			formatted = REM_DIFF.ReplaceAllString(formatted, "REMOVED: $1")
			lines = append(lines, formatted)
		}
	}
	return strings.Join(lines, "\n")
}

func diffCommand(oldText string, newText string) string {
	tempDir, err := os.MkdirTemp("", uuid.New().String())
	if err != nil {
		return ""
	}
	tmpOld := path.Join(tempDir, "old")
	tmpNew := path.Join(tempDir, "new")

	os.WriteFile(tmpOld, []byte(oldText), 0644)
	os.WriteFile(tmpNew, []byte(newText), 0644)
	cmd := exec.Command("diff", "--changed-group-format", "-%<+%>", "--unchanged-group-format", "", "--minimal", "--ignore-blank-lines", "--ignore-space-change", "--ignore-case", "--suppress-common-lines", tmpOld, tmpNew)
	output, _ := cmd.Output()
	os.RemoveAll(tempDir)
	return string(output)
}

func ignoredDiffLine(line string) bool {
	return IGNORED_DIFF_LINE.MatchString(line)
}
