package clock

import "fmt"

const testVersion = 4
const totalMinutes = 24 * 60

// You can find more details and hints in the test file.

type Clock int

func New(hour, minute int) Clock {
	return Clock(normalize(hour*60 + minute))
}

func (c Clock) String() string {
	hour := c / 60
	minutes := c - hour*60
	return fmt.Sprintf("%02d:%02d", hour, minutes)
}

func (c Clock) Add(minutes int) Clock {
	return Clock(normalize(int(c) + minutes))
}

func normalize(minutes int) int {
	minutes = minutes % totalMinutes
	if minutes < 0 {
		minutes += totalMinutes
	}
	return minutes
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
