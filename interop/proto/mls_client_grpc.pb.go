// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MLSClientClient is the client API for MLSClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MLSClientClient interface {
	// The human-readable name of the stack
	Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error)
	// List of supported ciphersuites
	SupportedCiphersuites(ctx context.Context, in *SupportedCiphersuitesRequest, opts ...grpc.CallOption) (*SupportedCiphersuitesResponse, error)
	// Test vector generation and verification
	GenerateTestVector(ctx context.Context, in *GenerateTestVectorRequest, opts ...grpc.CallOption) (*GenerateTestVectorResponse, error)
	VerifyTestVector(ctx context.Context, in *VerifyTestVectorRequest, opts ...grpc.CallOption) (*VerifyTestVectorResponse, error)
	// Ways to become a member of a group
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	CreateKeyPackage(ctx context.Context, in *CreateKeyPackageRequest, opts ...grpc.CallOption) (*CreateKeyPackageResponse, error)
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
	ExternalJoin(ctx context.Context, in *ExternalJoinRequest, opts ...grpc.CallOption) (*ExternalJoinResponse, error)
	// Operations using a group state
	PublicGroupState(ctx context.Context, in *PublicGroupStateRequest, opts ...grpc.CallOption) (*PublicGroupStateResponse, error)
	StateAuth(ctx context.Context, in *StateAuthRequest, opts ...grpc.CallOption) (*StateAuthResponse, error)
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
	Protect(ctx context.Context, in *ProtectRequest, opts ...grpc.CallOption) (*ProtectResponse, error)
	Unprotect(ctx context.Context, in *UnprotectRequest, opts ...grpc.CallOption) (*UnprotectResponse, error)
	StorePSK(ctx context.Context, in *StorePSKRequest, opts ...grpc.CallOption) (*StorePSKResponse, error)
	AddProposal(ctx context.Context, in *AddProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error)
	UpdateProposal(ctx context.Context, in *UpdateProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error)
	RemoveProposal(ctx context.Context, in *RemoveProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error)
	PSKProposal(ctx context.Context, in *PSKProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error)
	ReInitProposal(ctx context.Context, in *ReInitProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error)
	AppAckProposal(ctx context.Context, in *AppAckProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error)
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
	HandleCommit(ctx context.Context, in *HandleCommitRequest, opts ...grpc.CallOption) (*HandleCommitResponse, error)
	HandleExternalCommit(ctx context.Context, in *HandleExternalCommitRequest, opts ...grpc.CallOption) (*HandleExternalCommitResponse, error)
}

type mLSClientClient struct {
	cc grpc.ClientConnInterface
}

func NewMLSClientClient(cc grpc.ClientConnInterface) MLSClientClient {
	return &mLSClientClient{cc}
}

