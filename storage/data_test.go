package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRecords_ToRecordsById(t *testing.T) {
	var input = Records{
		{
			Id:        10,
			Time:      time.Time{},
			Label:     "name 10",
			Issuer:    "issuer 10",
			Secret:    "secret 10",
			Digits:    8,
			Period:    30,
			Algorithm: "",
		},
		{
			Id:        5,
			Time:      time.Time{},
			Label:     "name 5",
			Issuer:    "issuer 5",
			Secret:    "secret 5",
			Digits:    12,
			Period:    60,
			Algorithm: "sha1",
		},
	}

	var expectedOutput = RecordsById{
		input[0].Id: input[0],
		input[1].Id: input[1],
	}

	var output = input.ToRecordsById()
	assert.Equal(t, expectedOutput, output)
}

func TestRecords_ToRecordsByName(t *testing.T) {
	var input = Records{
		{
			Id:        10,
			Time:      time.Time{},
			Label:     "name 10",
			Issuer:    "issuer 10",
			Secret:    "secret 10",
			Digits:    6,
			Period:    30,
			Algorithm: "sha1",
		},
		{
			Id:        5,
			Time:      time.Time{},
			Label:     "name 5",
			Issuer:    "issuer 5",
			Secret:    "secret 5",
			Digits:    12,
			Period:    60,
			Algorithm: "",
		},
	}

	var expectedOutput = RecordsByLabel{
		input[0].Label: input[0],
		input[1].Label: input[1],
	}

	var output = input.ToRecordsByLabel()
	assert.Equal(t, expectedOutput, output)
}
