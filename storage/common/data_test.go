package common

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

func TestRecord_Validate(t *testing.T) {
	type fields struct {
		Id        uint64
		Time      time.Time
		Label     string
		Issuer    string
		Secret    string
		Digits    uint8
		Period    uint8
		Algorithm string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Empty fields",
			fields: fields{
				Id:        0,
				Time:      time.Time{},
				Label:     "label 1",
				Issuer:    "issuer 1",
				Secret:    "secret 1",
				Digits:    0,
				Period:    0,
				Algorithm: "",
			},
		},
		{
			name: "Empty fields",
			fields: fields{
				Id:        1,
				Time:      time.Time{},
				Label:     "label 2",
				Issuer:    "issuer 2",
				Secret:    "secret 2",
				Digits:    10,
				Period:    20,
				Algorithm: "md5",
			},
		},
	}

	var expectedRecords = Records{
		Record{
			Id:        0,
			Time:      time.Time{},
			Label:     "label 1",
			Issuer:    "issuer 1",
			Secret:    "secret 1",
			Digits:    6,
			Period:    30,
			Algorithm: "sha1",
		},
		Record{
			Id:        1,
			Time:      time.Time{},
			Label:     "label 2",
			Issuer:    "issuer 2",
			Secret:    "secret 2",
			Digits:    10,
			Period:    20,
			Algorithm: "md5",
		},
	}

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := &Record{
				Id:        tt.fields.Id,
				Time:      tt.fields.Time,
				Label:     tt.fields.Label,
				Issuer:    tt.fields.Issuer,
				Secret:    tt.fields.Secret,
				Digits:    tt.fields.Digits,
				Period:    tt.fields.Period,
				Algorithm: tt.fields.Algorithm,
			}
			record.Validate()
			assert.EqualValues(t, expectedRecords[index], *record)
		})
	}
}
