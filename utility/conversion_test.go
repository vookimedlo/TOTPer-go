package utility

import (
	"golang.org/x/exp/constraints"
	"math"
	"testing"
)

func TestCastSigned32ToUnsigned8(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[int32, uint8]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   -1,
			want:  0,
			want1: false,
		},
		{
			name:  "",
			arg:   math.MaxUint8,
			want:  math.MaxUint8,
			want1: true,
		},
		{
			name:  "",
			arg:   math.MaxUint8 + 1,
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[int32, uint8](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastSigned32ToUnsigned16(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[int32, uint16]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   -1,
			want:  0,
			want1: false,
		},
		{
			name:  "",
			arg:   math.MaxUint16,
			want:  math.MaxUint16,
			want1: true,
		},
		{
			name:  "",
			arg:   math.MaxUint16 + 1,
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[int32, uint16](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastSigned32ToUnsigned32(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[int32, uint32]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   -1,
			want:  0,
			want1: false,
		},
		{
			name:  "",
			arg:   math.MaxInt32,
			want:  math.MaxInt32,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[int32, uint32](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastSigned32ToUnsigned64(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[int, uint64]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   -1,
			want:  0,
			want1: false,
		},
		{
			name:  "",
			arg:   math.MaxInt64,
			want:  math.MaxInt64,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[int, uint64](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastUnsigned8ToSigned32(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[uint8, int32]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   math.MaxUint8,
			want:  math.MaxUint8,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[uint8, int32](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastUnsigned16ToSigned32(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[uint16, int32]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   math.MaxUint16,
			want:  math.MaxUint16,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[uint16, int32](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastUnsigned32ToSigned32(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[uint32, int32]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   math.MaxUint32,
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[uint32, int32](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCastUnsigned64ToSigned32(t *testing.T) {
	type testCase[V constraints.Integer, R constraints.Integer] struct {
		name  string
		arg   V
		want  R
		want1 bool
	}
	tests := []testCase[uint64, int32]{
		{
			name:  "",
			arg:   0,
			want:  0,
			want1: true,
		},
		{
			name:  "",
			arg:   1,
			want:  1,
			want1: true,
		},
		{
			name:  "",
			arg:   math.MaxUint64,
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CastInteger[uint64, int32](tt.arg)
			if got != tt.want {
				t.Errorf("CastInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CastInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
