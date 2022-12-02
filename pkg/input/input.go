package input

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
