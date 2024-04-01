package common

import "time"

type Record struct {
	Id        uint64
	Time      time.Time
	Label     string
	Issuer    string
	Secret    string
	Digits    uint8
	Period    uint8
	Algorithm string
}

func (record *Record) Validate() {
	if record.Algorithm == "" {
		record.Algorithm = "sha1"
	}
	if record.Digits == 0 {
		record.Digits = 6
	}
	if record.Period == 0 {
		record.Period = 30
	}
}

type Records []Record

type RecordsByLabel map[string]Record
type RecordsById map[uint64]Record

func (records Records) ToRecordsByLabel() RecordsByLabel {
	var recordsByLabel = make(RecordsByLabel)
	for _, record := range records {
		recordsByLabel[record.Label] = record
	}
	return recordsByLabel
}

func (records Records) ToRecordsById() RecordsById {
	var recordsById = make(RecordsById)
	for _, record := range records {
		recordsById[record.Id] = record
	}
	return recordsById
}
