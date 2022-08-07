package conversion

import (
	"github.com/DYSN-Project/gateway/internal/model/response"
	pb "github.com/DYSN-Project/gateway/internal/transport/grpc/pb/area"
)

func FromAreaPbToAreaModel(areaPb *pb.AreaFull, areaEnt *response.Area) {
	areaEnt.Id = areaPb.GetId()
	areaEnt.Name = areaPb.GetName()
	areaEnt.Description = areaPb.GetDescription()
	areaEnt.IsActive = areaPb.GetIsActive()
	areaEnt.CreatedAt = areaPb.GetCreatedAt().AsTime()
	areaEnt.UpdatedAt = areaPb.GetUpdatedAt().AsTime()
}

func FromAreaPbListToAreaModelList(areaPb *pb.AreaList, areaList response.AreaList) {
	for _, v := range areaPb.GetList() {
		area := response.NewAreaIngot()
		FromAreaPbToAreaModel(v, area)
		areaList = append(areaList, area)
	}
}
