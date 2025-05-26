package details

import (
	"log"
	"net"
	"os"
)

// hostanme from os package of golang
func GetHostname() (name string, err error) {
	hostname, err := os.Hostname()
	return hostname, err

}

// used stack overflow for this func
func GetIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, err
}
