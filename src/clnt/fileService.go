package clnt

import (
	"fmt"

	"github.com/betalixt/eventSourceUsers/clnt/models"
	"github.com/betalixt/eventSourceUsers/optn"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
)

type FileServiceClient struct {
  optn *optn.FileServiceClientOptions
}

func NewFileServiceClient(
  options *optn.FileServiceClientOptions,
) *FileServiceClient {
  return &FileServiceClient{
    optn: options,
  }
}

func (clnt *FileServiceClient) GetFileMeta(
  ctx *txcontext.TransactionContext,
  fileId string,
) (*models.FileMeta, error) {
	http := ctx.GetHttpClient()
	res, err := http.Get(
		nil,
		fmt.Sprintf("%s/api/files/{}/meta", clnt.optn.BaseUrl),
		nil,
		fileId,
	)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response from upstream")
	}
	meta := models.FileMeta{}
	err = res.Unmarshal(&meta)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
  return &meta, nil
}
