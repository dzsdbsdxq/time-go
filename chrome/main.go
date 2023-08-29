package main

import (
	"fmt"
	"sync"
	"time-go/chrome/edp"
)

func main() {

	result := []*edp.Result{
		{
			Name: "孤勇者",
		},
		{
			Name: "会呼吸的痛",
		},
	}
	chromeRdp := edp.ChromeRdp{}

	req, err := edp.NewQQMusic()
	if err != nil {
		panic(err)
		return
	}

	var wg sync.WaitGroup
	for i := 0; i <= len(result)-1; i++ {
		wg.Add(1)
		go func(res *edp.Result) {
			defer wg.Done()
			s, _ := req.GetQQMusicMid(res.Name)
			res.Mid = s.Req0.Data.Body.Song.List[0].Mid
			chromeRdp.NewChromeTab().OpenPlayerTag(res)
		}(result[i])
	}
	wg.Wait()

	for _, res := range result {
		fmt.Println(*res)
	}

}
