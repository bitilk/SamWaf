package request

type WafAllowUrlAddReq struct {
	HostCode    string `json:"host_code"`                        //网站唯一码（主要键）
	CompareType string `json:"compare_type" form:"compare_type"` //对比方式
	Url         string `json:"url"`                              //白名单url
	Remarks     string `json:"remarks"`                          //备注
}
