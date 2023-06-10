package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sumOfBirds := 0
	for _, birds := range birdsPerDay {
		sumOfBirds += birds
	}
	return sumOfBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sumOfBirds := 0
    startingIndex := (week - 1) * 7
    endingIndex := startingIndex + 7
	sliceLength := len(birdsPerDay)

    for i := startingIndex; i < endingIndex && i < sliceLength; i++ {
        sumOfBirds += birdsPerDay[i]
    }

    return sumOfBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	sliceLength := len(birdsPerDay)
	for i := 0; i < sliceLength; i += 2 {
        birdsPerDay[i]++
    }
    return birdsPerDay
}
