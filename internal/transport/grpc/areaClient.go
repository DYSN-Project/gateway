package grpc

import (
	"context"
	"github.com/DYSN-Project/gateway/internal/conversion"
	"github.com/DYSN-Project/gateway/internal/model/response"
	"github.com/DYSN-Project/gateway/internal/transport/grpc/pb/area"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"google.golang.org/grpc"
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

func (a *Area) List(limit, offset int) (list response.AreaList, err error) {
	data := pb.ListRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rsp, err := a.client.List(context.Background(), &data)
	if err != nil {
		return nil, err
	}
	conversion.FromAreaPbListToAreaModelList(rsp, list)

	return list, nil
}

func (a *Area) Show(id string) (area *response.Area, err error) {
	data := pb.Identity{
		Id: id,
	}

	rsp, err := a.client.Show(context.Background(), &data)
	if err != nil {
		return nil, err
	}
	conversion.FromAreaPbToAreaModel(rsp, area)

	return area, nil
}

func (a *Area) CreateArea(name, description string) (area *response.Area, err error) {
	data := pb.CreateAreaRequest{
		Name:        name,
		Description: description,
	}

	rsp, err := a.client.CreateArea(context.Background(), &data)
	if err != nil {
		return nil, err
	}
	conversion.FromAreaPbToAreaModel(rsp, area)

	return area, nil
}

func (a *Area) UpdateArea(id, name, description string, isActive bool) error {
	data := pb.UpdateAreaRequest{
		Id:          id,
		Name:        name,
		Description: description,
		IsActive:    isActive,
	}

	_, err := a.client.UpdateArea(context.Background(), &data)
	if err != nil {
		return err
	}

	return nil
}

func (a *Area) DeleteArea(id string) error {
	data := pb.Identity{
		Id: id,
	}

	_, err := a.client.DeleteArea(context.Background(), &data)
	if err != nil {
		return err
	}

	return nil
}
