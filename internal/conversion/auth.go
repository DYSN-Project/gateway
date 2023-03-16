package conversion

import (
	"github.com/DYSN-Project/gateway/internal/model/response"
	pb "github.com/DYSN-Project/gateway/internal/transport/grpc/pb/auth"
)

func FromUserPbToUserModel(userPb *pb.User, userEnt *response.User) {
	userEnt.ID = userPb.Id
	userEnt.Email = userPb.Email
	userEnt.CreatedAt = userPb.CreatedAt.AsTime()
	userEnt.UpdatedAt = userPb.UpdatedAt.AsTime()
}
