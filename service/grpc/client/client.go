package client

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/ptypes/any"
	bConst "github.com/rew3/rew3-internal/app/common/constants"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
	"github.com/rew3/rew3-internal/service/grpc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

/**
 * GRPC Client for Rew3 Service.
 */
type Client struct {
	serverUrl   string
	isInsecure  bool
	serviceName string
	grpcClient  pb.ServiceProtoClient
}

func NewClient(serverUrl string, isInsecureConnection bool, serviceName string) *Client {
	client := &Client{
		serverUrl:   serverUrl,
		isInsecure:  isInsecureConnection,
		serviceName: serviceName,
	}
	client.init()
	return client
}

/**
 * Initialize GRPC Client connection.
 */
func (client *Client) init() {
	url := client.serverUrl
	isInsecure := client.isInsecure
	connection, err := NewConn(url, isInsecure)
	if err != nil {
		logger.Log().Error("Unable to connect to " + client.serviceName + " Client.")
	}
	client.grpcClient = pb.NewServiceProtoClient(connection)
	logger.Log().Info("Registered : GRPC Client for " + client.serviceName)
}

/**
 * Execute Request.
 */
func (client *Client) ExecuteRequest(ctx context.Context, request grpc.RequestPayload) grpc.ResponsePayload {
	if client.grpcClient == nil {
		client.init()
	}
	// Marshal the json.RawMessage to a byte array
	dataBytes, err := json.Marshal(request.Data)
	if err != nil {
		logger.Log().Error("Error marshaling raw data from RequestPayload to proto:", err)
	}
	requestProto := pb.RequestPayloadProto{
		ApiOperation:   string(request.API),
		RequestContext: client.requestContextProto(request.Context),
		Data: &any.Any{
			TypeUrl: "json_data", // Provide a type URL to identify the data type
			Value:   dataBytes,   // The byte array containing the JSON data
		},
	}
	response, err := client.grpcClient.ExecuteRequest(ctx, &requestProto)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			code := grpcErr.Code()
			return grpc.ResponsePayload{
				API:           request.API,
				StatusMessage: client.serviceName + " Error: " + grpcErr.Message(),
				Status:        getErrorStatusFromGrpcCode(code),
			}
		}
		return grpc.ResponsePayload{
			API:           request.API,
			StatusMessage: client.serviceName + " Error: " + err.Error(),
			Status:        constants.INTERNAL_SERVER_ERROR,
		}
	} else {
		status := constants.INTERNAL_SERVER_ERROR
		if len(response.StatusType.String()) != 0 {
			status = constants.StatusType(response.StatusType.String())
		}
		data, _ := grpc.ToType[interface{}](response.Data.Value)
		return grpc.ResponsePayload{
			API:           api.APIOperation(response.ApiOperation),
			StatusMessage: response.StatusMessage,
			Status:        status,
			Data:          data,
			DataMeta:      grpc.DataMeta{Type: resolveDataType(response.DataMeta.Type)},
		}
	}
}

func resolveDataType(dType *pb.DataType) grpc.DataType {
	switch dType.DataType {
	case pb.DataTypeEnum_EMPTY:
		return grpc.Empty{}
	case pb.DataTypeEnum_BINARY:
		return grpc.Binary{}
	case pb.DataTypeEnum_LIST:
		return grpc.List{Type: bConst.Entity(string(dType.TypeMeta.Value))}
	case pb.DataTypeEnum_SCALAR_LIST:
		return grpc.ScalarList{Type: grpc.ScalarType(string(dType.TypeMeta.Value))}
	case pb.DataTypeEnum_OBJECT:
		return grpc.Object{Type: bConst.Entity(string(dType.TypeMeta.Value))}
	case pb.DataTypeEnum_SCALAR:
		return grpc.Scalar{Type: grpc.ScalarType(string(dType.TypeMeta.Value))}
	default:
		return grpc.Empty{}
	}
}

