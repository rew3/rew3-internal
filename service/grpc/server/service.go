package server

import (
	"context"

	pAny "github.com/golang/protobuf/ptypes/any"
	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
	"github.com/rew3/rew3-internal/service/grpc/proto/pb"

	ac "github.com/rew3/rew3-internal/app/account/constants"
	grpcLib "google.golang.org/grpc"
)

/**
 * Main GRPC Service to handle and execute request for server.
 */
type Service struct {
	pb.UnimplementedServiceProtoServer
	serviceMethodRegistry *ServiceMethodRegistry
}

func NewService(registry *ServiceMethodRegistry) *Service {
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
	isExists, method := service.serviceMethodRegistry.GetServiceMethod(requestPayload.API)
	if isExists {
		response := method.call(ctx, requestPayload)
		return service.responsePayloadProto(response), nil
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
func (service *Service) requestPayload(request *pb.RequestPayloadProto) grpc.RequestPayload {
	return grpc.RequestPayload{
		API:     api.ResolveOperation(request.ApiOperation),
		Context: service.requestContext(request.RequestContext),
		Data:    request.Data.Value,
	}
}

/**
 * Convert to proto response from Response Payload.
 */
func (service *Service) responsePayloadProto(response *grpc.ResponsePayload) *pb.ResponsePayloadProto {
	dataBytes, err := grpc.ToJson(response.Data)
	if err != nil {
		logger.Log().Error("Error marshaling raw data from ResponsePayload to proto:", err)
		return &pb.ResponsePayloadProto{
			ApiOperation:  string(response.API),
			StatusType:    pb.StatusTypeProto_INTERNAL_SERVER_ERROR,
			StatusMessage: "Error marshaling raw data from ResponsePayload to proto:: " + err.Error(),
		}
	}
	/*dataMetaTypeBytes, err := json.Marshal(response.DataMeta.Type)
	if err != nil {
		logger.Log().Error("Error marshaling raw data from ResponsePayload to proto:", err)
		return &pb.ResponsePayloadProto{
			ApiOperation:  string(response.API),
			StatusType:    pb.StatusTypeProto_INTERNAL_SERVER_ERROR,
			StatusMessage: "Error marshaling raw data from ResponsePayload to proto:: " + err.Error(),
		}
	}*/
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
		DataMeta: &pb.DataMeta{
			Type: resolveDataType(response.DataMeta.Type),
		},
	}
}

func resolveDataType(dataType grpc.DataType) *pb.DataType {
	switch t := dataType.GetType().(type) {
	case grpc.Empty:
		return &pb.DataType{DataType: pb.DataTypeEnum_EMPTY}
	case grpc.Binary:
		return &pb.DataType{DataType: pb.DataTypeEnum_BINARY}
	case grpc.List:
		return &pb.DataType{DataType: pb.DataTypeEnum_LIST, TypeMeta: &pAny.Any{
			TypeUrl: "json_data",
			Value:   []byte(t.Type),
		}}
	case grpc.ScalarList:
		return &pb.DataType{DataType: pb.DataTypeEnum_SCALAR_LIST, TypeMeta: &pAny.Any{
			TypeUrl: "json_data",
			Value:   []byte(t.Type),
		}}
	case grpc.Object:
		return &pb.DataType{DataType: pb.DataTypeEnum_OBJECT, TypeMeta: &pAny.Any{
			TypeUrl: "json_data",
			Value:   []byte(t.Type),
		}}
	case grpc.Scalar:
		return &pb.DataType{DataType: pb.DataTypeEnum_SCALAR, TypeMeta: &pAny.Any{
			TypeUrl: "json_data",
			Value:   []byte(t.Type),
		}}
	default:
		return nil
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
