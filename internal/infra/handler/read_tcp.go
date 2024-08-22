package handler

import (
	"net"
	"time"
)

func (h *TcpHandler) ReadTCP(conn TCPAddrInterface) ([]byte, error) {
	buffer := make([]byte, 1024)
	totalBytesRead := 0

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		bytesRead, err := conn.Read(buffer[totalBytesRead:])
		if err != nil {
			opErr, ok := err.(*net.OpError)
			if ok && opErr.Timeout() {
				h.Retries++
				if h.Retries >= h.MaxRetries {
					return nil, err
				}
				continue
			}
			return nil, err
		}
		totalBytesRead += bytesRead
		if bytesRead == 0 {
			break
		}
	}

	return buffer[:totalBytesRead], nil

}
