package provider

const (
	accessTokenCacheKey = "miniAT"
	sessionURL          = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	paidUnionIdURL      = "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s"
	accessTokenURL      = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	dailyRetainURL      = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=%s"
)
