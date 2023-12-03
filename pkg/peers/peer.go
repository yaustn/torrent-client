package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

const peerSize = 6 // 4 bytes for IP, 2 for port

// Unmarshal parses peer IP addresses and ports from a buffer
func Unmarshal(peersBytes []byte) ([]Peer, error) {
	if len(peersBytes)%peerSize != 0 {
		err := fmt.Errorf("Received malformed peers")
		return nil, err
	}

	numPeers := len(peersBytes) / peerSize
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peersBytes[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16(peersBytes[offset+4 : offset+6])
	}

	return peers, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
