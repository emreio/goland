package inmodulelib

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTcpConnectorPost(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start tcp listener: %v", err)
	}
	defer listener.Close()

	received := make(chan string, 1)
	serverErrors := make(chan error, 1)

	go func() {
		connection, err := listener.Accept()
		if err != nil {
			serverErrors <- err
			return
		}
		defer connection.Close()

		request, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			serverErrors <- err
			return
		}

		received <- request
		_, err = fmt.Fprint(connection, "posted\n")
		serverErrors <- err
	}()

	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		t.Fatalf("failed to split listener address: %v", err)
	}

	connector := NewTcpConnectorWithTimeout(host, port, time.Second)
	response, err := connector.Post("hello tcp\n")
	if err != nil {
		t.Fatalf("failed to post tcp data: %v", err)
	}

	if response != "posted\n" {
		t.Fatalf("expected response %q, got %q", "posted\n", response)
	}

	if request := <-received; request != "hello tcp\n" {
		t.Fatalf("expected request %q, got %q", "hello tcp\n", request)
	}

	if err := <-serverErrors; err != nil {
		t.Fatalf("server error: %v", err)
	}
}
