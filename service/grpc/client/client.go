package client

import (
	"context"
	"encoding/json"
	"log"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
	"github.com/rew3/rew3-internal/service/grpc/proto/pb"
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
		log.Println("Unable to connect to " + client.serviceName + " Client.")
	}
	client.grpcClient = pb.NewServiceProtoClient(connection)
	log.Println(": Registered : GRPC Client for " + client.serviceName)
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
		log.Fatal("Error marshaling raw data from RequestPayload to proto:", err)
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
		return grpc.ResponsePayload{
			API:           api.APIOperation(response.ApiOperation),
			StatusMessage: response.StatusMessage,
			Status:        status,
			Data:          response.Data.Value,
		}
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
