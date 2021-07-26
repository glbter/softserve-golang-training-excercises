package task4

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOccurancesInFile(t *testing.T) {
	var val io.Reader = bytes.NewReader([]byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."))
	tts := []struct {
		r    *io.Reader
		s    string
		want int
	}{
		{&val, "Lorem", 4},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, CountOccurancesInFile(tt.r, tt.s))
	}
}

func TestChangeOccurancesInFile(t *testing.T) {
	txt := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	var val io.Reader = bytes.NewReader([]byte(txt))
	file, err := ioutil.TempFile("", "test.txt")
	ioutil.WriteFile(file.Name(), []byte(txt), 0644)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	tts := []struct {
		r    *io.Reader
		f    string
		cs   *ChangedString
		want string
	}{
		{&val, file.Name(), &ChangedString{"Lorem", "LOREM"}, "LOREM Ipsum is simply dummy text of the printing and typesetting industry. LOREM Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing LOREM Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of LOREM Ipsum."},
	}
	for _, tt := range tts {
		ChangeOccurancesInFile(tt.r, tt.f, tt.cs)
		got, _ := ioutil.ReadAll(file)
		assert.Equal(t, tt.want, string(got))
	}
}

func TestGetParams(t *testing.T) {
	tts := []struct {
		in   string
		want []string
		err  error
	}{
		{"   test.txt	abracadabra ", []string{"test.txt", "abracadabra"}, nil},
		{"   test.txt	abracadabra    bebra", []string{"test.txt", "abracadabra", "bebra"}, nil},
		{"   test.txt	", nil, ErrWrongInputForm},
	}
	for _, tt := range tts {
		got, err := GetParams(tt.in)
		assert.EqualValues(t, tt.want, got)
		assert.EqualValues(t, tt.err, err)
	}
}
