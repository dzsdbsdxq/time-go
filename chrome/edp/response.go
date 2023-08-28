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

type QQMusicRequest struct {
	Comm struct {
		GTk         int    `json:"g_tk"`
		Uin         int    `json:"uin"`
		Format      string `json:"format"`
		InCharset   string `json:"inCharset"`
		OutCharset  string `json:"outCharset"`
		Notice      int    `json:"notice"`
		Platform    string `json:"platform"`
		NeedNewCode int    `json:"needNewCode"`
		Ct          int    `json:"ct"`
		Cv          int    `json:"cv"`
	} `json:"comm"`
	Req0 struct {
		Method string `json:"method"`
		Module string `json:"module"`
		Param  struct {
			Remoteplace string `json:"remoteplace"`
			Searchid    string `json:"searchid"`
			SearchType  int    `json:"search_type"`
			Query       string `json:"query"`
			PageNum     int    `json:"page_num"`
			NumPerPage  int    `json:"num_per_page"`
		} `json:"param"`
	} `json:"req_0"`
}

func NewQQMusicRequest(name string) *QQMusicRequest {
	return &QQMusicRequest{
		Comm: struct {
			GTk         int    `json:"g_tk"`
			Uin         int    `json:"uin"`
			Format      string `json:"format"`
			InCharset   string `json:"inCharset"`
			OutCharset  string `json:"outCharset"`
			Notice      int    `json:"notice"`
			Platform    string `json:"platform"`
			NeedNewCode int    `json:"needNewCode"`
			Ct          int    `json:"ct"`
			Cv          int    `json:"cv"`
		}{
			GTk:         632906770,
			Uin:         1335244575,
			Format:      "json",
			InCharset:   "utf-8",
			OutCharset:  "utf-8",
			Notice:      0,
			Platform:    "h5",
			NeedNewCode: 1,
			Ct:          23,
			Cv:          0,
		},
		Req0: struct {
			Method string `json:"method"`
			Module string `json:"module"`
			Param  struct {
				Remoteplace string `json:"remoteplace"`
				Searchid    string `json:"searchid"`
				SearchType  int    `json:"search_type"`
				Query       string `json:"query"`
				PageNum     int    `json:"page_num"`
				NumPerPage  int    `json:"num_per_page"`
			} `json:"param"`
		}{
			Method: "DoSearchForQQMusicDesktop",
			Module: "usic.search.SearchCgiService",
			Param: struct {
				Remoteplace string `json:"remoteplace"`
				Searchid    string `json:"searchid"`
				SearchType  int    `json:"search_type"`
				Query       string `json:"query"`
				PageNum     int    `json:"page_num"`
				NumPerPage  int    `json:"num_per_page"`
			}(struct {
				Remoteplace string
				Searchid    string
				SearchType  int
				Query       string
				PageNum     int
				NumPerPage  int
			}{
				Remoteplace: "txt.mqq.all",
				Searchid:    "69887720779077945",
				SearchType:  0,
				Query:       name,
				PageNum:     1,
				NumPerPage:  1,
			}),
		},
	}
}
