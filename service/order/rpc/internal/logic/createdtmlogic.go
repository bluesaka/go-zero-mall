package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDtmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDtmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDtmLogic {
	return &CreateDtmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDtmLogic) CreateDtm(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Uid,
		})
		if err != nil {
			return fmt.Errorf("用户不存在")
		}

		newOrder := model.Order{
			Uid:    in.Uid,
			Pid:    in.Pid,
			Amount: in.Amount,
			Status: 0,
		}
		// 创建订单
		_, err = l.svcCtx.OrderModel.TxInsert(tx, &newOrder)
		if err != nil {
			return fmt.Errorf("订单创建失败")
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{}, nil
}
