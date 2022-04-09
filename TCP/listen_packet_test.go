package main

import (
    "bytes"
    "context"
    "net"
    "testing"
)

func TestListenPacketUDP(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
    if err != nil {
        t.Fatal(err)
    }
    defer cancel()

    client, err := net.ListenPacket("udp", "127.0.0.1:")
    if err != nil {
        t.Fatal(err)
    }
    defer func() { _ = client.Close() }()
}