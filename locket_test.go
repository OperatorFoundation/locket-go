package locketgo

import (
	"fmt"
	"net"
	"os"
	"testing"
)

const data = "test"

func TestMain(m *testing.M) {
	listener, listenErr := net.Listen("tcp", "127.0.0.1:1234")
	if listenErr != nil {
		return
	}
	go acceptConnections(listener)

	os.Exit(m.Run())
}

func acceptConnections(listener net.Listener) {
	serverBuffer := make([]byte, 4)
	for {
		serverConn, err := listener.Accept()
		if err != nil {
			return
		}

		locketConn, locketErr := NewLocketConn(serverConn, "/Users/bluesaxorcist/Desktop")
		if locketErr != nil {
			fmt.Println(locketErr)
			return
		}

		go func() {
			//read on server side
			_, serverReadErr := locketConn.Read(serverBuffer)
			if serverReadErr != nil {
				return
			}

			//write data from serverConn for client to read
			_, serverWriteErr := locketConn.Write([]byte(data))
			if serverWriteErr != nil {
				return
			}
		}()
	}
}

func TestLocket(t *testing.T) {
	//create client buffer
	clientBuffer := make([]byte, 4)

	//call dial on client and check error
	conn, dialErr := net.Dial("tcp", "127.0.0.1:1234")
	if dialErr != nil {
		fmt.Println("clientConn Dial error")
		t.Fail()
		return
	}

	locketConn, locketErr := NewLocketConn(conn, "/Users/bluesaxorcist/Desktop")
	if locketErr != nil {
		fmt.Println(locketErr)
		t.Fail()
		return
	}

	//write data from clientConn for server to read
	_, clientWriteErr := locketConn.Write([]byte(data))
	if clientWriteErr != nil {
		fmt.Println("client write error")
		t.Fail()
		return
	}

	//read on client side
	_, clientReadErr := locketConn.Read(clientBuffer)
	if clientReadErr != nil {
		fmt.Println("client read error")
		t.Fail()
		return
	}
}

func TestTrim(t *testing.T) {
	passed := trimString("passed123134jhb4bkj24j2h3jhreiwibr", 6)

	if passed != "passed" {
		t.Fail()
	}
}