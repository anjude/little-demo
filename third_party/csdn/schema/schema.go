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

type DiggCommentReq struct {
	CommentId string `json:"commentId"`
	ArticleId string `json:"articleId"`
}

type DiggCommentResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TraceId string      `json:"traceId"`
	Data    interface{} `json:"data"`
}

type ArticleCommentListReq struct {
	Type string `json:"type"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

type ArticleCommentListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Data    struct {
		List []struct {
			CommentId             int    `json:"commentId"`
			ArticleId             int    `json:"articleId"`
			BlogId                int    `json:"blogId"`
			ParentId              int    `json:"parentId"`
			PostTime              string `json:"postTime"`
			Content               string `json:"content"`
			UserName              string `json:"userName"`
			Status                int    `json:"status"`
			OpenStatus            int    `json:"openStatus"`
			Digg                  int    `json:"digg"`
			Avatar                string `json:"avatar"`
			Title                 string `json:"title"`
			CommentAuth           int    `json:"commentAuth"`
			ArticleUrl            string `json:"articleUrl"`
			DateFormat            string `json:"dateFormat"`
			Nickname              string `json:"nickname"`
			ArticleAuthorNickname string `json:"articleAuthorNickname"`
		} `json:"list"`
		Total int `json:"total"`
		Page  int `json:"page"`
		Size  int `json:"size"`
	} `json:"data"`
}

type GetUserArticleListReq struct {
	BusinessType string `json:"businessType"`
	Page         int    `json:"page"`
	Size         int    `json:"size"`
	UserName     string `json:"userName"`
}

type GetUserArticleListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Data    struct {
		List []struct {
			ArticleId    int      `json:"articleId"`
			Title        string   `json:"title"`
			Description  string   `json:"description"`
			Url          string   `json:"url"`
			Type         int      `json:"type"`
			Top          bool     `json:"top"`
			ForcePlan    bool     `json:"forcePlan"`
			ViewCount    int      `json:"viewCount"`
			CommentCount int      `json:"commentCount"`
			EditUrl      string   `json:"editUrl"`
			PostTime     string   `json:"postTime"`
			DiggCount    int      `json:"diggCount"`
			FormatTime   string   `json:"formatTime"`
			PicList      []string `json:"picList"`
			CollectCount int      `json:"collectCount"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}
