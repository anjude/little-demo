package schema

type GetHotListReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Type     string `form:"type"`
}

type GetHotListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Data    []struct {
		Period            string      `json:"period"`
		HotRankScore      string      `json:"hotRankScore"`
		PcHotRankScore    string      `json:"pcHotRankScore"`
		LoginUserIsFollow bool        `json:"loginUserIsFollow"`
		NickName          string      `json:"nickName"`
		AvatarUrl         string      `json:"avatarUrl"`
		UserName          string      `json:"userName"`
		ArticleTitle      string      `json:"articleTitle"`
		ArticleDetailUrl  string      `json:"articleDetailUrl"`
		CommentCount      string      `json:"commentCount"`
		FavorCount        string      `json:"favorCount"`
		ViewCount         string      `json:"viewCount"`
		HotComment        interface{} `json:"hotComment"`
		PicList           []string    `json:"picList"`
		IsNew             interface{} `json:"isNew"`
		ProductId         string      `json:"productId"`
		ProductType       string      `json:"productType"`
		RecommendType     string      `json:"recommendType"`
		ReportData        interface{} `json:"report_data"`
	} `json:"data"`
}

type CommentReq struct {
	CommentId string `json:"commentId"`
	Content   string `json:"content"`
	ArticleId string `json:"articleId"`
}
type CommentResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Data    int    `json:"data"`
}
