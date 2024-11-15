package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"go-redis/config"
)

func RunSyncTCPServer() {
	log.Println("starting synchronous TCP server on ", config.Host, config.Port)

	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))

	var num_of_clients = 0

	if err != nil {
		panic(err)
	}

	for {
		c, err := lsnr.Accept()

		if err != nil {
			panic(err)
		}
		num_of_clients += 1

		log.Println("client connected with address:", c.RemoteAddr(), "concurrent clients", num_of_clients)

		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				num_of_clients -= 1
				if err == io.EOF {
					break
				}
				log.Println("err", err)
			}
			log.Println("command", cmd)
            
			if err = respond(cmd, c); err != nil {
				log.Println("err write:", err)
			}

		}
	}


}

func readCommand(c net.Conn) (string, error) {
    var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
    if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err 
	}
	return nil
}