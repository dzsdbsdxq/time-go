package edp

import (
	"fmt"
	"testing"
)

func TestNewQQMusicRequest(t *testing.T) {
	song := "会呼吸的痛"
	req, err := NewQQMusic()
	if err != nil {
		panic(err)
	}
	s, err := req.GetQQMusicMid(song)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
