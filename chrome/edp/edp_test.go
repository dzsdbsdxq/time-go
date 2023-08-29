package edp

import (
	"testing"
	"time"
)

func TestChromeRdp_OpenPlayerTag(t *testing.T) {

	//edpx, cancel := ed.NewChromeRdp()
	//defer cancel()
	resultChan := make(chan Result, 2)
	//,003UkWuI0E8U0l
	for i := 0; i < 4; i++ {
		ed := ChromeRdp{}
		edpx, _ := ed.NewChromeRdp()
		//defer cancel()
		edpx.OpenPlayerTag(resultChan, "004K6Ne61a1VA8")
	}

	time.Sleep(5 * time.Second)
}
