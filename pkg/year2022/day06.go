package year2022

type Day06 struct{}

func (p Day06) PartA(lines []string) any {
	return solve(lines[0], 4)
}

func (p Day06) PartB(lines []string) any {
	return solve(lines[0], 14)
}

func solve(datastream string, markerPosition int) int {
	start := 0
	end := markerPosition

	for i := 0; i < len(datastream); i += 1 {
		chunk := datastream[start:end]
		if unique(chunk) {
			return end
		}

		start++
		end++
	}

	panic("no result found")
}

func unique(s string) bool {
	m := make(map[rune]bool)
	for _, i := range s {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
