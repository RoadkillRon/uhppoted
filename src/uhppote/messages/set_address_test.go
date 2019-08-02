package messages

import (
	"net"
	"reflect"
	"testing"
	codec "uhppote/encoding/UTO311-L0x"
)

func TestMarshalSetAddressRequest(t *testing.T) {
	expected := []byte{
		0x17, 0x96, 0x00, 0x00, 0x2d, 0x55, 0x39, 0x19, 0xc0, 0xa8, 0x01, 0x7d, 0xff, 0xff, 0xff, 0x00,
		0xc0, 0xa8, 0x01, 0x00, 0x55, 0xaa, 0xaa, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	request := SetAddressRequest{
		SerialNumber: 423187757,
		Address:      net.IPv4(192, 168, 1, 125),
		Mask:         net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(192, 168, 1, 0),
		MagicNumber:  0x55aaaa55,
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