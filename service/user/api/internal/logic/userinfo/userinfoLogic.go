package userinfo

import (
	"context"
	"doushen_by_liujun/internal/common"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/user/api/internal/svc"
	"doushen_by_liujun/service/user/api/internal/types"
	"doushen_by_liujun/service/user/rpc/pb"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserinfoReq) (resp *types.UserinfoResp, err error) {
	// todo: add your logic here and delete this line
	logger, e := util.ParseToken(req.Token)
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.TOKEN_EXPIRE_ERROR,
			StatusMsg:  "无效token",
			User:       types.User{},
		}, err
	}
	info, e := l.svcCtx.UserRpcClient.GetUserinfoById(l.ctx, &pb.GetUserinfoByIdReq{
		Id: req.UserId,
	})
	var user types.User
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.DB_ERROR,
			StatusMsg:  e.Error(),
			User:       user,
		}, err
	}
	followCount, e := l.svcCtx.UserRpcClient.GetFollowsCountById(l.ctx, &pb.GetFollowsCountByIdReq{
		Id: req.UserId,
	})
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.DB_ERROR,
			StatusMsg:  "查询关注数量失败",
			User:       user,
		}, err
	}
	followerCount, e := l.svcCtx.UserRpcClient.GetFollowersCountById(l.ctx, &pb.GetFollowersCountByIdReq{
		Id: req.UserId,
	})
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.DB_ERROR,
			StatusMsg:  "查询粉丝数量失败",
			User:       user,
		}, err
	}
	isFollow, e := l.svcCtx.UserRpcClient.CheckIsFollow(l.ctx, &pb.CheckIsFollowReq{
		Userid:   logger.ID,
		Followid: strconv.Itoa(int(info.Userinfo.Id)),
	})
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.DB_ERROR,
			StatusMsg:  "查询粉丝数量失败",
			User:       user,
		}, err
	}
	user = types.User{
		UserId:          info.Userinfo.Id,
		Name:            info.Userinfo.Name,
		FollowCount:     followCount.Count,
		FollowerCount:   followerCount.Count,
		IsFollow:        isFollow.IsFollowed, //我对这个的理解就是当前用户对这条数据的用户是否关注
		Avatar:          info.Userinfo.Avatar,
		BackgroundImage: info.Userinfo.BackgroundImage,
		Signature:       info.Userinfo.Signature,
		WorkCount:       0, //查表
		FavoriteCount:   0, //查表
		TotalFavorited:  0, //查表
	}
	return &types.UserinfoResp{
		StatusCode: common.OK,
		StatusMsg:  "查询成功",
		User:       user,
	}, nil
}
