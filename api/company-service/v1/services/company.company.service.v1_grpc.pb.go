// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.1
// - protoc             v5.26.1
// source: api/company-service/v1/services/company.company.service.v1.proto

package companyservicev1

import (
	context "context"
	fmt "fmt"

	client "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/client"
	def "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config/def"
	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/company-service/v1/resources"
	fx "go.uber.org/fx"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7
const _ = fx.Version

var _ = new(fmt.Stringer)
var _ = new(def.Server)
var _ = new(client.RegisterGRPCClientFactoryType)

const (
	CompanyV1_GetEmployeeInfoByLogin_FullMethodName     = "/api.company.service.companyservicev1.CompanyV1/GetEmployeeInfoByLogin"
	CompanyV1_GetEmployeeByUid_FullMethodName           = "/api.company.service.companyservicev1.CompanyV1/GetEmployeeByUid"
	CompanyV1_UpdateEmployee_FullMethodName             = "/api.company.service.companyservicev1.CompanyV1/UpdateEmployee"
	CompanyV1_CheckAndUpdatePrivilege_FullMethodName    = "/api.company.service.companyservicev1.CompanyV1/CheckAndUpdatePrivilege"
	CompanyV1_GetPrivilegeLimit_FullMethodName          = "/api.company.service.companyservicev1.CompanyV1/GetPrivilegeLimit"
	CompanyV1_GetPartCompanyList_FullMethodName         = "/api.company.service.companyservicev1.CompanyV1/GetPartCompanyList"
	CompanyV1_GetCompanyById_FullMethodName             = "/api.company.service.companyservicev1.CompanyV1/GetCompanyById"
	CompanyV1_GetUsersByIds_FullMethodName              = "/api.company.service.companyservicev1.CompanyV1/GetUsersByIds"
	CompanyV1_GetSupplierCompanyById_FullMethodName     = "/api.company.service.companyservicev1.CompanyV1/GetSupplierCompanyById"
	CompanyV1_CreateDesignerORGForPerson_FullMethodName = "/api.company.service.companyservicev1.CompanyV1/CreateDesignerORGForPerson"
	CompanyV1_GetAndCheckLastCompany_FullMethodName     = "/api.company.service.companyservicev1.CompanyV1/GetAndCheckLastCompany"
)

// CompanyV1Client is the client API for CompanyV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyV1Client interface {
	GetEmployeeInfoByLogin(ctx context.Context, in *resources.GetEmployeeInfoByLoginRequest, opts ...grpc.CallOption) (*resources.GetEmployeeInfoByLoginReply, error)
	GetEmployeeByUid(ctx context.Context, in *resources.GetEmployeeByUidRequest, opts ...grpc.CallOption) (*resources.GetEmployeeByUidReply, error)
	UpdateEmployee(ctx context.Context, in *resources.UpdateEmployeeRequest, opts ...grpc.CallOption) (*resources.UpdateEmployeeReply, error)
	CheckAndUpdatePrivilege(ctx context.Context, in *resources.CheckAndUpdatePrivilegeRequest, opts ...grpc.CallOption) (*resources.CheckAndUpdatePrivilegeReply, error)
	GetPrivilegeLimit(ctx context.Context, in *resources.GetPrivilegeLimitRequest, opts ...grpc.CallOption) (*resources.GetPrivilegeLimitReply, error)
	GetPartCompanyList(ctx context.Context, in *resources.GetPartCompanyListRequest, opts ...grpc.CallOption) (*resources.GetPartCompanyListReply, error)
	GetCompanyById(ctx context.Context, in *resources.GetCompanyByIdRequest, opts ...grpc.CallOption) (*resources.GetCompanyByIdReply, error)
	GetUsersByIds(ctx context.Context, in *resources.GetUsersByIdsRequest, opts ...grpc.CallOption) (*resources.GetUsersByIdsReply, error)
	GetSupplierCompanyById(ctx context.Context, in *resources.GetSupplierCompanyByIdRequest, opts ...grpc.CallOption) (*resources.GetSupplierCompanyByIdReply, error)
	CreateDesignerORGForPerson(ctx context.Context, in *resources.CreateDesignerORGForPersonRequest, opts ...grpc.CallOption) (*resources.CreateDesignerORGForPersonReply, error)
	GetAndCheckLastCompany(ctx context.Context, in *resources.GetAndCheckLastCompanyRequest, opts ...grpc.CallOption) (*resources.GetAndCheckLastCompanyReply, error)
	RegisterNameForDiscover() string
}

