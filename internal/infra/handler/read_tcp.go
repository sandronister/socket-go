package handler

import (
	"net"
)

func (h *TcpHandler) tryError(err error) (int, error) {
	opErr, ok := err.(*net.OpError)

	if ok && opErr.Timeout() {
		h.Retries++
		if h.Retries >= h.MaxRetries {
			return 0, err
		}
	}

	return 0, nil
}

func (h *TcpHandler) ReadTCP(conn *net.TCPConn) (int, error) {
	n, err := conn.Read(h.ReadBuffer)

	if err != nil {
		res, errorOp := h.tryError(err)

		if errorOp != nil {
			return res, errorOp
		}
	}

	return n, nil

}
