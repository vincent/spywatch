package scrapper

import (
	"testing"
)

func TestHtmlContentPreFilters(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "removes empty HTML comment",
			input:  "Hello<!-- -->World",
			output: "HelloWorld",
		},
		{
			name:   "removes empty HTML comment with spaces",
			input:  "A<!--   -->B",
			output: "AB",
		},
		{
			name:   "removes CDATA section",
			input:  "foo<![CDATA[bar]]>baz",
			output: "foobaz",
		},
		{
			name:   "removes both comment and CDATA",
			input:  "<!-- --><![CDATA[abc]]>xyz",
			output: "xyz",
		},
		{
			name:   "no match",
			input:  "plain text",
			output: "plain text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HtmlContentPreFilters(tt.input)
			if got != tt.output {
				t.Errorf("HtmlContentPreFilters(%q) = %q; want %q", tt.input, got, tt.output)
			}
		})
	}
}

func TestHtmlContentPostFilters(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "removes line with only slash",
			input:  "/\nfoo\n",
			output: "foo",
		},
		{
			name:   "removes line with only dot",
			input:  ".\nbar\n",
			output: "bar",
		},
		{
			name:   "removes line with only dollar",
			input:  "$\nbaz\n",
			output: "baz",
		},
		{
			name:   "removes line with only space",
			input:  " \nabc\n",
			output: "abc",
		},
		{
			name:   "removes line with only asterisk",
			input:  "* \ndef\n",
			output: "def",
		},
		{
			name:   "removes line with only bullet",
			input:  "•  \nghi\n",
			output: "ghi",
		},
		{
			name:   "keeps normal lines",
			input:  "foo\n•  \nbar\nbaz\n",
			output: "foo\nbar\nbaz",
		},
		{
			name:   "mixed lines",
			input:  "/\n.\n$\n \n*\n•\nkeep\n",
			output: "keep",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TextContentPostFilters(tt.input)
			if got != tt.output {
				t.Errorf("HtmlContentPostFilters(%q) = %q; want %q", tt.input, got, tt.output)
			}
		})
	}
}
