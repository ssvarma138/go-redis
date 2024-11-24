package core

import (
	"errors"
	"fmt"
	"net"
)

func EvalAndRespond(cmd *RedisCmd, c net.Conn) (error){
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}

func evalPING(args []string, c net.Conn) (error) {
   if len(args) >= 2 {
	return errors.New("ERR wrong number of arguments for 'ping' command")
   }
   var b []byte
   if len(args) == 0 {
	b = Encode("PONG", true)
   } else {
	b = Encode(args[0], false)
   }
   _, err := c.Write(b)
   return err
}

func Encode(value interface{}, isSimple bool) ([]byte){
	switch v := value.(type) {
	case string:
		if isSimple {
			return []byte(fmt.Sprintf("+%s\r\n", v))
		} else {
			return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v))
		}
	}
	return []byte{}
}