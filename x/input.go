package x

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadString(fname string) string {
	s, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(s))
}
func ReadLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
