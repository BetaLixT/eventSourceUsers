package ctrl

import (
	"github.com/betalixt/eventSourceUsers/svc"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
	"github.com/gin-gonic/gin"
)

type AttachmentController struct {
  svc *svc.AttachmentService
}

func (ctrl *AttachmentController) getAttachmentMeta(ctx *gin.Context) {
  tctx := ctx.MustGet("tctx").(*txcontext.TransactionContext)
  fileId := ctx.Param("fileId")
  res, err := ctrl.svc.GetAttachmentMeta(tctx, fileId)
  if err != nil {
    ctx.Error(err)
  } else {
    ctx.JSON(200, res)
  }
}

func (ctrl *AttachmentController) RegisterRoutes(router *gin.RouterGroup) {
  router.GET(":fileId/meta", ctrl.getAttachmentMeta)
}

func NewAttachmentController(
  svc *svc.AttachmentService,
) *AttachmentController {
  return &AttachmentController{
    svc: svc,
  }
}
