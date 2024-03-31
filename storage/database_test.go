package storage

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

const filename = `file::memory:?cache=shared`

func TestOpen(t *testing.T) {
	database, err := Open(filename)
	assert.Nil(t, err)
	defer database.Close()

	assert.Equal(t, filename, database.fileName)
	assert.NotNil(t, database.database)
	assert.NotNil(t, database.insertCommand)
	assert.NotNil(t, database.updateCommand)
}

func TestDatabase_Close(t *testing.T) {
	database, err := Open(filename)
	assert.Nil(t, err)

	err = database.Close()
	assert.Nil(t, err)
}

func helperInsert(t *testing.T, database *Database, records Records) error {
	t.Helper()
	for _, record := range records {
		if err := database.Insert(record.Label, record.Issuer, record.Secret, record.Digits, record.Period, record.Algorithm); err != nil {
			return err
		}
	}
	return nil
}

func helperClearTime(t *testing.T, records *Records) {
	t.Helper()
	for i := 0; i < len(*records); i++ {
		record := &(*records)[i]
		record.Time = time.Time{}
	}
}

func TestDatabase_InsertQueryUpdate(t *testing.T) {
	var input = Records{
		{
			Id:        1,
			Time:      time.Time{},
			Label:     "name 1",
			Issuer:    "desc 1",
			Secret:    "code 1",
			Digits:    1,
			Period:    2,
			Algorithm: "",
		},
		{
			Id:        2,
			Time:      time.Time{},
			Label:     "name 2",
			Issuer:    "desc 2",
			Secret:    "code 2",
			Digits:    3,
			Period:    4,
			Algorithm: "",
		},
		{
			Id:        3,
			Time:      time.Time{},
			Label:     "name 3",
			Issuer:    "desc 3",
			Secret:    "code 3",
			Digits:    5,
			Period:    6,
			Algorithm: "",
		},
	}

	database, err := Open(filename)
	assert.Nil(t, err)
	defer database.Close()

	err = helperInsert(t, database, input)
	assert.Nil(t, err)

	records, err := database.Query()
	assert.Nil(t, err)
	assert.Len(t, records, len(input))

	// Time items are different
	assert.NotEqualValues(t, input, records)

	// Clear time items
	helperClearTime(t, &records)
	assert.ElementsMatch(t, input, records)

	input[0].Label = "Label Updated"
	input[1].Issuer = "Desc Updated"
	input[2].Secret = "Secret Updated"

	for _, value := range input {
		err := database.Update(value)
		assert.Nil(t, err)
	}

	records, err = database.Query()
	assert.Nil(t, err)
	assert.Len(t, records, len(input))

	// Time items are different
	assert.NotEqualValues(t, input, records)

	// Clear time items
	helperClearTime(t, &records)
	assert.ElementsMatch(t, input, records)
}

func TestDatabase_UpdateFailOnNonUniqueName(t *testing.T) {
	var input = Records{
		{
			Id:        1,
			Time:      time.Time{},
			Label:     "name 1",
			Issuer:    "desc 1",
			Secret:    "code 1",
			Algorithm: "",
		},
		{
			Id:        2,
			Time:      time.Time{},
			Label:     "name 2",
			Issuer:    "desc 2",
			Secret:    "code 2",
			Algorithm: "",
		},
		{
			Id:        3,
			Time:      time.Time{},
			Label:     "name 3",
			Issuer:    "desc 3",
			Secret:    "code 3",
			Algorithm: "",
		},
	}

	database, err := Open(filename)
	assert.Nil(t, err)
	defer database.Close()

	err = helperInsert(t, database, input)
	assert.Nil(t, err)

	records, err := database.Query()
	assert.Nil(t, err)
	assert.Len(t, records, len(input))

	// Time items are different
	assert.NotEqualValues(t, input, records)

	// Clear time items
	helperClearTime(t, &records)
	assert.ElementsMatch(t, input, records)

	input0BackupName := input[0].Label

	input[0].Label = input[1].Label
	input[1].Issuer = "Desc Updated"
	input[2].Secret = "Secret Updated"

	for index, value := range input {
		err := database.Update(value)
		if index == 0 {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}

	input[0].Label = input0BackupName

	records, err = database.Query()
	assert.Nil(t, err)
	assert.Len(t, records, len(input))

	// Time items are different
	assert.NotEqualValues(t, input, records)

	// Clear time items
	helperClearTime(t, &records)
	assert.ElementsMatch(t, input, records)
}

func Fuzz_sqlCommand(f *testing.F) {
	f.Add("prefix", "table", "suffix")
	f.Add("", "", "")
	f.Add("prefix", "", "")
	f.Add("", "", "suffix")
	f.Add(" ", " ", " ")
	f.Add(" another prefix", "another table", "my suffix ")

	f.Fuzz(func(t *testing.T, prefix string, table string, suffix string) {
		expected := strings.TrimSpace(prefix + " " + table + " " + suffix)
		got := sqlCommand(prefix, table, suffix)
		assert.Equal(t, expected, got)
	})
}
