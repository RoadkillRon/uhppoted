package messages

import (
	"reflect"
	"testing"
	"time"
	codec "uhppote/encoding/UTO311-L0x"
	"uhppote/types"
)

func TestMarshalGetTimeRequest(t *testing.T) {
	expected := []byte{
		0x17, 0x32, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	request := GetTimeRequest{
		SerialNumber: 423187757,
	}

	m, err := codec.Marshal(request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Invalid byte array:\nExpected:\n%s\nReturned:\n%s", dump(expected, ""), dump(m, ""))
		return
	}
}

func TestFactoryUnmarshalGetTimeRequest(t *testing.T) {
	message := []byte{
		0x17, 0x32, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	request, err := UnmarshalRequest(message)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if request == nil {
		t.Fatalf("Unexpected request: %v\n", request)
	}

	rq, ok := request.(*GetTimeRequest)
	if !ok {
		t.Fatalf("Invalid request type - expected:%T, got: %T\n", &GetTimeRequest{}, request)
	}

	if rq.MsgType != 0x32 {
		t.Errorf("Incorrect 'message type' from valid message: %02x\n", rq.MsgType)
	}
}

func TestUnmarshalGetTimeResponse(t *testing.T) {
	message := []byte{
		0x17, 0x32, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x20, 0x19, 0x12, 0x29, 0x12, 0x34, 0x56, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	reply := GetTimeResponse{}

	err := codec.Unmarshal(message, &reply)

	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	if reply.MsgType != 0x32 {
		t.Errorf("Incorrect 'message type' - expected:%02X, got:%02x\n", 0x32, reply.MsgType)
	}

	if reply.SerialNumber != 423187757 {
		t.Errorf("Incorrect 'serial number' from valid message: %v\n", reply.SerialNumber)
	}

	datetime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-29 12:34:56", time.Local)
	if reply.DateTime != types.DateTime(datetime) {
		t.Errorf("Incorrect 'date/time' - expected:%s, got:%s\n", datetime.Format("2006-01-02 15:04:05"), reply.DateTime)
	}
}

func TestFactoryUnmarshalGetTimeResponse(t *testing.T) {
	message := []byte{
		0x17, 0x32, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x20, 0x19, 0x12, 0x29, 0x12, 0x34, 0x56, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	response, err := UnmarshalResponse(message)

	if err != nil {
		t.Fatalf("Unexpected error: %v\n", err)
	}

	if response == nil {
		t.Fatalf("Unexpected response: %v\n", response)
	}

	reply, ok := response.(*GetTimeResponse)
	if !ok {
		t.Fatalf("Invalid response type - expected:%T, got: %T\n", &GetTimeResponse{}, response)
	}

	if reply.MsgType != 0x32 {
		t.Errorf("Incorrect 'message type' - expected:%02X, got:%02x\n", 0x32, reply.MsgType)
	}

	if reply.SerialNumber != 423187757 {
		t.Errorf("Incorrect 'serial number' from valid message: %v\n", reply.SerialNumber)
	}

	datetime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-29 12:34:56", time.Local)
	if reply.DateTime != types.DateTime(datetime) {
		t.Errorf("Incorrect 'date/time' - expected:%s, got:%s\n", datetime.Format("2006-01-02 15:04:05"), reply.DateTime)
	}
}

func TestUnmarshalGetTimeResponseWithInvalidMsgType(t *testing.T) {
	message := []byte{
		0x17, 0x94, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0x20, 0x19, 0x12, 0x29, 0x12, 0x34, 0x56, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	reply := GetTimeResponse{}

	err := codec.Unmarshal(message, &reply)

	if err == nil {
		t.Errorf("Expected error: '%v'", "Invalid value in message - expected 0x30, received 0x94")
		return
	}
}
