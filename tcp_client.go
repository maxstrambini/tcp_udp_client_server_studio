/*

TCP SERVER for studio

[2017-09-05] Max

to test:
>echo -n "test out the server" | nc localhost 3333

*/

package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	serverAddress := "127.0.0.1"
	serverPortString := "3333"
	message := ""

	if len(os.Args) > 1 {
		//address specified on command line
		serverAddress = os.Args[1]
	} else {
		fmt.Printf("Missing destination address (params: 'Address', 'Port', 'Message')\n")
		return
	}

	if len(os.Args) > 2 {
		//port specified on command line
		serverPortString = os.Args[2]
	} else {
		fmt.Printf("Missing destination port (params: 'Address', 'Port', 'Message')\n")
		return
	}
	if len(os.Args) > 3 {
		//message specified on command line
		message = os.Args[3]
	}

	serverAddrAndPort := serverAddress + ":" + serverPortString
	fmt.Printf("Connecting to: '%s'\n", serverAddrAndPort)

	// connect to this socket
	conn, err := net.Dial("tcp", serverAddrAndPort)
	if err != nil {
		fmt.Printf("Error connecting to: '%s'\nError details: %v\n", serverAddrAndPort, err)
		return
	}

	for {
		if len(message) == 0 {
			fmt.Printf("Type your message:\n")
			// read in input from stdin
			reader := bufio.NewReader(os.Stdin)
			//fmt.Printf("Type the message to send: \n")
			message, _ = reader.ReadString('\n')
		}
		// send to socket
		fmt.Fprintf(conn, message+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Reply from server: [%s]\n", message)

		break
	}
}
