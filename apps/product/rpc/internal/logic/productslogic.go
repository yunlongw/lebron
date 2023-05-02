package logic

import (
	"context"
	"strconv"
	"strings"

	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/model"
	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	products := make(map[int64]*product.ProductItem)

	pdis := strings.Split(in.ProductIds, ",")
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		// 将 PID 挨个放入管道
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		pidStr := item.(string)
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			return
		}
		// 查询 item 对应的商品数据
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, pid)
		if err != nil {
			return
		}

		// 将查询结果写入到 writer
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var r []*model.Product
		// 将缓冲的商品信息放到一个 product 集合中
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}

	// ps 返回的就是整个 reducerFunc 组装的 product 集合
	for _, p := range ps.([]*model.Product) {
		products[p.Id] = &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		}
	}
	return &product.ProductResponse{Products: products}, nil
}
