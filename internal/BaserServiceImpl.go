package internal

import (
	"context"
	v1 "go-micro-base-template/api/provider/v1"
)

type ProviderServiceImpl struct {
}

func (c *ProviderServiceImpl) BaserService(ctx context.Context, request *v1.Request, response *v1.Response) error {
	response.Event = request.GetEvent()
	response.ReqId = request.GetReqId()
	response.Body = request.GetBody()
	response.Desc = "success"
	//fmt.Println("%#v", request)
	return nil
}
