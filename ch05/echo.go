// Listing 5-1. Dette er en enkel UDP echo server implementasjon fra Chapter 5 - Unreliable UDP Communication
/*
	Denne koden skal la deg "spinne opp" en UDP server som vil "echoe" alle UDP packets den mottar til senderen.
*/
package ch05

import (
	"context"
	"fmt"
	"net"
)

func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) { // 1.Funksjonen får kontekst som tillater kansellering av echo serveren fra caller+string address
	s, err := net.ListenPacket("udp", addr) // 2. Lager en UDP connection for serveren med et call til net.ListenPacket, som returnerer net.PacketConn IF og error IF
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
		}()

		buf := make([]byte, 1024)
	for {
			n, clientAddr, err := s.ReadFrom(buf) // klient til server
			if err != nil {
				return
			}
			_, err = s.WriteTo(buf[:n], clientAddr) // server til klient
			if err != nil {
				return
			}
		}

	}()
	return s.LocalAddr(), nil
}