// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	order "github.com/zhoushuguang/lebron/apps/app/api/internal/handler/order"
	user "github.com/zhoushuguang/lebron/apps/app/api/internal/handler/user"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 路由
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/home/banner",
				Handler: HomeBannerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/flashsale",
				Handler: FlashSaleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/recommend",
				Handler: RecommendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/category/list",
				Handler: CategoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/cart/list",
				Handler: CartListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/product/comment",
				Handler: ProductCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/order/list",
				Handler: OrderListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/product/detail",
				Handler: ProductDetailHandler(serverCtx),
			},
		},
	)


	// 登录操作
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	// 登录后操作
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addReceiveAddress",
				Handler: user.AddReceiveAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/editReceiveAddress",
				Handler: user.EditReceiveAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delReceiveAddress",
				Handler: user.DelReceiveAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getReceiveAddressList",
				Handler: user.UserReceiveAddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addCollection",
				Handler: user.UserCollectionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delCollection",
				Handler: user.UserCollectionDelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getCollectionList",
				Handler: user.UserCollectionListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/user"),
	)

	// 添加购物车
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: order.AddOrderHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/order"),
	)
}
