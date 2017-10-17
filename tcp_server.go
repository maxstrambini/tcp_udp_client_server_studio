/*

TCP SERVER for studio

[2017-09-05] Max

to test:
>echo -n "test out the server" | nc localhost 3333

*/

package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var version = "1.0.1"
var start = time.Now()

var progressiveId = 10000

func main() {

	//setting up the log:
	log.Printf("SETTING LOG 'tcp_server.log' ...\n")
	f, err := os.OpenFile("tcp_server.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()
	////to log to file only:
	//log.SetOutput(f)
	//to log to stdout AND file:
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	//log set up completed.

	log.Printf("Version %s\n", version)

	/*
		log.Println("Reading config ...")
		ReadConfig() //read into var conf
		log.Println("Done with config")
	*/

	serverPortString := CONN_PORT //occhio che e' string

	if len(os.Args) > 1 {
		//port specified on command line
		serverPortString = os.Args[1]
	}

	serverAddrAndPort := CONN_HOST + ":" + serverPortString

	log.Printf("Serving on %s\n", serverAddrAndPort)

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, serverAddrAndPort)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Println("Listening on " + serverAddrAndPort)
	for {
		// Listen for an incoming connection.
		log.Println("Accepting ...")
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		log.Println("Client connected ...")
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	log.Println("BEGIN handle request ...")

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	log.Printf("Bytes received: [%v]", reqLen)
	message := string(buf)
	log.Printf("Message received: [%s]", message)
	// Send a response back to person contacting us.
	conn.Write([]byte("OK"))
	// Close the connection when you're done with it.
	conn.Close()
	log.Println("END handle request")

}
