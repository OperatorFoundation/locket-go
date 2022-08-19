package locketgo

import (
	"encoding/hex"
	"net"
	"os"
	"path"
	"time"

	"github.com/kataras/golog"
)

type LocketConn struct {
	conn    net.Conn
}

func NewLocketConn(conn net.Conn, logDir string) (*LocketConn, error) {
	path := path.Join(logDir, "locket.log")
	logFile, logFileErr := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if logFileErr != nil {
		return nil, logFileErr
	}

	golog.SetOutput(logFile)

	return &LocketConn{conn: conn}, nil
}

func (locket LocketConn) Read(b []byte) (int, error) {
	bytesRead, readErr := locket.conn.Read(b)
	if readErr != nil {
		golog.Errorf("Read(b []byte): Error: %s", readErr)
		return bytesRead, readErr
	}

	bString := string(b[:])
	bHex := hex.EncodeToString(b)
	golog.Infof("Read(b []byte): \"%s\" - %d - %s", bString, bytesRead, bHex)
	
	return bytesRead, readErr
}

func (locket LocketConn) Write(b []byte) (int, error) {
	bytesWritten, writeErr := locket.conn.Write(b)
	if writeErr != nil {
		golog.Errorf("Write(b []byte): Error: %s", writeErr)
		return bytesWritten, writeErr
	}

	bString := string(b[:])
	bHex := hex.EncodeToString(b)
	golog.Infof("Write(b []byte): \"%s\" - %d - %s", bString, bytesWritten, bHex)

	return bytesWritten, writeErr
}

func (locket LocketConn) Close() error {
	closeError := locket.conn.Close()
	if closeError != nil {
		golog.Errorf("Close(): Error: %s", closeError)
	}

	golog.Info("Close() called successfully")

	return closeError
}

func (locket LocketConn) LocalAddr() net.Addr {
	return locket.conn.LocalAddr()
}

func (locket LocketConn) RemoteAddr() net.Addr {
	return locket.conn.RemoteAddr()
}

func (locket LocketConn) SetDeadline(t time.Time) error {
	return locket.conn.SetDeadline(t)
}

func (locket LocketConn) SetReadDeadline(t time.Time) error {
	return locket.conn.SetReadDeadline(t)
}

func (locket LocketConn) SetWriteDeadline(t time.Time) error {
	return locket.conn.SetWriteDeadline(t)
}