type companyV1Client struct {
	cc grpc.ClientConnInterface
}

func (c *companyV1Client) RegisterNameForDiscover() string {
	return "/company-service/v1"
}

func newCompanyV1Client(cc grpc.ClientConnInterface) CompanyV1Client {
	return &companyV1Client{cc}
}

func registerCompanyV1ClientGRPCNameProvider() []string {
	return []string{"/company-service/v1", "grpc"}
}

// RegisterCompanyV1ClientGRPCProvider is the provider for injection framework
// creator is the factory function which use to create the CompanyV1Client instance/implement
// the creator function receive dependency provided by fx to create ClientInterface,
// and returns the new dependency can use by others functions
func RegisterCompanyV1ClientGRPCProvider(creator interface{}) []interface{} {
	return []interface{}{
		fx.Annotate(
			newCompanyV1Client,
			fx.As(new(CompanyV1Client)),
			fx.ParamTags(`name:"/company-service/v1/grpc/companyV1"`),
		),
		fx.Annotate(
			creator,
			fx.As(new(grpc.ClientConnInterface)),
			fx.ParamTags(`name:"/company-service/v1/grpc/name/companyV1"`),
			fx.ResultTags(`name:"/company-service/v1/grpc/companyV1"`),
		),
		fx.Annotate(
			registerCompanyV1ClientGRPCNameProvider,
			fx.ResultTags(`name:"/company-service/v1/grpc/name/companyV1"`),
		),
	}
}

type CompanyV1ClientGRPCFactory interface {
	New(conf *def.Server) (CompanyV1Client, error)
}

type companyV1ClientGRPCFactoryImpl struct {
	factory client.RegisterGRPCClientFactoryType
}

func (p *companyV1ClientGRPCFactoryImpl) New(conf *def.Server) (CompanyV1Client, error) {
	cc, err := p.factory(conf)
	if err != nil {
		return nil, fmt.Errorf("create CompanyV1Client failed cause %s", err)
	}
	return &companyV1Client{cc: cc}, nil
}

func RegisterCompanyV1ClientGRPCFactoryProvider(factory client.RegisterGRPCClientFactoryType) CompanyV1ClientGRPCFactory {
	return &companyV1ClientGRPCFactoryImpl{factory: factory}
}

