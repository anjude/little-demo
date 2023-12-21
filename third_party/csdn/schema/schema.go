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

type CommentListReq struct {
	CommentId string `json:"commentId"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

type CommentListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Data    struct {
		Count      int `json:"count"`
		PageCount  int `json:"pageCount"`
		FloorCount int `json:"floorCount"`
		FoldCount  int `json:"foldCount"`
		List       []struct {
			Info struct {
				CommentId             int           `json:"commentId"`
				ArticleId             int           `json:"articleId"`
				ParentId              int           `json:"parentId"`
				PostTime              string        `json:"postTime"`
				Content               string        `json:"content"`
				UserName              string        `json:"userName"`
				Digg                  int           `json:"digg"`
				DiggArr               []interface{} `json:"diggArr"`
				LoginUserDigg         bool          `json:"loginUserDigg"`
				ParentUserName        string        `json:"parentUserName"`
				ParentNickName        string        `json:"parentNickName"`
				Avatar                string        `json:"avatar"`
				NickName              string        `json:"nickName"`
				DateFormat            string        `json:"dateFormat"`
				Years                 int           `json:"years"`
				Vip                   bool          `json:"vip"`
				VipIcon               string        `json:"vipIcon"`
				VipUrl                string        `json:"vipUrl"`
				CompanyBlog           bool          `json:"companyBlog"`
				CompanyBlogIcon       string        `json:"companyBlogIcon"`
				Flag                  bool          `json:"flag"`
				FlagIcon              string        `json:"flagIcon"`
				LevelIcon             string        `json:"levelIcon"`
				CommentFromTypeResult struct {
					Index int    `json:"index"`
					Key   string `json:"key"`
					Title string `json:"title"`
				} `json:"commentFromTypeResult"`
				IsTop           bool        `json:"isTop"`
				IsBlack         bool        `json:"isBlack"`
				Region          string      `json:"region"`
				OrderNo         string      `json:"orderNo"`
				RedEnvelopeInfo interface{} `json:"redEnvelopeInfo"`
				GptInfo         interface{} `json:"gptInfo"`
			} `json:"info"`
			Sub []struct {
				CommentId             int           `json:"commentId"`
				ArticleId             int           `json:"articleId"`
				ParentId              int           `json:"parentId"`
				PostTime              string        `json:"postTime"`
				Content               string        `json:"content"`
				UserName              string        `json:"userName"`
				Digg                  int           `json:"digg"`
				DiggArr               []interface{} `json:"diggArr"`
				LoginUserDigg         bool          `json:"loginUserDigg"`
				ParentUserName        string        `json:"parentUserName"`
				ParentNickName        string        `json:"parentNickName"`
				Avatar                string        `json:"avatar"`
				NickName              string        `json:"nickName"`
				DateFormat            string        `json:"dateFormat"`
				Years                 int           `json:"years"`
				Vip                   bool          `json:"vip"`
				VipIcon               string        `json:"vipIcon"`
				VipUrl                string        `json:"vipUrl"`
				CompanyBlog           bool          `json:"companyBlog"`
				CompanyBlogIcon       string        `json:"companyBlogIcon"`
				Flag                  bool          `json:"flag"`
				FlagIcon              string        `json:"flagIcon"`
				LevelIcon             string        `json:"levelIcon"`
				CommentFromTypeResult struct {
					Index int    `json:"index"`
					Key   string `json:"key"`
					Title string `json:"title"`
				} `json:"commentFromTypeResult"`
				IsTop           bool        `json:"isTop"`
				IsBlack         bool        `json:"isBlack"`
				Region          string      `json:"region"`
				OrderNo         string      `json:"orderNo"`
				RedEnvelopeInfo interface{} `json:"redEnvelopeInfo"`
				GptInfo         interface{} `json:"gptInfo"`
			} `json:"sub"`
			PointCommentId interface{} `json:"pointCommentId"`
		} `json:"list"`
	} `json:"data"`
}
