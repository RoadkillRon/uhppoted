package simulator

import (
	"uhppote"
	"uhppote-simulator/simulator/entities"
)

func (s *Simulator) SetDoorDelay(request *uhppote.SetDoorDelayRequest) (*uhppote.SetDoorDelayResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	if request.Unit != 0x03 {
		return nil, nil
	}

	door := request.Door
	if door < 1 || door > 4 {
		return nil, nil
	}

	s.Doors[door].Delay = entities.Delay(uint64(request.Delay) * 1000000000)

	err := s.Save()
	if err != nil {
		return nil, err
	}

	response := uhppote.SetDoorDelayResponse{
		SerialNumber: s.SerialNumber,
		Door:         door,
		Unit:         0x03,
		Delay:        s.Doors[door].Delay.Seconds(),
	}

	return &response, nil
}