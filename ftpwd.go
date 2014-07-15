/*
Ftpwd is a small Go tool that can recover your FTP password for your saved ftp
connections.

Usage:
	Run ftpwd. It listens on port 2121
	In your FTP client, change the "server/host" to 127.0.0.1 and the "port" to 2121.
	Try to connect and see the results in the ftpwd terminal

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Styling!
var (
	red   = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	bold  = string([]byte{27, 91, 49, 109})
	rbold = string([]byte{27, 91, 50, 50, 109})
	reset = string([]byte{27, 91, 48, 109})
)

// ftpConn struct
type ftpConn struct {
	conn     net.Conn
	username string
	password string
}

/*
	Main application
	Start a new listener on port 2121
*/
func main() {
	ftpserver, err := net.Listen("tcp", "0.0.0.0:2121")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening to 0.0.0.0:2121")

	for {
		ftpconnection, err := ftpserver.Accept()
		if err != nil {
			log.Fatal(err)
		}

		c := ftpConn{conn: ftpconnection}
		go handler(&c)
	}
}

/*
	Handle the FTP Connection
	Just keep listening on the bufio reader
*/
func handler(c *ftpConn) {

	// Send welcome message
	fmt.Fprintf(c.conn, "220 Welcome to ftpwd v1.337\r\n")

	// Listen for input
	reader := bufio.NewReader(c.conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		params := strings.SplitN(strings.Trim(line, "\r\n"), " ", 2)
		cmd := params[0]
		param := ""
		if len(params) == 1 {
			param = ""
		} else {
			param = strings.TrimSpace(params[1])
		}

		// We just want to handle the user and pass commands when an FTP
		// server is connecting.
		switch cmd {
		case "AUTH":
			if param == "TLS" {
				fmt.Fprintf(c.conn, "421 For now no TLS support.\r\n")
			}
		case "USER":
			fmt.Fprintf(c.conn, "331 User name okay, need password.\r\n")
			c.username = param
		case "PASS":
			fmt.Fprintf(c.conn, "421 Thanks and bye.\r\n")
			c.password = param

			// After we received the pass, present it to the user
			log.Printf("Your username is %s%s%s and your password is %s%s%s%s%s",
				bold, c.username, rbold, red, bold, c.password, rbold, reset)

			// Close the FTP connection
			c.conn.Close()
		default:
			fmt.Fprintf(c.conn, "\r\n")
		}
	}
}
