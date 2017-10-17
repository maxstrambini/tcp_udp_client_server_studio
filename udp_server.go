/*

UDP SERVER for studio

[2017-08-21] Max

*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var version = "1.0.1"
var start = time.Now()

var progressiveId = 10000

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	//setting up the log:
	log.Printf("SETTING LOG 'udpserver.log' ...\n")
	f, err := os.OpenFile("udpserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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

	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

	}

}
