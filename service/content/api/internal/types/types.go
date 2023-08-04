// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量
	WorkCount       int64  `json:"work_count"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count"`   //点赞数量
}

type Video struct {
	Id            int64  `json:"id"`             // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 视频标题
}

type FeedReq struct {
	Token       string `json:"token,optional"`       // 可选参数，登录用户设置
	Latest_time int64  `json:"latest_time,optional"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

type FeedResp struct {
	StatusCode int32   `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,optional"` // 返回状态描述
	Video_list []Video `json:"video_list"`          // 视频列表
	Next_time  int64   `json:"next_time,optional"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type PublishListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type PublishListResp struct {
	StatusCode int32   `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,optional"` // 返回状态描述
	Video_list []Video `json:"video_list"`          // 视频列表
}

type FavoriteActionReq struct {
	Token      string `json:"token"`       // 用户鉴权token
	VideoId    int64  `json:"video_id"`    // 视频id
	ActionType int32  `json:"action_type"` // 1-点赞，2-取消点赞
}

type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,optional"` // 返回状态描述
}

type FavoriteListReq struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserId int64  `json:"user_id"` // 用户id
}

type FavoriteListResp struct {
	StatusCode int32   `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,optional"` // 返回状态描述
	Video_list []Video `json:"video_list"`          // 视频列表
}

type Comment struct {
	User       User   `json:"user"`        // 评论用户信息
	Id         int64  `json:"id"`          // 评论id
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type CommentActionReq struct {
	Token       string `json:"token"`                 // 用户鉴权token
	VideoId     int64  `json:"video_id"`              // 视频id
	ActionType  int32  `json:"action_type"`           // 1-评论，2-删除评论
	CommentText string `json:"comment_text,optional"` // 评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id,optional"`   // 要删除的评论id，在action_type=2的时候使用
}

type CommentActionResp struct {
	StatusCode int32   `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,optional"` // 返回状态描述
	Comment    Comment `json:"comment,optional"`    // 评论内容，action_type=1时返回
}

type CommentListReq struct {
	Token   string `json:"token"`    // 用户鉴权token
	VideoId int64  `json:"video_id"` // 视频id
}

type CommentListResp struct {
	StatusCode  int32     `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg,optional"` // 返回状态描述
	CommentList []Comment `json:"comment_list"`        // 评论列表
}
