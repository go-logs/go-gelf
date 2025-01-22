package gelf

import (
	"crypto/tls"
	"net"
	"os"
)

type TLSWriter struct {
	TCPWriter
	TlsConfig *tls.Config
}

func (w *TLSWriter) Dial(addr string) (net.Conn, error) {
	return tls.Dial("tcp", addr, w.TlsConfig)
}

func NewTLSWriter(addr string, tlsConfig *tls.Config) (*TLSWriter, error) {
	w := new(TLSWriter)
	w.MaxReconnect = DefaultMaxReconnect
	w.ReconnectDelay = DefaultReconnectDelay
	w.proto = "tls"
	w.addr = addr
	w.TlsConfig = tlsConfig

	var err error
	if w.conn, err = w.Dial(w.addr); err != nil {
		return nil, err
	}
	if w.hostname, err = os.Hostname(); err != nil {
		return nil, err
	}

	return w, nil
}
