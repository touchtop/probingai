package api

import (
	"adams549659584/go-proxy-bingai/api/helper"
	"adams549659584/go-proxy-bingai/common"
	"net/http"
	"os"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	sydneyURL := os.Getenv("SYDNEY_PROXY_DM")
	if sydneyURL == "" {
		sydneyURL = common.BING_SYDNEY_URL
	}
	common.NewSingleHostReverseProxy(sydneyURL).ServeHTTP(w, r)
}
}
