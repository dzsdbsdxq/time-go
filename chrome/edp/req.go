package edp

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"io"
	"strconv"
	"time"
)

const qqMusicBaseUrl = "https://u.y.qq.com"

type QQMusicClient struct {
	client *req.Client
}

func NewQQMusic() (*QQMusicClient, error) {
	cli := &QQMusicClient{
		client: req.C().SetBaseURL(qqMusicBaseUrl),
	}
	cli.client.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		return nil
	}).OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		return nil
	})
	return cli, nil
}
func (cli *QQMusicClient) GetQQMusicMid(songName string) (*QQMusicResponse, error) {
	fmt.Println(songName)
	bodyJsonString := `{"comm":  { "format": "json","inCharset": "utf-8","outCharset": "utf-8"},"req_0": {"method": "DoSearchForQQMusicDesktop","module": "music.search.SearchCgiService","param": {"remoteplace": "txt.mqq.all","query": "` + songName + `","page_num": 1,"num_per_page": 1}}}`
	response := new(QQMusicResponse)
	resp, err := cli.client.R().SetHeader("Content-Type", "application/json").SetBodyJsonString(bodyJsonString).SetSuccessResult(response).Post("/cgi-bin/musicu.fcg?_webcgikey=DoSearchForQQMusicDesktop&_=" + strconv.FormatInt(time.Now().UnixMilli(), 10))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}
	fmt.Println(response)
	return response, err
}
