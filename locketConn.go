package locketgo

import (
	"encoding/base64"
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
	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("Read(b []byte): \"%s\" - %d - %x - ", bString, bytesRead, b)
	
	return bytesRead, readErr
}

func (locket LocketConn) Write(b []byte) (int, error) {
	bytesWritten, writeErr := locket.conn.Write(b)
	bString := base64.StdEncoding.EncodeToString(b[:])
	golog.Infof("Write(b []byte): \"%s\" - %d - %x - ", bString, bytesWritten, b)

	return bytesWritten, writeErr
}

func (locket LocketConn) Close() error {
	golog.Info("Close() called")
	return locket.conn.Close()
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