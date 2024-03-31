package utility

import (
	"golang.org/x/exp/constraints"
	"math"
	"reflect"
	"strconv"
)

type isSignedInteger = bool

type integerType struct {
	isSigned isSignedInteger
	maxValue uint64
	minValue int64
}

var integerTypes = map[reflect.Kind]integerType{
	reflect.Uint: {
		isSigned: false,
		maxValue: math.MaxUint,
		minValue: 0,
	},
	reflect.Uint8: {
		isSigned: false,
		maxValue: math.MaxUint8,
		minValue: 0,
	},
	reflect.Uint16: {
		isSigned: false,
		maxValue: math.MaxUint16,
		minValue: 0,
	},
	reflect.Uint32: {
		isSigned: false,
		maxValue: math.MaxUint32,
		minValue: 0,
	},
	reflect.Uint64: {
		isSigned: false,
		maxValue: math.MaxUint64,
		minValue: 0,
	},
	reflect.Int: {
		isSigned: true,
		maxValue: math.MaxInt,
		minValue: math.MinInt,
	},
	reflect.Int8: {
		isSigned: true,
		maxValue: math.MaxInt8,
		minValue: math.MinInt8,
	},
	reflect.Int16: {
		isSigned: true,
		maxValue: math.MaxInt16,
		minValue: math.MinInt16,
	},
	reflect.Int32: {
		isSigned: true,
		maxValue: math.MaxInt32,
		minValue: math.MinInt32,
	},
	reflect.Int64: {
		isSigned: true,
		maxValue: math.MaxInt64,
		minValue: math.MinInt64,
	},
}

func checkRange[S constraints.Integer, T constraints.Integer](value S) bool {

	sourceType := integerTypes[reflect.TypeOf(value).Kind()]
	targetType := integerTypes[reflect.TypeOf(T(0)).Kind()]

	if sourceType.isSigned && value < 0 {
		return targetType.minValue < int64(value)
	}

	return uint64(value) <= targetType.maxValue
}

func CastInteger[V constraints.Integer, R constraints.Integer](value V) (R, bool) {
	if checkRange[V, R](value) {
		return R(value), true
	}

	return 0, false
}

func ConvertToIntegerOrEmptyToZero[R constraints.Integer](value string) (R, bool) {
	if value != "" {
		if intValue, err := strconv.Atoi(value); err != nil {
			if u, ok := CastInteger[int, R](intValue); ok {
				return u, true
			}
		}
		return 0, false
	}
	return 0, true
}
