package video

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"toktik-api/internal/api"
	"toktik-api/internal/model/request"
	"toktik-api/internal/model/response"
	"toktik-common/errcode"
	res2 "toktik-common/response"
	"toktik-common/token"
	"toktik-rpc/kitex_gen/favor"
)

type HandlerFavor struct {
}

func NewHandlerVideo() *HandlerFavor {
	return &HandlerFavor{}
}

func (v *HandlerFavor) FavoriteAction(ctx context.Context, c *app.RequestContext) {
	res := res2.NewResponse(c)
	// 1.接收参数 参数模型
	req := &request.FavoriteActionRequest{}
	if err := c.Bind(req); err != nil {
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 2.校验参数
	content, ok := token.GetTokenContent(c)
	if !ok {
		res.Reply(errcode.ErrServer)
		return
	}
	req.UserId = content.ID
	// 3.调用rpc服务获取响应
	params := &favor.FavoriteActionRequest{}
	_ = copier.Copy(params, req)
	result, err := api.FavorClient.FavoriteAction(ctx, params)
	if err != nil {
		res.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	// 4.返回结果
	resp := &response.FollowActionResponse{}
	_ = copier.Copy(resp, result)
	res.Reply(nil, resp)
}
func (v *HandlerFavor) FavoriteList(ctx context.Context, c *app.RequestContext) {
	res := res2.NewResponse(c)
	// 1.接收参数 参数模型
	req := &request.FavoriteListRequest{}
	if err := c.Bind(req); err != nil {
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 2.校验参数
	content, ok := token.GetTokenContent(c)
	if !ok {
		res.Reply(errcode.ErrServer)
		return
	}
	req.MyUserId = content.ID
	// 3.调用rpc服务获取响应
	params := &favor.FavoriteListRequest{}
	_ = copier.Copy(params, req)
	result, err := api.FavorClient.FavoriteList(ctx, params)
	if err != nil {
		res.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	// 4.返回结果
	resp := &response.FavoriteListResponse{}
	_ = copier.Copy(resp, result)
	res.Reply(nil, resp)
}
