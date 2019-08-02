package simulator

import (
	"uhppote/messages"
)

func (s *Simulator) DeleteCards(request *messages.DeleteCardsRequest) (*messages.DeleteCardsResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	deleted := false
	saved := false

	if request.MagicNumber == 0x55aaaa55 {
		if deleted = s.Cards.DeleteAll(); deleted {
			if err := s.Save(); err == nil {
				saved = true
			}
		}
	}

	response := messages.DeleteCardsResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    deleted && saved,
	}

	return &response, nil
}
