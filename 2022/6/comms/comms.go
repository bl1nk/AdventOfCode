package comms

import "os"

type Comms struct {
	data string
}

func New(fileName string) (Comms, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return Comms{}, err
	}

	return Comms{data: string(b)}, nil
}

func (c *Comms) FindStartOfPacketMarker() int {
	return findMarker(c.data, 4)
}

func (c *Comms) FindStartOfMessageMarker() int {
	return findMarker(c.data, 14)
}

func findMarker(data string, numUnique int) int {
	n := 0
	for n < len(data) {
		buf := make(map[byte]struct{})
		for i := n; i < (n + numUnique); i++ {
			buf[data[i]] = struct{}{}
		}

		if len(buf) == numUnique {
			return n + numUnique
		}

		n++
	}
	return -1
}
