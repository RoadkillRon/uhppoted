package uhppote

import (
	"net"
	"reflect"
	"testing"
	codec "uhppote/encoding/UTO311-L0x"
)

func TestMarshalSetListenerRequest(t *testing.T) {
	expected := []byte{
		0x17, 0x90, 0x00, 0x00, 0x2D, 0x55, 0x39, 0x19, 0xc0, 0xa8, 0x01, 0x64, 0x40, 0x9c, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	request := SetListenerRequest{
		SerialNumber: 423187757,
		Address:      net.IPv4(192, 168, 1, 100),
		Port:         40000,
	}

	m, err := codec.Marshal(request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Invalid byte array:\nExpected:\n%s\nReturned:\n%s", print(expected), print(m))
		return
	}
}

func TestUnmarshalSetListenerResponse(t *testing.T) {
	message := []byte{
		0x17, 0x90, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	reply := SetListenerResponse{}

	err := codec.Unmarshal(message, &reply)

	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	if reply.MsgType != 0x90 {
		t.Errorf("Incorrect 'message type' - expected:%02X, got:%02x\n", 0x90, reply.MsgType)
	}

	if reply.SerialNumber != 423187757 {
		t.Errorf("Incorrect 'serial number' - expected:%v, got:%v\n", 423187757, reply.SerialNumber)
	}

	if reply.Success {
		t.Errorf("Incorrect 'result' - expected:%v, got:%v\n", true, reply.Success)
	}
}

func TestUnmarshalSetListenerResponseWithInvalidMsgType(t *testing.T) {
	message := []byte{
		0x17, 0x94, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0xc0, 0xa8, 0x00, 0x00, 0xff, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x19, 0x39, 0x55, 0x2d, 0x2d, 0x55, 0x39, 0x19, 0x08, 0x92,
		0x20, 0x18, 0x08, 0x16, 0x20, 0x18, 0x12, 0x31, 0x12, 0x23, 0x34, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	reply := SetListenerResponse{}

	err := codec.Unmarshal(message, &reply)

	if err == nil {
		t.Errorf("Expected error: '%v'", "Invalid value in message - expected 0x90, received 0x94")
		return
	}
}
