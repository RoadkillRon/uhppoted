package uhppoted

import (
	"fmt"
	"time"
	"uhppote/types"
)

type GetEventsRequest struct {
	DeviceID DeviceID
	Start    *types.DateTime
	End      *types.DateTime
}

type GetEventsResponse struct {
	DeviceID DeviceID    `json:"device-id,omitempty"`
	Dates    *DateRange  `json:"dates,omitempty"`
	Events   *EventRange `json:"events,omitempty"`
}

type GetEventRequest struct {
	DeviceID DeviceID
	EventID  uint32
}

type GetEventResponse struct {
	DeviceID DeviceID `json:"device-id"`
	Event    event    `json:"event"`
}

type DateRange struct {
	Start *types.DateTime `json:"start,omitempty"`
	End   *types.DateTime `json:"end,omitempty"`
}

func (d *DateRange) String() string {
	if d.Start != nil && d.End != nil {
		return fmt.Sprintf("{ Start:%v, End:%v }", d.Start, d.End)
	}

	if d.Start != nil {
		return fmt.Sprintf("{ Start:%v }", d.Start)
	}

	if d.End != nil {
		return fmt.Sprintf("{ End:%v }", d.End)
	}

	return "{}"
}

type EventRange struct {
	First uint32 `json:"first"`
	Last  uint32 `json:"last"`
}

func (e *EventRange) String() string {
	return fmt.Sprintf("{ First:%v, Last:%v }", e.First, e.Last)
}

type event struct {
	Index      uint32         `json:"event-id"`
	Type       uint8          `json:"event-type"`
	Granted    bool           `json:"access-granted"`
	Door       uint8          `json:"door-id"`
	DoorOpened bool           `json:"door-opened"`
	UserID     uint32         `json:"user-id"`
	Timestamp  types.DateTime `json:"timestamp"`
	Result     uint8          `json:"event-result"`
}

func (u *UHPPOTED) GetEvents(request GetEventsRequest) (*GetEventsResponse, error) {
	u.debug("get-events", fmt.Sprintf("request  %+v", request))

	device := uint32(request.DeviceID)
	start := request.Start
	end := request.End

	event, err := u.Uhppote.GetEvent(device, 0xffffffff)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", InternalServerError, fmt.Errorf("Error getting last event index from %v (%w)", device, err))
	}

	first := uint32(0)
	last := uint32(0)
	if event != nil {
		first = 1
		last = event.Index

		if start != nil {
			first = last
		}

		if end != nil {
			last = 1
		}

		// TODO: this logic doesn't handle wrap around i.e. when the UHPPOTE controller event index
		// increments from 100000 to 0
		if start != nil || end != nil {
			for index := event.Index; index > 0; index-- {
				record, err := u.Uhppote.GetEvent(device, index)
				if err != nil {
					return nil, fmt.Errorf("%w: %v", InternalServerError, fmt.Errorf("Error getting event for index %v from %v (%w)", index, device, err))
				}

				if start != nil && !time.Time(record.Timestamp).Before(time.Time(*start)) && record.Index < first {
					first = record.Index
				}

				if end != nil && !time.Time(*end).Before(time.Time(record.Timestamp)) && record.Index > last {
					last = record.Index
				}
			}
		}
	}

	dates := (*DateRange)(nil)
	if start != nil || end != nil {
		dates = &DateRange{
			Start: start,
			End:   end,
		}
	}

	events := (*EventRange)(nil)
	if first != 0 || last != 0 {
		events = &EventRange{
			First: first,
			Last:  last,
		}
	}

	response := GetEventsResponse{
		DeviceID: DeviceID(device),
		Dates:    dates,
		Events:   events,
	}

	u.debug("get-events", fmt.Sprintf("response %+v", response))

	return &response, nil
}

func (u *UHPPOTED) GetEvent(request GetEventRequest) (*GetEventResponse, error) {
	u.debug("get-events", fmt.Sprintf("request  %+v", request))

	device := uint32(request.DeviceID)
	eventID := request.EventID

	record, err := u.Uhppote.GetEvent(device, eventID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", InternalServerError, fmt.Errorf("Error getting event for ID %v from %v (%w)", eventID, device, err))
	}

	if record == nil {
		return nil, fmt.Errorf("%w: %v", NotFound, fmt.Errorf("No event record for ID %v for %v", eventID, device))
	}

	if record.Index != eventID {
		return nil, fmt.Errorf("%w: %v", NotFound, fmt.Errorf("No event record for ID %v for %v", eventID, device))
	}

	response := GetEventResponse{
		DeviceID: DeviceID(record.SerialNumber),
		Event: event{
			Index:      record.Index,
			Type:       record.Type,
			Granted:    record.Granted,
			Door:       record.Door,
			DoorOpened: record.DoorOpened,
			UserID:     record.UserID,
			Timestamp:  record.Timestamp,
			Result:     record.Result,
		},
	}

	u.debug("get-event", fmt.Sprintf("response %+v", response))

	return &response, nil
}
