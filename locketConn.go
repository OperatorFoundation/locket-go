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
	conn       net.Conn
	Identifier string
}

func NewLocketConn(conn net.Conn, logDir string, identifier string) (*LocketConn, error) {
	path := path.Join(logDir, "locket.log")
	logFile, logFileErr := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if logFileErr != nil {
		return nil, logFileErr
	}

	golog.SetOutput(logFile)

	return &LocketConn{conn: conn, Identifier: identifier}, nil
}

func (locket LocketConn) Read(b []byte) (int, error) {
	bytesRead, readErr := locket.conn.Read(b)
	if readErr != nil {
		golog.Errorf("%s : Read(b []byte): Error: %s", locket.Identifier, readErr)
		return bytesRead, readErr
	}

	bString := trimString(string(b[:]), bytesRead)
	bHex := trimString(hex.EncodeToString(b), bytesRead * 2)
	golog.Infof("%s : Read(b []byte): \"%s\" - %d - %s", locket.Identifier, bString, bytesRead, bHex)
	
	return bytesRead, readErr
}

func (locket LocketConn) Write(b []byte) (int, error) {
	bytesWritten, writeErr := locket.conn.Write(b)
	if writeErr != nil {
		golog.Errorf("%s : Write(b []byte): Error: %s", locket.Identifier, writeErr)
		return bytesWritten, writeErr
	}

	bString := trimString(string(b[:]), bytesWritten)
	bHex := trimString(hex.EncodeToString(b), bytesWritten * 2)
	
	golog.Infof("%s : Write(b []byte): \"%s\" - %d - %s", locket.Identifier, bString, bytesWritten, bHex)

	return bytesWritten, writeErr
}

func (locket LocketConn) Close() error {
	closeError := locket.conn.Close()
	if closeError != nil {
		golog.Errorf("%s : Close(): Error: %s", locket.Identifier, closeError)
	}

	golog.Infof("%s : Close() called successfully", locket.Identifier)

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

func trimString(str string, expectedLength int) string{
	if len(str) > (expectedLength) {
		return str[0:expectedLength]
	} else {
		return str
	}
}