package main
// я не успел сделать это задание.

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type ClientJob struct {
	cmd string
	conn net.Conn
}

func generateResponses(cj <-chan ClientJob) {
	for {
		job := <- cj

		if strings.Compare("logout", job.cmd) == 0 {
			job.conn.Write([]byte("Bye"))
			return
		}
		fmt.Println(job.cmd)
		job.conn.Write([]byte(job.cmd))

	}
}

func main() {
	clientJobs := make(chan ClientJob)

	go generateResponses(clientJobs)

	ln, err := net.Listen("tcp", ":2323")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		go func() {
			buf := bufio.NewReader(conn)

			conn.Write([]byte("Hello"))

			for {
				conn.Write([]byte(">"))

				cmd, err := buf.ReadString('\n')

				cmd = strings.Replace(cmd, "\r\n", "", -1)
				if err != nil {
					break
				}

				clientJobs <- ClientJob{
					cmd, 
					conn,
				}
			}
		}()
	}
}
