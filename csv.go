package csvhelper

import (
	"encoding/csv"
	"reflect"
	"strconv"
)

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type reflect.Kind
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type.String()
}

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		switch f.Kind() {
		case reflect.String:
			f.SetString(record[i])
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ival, err := strconv.ParseInt(record[i], 10, 64)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			ival, err := strconv.ParseUint(record[i], 10, 64)
			if err != nil {
				return err
			}
			f.SetUint(ival)
		case reflect.Float32, reflect.Float64:
			ival, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return err
			}
			f.SetFloat(ival)
		default:
			return &UnsupportedType{f.Kind()}
		}
	}
	return nil
}

func UnmarshalFieldsByIndex(reader *csv.Reader, v interface{}, indices ...int) error {
	if len(indices) == 0 {
		return Unmarshal(reader, v)
	}
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	if len(indices) > len(record) {
		return nil
	}
	for _, i := range indices {
		f := s.Field(i)
		switch f.Kind() {
		case reflect.String:
			f.SetString(record[i])
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ival, err := strconv.ParseInt(record[i], 10, 64)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			ival, err := strconv.ParseUint(record[i], 10, 64)
			if err != nil {
				return err
			}
			f.SetUint(ival)
		case reflect.Float32, reflect.Float64:
			ival, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return err
			}
			f.SetFloat(ival)
		default:
			return &UnsupportedType{f.Kind()}
		}
	}
	return nil
}
