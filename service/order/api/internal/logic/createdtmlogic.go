package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"go-zero-mall/service/order/rpc/order"
	"go-zero-mall/service/product/rpc/product"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDtmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDtmLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateDtmLogic {
	return CreateDtmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDtmLogic) CreateDtm(req types.CreateRequest) (resp *types.CreateResponse, err error) {
	// 获取OrderRpc BuildTarget
	orderRpcBuilder, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// 获取ProductRpc BuildTarget，类似HTTP的Host
	productRpcBuilder, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://etcd:2379/dtmservice"

	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)

	// dtm子事务屏障数据库表，默认为 dtm_barrier.barrier，可通过方法自定义
	//dtmcli.SetBarrierTableName("dtm_xxx.test")

	orderReq := &order.CreateRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: 0,
	}

	productReq := &product.DecrStockRequest{
		Id:  req.Pid,
		Num: 1,
	}

	// 创建一个SAGA协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		// URL可以从 `c.cc.Invoke` 中找到
		Add(orderRpcBuilder+"/orderclient.Order/CreateDtm", orderRpcBuilder+"/orderclient.Order/CreateRevert", orderReq).
		Add(productRpcBuilder+"/productclient.Product/DecrStock", productRpcBuilder+"/productclient.Product/DecrStockRevert", productReq)

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{}, nil
}
