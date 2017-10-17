/*

UDP CLIENT for studio

[2017-08-21] Max

*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
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

	//setting up the log:
	log.Printf("SETTING LOG 'udpclient.log' ...\n")
	f, err := os.OpenFile("udpclient.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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

	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)

	}

}
