package main

import (
	"fmt"
	"sync"
	"time-go/chrome/edp"
)

func main() {

	//resultChan := make(chan edp.Result)
	//defer close(resultChan)
	result := []*edp.Result{
		{
			Name:  "孤勇者",
			Mid:   "",
			Url:   "",
			Lyric: "",
			//LyricChan: make(chan string),
			//UrlChan:   make(chan string),
		},
		{
			Name:  "会呼吸的痛",
			Mid:   "",
			Url:   "",
			Lyric: "",
			//LyricChan: make(chan string),
			//UrlChan:   make(chan string),
		},
	}
	req, err := edp.NewQQMusic()
	if err != nil {
		panic(err)
		return
	}

	//chromeRdp := edp.ChromeRdp{}
	//rdp, cancel := chromeRdp.NewChromeRdp()
	//defer cancel()

	var wg sync.WaitGroup
	for i := 0; i <= len(result)-1; i++ {
		wg.Add(1)
		go func(i int, name string) {
			defer wg.Done()
			s, _ := req.GetQQMusicMid(name)
			result[i].Mid = s.Req0.Data.Body.Song.List[0].Mid

			chromeRdp := edp.ChromeRdp{}
			rdp, _ := chromeRdp.NewChromeRdp()
			rdp.OpenPlayerTag(result[i])

		}(i, result[i].Name)
	}
	wg.Wait()

	for _, res := range result {
		fmt.Println(*res)
	}

}
