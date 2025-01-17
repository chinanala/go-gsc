package main

import (
	"fmt"
	"os"

	"github.com/bujnlc8/go-gsc/controller"
	"github.com/bujnlc8/go-gsc/util"
	"github.com/gin-gonic/gin"
)

func setRoute(r *gin.Engine) {
	r.GET("/songci/index/:id/:open_id", controller.HandleIndex)
	r.GET("/songci/query/:q/:page/:open_id", controller.HandleQuery)
	r.GET("/gsc/index/:id/:open_id", controller.HandleIndex)
	r.GET("/gsc/query/:q/:page/:open_id", controller.HandleQuery)
	r.GET("/gsc/query_by_page/:q/:page/:open_id", controller.HandleQueryByPage)
	r.GET("/user/auth/:code", controller.Code2Session)
	r.GET("/user/like/:open_id/:gsc_id", controller.SetUserLike)
	r.GET("/user/dislike/:open_id/:gsc_id", controller.SetUserDisLike)
	r.GET("/songci/mylike/:open_id", controller.QueryMyLike)
	r.GET("/gsc/mylike/:open_id", controller.QueryMyLike)
	r.GET("/gsc/mylike_by_page/:open_id", controller.QueryMyLikeByPage)
	r.GET("/gsc/short_index", controller.HandleShortIndex)
	r.GET("/gsc/query_by_page_a/:page/:q/:open_id", controller.HandleQueryByPage)
	r.POST("/gsc/feedback/:open_id/:gsc_id", controller.HandleUserFeedBack)
	r.GET("/user/:open_id/captcha", controller.HandleCaptcha)
	r.GET("/user/:open_id/ad", controller.HandleAd)
	r.GET("/user/auth_alipay/:code", controller.Code2SessionAliPay)
	r.POST("/gsc/audio_url", controller.GetSignUrlForAudio)
}
func main() {
	if os.Getenv("GSC_DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	setRoute(r)
	listenAddr := fmt.Sprintf("%v", util.GetConfStr("listenAddr"))
	r.Run(listenAddr)
}
