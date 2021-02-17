package models

type Config struct {
	Database struct {
		Init bool `json:"init"`
	} `json:"database"`
	Server struct {
		Enable bool   `json:"enable"`
		IP     string `json:"ip"`
		Port   int    `json:"port"`
	} `json:"server"`
	HTTPClient struct {
		Timeout int `json:"timeout"`
	} `json:"http-client"`
	Crawler bool `json:"crawler"`
	Rss     bool `json:"rss"`
	JSON    bool `json:"json"`
	Plugin  bool `json:"plugin"`
}
