package logs


// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		switch c {
		case '❗':
			return "recommendation"

		case '🔍':
			return "search"

		case '☀':
			return "weather"
		}
	}

	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newlog []rune

	for _, c := range log {
		if c == oldRune {
			newlog = append(newlog, newRune)
		} else {
			newlog = append(newlog, c)
		}
	}

	return string(newlog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	counter :=0

	for range log {
		counter++
	}

	return counter <= limit
}
