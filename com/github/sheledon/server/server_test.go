package server

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestRpcServer_Listener(t *testing.T) {
	t.Run("server", func(t *testing.T) {
		server := NewRpcServer("127.0.0.1:8080")
		server.Listener()
	})
	t.Run("client", func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			t.Fatal("fail connect")
		}
		defer conn.Close()
		conn.Write([]byte("Hello Rpc Server"))
		time.Sleep(10*time.Second)
		conn.Write([]byte("exceuse me"))
		time.Sleep(100*time.Second)
	})
	t.Run("test", func(t *testing.T) {
		s1 := make([]byte,0)
		s1 = append(s1, 1)
		fmt.Println(s1)
	})
}
func TestTmp(t *testing.T)  {
	go func() {
		time.Sleep(3*time.Second)
		panic("fdsa")
	}()
	time.Sleep(10*time.Second)
}