package logic

import (
	"context"
	"database/sql"
	"strconv"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UploadContractPhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadContractPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadContractPhotoLogic {
	return &UploadContractPhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadContractPhotoLogic) UploadContractPhoto(req *types.UploadContractorPhotoRequest) (resp *types.UploadContractorPhotoResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var contractorId int64
	if role == variables.Company {
		contractorId = req.Contractor_id
	} else if role == variables.Contractor {
		contractorId = uid
	}

	// Get contractor details
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, contractorId)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Verify company
	if role == variables.Company {
		if contractor_item.FinanceId != uid {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
	}

	// Save image
	photoPath, err := util.SaveImage(req.Contractor_photo, "contractor", strconv.FormatInt(contractorId, 10))
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update contractor details
	contractor_item.ContractorPhoto = sql.NullString{String: photoPath, Valid: photoPath != strconv.FormatInt(contractorId, 10)}
	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.UploadContractorPhotoResponse{
		Contractor_photo: photoPath}, nil
}
