package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
	"github.com/rew3/rew3-internal/service/grpc/proto/pb"

	ac "github.com/rew3/rew3-internal/app/account/constants"
)

type GRPCService struct {
	pb.UnimplementedServiceProtoServer
	requestHandler RequestHandler
}

func NewGRPCService(handler RequestHandler) *GRPCService {
	return &GRPCService{requestHandler: handler}
}

func (service *GRPCService) ExecuteRequest(ctx context.Context, request *pb.RequestPayloadProto) (*pb.ResponsePayloadProto, error) {
	requestPayload := service.requestPayload(request)
	response := service.requestHandler.Handle(requestPayload)
	return service.responsePayloadProto(&response), nil
}

func (service *GRPCService) requestPayload(request *pb.RequestPayloadProto) grpc.RequestPayload {
	return grpc.RequestPayload{
		API:     api.ResolveOperation(request.ApiOperation),
		Context: service.requestContext(request.RequestContext),
		Data:    request.Data.Value,
	}
}

func (service *GRPCService) responsePayloadProto(response *grpc.ResponsePayload) *pb.ResponsePayloadProto {
	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		log.Fatal("Error marshaling raw data from ResponsePayload to proto:", err)
		return &pb.ResponsePayloadProto{
			ApiOperation:  string(response.API),
			StatusType:    pb.StatusTypeProto_INTERNAL_SERVER_ERROR,
			StatusMessage: "Error marshaling raw data from ResponsePayload to proto:: " + err.Error(),
		}
	}
	statusMap := map[constants.StatusType]pb.StatusTypeProto{
		constants.OK:                    pb.StatusTypeProto_OK,
		constants.CREATED:               pb.StatusTypeProto_CREATED,
		constants.DELETED:               pb.StatusTypeProto_DELETED,
		constants.ACCEPTED:              pb.StatusTypeProto_ACCEPTED,
		constants.BAD_REQUEST:           pb.StatusTypeProto_BAD_REQUEST,
		constants.UNAUTHORIZED:          pb.StatusTypeProto_UNAUTHORIZED,
		constants.FORBIDDEN:             pb.StatusTypeProto_FORBIDDEN,
		constants.INTERNAL_SERVER_ERROR: pb.StatusTypeProto_INTERNAL_SERVER_ERROR,
		constants.BAD_GATEWAY:           pb.StatusTypeProto_BAD_GATEWAY,
		constants.SERVICE_UNAVAILABLE:   pb.StatusTypeProto_SERVICE_UNAVAILABLE,
		constants.GATEWAY_TIMEOUT:       pb.StatusTypeProto_GATEWAY_TIMEOUT,
		constants.NOT_FOUND:             pb.StatusTypeProto_NOT_FOUND,
	}
	return &pb.ResponsePayloadProto{
		ApiOperation:  string(response.API),
		StatusType:    statusMap[response.Status],
		StatusMessage: response.StatusMessage,
		Data: &any.Any{
			TypeUrl: "json_data", // Provide a type URL to identify the data type
			Value:   dataBytes,   // The byte array containing the JSON data
		},
	}
}

func (service *GRPCService) requestContext(proto *pb.RequestContextProto) request.RequestContext {
	if proto == nil {
		return request.RequestContext{}
	}
	return request.RequestContext{
		Member: proto.Member,
		User: account.MiniUser{
			Id:        proto.User.XId.Value,
			FirstName: proto.User.FirstName.Value,
			LastName:  proto.User.LastName.Value,
		},
		FullName:     proto.User.FirstName.Value + " " + proto.User.LastName.Value,
		Lang:         proto.Lang,
		Entity:       proto.Entity.Value,
		Module:       proto.Module.Value,
		IsExternal:   proto.IsExternal,
		IsAdmin:      proto.IsAdmin,
		AccountType:  ac.AccountTypeAlias(proto.AccountType.String()),
		SecurityType: ac.SecurityTypeAlias(proto.SecurityType.String()),
	}
}
