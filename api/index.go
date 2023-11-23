package api

import (
	"jokyo3/probingai/api/helper"
	"jokyo3/probingai/common"
	"net/http"
	"strings"
//	"os"
//	"net/url" 
)

//var bingURL = os.Getenv("BING_PROXY_DM")

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, common.PROXY_WEB_PAGE_PATH, http.StatusFound)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/turing") {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
	}
//	if bingURL == "" {
//		bingURL = common.BING_URL.String()
//	}
//	proxyurl, _ := url.Parse(bingURL)
	common.NewSingleHostReverseProxy(common.bingURL).ServeHTTP(w, r)
}

