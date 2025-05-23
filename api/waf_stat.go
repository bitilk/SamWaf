package api

import (
	"SamWaf/model/common/response"
	"SamWaf/model/request"
	"github.com/gin-gonic/gin"
)

type WafStatApi struct {
}

// 获取系统基本信息
func (w *WafStatApi) StatSysinfoApi(c *gin.Context) {

	response.OkWithDetailed(wafStatService.StatHomeSysinfo(c), "获取成功", c)
}

// 获取运行系统基本信息
func (w *WafStatApi) StatRumtimeSysinfoApi(c *gin.Context) {

	response.OkWithDetailed(wafStatService.StatHomeRumtimeSysinfo(), "获取成功", c)
}

func (w *WafStatApi) StatHomeSumDayApi(c *gin.Context) {

	wafStat, _ := wafStatService.StatHomeSumDayApi()
	response.OkWithDetailed(wafStat, "获取成功", c)
}
func (w *WafStatApi) StatHomeSumDayRangeApi(c *gin.Context) {
	var req request.WafStatsDayRangeReq
	err := c.ShouldBind(&req)
	if err == nil {
		wafStat, _ := wafStatService.StatHomeSumDayRangeApi(req)
		response.OkWithDetailed(wafStat, "获取成功", c)
	} else {

		response.FailWithMessage("解析失败", c)
	}
}
func (w *WafStatApi) StatHomeSumDayTopIPRangeApi(c *gin.Context) {
	var req request.WafStatsDayRangeReq
	err := c.ShouldBind(&req)
	if err == nil {
		wafStat, _ := wafStatService.StatHomeSumDayTopIPRangeApi(req)
		response.OkWithDetailed(wafStat, "获取成功", c)
	} else {

		response.FailWithMessage("解析失败", c)
	}
}
