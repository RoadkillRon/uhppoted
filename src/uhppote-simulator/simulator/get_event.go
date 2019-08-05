package simulator

import (
	"net"
	"uhppote/messages"
)

func (s *Simulator) getEvent(addr *net.UDPAddr, request *messages.GetEventRequest) {
	if s.SerialNumber == request.SerialNumber {
		index := request.Index
		if index > s.Events.LastIndex() {
			index = s.Events.LastIndex()
		}

		if event := s.Events.Get(index); event != nil {
			response := messages.GetEventResponse{
				SerialNumber: s.SerialNumber,
				Index:        index,
				Type:         event.Type,
				Granted:      event.Granted,
				Door:         event.Door,
				DoorOpened:   event.DoorOpened,
				UserId:       event.UserId,
				Timestamp:    event.Timestamp,
				RecordType:   event.RecordType,
			}

			s.send(addr, &response)
		}
	}
}
