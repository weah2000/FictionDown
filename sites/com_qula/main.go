package com_qula

import (
	"net/http"
	"net/url"

	"github.com/ma6254/FictionDown/site"
)

func Site() site.SiteA {
	return site.SiteA{
		Name:     "笔趣阁",
		HomePage: "https://www.qu-la.com/",
		Tags: func() []string {
			return []string{
				"盗版",
				"一般书源",
				"PTCMS",
				"笔趣阁",
			}
		},
		Match: []string{
			`https://www\.qu-la\.com/booktxt/\d+/*`,
			`https://www\.qu-la\.com/booktxt/\d+/\d+\.html/*`,
		},
		BookInfo: site.Type1BookInfo(
			`//meta[@property="og:novel:book_name"]/@content`,
			`//meta[@property="og:image"]/@content`,
			`//meta[@property="og:novel:author"]/@content`,
			`//*[@id="list"]/div[3]/ul[2]/li/a`,
			// `//*[@id="list"]/dl/dd/a`,
		),
		// *[@id="chapter"]/div[1]/div[2]
		Chapter: site.Type1Chapter(`//*[@id="txt"]/text()`),
		Search: site.Type1Search("",
			func(s string) *http.Request {
				baseurl, err := url.Parse("https://so.biqusoso.com/s1.php")
				if err != nil {
					panic(err)
				}
				value := baseurl.Query()
				value.Add("ie", "utf-8")
				value.Add("siteid", "qu-la.com")
				value.Add("q", s)
				baseurl.RawQuery = value.Encode()

				req, err := http.NewRequest("GET", baseurl.String(), nil)
				if err != nil {
					panic(err)
				}
				return req
			},
			`//div[@class="search-list"]/ul/li[position()>1]`,
			`*[@class="s2"]/a`,
			`*[@class="s4"]`,
		),
	}
}
