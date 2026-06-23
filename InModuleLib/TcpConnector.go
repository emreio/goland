package inmodulelib

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const defaultTcpConnectorTimeout = 5 * time.Second

// TcpConnector is a small helper for posting data to a TCP endpoint.
type TcpConnector struct {
	host    string
	port    string
	timeout time.Duration
}

// NewTcpConnector creates a TCP connector with the default timeout.
func NewTcpConnector(host string, port string) *TcpConnector {
	return NewTcpConnectorWithTimeout(host, port, defaultTcpConnectorTimeout)
}

// NewTcpConnectorWithTimeout creates a TCP connector with a caller-provided timeout.
func NewTcpConnectorWithTimeout(host string, port string, timeout time.Duration) *TcpConnector {
	return &TcpConnector{host: host, port: port, timeout: timeout}
}

// Post sends data to the configured TCP endpoint and returns the first response line.
func (tcpConnector *TcpConnector) Post(data string) (string, error) {
	if tcpConnector == nil {
		return "", fmt.Errorf("tcp connector is nil")
	}

	address := net.JoinHostPort(tcpConnector.host, tcpConnector.port)
	connection, err := net.DialTimeout("tcp", address, tcpConnector.timeout)
	if err != nil {
		return "", err
	}
	defer connection.Close()

	if tcpConnector.timeout > 0 {
		deadline := time.Now().Add(tcpConnector.timeout)
		if err := connection.SetDeadline(deadline); err != nil {
			return "", err
		}
	}

	if _, err := fmt.Fprint(connection, data); err != nil {
		return "", err
	}

	response, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		return "", err
	}

	return response, nil
}

// PostTcp sends data to a TCP endpoint using the default connector settings.
func PostTcp(host string, port string, data string) (string, error) {
	return NewTcpConnector(host, port).Post(data)
}