func (c *mLSClientClient) Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) SupportedCiphersuites(ctx context.Context, in *SupportedCiphersuitesRequest, opts ...grpc.CallOption) (*SupportedCiphersuitesResponse, error) {
	out := new(SupportedCiphersuitesResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/SupportedCiphersuites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) GenerateTestVector(ctx context.Context, in *GenerateTestVectorRequest, opts ...grpc.CallOption) (*GenerateTestVectorResponse, error) {
	out := new(GenerateTestVectorResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/GenerateTestVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) VerifyTestVector(ctx context.Context, in *VerifyTestVectorRequest, opts ...grpc.CallOption) (*VerifyTestVectorResponse, error) {
	out := new(VerifyTestVectorResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/VerifyTestVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) CreateKeyPackage(ctx context.Context, in *CreateKeyPackageRequest, opts ...grpc.CallOption) (*CreateKeyPackageResponse, error) {
	out := new(CreateKeyPackageResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/CreateKeyPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/JoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) ExternalJoin(ctx context.Context, in *ExternalJoinRequest, opts ...grpc.CallOption) (*ExternalJoinResponse, error) {
	out := new(ExternalJoinResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/ExternalJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) PublicGroupState(ctx context.Context, in *PublicGroupStateRequest, opts ...grpc.CallOption) (*PublicGroupStateResponse, error) {
	out := new(PublicGroupStateResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/PublicGroupState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) StateAuth(ctx context.Context, in *StateAuthRequest, opts ...grpc.CallOption) (*StateAuthResponse, error) {
	out := new(StateAuthResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/StateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) Protect(ctx context.Context, in *ProtectRequest, opts ...grpc.CallOption) (*ProtectResponse, error) {
	out := new(ProtectResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/Protect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) Unprotect(ctx context.Context, in *UnprotectRequest, opts ...grpc.CallOption) (*UnprotectResponse, error) {
	out := new(UnprotectResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/Unprotect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) StorePSK(ctx context.Context, in *StorePSKRequest, opts ...grpc.CallOption) (*StorePSKResponse, error) {
	out := new(StorePSKResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/StorePSK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) AddProposal(ctx context.Context, in *AddProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/AddProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) UpdateProposal(ctx context.Context, in *UpdateProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/UpdateProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) RemoveProposal(ctx context.Context, in *RemoveProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/RemoveProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) PSKProposal(ctx context.Context, in *PSKProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/PSKProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) ReInitProposal(ctx context.Context, in *ReInitProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/ReInitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) AppAckProposal(ctx context.Context, in *AppAckProposalRequest, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/AppAckProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) HandleCommit(ctx context.Context, in *HandleCommitRequest, opts ...grpc.CallOption) (*HandleCommitResponse, error) {
	out := new(HandleCommitResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/HandleCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mLSClientClient) HandleExternalCommit(ctx context.Context, in *HandleExternalCommitRequest, opts ...grpc.CallOption) (*HandleExternalCommitResponse, error) {
	out := new(HandleExternalCommitResponse)
	err := c.cc.Invoke(ctx, "/mls_client.MLSClient/HandleExternalCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MLSClientServer is the server API for MLSClient service.
// All implementations must embed UnimplementedMLSClientServer
// for forward compatibility
type MLSClientServer interface {
	// The human-readable name of the stack
	Name(context.Context, *NameRequest) (*NameResponse, error)
	// List of supported ciphersuites
	SupportedCiphersuites(context.Context, *SupportedCiphersuitesRequest) (*SupportedCiphersuitesResponse, error)
	// Test vector generation and verification
	GenerateTestVector(context.Context, *GenerateTestVectorRequest) (*GenerateTestVectorResponse, error)
	VerifyTestVector(context.Context, *VerifyTestVectorRequest) (*VerifyTestVectorResponse, error)
	// Ways to become a member of a group
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	CreateKeyPackage(context.Context, *CreateKeyPackageRequest) (*CreateKeyPackageResponse, error)
	JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error)
	ExternalJoin(context.Context, *ExternalJoinRequest) (*ExternalJoinResponse, error)
	// Operations using a group state
	PublicGroupState(context.Context, *PublicGroupStateRequest) (*PublicGroupStateResponse, error)
	StateAuth(context.Context, *StateAuthRequest) (*StateAuthResponse, error)
	Export(context.Context, *ExportRequest) (*ExportResponse, error)
	Protect(context.Context, *ProtectRequest) (*ProtectResponse, error)
	Unprotect(context.Context, *UnprotectRequest) (*UnprotectResponse, error)
	StorePSK(context.Context, *StorePSKRequest) (*StorePSKResponse, error)
	AddProposal(context.Context, *AddProposalRequest) (*ProposalResponse, error)
	UpdateProposal(context.Context, *UpdateProposalRequest) (*ProposalResponse, error)
	RemoveProposal(context.Context, *RemoveProposalRequest) (*ProposalResponse, error)
	PSKProposal(context.Context, *PSKProposalRequest) (*ProposalResponse, error)
	ReInitProposal(context.Context, *ReInitProposalRequest) (*ProposalResponse, error)
	AppAckProposal(context.Context, *AppAckProposalRequest) (*ProposalResponse, error)
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
	HandleCommit(context.Context, *HandleCommitRequest) (*HandleCommitResponse, error)
	HandleExternalCommit(context.Context, *HandleExternalCommitRequest) (*HandleExternalCommitResponse, error)
	mustEmbedUnimplementedMLSClientServer()
}

// UnimplementedMLSClientServer must be embedded to have forward compatible implementations.
type UnimplementedMLSClientServer struct {
}

func (UnimplementedMLSClientServer) Name(context.Context, *NameRequest) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedMLSClientServer) SupportedCiphersuites(context.Context, *SupportedCiphersuitesRequest) (*SupportedCiphersuitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupportedCiphersuites not implemented")
}
func (UnimplementedMLSClientServer) GenerateTestVector(context.Context, *GenerateTestVectorRequest) (*GenerateTestVectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTestVector not implemented")
}
func (UnimplementedMLSClientServer) VerifyTestVector(context.Context, *VerifyTestVectorRequest) (*VerifyTestVectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTestVector not implemented")
}
func (UnimplementedMLSClientServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedMLSClientServer) CreateKeyPackage(context.Context, *CreateKeyPackageRequest) (*CreateKeyPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeyPackage not implemented")
}
func (UnimplementedMLSClientServer) JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedMLSClientServer) ExternalJoin(context.Context, *ExternalJoinRequest) (*ExternalJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExternalJoin not implemented")
}
func (UnimplementedMLSClientServer) PublicGroupState(context.Context, *PublicGroupStateRequest) (*PublicGroupStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicGroupState not implemented")
}
func (UnimplementedMLSClientServer) StateAuth(context.Context, *StateAuthRequest) (*StateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StateAuth not implemented")
}
func (UnimplementedMLSClientServer) Export(context.Context, *ExportRequest) (*ExportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedMLSClientServer) Protect(context.Context, *ProtectRequest) (*ProtectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Protect not implemented")
}
func (UnimplementedMLSClientServer) Unprotect(context.Context, *UnprotectRequest) (*UnprotectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unprotect not implemented")
}
func (UnimplementedMLSClientServer) StorePSK(context.Context, *StorePSKRequest) (*StorePSKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePSK not implemented")
}
func (UnimplementedMLSClientServer) AddProposal(context.Context, *AddProposalRequest) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProposal not implemented")
}
func (UnimplementedMLSClientServer) UpdateProposal(context.Context, *UpdateProposalRequest) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProposal not implemented")
}
func (UnimplementedMLSClientServer) RemoveProposal(context.Context, *RemoveProposalRequest) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProposal not implemented")
}
func (UnimplementedMLSClientServer) PSKProposal(context.Context, *PSKProposalRequest) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PSKProposal not implemented")
}
func (UnimplementedMLSClientServer) ReInitProposal(context.Context, *ReInitProposalRequest) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReInitProposal not implemented")
}
func (UnimplementedMLSClientServer) AppAckProposal(context.Context, *AppAckProposalRequest) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppAckProposal not implemented")
}
func (UnimplementedMLSClientServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedMLSClientServer) HandleCommit(context.Context, *HandleCommitRequest) (*HandleCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCommit not implemented")
}
func (UnimplementedMLSClientServer) HandleExternalCommit(context.Context, *HandleExternalCommitRequest) (*HandleExternalCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleExternalCommit not implemented")
}
func (UnimplementedMLSClientServer) mustEmbedUnimplementedMLSClientServer() {}

// UnsafeMLSClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MLSClientServer will
// result in compilation errors.
type UnsafeMLSClientServer interface {
	mustEmbedUnimplementedMLSClientServer()
}

func RegisterMLSClientServer(s grpc.ServiceRegistrar, srv MLSClientServer) {
	s.RegisterService(&MLSClient_ServiceDesc, srv)
}

func _MLSClient_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).Name(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_SupportedCiphersuites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportedCiphersuitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).SupportedCiphersuites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/SupportedCiphersuites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).SupportedCiphersuites(ctx, req.(*SupportedCiphersuitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_GenerateTestVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTestVectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).GenerateTestVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/GenerateTestVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).GenerateTestVector(ctx, req.(*GenerateTestVectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_VerifyTestVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTestVectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).VerifyTestVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/VerifyTestVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).VerifyTestVector(ctx, req.(*VerifyTestVectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_CreateKeyPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeyPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).CreateKeyPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/CreateKeyPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).CreateKeyPackage(ctx, req.(*CreateKeyPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/JoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).JoinGroup(ctx, req.(*JoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_ExternalJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExternalJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).ExternalJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/ExternalJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).ExternalJoin(ctx, req.(*ExternalJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_PublicGroupState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicGroupStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).PublicGroupState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/PublicGroupState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).PublicGroupState(ctx, req.(*PublicGroupStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_StateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).StateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/StateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).StateAuth(ctx, req.(*StateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).Export(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_Protect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).Protect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/Protect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).Protect(ctx, req.(*ProtectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_Unprotect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnprotectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).Unprotect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/Unprotect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).Unprotect(ctx, req.(*UnprotectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_StorePSK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePSKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).StorePSK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/StorePSK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).StorePSK(ctx, req.(*StorePSKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_AddProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).AddProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/AddProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).AddProposal(ctx, req.(*AddProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_UpdateProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).UpdateProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/UpdateProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).UpdateProposal(ctx, req.(*UpdateProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_RemoveProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).RemoveProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/RemoveProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).RemoveProposal(ctx, req.(*RemoveProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_PSKProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PSKProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).PSKProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/PSKProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).PSKProposal(ctx, req.(*PSKProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_ReInitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReInitProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).ReInitProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/ReInitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).ReInitProposal(ctx, req.(*ReInitProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_AppAckProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppAckProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).AppAckProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/AppAckProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).AppAckProposal(ctx, req.(*AppAckProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_HandleCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).HandleCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/HandleCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).HandleCommit(ctx, req.(*HandleCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MLSClient_HandleExternalCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleExternalCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLSClientServer).HandleExternalCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mls_client.MLSClient/HandleExternalCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLSClientServer).HandleExternalCommit(ctx, req.(*HandleExternalCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MLSClient_ServiceDesc is the grpc.ServiceDesc for MLSClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MLSClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mls_client.MLSClient",
	HandlerType: (*MLSClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _MLSClient_Name_Handler,
		},
		{
			MethodName: "SupportedCiphersuites",
			Handler:    _MLSClient_SupportedCiphersuites_Handler,
		},
		{
			MethodName: "GenerateTestVector",
			Handler:    _MLSClient_GenerateTestVector_Handler,
		},
		{
			MethodName: "VerifyTestVector",
			Handler:    _MLSClient_VerifyTestVector_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _MLSClient_CreateGroup_Handler,
		},
		{
			MethodName: "CreateKeyPackage",
			Handler:    _MLSClient_CreateKeyPackage_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _MLSClient_JoinGroup_Handler,
		},
		{
			MethodName: "ExternalJoin",
			Handler:    _MLSClient_ExternalJoin_Handler,
		},
		{
			MethodName: "PublicGroupState",
			Handler:    _MLSClient_PublicGroupState_Handler,
		},
		{
			MethodName: "StateAuth",
			Handler:    _MLSClient_StateAuth_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _MLSClient_Export_Handler,
		},
		{
			MethodName: "Protect",
			Handler:    _MLSClient_Protect_Handler,
		},
		{
			MethodName: "Unprotect",
			Handler:    _MLSClient_Unprotect_Handler,
		},
		{
			MethodName: "StorePSK",
			Handler:    _MLSClient_StorePSK_Handler,
		},
		{
			MethodName: "AddProposal",
			Handler:    _MLSClient_AddProposal_Handler,
		},
		{
			MethodName: "UpdateProposal",
			Handler:    _MLSClient_UpdateProposal_Handler,
		},
		{
			MethodName: "RemoveProposal",
			Handler:    _MLSClient_RemoveProposal_Handler,
		},
		{
			MethodName: "PSKProposal",
			Handler:    _MLSClient_PSKProposal_Handler,
		},
		{
			MethodName: "ReInitProposal",
			Handler:    _MLSClient_ReInitProposal_Handler,
		},
		{
			MethodName: "AppAckProposal",
			Handler:    _MLSClient_AppAckProposal_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _MLSClient_Commit_Handler,
		},
		{
			MethodName: "HandleCommit",
			Handler:    _MLSClient_HandleCommit_Handler,
		},
		{
			MethodName: "HandleExternalCommit",
			Handler:    _MLSClient_HandleExternalCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mls_client.proto",
}
