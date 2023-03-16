package grpc

import (
	"context"
	"github.com/DYSN-Project/gateway/internal/conversion"
	"github.com/DYSN-Project/gateway/internal/model/form"
	"github.com/DYSN-Project/gateway/internal/model/response"
	"github.com/DYSN-Project/gateway/internal/transport/grpc/pb/area"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Area struct {
	client pb.AreaClient
}

func NewArea(port string, errLog *logger.Logger) *Area {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(port, opt...)
	if err != nil {
		errLog.ErrorLog.Fatal(err)
	}

	return &Area{
		client: pb.NewAreaClient(conn),
	}
}

func (a *Area) List() (list response.AreaList, err error) {
	rsp, err := a.client.ListArea(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	conversion.FromAreaPbListToAreaModelList(rsp, list)

	return list, nil
}

func (a *Area) CreateArea(frm *form.CreateArea) (area *response.Area, err error) {
	data := &pb.CreateRequest{
		Name:        frm.Name,
		Description: frm.Description,
		Image:       frm.Image,
		Status:      frm.Status,
	}

	rsp, err := a.client.CreateArea(context.Background(), data)
	if err != nil {
		return nil, err
	}
	conversion.FromAreaPbToAreaModel(rsp, area)

	return area, nil
}

func (a *Area) UpdateArea(frm *form.UpdateArea) error {
	data := &pb.UpdateRequest{
		Id:          frm.ID,
		Name:        frm.Name,
		Description: frm.Description,
		Status:      frm.Status,
		Image:       frm.Image,
	}

	_, err := a.client.UpdateArea(context.Background(), data)
	if err != nil {
		return err
	}

	return nil
}

func (a *Area) DeleteArea(id string) error {
	data := &pb.AreaId{
		Id: id,
	}

	_, err := a.client.DeleteArea(context.Background(), data)
	if err != nil {
		return err
	}

	return nil
}
