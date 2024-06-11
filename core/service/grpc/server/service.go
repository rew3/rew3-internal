package server

import (
	"context"

	pAny "github.com/golang/protobuf/ptypes/any"
	"github.com/rew3/rew3-pkg/common/account"
	api "github.com/rew3/rew3-pkg/core/service/api"
	"github.com/rew3/rew3-pkg/core/service/grpc/payload"
	"github.com/rew3/rew3-pkg/core/service/grpc/proto/pb"
	"github.com/rew3/rew3-pkg/core/service/shared/request"
	"github.com/rew3/rew3-pkg/core/service/shared/response/constants"
	"github.com/rew3/rew3-pkg/utils/logger"

	ac "github.com/rew3/rew3-pkg/common/account/constants"
	grpcLib "google.golang.org/grpc"
)

/**
 * Main GRPC Service to handle and execute request for server.
 */
type Service struct {
	pb.UnimplementedServiceProtoServer
	serviceMethodRegistry *api.APIRegistry
}

func NewService(registry *api.APIRegistry) *Service {
	return &Service{serviceMethodRegistry: registry}
}

/**
 * Register this GRPC Service to Server.
 */
func (service *Service) RegisterToServer(server *grpcLib.Server) {
	pb.RegisterServiceProtoServer(server, service)
}

/**
 * Execute Request.
 */
func (service *Service) ExecuteRequest(ctx context.Context, request *pb.RequestPayloadProto) (*pb.ResponsePayloadProto, error) {
	if request.ApiOperation == "" {
		return &pb.ResponsePayloadProto{
			ApiOperation:  request.ApiOperation,
			StatusType:    pb.StatusTypeProto_SERVICE_UNAVAILABLE,
			StatusMessage: "API Method not found.",
		}, nil
	}
	requestPayload := service.requestPayload(request)
	isExists, method := service.serviceMethodRegistry.GetServiceAPI(requestPayload.API)
	if isExists {
		input, err := payload.ToType[map[string]interface{}](requestPayload.Data)
		if err != nil {
			return &pb.ResponsePayloadProto{
				ApiOperation:  string(requestPayload.API),
				StatusType:    pb.StatusTypeProto_SERVICE_UNAVAILABLE,
				StatusMessage: "Unable to parse request - " + err.Error(),
			}, nil
		}
		response := method.Call(ctx, requestPayload.Context, input)
		resPayload := payload.ToResponsePayload(requestPayload.API, response.Status, response.Message, response.Data)
		return service.responsePayloadProto(resPayload), nil
	} else {
		return &pb.ResponsePayloadProto{
			ApiOperation:  string(requestPayload.API),
			StatusType:    pb.StatusTypeProto_SERVICE_UNAVAILABLE,
			StatusMessage: "API Method not found.",
		}, nil
	}
}

/**
 * Convert proto input to Request Payload.
 */
func (service *Service) requestPayload(request *pb.RequestPayloadProto) payload.RequestPayload {
	return payload.RequestPayload{
		API:     api.ResolveEndpoint(request.ApiOperation),
		Context: service.requestContext(request.RequestContext),
		Data:    request.Data.Value,
	}
}

/**
 * Convert to proto response from Response Payload.
 */
func (service *Service) responsePayloadProto(response *payload.ResponsePayload) *pb.ResponsePayloadProto {
	dataBytes, err := payload.ToJson(response.Data)
	if err != nil {
		logger.Log().Error("Error marshaling raw data from ResponsePayload to proto:", err)
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
		Data: &pAny.Any{
			TypeUrl: "json_data", // Provide a type URL to identify the data type
			Value:   dataBytes,   // The byte array containing the JSON data
		},
	}
}

/**
 * Convert proto input to Request Context.
 */
func (service *Service) requestContext(proto *pb.RequestContextProto) request.RequestContext {
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
