package uhppote

import (
	"reflect"
	"testing"
	"uhppote/encoding"
)

func TestMarshalGetEventIndexRequest(t *testing.T) {
	expected := []byte{
		0x17, 0xb4, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	request := GetEventIndexRequest{
		MsgType:      0xb4,
		SerialNumber: 423187757,
	}

	m, err := uhppote.Marshal(request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Invalid byte array:\nExpected:\n%s\nReturned:\n%s", print(expected), print(m))
		return
	}
}

func TestUnmarshalGetEventIndexResponse(t *testing.T) {
	message := []byte{
		0x17, 0xb4, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	reply := GetEventIndexResponse{}

	err := uhppote.Unmarshal(message, &reply)

	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	if reply.MsgType != 0xb4 {
		t.Errorf("Incorrect 'message type' - expected:%02X, got:%02x\n", 0xb4, reply.MsgType)
	}

	if reply.SerialNumber != 423187757 {
		t.Errorf("Incorrect 'serial number' - expected:%d, got: %v\n", 423187757, reply.SerialNumber)
	}

	if reply.Index != 17 {
		t.Errorf("Incorrect 'index' - expected:%d, got:%d\n", 17, reply.Index)
	}
}