package edp

import (
	"errors"
	"github.com/imroc/req/v3"
	"io"
	"strconv"
	"time"
)

type QQMusicClient struct {
	client *req.Client
}

func NewQQMusic() (*QQMusicClient, error) {
	cli := &QQMusicClient{
		client: req.C(),
	}
	cli.client.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		return nil
	}).OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		return nil
	})
	return cli, nil
}
func (cli *QQMusicClient) GetQQMusicMid(songName string) (*QQMusicResponse, error) {
	bodyJsonString := `{"comm":  { "format": "json","inCharset": "utf-8","outCharset": "utf-8"},"req_0": {"method": "DoSearchForQQMusicDesktop","module": "music.search.SearchCgiService","param": {"remoteplace": "txt.mqq.all","query": "` + songName + `","page_num": 1,"num_per_page": 1}}}`
	response := new(QQMusicResponse)
	resp, err := cli.client.R().SetHeader("Content-Type", "application/json").SetBodyJsonString(bodyJsonString).SetSuccessResult(response).Post("https://u.y.qq.com/cgi-bin/musicu.fcg?_webcgikey=DoSearchForQQMusicDesktop&_=" + strconv.FormatInt(time.Now().UnixMilli(), 10))
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
	return response, err
}

func (cli *QQMusicClient) GetQQMusicDetail(result *Result, url, postData string) (*QQMusicDetail, error) {
	response := new(QQMusicDetail)
	resp, err := cli.client.R().
		SetHeader("Content-Type", "application/json").
		SetBodyJsonString(postData).SetSuccessResult(response).
		Post(url)
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
	//fmt.Println(response)
	var lyric string
	for _, info := range response.Req0.Data.Info {
		if info.Type == "lyric" {
			lyric = info.Content[0].Value
			break
		}
	}
	//result.LyricChan <- lyric
	result.Lyric = lyric
	return response, nil
}
