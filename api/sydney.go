package api

import (
	"jokyo3/probingai/api/helper"
	"jokyo3/probingai/common"
	"net/http"
//	"os"
//	"net/url" 
)

//var sydneyURL = os.Getenv("SYDNEY_PROXY_DM")

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
//	if sydneyURL == "" {
//		sydneyURL = common.BING_SYDNEY_URL.String()
//	}
//	sydproxyurl, _ := url.Parse(sydneyURL)
	common.NewSingleHostReverseProxy(common.SydneyURL).ServeHTTP(w, r)
}