func (c *companyV1Client) GetEmployeeInfoByLogin(ctx context.Context, in *resources.GetEmployeeInfoByLoginRequest, opts ...grpc.CallOption) (*resources.GetEmployeeInfoByLoginReply, error) {
	out := new(resources.GetEmployeeInfoByLoginReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetEmployeeInfoByLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetEmployeeByUid(ctx context.Context, in *resources.GetEmployeeByUidRequest, opts ...grpc.CallOption) (*resources.GetEmployeeByUidReply, error) {
	out := new(resources.GetEmployeeByUidReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetEmployeeByUid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) UpdateEmployee(ctx context.Context, in *resources.UpdateEmployeeRequest, opts ...grpc.CallOption) (*resources.UpdateEmployeeReply, error) {
	out := new(resources.UpdateEmployeeReply)
	err := c.cc.Invoke(ctx, CompanyV1_UpdateEmployee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) CheckAndUpdatePrivilege(ctx context.Context, in *resources.CheckAndUpdatePrivilegeRequest, opts ...grpc.CallOption) (*resources.CheckAndUpdatePrivilegeReply, error) {
	out := new(resources.CheckAndUpdatePrivilegeReply)
	err := c.cc.Invoke(ctx, CompanyV1_CheckAndUpdatePrivilege_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetPrivilegeLimit(ctx context.Context, in *resources.GetPrivilegeLimitRequest, opts ...grpc.CallOption) (*resources.GetPrivilegeLimitReply, error) {
	out := new(resources.GetPrivilegeLimitReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetPrivilegeLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetPartCompanyList(ctx context.Context, in *resources.GetPartCompanyListRequest, opts ...grpc.CallOption) (*resources.GetPartCompanyListReply, error) {
	out := new(resources.GetPartCompanyListReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetPartCompanyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetCompanyById(ctx context.Context, in *resources.GetCompanyByIdRequest, opts ...grpc.CallOption) (*resources.GetCompanyByIdReply, error) {
	out := new(resources.GetCompanyByIdReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetCompanyById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetUsersByIds(ctx context.Context, in *resources.GetUsersByIdsRequest, opts ...grpc.CallOption) (*resources.GetUsersByIdsReply, error) {
	out := new(resources.GetUsersByIdsReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetUsersByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetSupplierCompanyById(ctx context.Context, in *resources.GetSupplierCompanyByIdRequest, opts ...grpc.CallOption) (*resources.GetSupplierCompanyByIdReply, error) {
	out := new(resources.GetSupplierCompanyByIdReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetSupplierCompanyById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) CreateDesignerORGForPerson(ctx context.Context, in *resources.CreateDesignerORGForPersonRequest, opts ...grpc.CallOption) (*resources.CreateDesignerORGForPersonReply, error) {
	out := new(resources.CreateDesignerORGForPersonReply)
	err := c.cc.Invoke(ctx, CompanyV1_CreateDesignerORGForPerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyV1Client) GetAndCheckLastCompany(ctx context.Context, in *resources.GetAndCheckLastCompanyRequest, opts ...grpc.CallOption) (*resources.GetAndCheckLastCompanyReply, error) {
	out := new(resources.GetAndCheckLastCompanyReply)
	err := c.cc.Invoke(ctx, CompanyV1_GetAndCheckLastCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyV1Server is the server API for CompanyV1 service.
// All implementations must embed UnimplementedCompanyV1Server
// for forward compatibility
type CompanyV1Server interface {
	GetEmployeeInfoByLogin(context.Context, *resources.GetEmployeeInfoByLoginRequest) (*resources.GetEmployeeInfoByLoginReply, error)
	GetEmployeeByUid(context.Context, *resources.GetEmployeeByUidRequest) (*resources.GetEmployeeByUidReply, error)
	UpdateEmployee(context.Context, *resources.UpdateEmployeeRequest) (*resources.UpdateEmployeeReply, error)
	CheckAndUpdatePrivilege(context.Context, *resources.CheckAndUpdatePrivilegeRequest) (*resources.CheckAndUpdatePrivilegeReply, error)
	GetPrivilegeLimit(context.Context, *resources.GetPrivilegeLimitRequest) (*resources.GetPrivilegeLimitReply, error)
	GetPartCompanyList(context.Context, *resources.GetPartCompanyListRequest) (*resources.GetPartCompanyListReply, error)
	GetCompanyById(context.Context, *resources.GetCompanyByIdRequest) (*resources.GetCompanyByIdReply, error)
	GetUsersByIds(context.Context, *resources.GetUsersByIdsRequest) (*resources.GetUsersByIdsReply, error)
	GetSupplierCompanyById(context.Context, *resources.GetSupplierCompanyByIdRequest) (*resources.GetSupplierCompanyByIdReply, error)
	CreateDesignerORGForPerson(context.Context, *resources.CreateDesignerORGForPersonRequest) (*resources.CreateDesignerORGForPersonReply, error)
	GetAndCheckLastCompany(context.Context, *resources.GetAndCheckLastCompanyRequest) (*resources.GetAndCheckLastCompanyReply, error)
	mustEmbedUnimplementedCompanyV1Server()
}

// Generate Injection
type registerCompanyV1ServerGRPCResult struct{}

func (*registerCompanyV1ServerGRPCResult) String() string {
	return "CompanyV1ServerGRPCServer"
}

func RegisterCompanyV1ServerGRPCProvider(newer interface{}) []interface{} {
	return []interface{}{
		// For provide dependency
		fx.Annotate(
			newer,
			fx.As(new(CompanyV1Server)),
		),
		// For create instance
		fx.Annotate(
			registerCompanyV1ServerProviderImpl,
			fx.As(new(fmt.Stringer)),
			fx.ResultTags(`group:"grpc_register"`),
		),
	}
}

// registerCompanyV1ServerProviderImpl use to trigger register
func registerCompanyV1ServerProviderImpl(s grpc.ServiceRegistrar, srv CompanyV1Server) *registerCompanyV1ServerGRPCResult {
	registerCompanyV1Server(s, srv)
	return &registerCompanyV1ServerGRPCResult{}
}

// UnimplementedCompanyV1Server must be embedded to have forward compatible implementations.
type UnimplementedCompanyV1Server struct {
}

func (UnimplementedCompanyV1Server) GetEmployeeInfoByLogin(context.Context, *resources.GetEmployeeInfoByLoginRequest) (*resources.GetEmployeeInfoByLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeInfoByLogin not implemented")
}
func (UnimplementedCompanyV1Server) GetEmployeeByUid(context.Context, *resources.GetEmployeeByUidRequest) (*resources.GetEmployeeByUidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeByUid not implemented")
}
func (UnimplementedCompanyV1Server) UpdateEmployee(context.Context, *resources.UpdateEmployeeRequest) (*resources.UpdateEmployeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedCompanyV1Server) CheckAndUpdatePrivilege(context.Context, *resources.CheckAndUpdatePrivilegeRequest) (*resources.CheckAndUpdatePrivilegeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndUpdatePrivilege not implemented")
}
func (UnimplementedCompanyV1Server) GetPrivilegeLimit(context.Context, *resources.GetPrivilegeLimitRequest) (*resources.GetPrivilegeLimitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivilegeLimit not implemented")
}
func (UnimplementedCompanyV1Server) GetPartCompanyList(context.Context, *resources.GetPartCompanyListRequest) (*resources.GetPartCompanyListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartCompanyList not implemented")
}
func (UnimplementedCompanyV1Server) GetCompanyById(context.Context, *resources.GetCompanyByIdRequest) (*resources.GetCompanyByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyById not implemented")
}
func (UnimplementedCompanyV1Server) GetUsersByIds(context.Context, *resources.GetUsersByIdsRequest) (*resources.GetUsersByIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByIds not implemented")
}
func (UnimplementedCompanyV1Server) GetSupplierCompanyById(context.Context, *resources.GetSupplierCompanyByIdRequest) (*resources.GetSupplierCompanyByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupplierCompanyById not implemented")
}
func (UnimplementedCompanyV1Server) CreateDesignerORGForPerson(context.Context, *resources.CreateDesignerORGForPersonRequest) (*resources.CreateDesignerORGForPersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDesignerORGForPerson not implemented")
}
func (UnimplementedCompanyV1Server) GetAndCheckLastCompany(context.Context, *resources.GetAndCheckLastCompanyRequest) (*resources.GetAndCheckLastCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAndCheckLastCompany not implemented")
}
func (UnimplementedCompanyV1Server) mustEmbedUnimplementedCompanyV1Server() {}

// UnsafeCompanyV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyV1Server will
// result in compilation errors.
type UnsafeCompanyV1Server interface {
	mustEmbedUnimplementedCompanyV1Server()
}

func registerCompanyV1Server(s grpc.ServiceRegistrar, srv CompanyV1Server) {
	s.RegisterService(&CompanyV1_ServiceDesc, srv)
}

func _CompanyV1_GetEmployeeInfoByLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetEmployeeInfoByLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetEmployeeInfoByLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetEmployeeInfoByLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetEmployeeInfoByLogin(ctx, req.(*resources.GetEmployeeInfoByLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetEmployeeByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetEmployeeByUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetEmployeeByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetEmployeeByUid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetEmployeeByUid(ctx, req.(*resources.GetEmployeeByUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_UpdateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).UpdateEmployee(ctx, req.(*resources.UpdateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_CheckAndUpdatePrivilege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.CheckAndUpdatePrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).CheckAndUpdatePrivilege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_CheckAndUpdatePrivilege_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).CheckAndUpdatePrivilege(ctx, req.(*resources.CheckAndUpdatePrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetPrivilegeLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetPrivilegeLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetPrivilegeLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetPrivilegeLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetPrivilegeLimit(ctx, req.(*resources.GetPrivilegeLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetPartCompanyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetPartCompanyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetPartCompanyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetPartCompanyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetPartCompanyList(ctx, req.(*resources.GetPartCompanyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetCompanyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetCompanyByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetCompanyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetCompanyById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetCompanyById(ctx, req.(*resources.GetCompanyByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetUsersByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetUsersByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetUsersByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetUsersByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetUsersByIds(ctx, req.(*resources.GetUsersByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetSupplierCompanyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetSupplierCompanyByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetSupplierCompanyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetSupplierCompanyById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetSupplierCompanyById(ctx, req.(*resources.GetSupplierCompanyByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_CreateDesignerORGForPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.CreateDesignerORGForPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).CreateDesignerORGForPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_CreateDesignerORGForPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).CreateDesignerORGForPerson(ctx, req.(*resources.CreateDesignerORGForPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyV1_GetAndCheckLastCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetAndCheckLastCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyV1Server).GetAndCheckLastCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyV1_GetAndCheckLastCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyV1Server).GetAndCheckLastCompany(ctx, req.(*resources.GetAndCheckLastCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyV1_ServiceDesc is the grpc.ServiceDesc for CompanyV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.company.service.companyservicev1.CompanyV1",
	HandlerType: (*CompanyV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployeeInfoByLogin",
			Handler:    _CompanyV1_GetEmployeeInfoByLogin_Handler,
		},
		{
			MethodName: "GetEmployeeByUid",
			Handler:    _CompanyV1_GetEmployeeByUid_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _CompanyV1_UpdateEmployee_Handler,
		},
		{
			MethodName: "CheckAndUpdatePrivilege",
			Handler:    _CompanyV1_CheckAndUpdatePrivilege_Handler,
		},
		{
			MethodName: "GetPrivilegeLimit",
			Handler:    _CompanyV1_GetPrivilegeLimit_Handler,
		},
		{
			MethodName: "GetPartCompanyList",
			Handler:    _CompanyV1_GetPartCompanyList_Handler,
		},
		{
			MethodName: "GetCompanyById",
			Handler:    _CompanyV1_GetCompanyById_Handler,
		},
		{
			MethodName: "GetUsersByIds",
			Handler:    _CompanyV1_GetUsersByIds_Handler,
		},
		{
			MethodName: "GetSupplierCompanyById",
			Handler:    _CompanyV1_GetSupplierCompanyById_Handler,
		},
		{
			MethodName: "CreateDesignerORGForPerson",
			Handler:    _CompanyV1_CreateDesignerORGForPerson_Handler,
		},
		{
			MethodName: "GetAndCheckLastCompany",
			Handler:    _CompanyV1_GetAndCheckLastCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/company-service/v1/services/company.company.service.v1.proto",
}