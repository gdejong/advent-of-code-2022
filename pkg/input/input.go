package input

import "strconv"

func ArrayMapAndSumOverNonEmptyLines(lines []string, fn func(line string) int) int {
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		result += fn(line)
	}

	return result
}

func Chunk(lines []string, chunkSize int) [][]string {
	// Remove last element if it is empty (newline)
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var chunks [][]string
	for i := 0; i < len(lines); i += chunkSize {
		end := i + chunkSize

		chunks = append(chunks, lines[i:end])
	}

	return chunks
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}

func MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StringSliceToIntSlice(input []string) []int {
	var output []int
	for _, s := range input {
		output = append(output, MustInt(s))
	}

	return output
}
