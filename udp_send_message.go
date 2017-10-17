/*

UDP send message test

[2017-08-22] Max

*/

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var version = "1.0.1"
var start = time.Now()

var progressiveId = 10000

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	destIP := os.Args[1]
	destPort := os.Args[2]
	message := os.Args[3]

	//setting up the log:
	//log.Printf("SETTING LOG 'udp_send_message.log' ...\n")
	//f, err := os.OpenFile("udp_send_message.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	//if err != nil {
	//	log.Printf("error opening file: %v", err)
	//}
	//defer f.Close()
	////to log to file only:
	////log.SetOutput(f)
	//to log to stdout AND file:
	//mw := io.MultiWriter(os.Stdout, f)
	//log.SetOutput(mw)
	//log set up completed.

	log.Printf("Version %s\n", version)

	address := fmt.Sprintf("%s:%v", destIP, destPort)
	log.Printf("Sending to '%s' the message: '%s'", address, message)

	ServerAddr, err := net.ResolveUDPAddr("udp", address)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()

	buf := []byte(message)
	_, err = Conn.Write(buf)
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	log.Println("Message sent by UDP")
	return




}
