package API

import (
	"net"
	"strconv"
	"time"
)

var activeThreads = 0
var openPorts []int

type Scanner struct {
	Ip      string
	MinPort int
	MaxPort int
}

func (s *Scanner) Scan() []int {

	var port int
	var doneChannel chan bool

	openPorts = []int{}
	doneChannel = make(chan bool)

	for port = s.MinPort; port <= s.MaxPort; port++ {
		go testTCPConnection(s.Ip, port, doneChannel)
		activeThreads++
	}
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
	return openPorts
}
func testTCPConnection(ip string, port int, doneChannel chan bool) {
	var err error

	_, err = net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		openPorts = append(openPorts, port)
	}
	doneChannel <- true
}
