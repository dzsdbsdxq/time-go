package edp

type QQMusicResponse struct {
	Code int `json:"code"`
	Req0 struct {
		Code int `json:"code"`
		Data struct {
			Body struct {
				Song struct {
					List []struct {
						Mid      string `json:"mid"`
						Status   int    `json:"status"`
						Subtitle string `json:"subtitle"`
						Title    string `json:"title"`
					} `json:"list"`
				} `json:"song"`
			} `json:"body"`
			Code int `json:"code"`
		} `json:"data"`
	} `json:"req_0"`
}
