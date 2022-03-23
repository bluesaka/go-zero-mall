// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package server

import (
	"context"

	"go-zero-mall/service/product/rpc/internal/logic"
	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Create(ctx context.Context, in *product.CreateRequest) (*product.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *ProductServer) Update(ctx context.Context, in *product.UpdateRequest) (*product.UpdateResponse, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *ProductServer) Remove(ctx context.Context, in *product.RemoveRequest) (*product.RemoveResponse, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

func (s *ProductServer) Detail(ctx context.Context, in *product.DetailRequest) (*product.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *ProductServer) DecrStock(ctx context.Context, in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	l := logic.NewDecrStockLogic(ctx, s.svcCtx)
	return l.DecrStock(in)
}

func (s *ProductServer) DecrStockRevert(ctx context.Context, in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	l := logic.NewDecrStockRevertLogic(ctx, s.svcCtx)
	return l.DecrStockRevert(in)
}
