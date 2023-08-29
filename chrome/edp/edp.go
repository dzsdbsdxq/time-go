package edp

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"strings"
	"time"
)

const QqCookies = "pac_uid=0_ef341d4a01879; iip=0; pgv_pvid=5878963178; RK=0tuhRsENRj; ptcz=60a7f94bb84396401068c8e78e689f61260930054701d917e41739ac510de5d1; _clck=ofj9qx|1|fdz|0; fqm_pvqid=8fd1c33e-3019-4969-a1dd-823517e3ffbd; ts_uid=9193355500; music_ignore_pskey=202306271436Hn@vBj; ts_uid=9193355500; tmeLoginType=2; euin=oKoi7K-P7e4l7v**; ts_refer=ADTAGmyqq; fqm_sessionid=dfc068e1-7d57-45e7-bf18-606bcf5606b9; pgv_info=ssid=s8614525810; ts_last=y.qq.com/; ts_refer=music.qq.com/; ptui_loginuin=1335244575; login_type=1; psrf_qqunionid=0214AE5747295CDBD218E2BB9BF7CC8E; psrf_qqopenid=0A5FC04A1C80543131CCBB8CBDF13A1C; psrf_qqaccess_token=61C0D2D26920F0A856D7DD6EB97673F6; wxrefresh_token=; wxunionid=; psrf_musickey_createtime=1693318582; psrf_qqrefresh_token=88AD9F4A0BAF1A87D88B8C90EF9EBFB1; uin=1335244575; wxopenid=; qqmusic_key=Q_H_L_5-tSil0rlqxY46qjtEGtsdeMN-kGrVFhjU4qZN6wagGaR9anFTk5m8w; qm_keyst=Q_H_L_5-tSil0rlqxY46qjtEGtsdeMN-kGrVFhjU4qZN6wagGaR9anFTk5m8w; psrf_access_token_expiresAt=1701094582; ts_last=i.y.qq.com/n2/m/index.html"

var ChromeRdpCtx context.Context

type ChromeRdp struct {
	Ctx context.Context
}
type QQCookie struct {
	Name  string
	Value interface{}
}

func init() {
	// 禁用chrome headless
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck, //不检查默认浏览器
		chromedp.Flag("headless", true),
		//chromedp.Flag("blink-settings", "imagesEnabled=true"), //开启图像界面,重点是开启这个
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		chromedp.Flag("disable-extensions", true),        //开启插件支持
		chromedp.Flag("disable-default-apps", true),
		//chromedp.WindowSize(1920, 1080),    // 设置浏览器分辨率（窗口大小）
		chromedp.Flag("disable-gpu", true), //开启gpu渲染
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("mute-audio", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("no-default-browser-check", true),
		//chromedp.NoFirstRun, //设置网站不是首次运行
		//chromedp.UserAgent("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Mobile Safari/537.36"), //设置UserAgent
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ChromeRdpCtx = allocCtx
}
func (cr *ChromeRdp) NewChromeTab() *ChromeRdp {
	ctx, _ := chromedp.NewContext(
		ChromeRdpCtx,
	)
	ctx, _ = context.WithTimeout(ctx, 15*time.Second)
	cr.Ctx = ctx

	return cr
}
func (cr *ChromeRdp) ParseCookie() []*QQCookie {
	cookieArr := strings.Split(QqCookies, ";")
	var qc []*QQCookie
	for i := 0; i < len(cookieArr)-1; i++ {
		m := strings.Split(cookieArr[i], "=")
		qc = append(qc, &QQCookie{
			Name:  m[0],
			Value: m[1],
		})
	}
	return qc
}

func (cr *ChromeRdp) setCookies(cookies []*QQCookie, result *Result) chromedp.Tasks {

	return chromedp.Tasks{
		network.Enable(),
		chromedp.Emulate(device.IPhone12),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var cookiesParam []*network.CookieParam
			for _, v := range cookies {
				cookiesParam = append(cookiesParam, &network.CookieParam{Name: v.Name, Value: v.Value.(string)})
				err := network.SetCookie(strings.Trim(v.Name, " "), v.Value.(string)).WithPath("/").WithDomain(".qq.com").Do(ctx)
				if err != nil {
					panic(err)
				}
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			//监听
			chromedp.ListenTarget(ctx, func(ev interface{}) {
				switch ev := ev.(type) {
				case *network.EventRequestWillBeSent:
					if strings.Index(ev.Request.URL, "get_song_detail") != -1 {
						go func() {
							client, _ := NewQQMusic()
							client.GetQQMusicDetail(result, ev.Request.URL, ev.Request.PostData)
						}()

					}
					return
				case *network.EventResponseReceived:
					resp := ev.Response
					if resp.MimeType == "audio/mp4" {
						result.Url = resp.URL
						//log.Printf("success:%d, %s", resp.Status, resp.URL)
					}
					return
				}
			})
			return nil
		}),
		// 到网址
		chromedp.Navigate(`https://i.y.qq.com/v8/playsong.html?songmid=` + result.Mid),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Println("Chromedp Close")
			return chromedp.Cancel(ctx)
		}),
	}
}

func (cr *ChromeRdp) OpenPlayerTag(result *Result) {
	chromedp.Run(cr.Ctx, cr.setCookies(cr.ParseCookie(), result))
	fmt.Println("OpenPlayerTag cancel")
}
