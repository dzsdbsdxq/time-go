package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"time-go/chrome/edp"
)

func main() {
	result := []edp.Result{
		{
			Name: "孤勇者",
			Mid:  "",
			Url:  "",
		},
		{
			Name: "会呼吸的痛",
			Mid:  "",
			Url:  "",
		},
	}
	req, err := edp.NewQQMusic()
	if err != nil {
		panic(err)
		return
	}

	chromeRdp := edp.ChromeRdp{}
	rdp, cancel := chromeRdp.NewChromeRdp()
	// create a timeout
	//NewChromeRdp(ctx, 10*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for i := 0; i <= len(result)-1; i++ {
		wg.Add(1)
		go func(i int, name string) {
			s, _ := req.GetQQMusicMid(name, &wg)
			result[i].Mid = s.Req0.Data.Body.Song.List[0].Mid
		}(i, result[i].Name)
	}
	wg.Wait()

	for i := 0; i <= len(result)-1; i++ {
		go rdp.OpenPlayerTag(result[i].Mid)
	}
	select {
	case rr:

	}

}

// 监听
//func listenForNetworkEvent(ctx context.Context) {
//	chromedp.ListenTarget(ctx, func(ev interface{}) {
//		switch ev := ev.(type) {
//		case *network.EventResponseReceived:
//			resp := ev.Response
//			//fmt.Println(resp.MimeType, reflect.TypeOf(resp.MimeType), resp.MimeType == "audio/mp4")
//			if resp.MimeType == "audio/mp4" {
//				log.Printf("success:%d, %s", resp.Status, resp.URL)
//				return
//			}
//		}
//		return
//		// other needed network Event
//	})
//}

func genIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}
