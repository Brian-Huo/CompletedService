package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCompanyLogic {
	return &ListCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCompanyLogic) ListCompany(req *types.ListCompanyRequest) (resp *types.ListCompanyResponse, err error) {
	res, err := l.svcCtx.BCompanyModel.List(l.ctx)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Invalid, Company not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailCompanyResponse{}

	for _, item := range res {
		newItem := types.DetailCompanyResponse{
			Company_id:      item.CompanyId,
			Company_name:    item.CompanyName,
			Director_name:   item.DirectorName.String,
			Contact_details: item.ContactDetails,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListCompanyResponse{
		Items: allItems,
	}, nil
}
