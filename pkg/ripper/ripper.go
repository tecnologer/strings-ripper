package ripper

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Do(input string, length int) string {
	if utf8.RuneCountInString(input) <= length {
		return input
	}

	var output []string

	for len(input) > 0 {
		if utf8.RuneCountInString(input) <= length {
			output = append(output, formatChunkEnd(input))
			break
		}

		output = append(output, formatChunk(input[:length]))
		input = input[length:]
	}

	return strings.Join(output, "\n")
}

func formatChunk(chunk string) string {
	return fmt.Sprintf("%q+", chunk)
}

func formatChunkEnd(chunk string) string {
	return fmt.Sprintf("%q", chunk)
}
