package edp

import (
	"fmt"
	"testing"
)

func TestChromeRdp_OpenPlayerTag(t *testing.T) {
	result := &Result{
		Name: "會呼吸的痛",
		Mid:  "004K6Ne61a1VA8",
	}
	chromeRdp := ChromeRdp{}
	chromeRdp.NewChromeTab().OpenPlayerTag(result)
	fmt.Println(result)
}