func getErrorStatusFromGrpcCode(code codes.Code) constants.StatusType {
	switch code {
	case codes.OK:
		return constants.OK
	case codes.Canceled:
		return constants.SERVICE_UNAVAILABLE
	case codes.Unknown:
		return constants.INTERNAL_SERVER_ERROR
	case codes.InvalidArgument:
		return constants.BAD_REQUEST
	case codes.DeadlineExceeded:
		return constants.SERVICE_UNAVAILABLE
	case codes.NotFound:
		return constants.NOT_FOUND
	case codes.AlreadyExists:
		return constants.BAD_REQUEST
	case codes.PermissionDenied:
		return constants.FORBIDDEN
	case codes.ResourceExhausted:
		return constants.SERVICE_UNAVAILABLE
	case codes.FailedPrecondition:
		return constants.SERVICE_UNAVAILABLE
	case codes.Aborted:
		return constants.SERVICE_UNAVAILABLE
	case codes.OutOfRange:
		return constants.INTERNAL_SERVER_ERROR
	case codes.Unimplemented:
		return constants.SERVICE_UNAVAILABLE
	case codes.Internal:
		return constants.INTERNAL_SERVER_ERROR
	case codes.Unavailable:
		return constants.SERVICE_UNAVAILABLE
	case codes.DataLoss:
		return constants.INTERNAL_SERVER_ERROR
	case codes.Unauthenticated:
		return constants.UNAUTHORIZED
	default:
		return constants.OK
	}
}

/**
 * Convert request context to proto.
 */
func (client *Client) requestContextProto(rc request.RequestContext) *pb.RequestContextProto {
	// TODO - do auto mapping.
	return &pb.RequestContextProto{
		Member: rc.Member,
		User: &pb.MiniUserProto{
			XId: &wrapperspb.StringValue{
				Value: rc.User.Id,
			},
			FirstName: &wrapperspb.StringValue{
				Value: rc.User.FirstName,
			},
			LastName: &wrapperspb.StringValue{
				Value: rc.User.LastName,
			},
			Userpic: &wrapperspb.StringValue{
				Value: rc.User.Userpic,
			},
			RoleName: &wrapperspb.StringValue{
				Value: rc.User.RoleName,
			},
			Title: &wrapperspb.StringValue{
				Value: rc.User.Title,
			},
			JobTitle: &wrapperspb.StringValue{
				Value: rc.User.JobTitle,
			},
			CompanyName: &wrapperspb.StringValue{
				Value: rc.User.CompanyName,
			},
		},
		FullName: rc.FullName,
		Lang:     rc.Lang,
		Command: &wrapperspb.StringValue{
			Value: rc.Command,
		},
		ETag: &wrapperspb.StringValue{
			Value: rc.ETag,
		},
		Entity: &wrapperspb.StringValue{
			Value: rc.Entity,
		},
		Module: &wrapperspb.StringValue{
			Value: rc.Module,
		},
		IsExternal: rc.IsExternal,
		IsAdmin:    rc.IsAdmin,
		Info: &pb.UserInfoProto{
			Email:       []*pb.EmailProto{},
			PhoneNumber: []*pb.PhoneProto{},
			Personal: &pb.UserPersonalProto{
				FirstName: &wrapperspb.StringValue{
					Value: "",
				},
				MiddleName: &wrapperspb.StringValue{
					Value: "",
				},
				LastName: &wrapperspb.StringValue{
					Value: "",
				},
				FullName: &wrapperspb.StringValue{
					Value: "",
				},
				Salutation: &wrapperspb.StringValue{
					Value: "",
				},
				Gender: &wrapperspb.StringValue{
					Value: "",
				},
			},
			Company: &pb.CompanyProto{
				XId: &wrapperspb.StringValue{
					Value: "",
				},
				NoOfEmployees: &wrapperspb.StringValue{
					Value: "",
				},
				Website: &wrapperspb.StringValue{
					Value: "",
				},
				AnnualRevenue: &wrapperspb.Int64Value{
					Value: 0,
				},
				Industry: &wrapperspb.StringValue{
					Value: "",
				},
				Name: &wrapperspb.StringValue{
					Value: "",
				},
			},
		},
		Timezone: &wrapperspb.StringValue{
			Value: rc.Timezone,
		},
		Teams:            []*pb.TeamMiniAliasProto{},
		SubordinateUsers: rc.SubordinateUsers,
		RulesContext:     []*pb.SharingRuleContextProto{},
		Actions: &pb.UserActionsProto{
			TeamActions: &pb.TeamActionsProto{
				PostSignUpCompleted: false,
			},
		},
		AccountType:  1,
		SecurityType: 1,
	}
}
