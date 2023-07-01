package icmp

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"icmp_message/utils"
	"net"
	"os"
)

var (
	icmpBufferSize           = 2048
	FailedToSendICMPMessage  = fmt.Errorf("failed to send ICMP message")
	FailedToParseICMPMessage = fmt.Errorf("failed to parse ICMP message")
	UnknownICMPMessageType   = fmt.Errorf("unknown ICMP message type")
	DestinationUnreachable   = fmt.Errorf("destination unreachable")
)

func Send(packetConn *icmp.PacketConn, destAddress string, seq int, data []byte) error {
	dest, err := net.ResolveIPAddr("ip4", destAddress)
	if err != nil {
		return err
	}

	icmpMessage := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  seq,
			Data: data,
		},
	}
	icmpMessageBytes, err := icmpMessage.Marshal(nil)
	if err != nil {
		return err
	}
	succeedBytesCount, err := packetConn.WriteTo(icmpMessageBytes, dest)
	if err != nil {
		return err
	}
	if succeedBytesCount != len(icmpMessageBytes) {
		return FailedToSendICMPMessage
	}
	return nil
}

func Receive(packetConn *icmp.PacketConn) error {
	icmpMessageBytes := make([]byte, icmpBufferSize)
	succeedBytesCount, _, err := packetConn.ReadFrom(icmpMessageBytes)
	if err != nil {
		return err
	}
	icmpMessage, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), icmpMessageBytes[:utils.Min(icmpBufferSize, succeedBytesCount)])
	if err != nil {
		return err
	}
	switch icmpMessage.Type {
	case ipv4.ICMPTypeEcho, ipv4.ICMPTypeEchoReply:
		if body, ok := icmpMessage.Body.(*icmp.Echo); ok {
			fmt.Printf("[Receive] %s\n", string(body.Data))
			return nil
		}
		return FailedToParseICMPMessage
	case ipv4.ICMPTypeDestinationUnreachable:
		return DestinationUnreachable
	default:
		return UnknownICMPMessageType
	}
}

func InteractiveSendAndReceive(destAddress string) {
	packetConn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer packetConn.Close()

	quitChan := make(chan struct{})
	defer close(quitChan)

	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				err := Receive(packetConn)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	var message string
	for {
		_, err = fmt.Scanln(&message)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = Send(packetConn, destAddress, 0, []byte(message))
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("[Send] %s\n", message)
	}
}
