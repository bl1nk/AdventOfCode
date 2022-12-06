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
	numUnique := 4
	n := 0
	for n < len(c.data) {
		buf := make(map[byte]struct{})
		for i := n; i < (n + numUnique); i++ {
			buf[c.data[i]] = struct{}{}
		}

		if len(buf) == numUnique {
			return n + numUnique
		}

		n++
	}
	return -1
}
