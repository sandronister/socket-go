package handler

import (
	"net"
	"time"
)

func (h *TcpHandler) ReadTCP(conn TCPAddrInterface) (int, error) {
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
					return 0, err
				}
				continue
			}
			return 0, err
		}
		totalBytesRead += bytesRead
		if bytesRead == 0 {
			break
		}
	}

	h.ReadBuffer = buffer[:totalBytesRead]
	return totalBytesRead, nil

}
