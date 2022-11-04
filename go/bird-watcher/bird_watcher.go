package birdwatcher


// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0

	for _, birds := range birdsPerDay {
		sum += birds
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	chunks := chunk(birdsPerDay, 7)

	return TotalBirdCount(chunks[week - 1])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}


// Utils
func chunk[T comparable](slice []T, interval int) [][]T {
	var chunks [][]T

	index := -1

	for i, num := range slice {
		if i % interval == 0 {
			chunks = append(chunks, []T{})
			index += 1
		}

		chunks[index] = append(chunks[index], num)
	}

	return chunks
}
