package main

import (
	"strconv"
	"strings"
)

type P uint32

func (p P) pid() int {
	return int(p&0xFF000000) >> 24
}

func (p P) timestamp() int {
	return int(p & 0xFFFFFF)
}

func New(pid, timestamp int) P {
	return P(pid<<24 + timestamp)
}

type Stack []P

func (s Stack) len() int {
	return len(s)
}

func (s Stack) peek() P {
	return s[s.len()-1]
}

func (s *Stack) push(x P) {
	*s = append(*s, x)
}

func (s *Stack) pop() (r P) {
	r = s.peek()
	*s = (*s)[:s.len()-1]
	return
}

func exclusiveTime(n int, logs []string) []int {

	s := make(Stack, 0)
	r := make([]int, n)

	pre := 0
	for _, log := range logs {
		segments := strings.Split(log, ":")
		if len(segments) != 3 {
			panic("invalid log format.")
		}

		pid, _ := strconv.Atoi(segments[0])
		event := segments[1]
		timestamp, _ := strconv.Atoi(segments[2])

		if event == "start" {
			if s.len() > 0 {
				r[s.peek().pid()] += timestamp - s.peek().timestamp()
			}
			s.push(New(pid, timestamp))
		} else {
			timestamp += 1
			p := s.pop()
			r[p.pid()] += timestamp - pre
		}
		pre = timestamp
	}
	return r
}
