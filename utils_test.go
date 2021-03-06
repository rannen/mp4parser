package mp4parser

import "testing"
import "math"
import "time"

func TestDottedNotationToF(t *testing.T) {
	const eps = 1e-7

	tests := [...]struct {
		input []byte
		want  float64
	}{
		{[]byte{0xFF, 0x11}, 255.17},
		{[]byte{0x01, 0x00}, 1.0},
		{[]byte{0x01, 0x04}, 1.4},
		{[]byte{0x23, 0x56}, 35.86},
		{[]byte{0x23, 0x56, 0xff, 0x01}, 9046.65281},
	}

	for _, test := range tests {
		got, _ := dottedNotationToF(test.input)
		if math.Abs(got-test.want) > eps {
			t.Errorf("intput %v want %v , got %v", test.input, test.want, got)
		}
	}
}

func BenchmarkDottedNotationToF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dottedNotationToF([]byte{0x23, 0x56, 0xff, 0x01})
	}
}

func TestByteToUint(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Errorf("got pannic")
		}
	}()

	tests := [...]struct {
		input []byte
		want  uint64
	}{
		{[]byte{0x00}, 0},
		{[]byte{0x00, 0x80}, 128},
		{[]byte{0xFF}, 255},
		{[]byte{0xFF, 0x11}, 65297},
		{[]byte{0x01, 0x00}, 256},
		{[]byte{0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, math.MaxUint64},
	}

	for _, test := range tests {
		got := byteToUint(test.input)
		if got != test.want {
			t.Errorf("intput %v want %v , got %v", test.input, test.want, got)
		}

	}
}

func BenchmarkByteToUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteToUint([]byte{0x01, 0x00})
	}
}

func TestByteToUintOverFlow(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error("should panic when overflow")
		}
	}()

	byteToUint([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func TestGetFixTime(t *testing.T) {
	want := time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC)
	got, _ := getFixTime(0)
	if !got.Equal(want) {
		t.Errorf("want: %v \t got: %v", want, got)
	}
}
