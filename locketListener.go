package locketgo

import (
	"net"
	"os"
	"path"

	"github.com/kataras/golog"
)

type LocketListener struct {
	listener   net.Listener
	Identifier string
	logDir    string
}

func NewLocketListener(listener net.Listener, logDir string, identifier string) (*LocketListener, error) {
	path := path.Join(logDir, "locket.log")
	logFile, logFileErr := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if logFileErr != nil {
		return nil, logFileErr
	}

	golog.SetOutput(logFile)

	return &LocketListener{listener: listener, Identifier: identifier, logDir: logDir}, nil
}

func (locket LocketListener) Accept() (net.Conn, error) {
	conn, connError := locket.listener.Accept()
	if connError != nil {
		golog.Errorf("%s: Accept(): Error: %s", locket.Identifier, connError)
		return conn, connError
	}

	golog.Infof("%s: Accept() called", locket.Identifier)

	return NewLocketConn(conn, locket.logDir, locket.Identifier)
}

func (locket LocketListener) Close() error {
	closeErr := locket.listener.Close()
	if closeErr != nil {
		golog.Errorf("%s: Close(): Error: %s", locket.Identifier, closeErr)
	}

	golog.Infof("%s: Close() called", locket.Identifier)

	return closeErr
}

func (locket LocketListener) Addr() net.Addr {
	return locket.listener.Addr()
}