package edp

import (
	"testing"
)

func TestChromeRdp_OpenPlayerTag(t *testing.T) {
	ed := ChromeRdp{}
	edpx, cancel := ed.NewChromeRdp()
	defer cancel()
	edpx.OpenPlayerTag("003UkWuI0E8U0l")
	// 遍历监听到的事件，查找XHR请求并获取其响应数据
}
