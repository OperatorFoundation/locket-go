package locketgo

import (
	"encoding/base64"
	"net"
	"os"
	"path"
	"syscall"

	"github.com/kataras/golog"
)

type LocketUDPConn struct {
	conn net.UDPConn
}

func NewLocketUDPConn(conn net.UDPConn, logDir string) (*LocketUDPConn, error) {
	path := path.Join(logDir, "locket.log")
	logFile, logFileErr := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if logFileErr != nil {
		return nil, logFileErr
	}

	golog.SetOutput(logFile)

	return &LocketUDPConn{conn: conn}, nil
}

func (locket LocketUDPConn) SyscallConn() (syscall.RawConn, error) {
	return locket.conn.SyscallConn()
}

func (locket LocketUDPConn) ReadFromUDP(b []byte) (n int, addr *net.UDPAddr, err error) {
	bytesRead, readAddr, readErr := locket.conn.ReadFromUDP(b)
	if readErr != nil {
		golog.Errorf("ReadFromUDP(b []byte): Error: %s", readErr)
		return bytesRead, readAddr, readErr
	}
	
	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("ReadFromUDP(b []byte): \"%s\" - %d - %x", bString, bytesRead, b)

	return bytesRead, readAddr, readErr
}

func (locket LocketUDPConn) ReadFrom(b []byte) (int, net.Addr, error) {
	bytesRead, readAddr, readErr := locket.conn.ReadFrom(b)
	if readErr != nil {
		golog.Errorf("ReadFrom(b []byte): Error: %s", readErr)
		return bytesRead, readAddr, readErr
	}
	
	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("ReadFrom(b []byte): \"%s\" - %d - %x", bString, bytesRead, b)
	
	return bytesRead, readAddr, readErr
}

func (locket LocketUDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *net.UDPAddr, err error) {
	bytesRead, readOobn, readFlags, readAddr, readErr := locket.conn.ReadMsgUDP(b, oob)
	if readErr != nil {
		golog.Errorf("ReadMsgUDP(b, oob []byte): Error: %s", readErr)
		return bytesRead, readOobn, readFlags, readAddr, readErr
	}

	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("ReadMsgUDP(b, oob []byte): \"%s\" - %d - %x", bString, bytesRead, b)

	return bytesRead, readOobn, readFlags, readAddr, readErr
}

func (locket LocketUDPConn) WriteToUDP(b []byte, addr *net.UDPAddr) (int, error) {
	bytesWritten, writeErr := locket.conn.WriteToUDP(b, addr)
	if writeErr != nil {
		golog.Errorf("WriteToUDP(b []byte, addr *net.UDPAddr): Error: %s", writeErr)
		return bytesWritten, writeErr
	}

	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("WriteToUDP(b []byte, addr *net.UDPAddr): \"%s\" - %d - %x", bString, bytesWritten, b)

	return bytesWritten, writeErr
}

func (locket LocketUDPConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	bytesWritten, writeErr := locket.conn.WriteTo(b, addr)
	if writeErr != nil {
		golog.Errorf("WriteTo(b []byte, addr net.Addr): Error: %s", writeErr)
		return bytesWritten, writeErr
	}

	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("WriteTo(b []byte, addr net.Addr): \"%s\" - %d - %x", bString, bytesWritten, b)

	return bytesWritten, writeErr
}

func (locket LocketUDPConn) WriteMsgUDP(b, oob []byte, addr *net.UDPAddr) (n, oobn int, err error) {
	bytesWritten, writeOobn, writeErr := locket.conn.WriteMsgUDP(b, oob, addr)
	if writeErr != nil {
		golog.Errorf("WriteMsgUDP(b, oob []byte, addr *net.UDPAddr): Error: %s", writeErr)
		return bytesWritten, writeOobn, writeErr
	}

	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("WriteMsgUDP(b, oob []byte, addr *net.UDPAddr): \"%s\" - %d - %x", bString, bytesWritten, b)

	return bytesWritten, writeOobn, writeErr
}