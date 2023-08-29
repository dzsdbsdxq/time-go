package edp

type QQMusicDetail struct {
	Code    int    `json:"code"`
	Ts      int64  `json:"ts"`
	StartTs int64  `json:"start_ts"`
	Traceid string `json:"traceid"`
	Req0    Req0   `json:"req_0"`
}
type Singer struct {
	ID    int    `json:"id"`
	Mid   string `json:"mid"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Type  int    `json:"type"`
	Uin   int    `json:"uin"`
}
type Album struct {
	ID         int    `json:"id"`
	Mid        string `json:"mid"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	TimePublic string `json:"time_public"`
	Pmid       string `json:"pmid"`
}
type Mv struct {
	ID    int    `json:"id"`
	Vid   string `json:"vid"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Vt    int    `json:"vt"`
}
type File struct {
	MediaMid      string        `json:"media_mid"`
	Size24Aac     int           `json:"size_24aac"`
	Size48Aac     int           `json:"size_48aac"`
	Size96Aac     int           `json:"size_96aac"`
	Size192Ogg    int           `json:"size_192ogg"`
	Size192Aac    int           `json:"size_192aac"`
	Size128Mp3    int           `json:"size_128mp3"`
	Size320Mp3    int           `json:"size_320mp3"`
	SizeApe       int           `json:"size_ape"`
	SizeFlac      int           `json:"size_flac"`
	SizeDts       int           `json:"size_dts"`
	SizeTry       int           `json:"size_try"`
	TryBegin      int           `json:"try_begin"`
	TryEnd        int           `json:"try_end"`
	URL           string        `json:"url"`
	SizeHires     int           `json:"size_hires"`
	HiresSample   int           `json:"hires_sample"`
	HiresBitdepth int           `json:"hires_bitdepth"`
	B30S          int           `json:"b_30s"`
	E30S          int           `json:"e_30s"`
	Size96Ogg     int           `json:"size_96ogg"`
	Size360Ra     []interface{} `json:"size_360ra"`
	SizeDolby     int           `json:"size_dolby"`
	SizeNew       []int         `json:"size_new"`
}
type Pay struct {
	PayMonth   int `json:"pay_month"`
	PriceTrack int `json:"price_track"`
	PriceAlbum int `json:"price_album"`
	PayPlay    int `json:"pay_play"`
	PayDown    int `json:"pay_down"`
	PayStatus  int `json:"pay_status"`
	TimeFree   int `json:"time_free"`
}
type Action struct {
	Switch   int `json:"switch"`
	Msgid    int `json:"msgid"`
	Alert    int `json:"alert"`
	Icons    int `json:"icons"`
	Msgshare int `json:"msgshare"`
	Msgfav   int `json:"msgfav"`
	Msgdown  int `json:"msgdown"`
	Msgpay   int `json:"msgpay"`
	Switch2  int `json:"switch2"`
	Icon2    int `json:"icon2"`
}
type Ksong struct {
	ID  int    `json:"id"`
	Mid string `json:"mid"`
}
type Volume struct {
	Gain float64 `json:"gain"`
	Peak int     `json:"peak"`
	Lra  float64 `json:"lra"`
}
type TrackInfo struct {
	ID          int      `json:"id"`
	Type        int      `json:"type"`
	Mid         string   `json:"mid"`
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Singer      []Singer `json:"singer"`
	Album       Album    `json:"album"`
	Mv          Mv       `json:"mv"`
	Interval    int      `json:"interval"`
	Isonly      int      `json:"isonly"`
	Language    int      `json:"language"`
	Genre       int      `json:"genre"`
	IndexCd     int      `json:"index_cd"`
	IndexAlbum  int      `json:"index_album"`
	TimePublic  string   `json:"time_public"`
	Status      int      `json:"status"`
	Fnote       int      `json:"fnote"`
	File        File     `json:"file"`
	Pay         Pay      `json:"pay"`
	Action      Action   `json:"action"`
	Ksong       Ksong    `json:"ksong"`
	Volume      Volume   `json:"volume"`
	Label       string   `json:"label"`
	URL         string   `json:"url"`
	Bpm         int      `json:"bpm"`
	Version     int      `json:"version"`
	Trace       string   `json:"trace"`
	DataType    int      `json:"data_type"`
	ModifyStamp int      `json:"modify_stamp"`
	Pingpong    string   `json:"pingpong"`
	Ppurl       string   `json:"ppurl"`
	Tid         int      `json:"tid"`
	Ov          int      `json:"ov"`
	Sa          int      `json:"sa"`
	Es          string   `json:"es"`
	Vs          []string `json:"vs"`
	Vi          []int    `json:"vi"`
}
type Content struct {
	ID        int    `json:"id"`
	Value     string `json:"value"`
	Mid       string `json:"mid"`
	Type      int    `json:"type"`
	ShowType  int    `json:"show_type"`
	IsParent  int    `json:"is_parent"`
	Picurl    string `json:"picurl"`
	ReadCnt   int    `json:"read_cnt"`
	Author    string `json:"author"`
	Jumpurl   string `json:"jumpurl"`
	OriPicurl string `json:"ori_picurl"`
}
type Info struct {
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Content     []Content `json:"content"`
	Pos         int       `json:"pos"`
	More        int       `json:"more"`
	Selected    string    `json:"selected"`
	UsePlatform int       `json:"use_platform"`
}
type Extras struct {
	Name      string `json:"name"`
	Transname string `json:"transname"`
	Subtitle  string `json:"subtitle"`
	From      string `json:"from"`
	Wikiurl   string `json:"wikiurl"`
}
type Data struct {
	TrackInfo TrackInfo `json:"track_info"`
	Info      []Info    `json:"info"`
	Extras    Extras    `json:"extras"`
}
type Req0 struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}
