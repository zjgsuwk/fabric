// Code generated by protoc-gen-go.
// source: devops.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuildResult_StatusCode int32

const (
	BuildResult_UNDEFINED BuildResult_StatusCode = 0
	BuildResult_SUCCESS   BuildResult_StatusCode = 1
	BuildResult_FAILURE   BuildResult_StatusCode = 2
)

var BuildResult_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "SUCCESS",
	2: "FAILURE",
}
var BuildResult_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"SUCCESS":   1,
	"FAILURE":   2,
}

func (x BuildResult_StatusCode) String() string {
	return proto.EnumName(BuildResult_StatusCode_name, int32(x))
}

// Secret is a temporary object to establish security with the Devops.
// A better solution using certificate will be introduced later
type Secret struct {
	EnrollId     string `protobuf:"bytes,1,opt,name=enrollId" json:"enrollId,omitempty"`
	EnrollSecret string `protobuf:"bytes,2,opt,name=enrollSecret" json:"enrollSecret,omitempty"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}

type SigmaInput struct {
	Secret   *Secret `protobuf:"bytes,1,opt,name=secret" json:"secret,omitempty"`
	AppTCert []byte  `protobuf:"bytes,2,opt,name=appTCert,proto3" json:"appTCert,omitempty"`
	Data     []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SigmaInput) Reset()         { *m = SigmaInput{} }
func (m *SigmaInput) String() string { return proto.CompactTextString(m) }
func (*SigmaInput) ProtoMessage()    {}

func (m *SigmaInput) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

type ExecuteWithBinding struct {
	ChaincodeInvocationSpec *ChaincodeInvocationSpec `protobuf:"bytes,1,opt,name=chaincodeInvocationSpec" json:"chaincodeInvocationSpec,omitempty"`
	Binding                 []byte                   `protobuf:"bytes,2,opt,name=binding,proto3" json:"binding,omitempty"`
}

func (m *ExecuteWithBinding) Reset()         { *m = ExecuteWithBinding{} }
func (m *ExecuteWithBinding) String() string { return proto.CompactTextString(m) }
func (*ExecuteWithBinding) ProtoMessage()    {}

func (m *ExecuteWithBinding) GetChaincodeInvocationSpec() *ChaincodeInvocationSpec {
	if m != nil {
		return m.ChaincodeInvocationSpec
	}
	return nil
}

type SigmaOutput struct {
	Tcert        []byte `protobuf:"bytes,1,opt,name=tcert,proto3" json:"tcert,omitempty"`
	Sigma        []byte `protobuf:"bytes,2,opt,name=sigma,proto3" json:"sigma,omitempty"`
	Asn1Encoding []byte `protobuf:"bytes,3,opt,name=asn1Encoding,proto3" json:"asn1Encoding,omitempty"`
}

func (m *SigmaOutput) Reset()         { *m = SigmaOutput{} }
func (m *SigmaOutput) String() string { return proto.CompactTextString(m) }
func (*SigmaOutput) ProtoMessage()    {}

type BuildResult struct {
	Status         BuildResult_StatusCode   `protobuf:"varint,1,opt,name=status,enum=protos.BuildResult_StatusCode" json:"status,omitempty"`
	Msg            string                   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	DeploymentSpec *ChaincodeDeploymentSpec `protobuf:"bytes,3,opt,name=deploymentSpec" json:"deploymentSpec,omitempty"`
}

func (m *BuildResult) Reset()         { *m = BuildResult{} }
func (m *BuildResult) String() string { return proto.CompactTextString(m) }
func (*BuildResult) ProtoMessage()    {}

func (m *BuildResult) GetDeploymentSpec() *ChaincodeDeploymentSpec {
	if m != nil {
		return m.DeploymentSpec
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.BuildResult_StatusCode", BuildResult_StatusCode_name, BuildResult_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Devops service

type DevopsClient interface {
	// Log in - passed Secret object and returns Response object, where
	// msg is the security context to be used in subsequent invocations
	Login(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error)
	// Build the chaincode package.
	Build(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error)
	// Deploy the chaincode package to the chain.
	Deploy(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error)
	// Invoke chaincode.
	Invoke(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error)
	// Invoke chaincode.
	Query(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error)
	// Retrieve a TCert.
	EXP_GetApplicationTCert(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_PrepareForTx(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_ProduceSigma(ctx context.Context, in *SigmaInput, opts ...grpc.CallOption) (*Response, error)
	// Execute a transaction with a specific binding
	EXP_ExecuteWithBinding(ctx context.Context, in *ExecuteWithBinding, opts ...grpc.CallOption) (*Response, error)
}

type devopsClient struct {
	cc *grpc.ClientConn
}

func NewDevopsClient(cc *grpc.ClientConn) DevopsClient {
	return &devopsClient{cc}
}

func (c *devopsClient) Login(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Build(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error) {
	out := new(ChaincodeDeploymentSpec)
	err := grpc.Invoke(ctx, "/protos.Devops/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Deploy(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error) {
	out := new(ChaincodeDeploymentSpec)
	err := grpc.Invoke(ctx, "/protos.Devops/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Invoke(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Query(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_GetApplicationTCert(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_GetApplicationTCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_PrepareForTx(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_PrepareForTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_ProduceSigma(ctx context.Context, in *SigmaInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_ProduceSigma", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_ExecuteWithBinding(ctx context.Context, in *ExecuteWithBinding, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_ExecuteWithBinding", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Devops service

type DevopsServer interface {
	// Log in - passed Secret object and returns Response object, where
	// msg is the security context to be used in subsequent invocations
	Login(context.Context, *Secret) (*Response, error)
	// Build the chaincode package.
	Build(context.Context, *ChaincodeSpec) (*ChaincodeDeploymentSpec, error)
	// Deploy the chaincode package to the chain.
	Deploy(context.Context, *ChaincodeSpec) (*ChaincodeDeploymentSpec, error)
	// Invoke chaincode.
	Invoke(context.Context, *ChaincodeInvocationSpec) (*Response, error)
	// Invoke chaincode.
	Query(context.Context, *ChaincodeInvocationSpec) (*Response, error)
	// Retrieve a TCert.
	EXP_GetApplicationTCert(context.Context, *Secret) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_PrepareForTx(context.Context, *Secret) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_ProduceSigma(context.Context, *SigmaInput) (*Response, error)
	// Execute a transaction with a specific binding
	EXP_ExecuteWithBinding(context.Context, *ExecuteWithBinding) (*Response, error)
}

func RegisterDevopsServer(s *grpc.Server, srv DevopsServer) {
	s.RegisterService(&_Devops_serviceDesc, srv)
}

func _Devops_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Build(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Deploy(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Invoke(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Query(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_EXP_GetApplicationTCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).EXP_GetApplicationTCert(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_EXP_PrepareForTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).EXP_PrepareForTx(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_EXP_ProduceSigma_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SigmaInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).EXP_ProduceSigma(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_EXP_ExecuteWithBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ExecuteWithBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).EXP_ExecuteWithBinding(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Devops_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Devops",
	HandlerType: (*DevopsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Devops_Login_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Devops_Build_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Devops_Deploy_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _Devops_Invoke_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Devops_Query_Handler,
		},
		{
			MethodName: "EXP_GetApplicationTCert",
			Handler:    _Devops_EXP_GetApplicationTCert_Handler,
		},
		{
			MethodName: "EXP_PrepareForTx",
			Handler:    _Devops_EXP_PrepareForTx_Handler,
		},
		{
			MethodName: "EXP_ProduceSigma",
			Handler:    _Devops_EXP_ProduceSigma_Handler,
		},
		{
			MethodName: "EXP_ExecuteWithBinding",
			Handler:    _Devops_EXP_ExecuteWithBinding_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
