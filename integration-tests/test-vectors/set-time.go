package test_vectors

var SetTime = Test{
	TestName: "set-time",

	Request: Request{
		Values: map[string]any{
			"controller": 405419896,
			"datetime":   time.Date(2021, time.Month.May, 28, 14, 56, 14, 0, time.Location.Local),
		},
		Message: []byte{
			0x17, 0x30, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x21, 0x05, 0x28, 0x14, 0x56, 0x14, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	},

	Responses: []Response{
		{
			Values: map[string]any{
				"controller": 405419896,
				"datetime":   time.Date(2021, time.Month.May, 28, 14, 56, 14, 0, time.Location.Local),
			},
			Message: []byte{
				0x17, 0x30, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x21, 0x05, 0x28, 0x14, 0x56, 0x14, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	},
}
