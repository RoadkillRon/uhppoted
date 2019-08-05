package simulator

import (
	"net"
	"time"
	"uhppote/messages"
	"uhppote/types"
)

func (s *Simulator) find(addr *net.UDPAddr, request *messages.FindDevicesRequest) {
	utc := time.Now().UTC()
	datetime := utc.Add(time.Duration(s.TimeOffset))

	response := messages.FindDevicesResponse{
		SerialNumber: s.SerialNumber,
		IpAddress:    s.IpAddress,
		SubnetMask:   s.SubnetMask,
		Gateway:      s.Gateway,
		MacAddress:   s.MacAddress,
		Version:      types.Version(s.Version),
		Date:         types.Date(datetime),
	}

	s.send(addr, &response)
}
