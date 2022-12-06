package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComms_FindStartOfPacketMarker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example 1",
			input:    `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
			expected: 7,
		},
		{
			name:     "example 2",
			input:    `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			expected: 5,
		},
		{
			name:     "example 3",
			input:    `nppdvjthqldpwncqszvftbrmjlhg`,
			expected: 6,
		},
		{
			name:     "example 4",
			input:    `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			expected: 10,
		},
		{
			name:     "example 5",
			input:    `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
			expected: 11,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Comms{
				data: test.input,
			}
			actual := c.FindStartOfPacketMarker()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestComms_FindStartOfMessageMarker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example 1",
			input:    `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
			expected: 19,
		},
		{
			name:     "example 2",
			input:    `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			expected: 23,
		},
		{
			name:     "example 3",
			input:    `nppdvjthqldpwncqszvftbrmjlhg`,
			expected: 23,
		},
		{
			name:     "example 4",
			input:    `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			expected: 29,
		},
		{
			name:     "example 5",
			input:    `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
			expected: 26,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Comms{
				data: test.input,
			}
			actual := c.FindStartOfMessageMarker()
			assert.Equal(t, test.expected, actual)
		})
	}
}
