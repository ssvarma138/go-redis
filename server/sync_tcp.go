package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"go-redis/config"
	"go-redis/core"
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
            
			respond(cmd, c)

		}
	}


}


func readCommand(c net.Conn) (*core.RedisCmd, error) {
    var buf []byte
	buf = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}
	tokens, err := core.DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}
	return &core.RedisCmd{
		Cmd: strings.ToUpper(tokens[0]),
		Args: tokens[1:],
		}, nil
}

func respond(cmd *core.RedisCmd, c net.Conn) {
	err := core.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(c, err)
	}
}

func respondError(c net.Conn, err error) {
	c.Write([]byte(fmt.Sprintf("-%s\r\n", err)))
}