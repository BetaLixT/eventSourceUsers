package svc

import (
	"github.com/betalixt/eventSourceUsers/clnt"
	"github.com/betalixt/eventSourceUsers/clnt/models"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
)

type AttachmentService struct {
	fsclnt *clnt.FileServiceClient
}

func (svc *AttachmentService) GetAttachmentMeta(
	ctx *txcontext.TransactionContext,
	fileId string,
) (*models.FileMeta, error) {
  return svc.fsclnt.GetFileMeta(ctx, fileId)
}

func NewAttachmentService(
  fsclnt *clnt.FileServiceClient,
) *AttachmentService {
  return &AttachmentService{
    fsclnt: fsclnt,
  }
}
