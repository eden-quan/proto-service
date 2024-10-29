package filesystem_service

import (
	"bytes"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	filesystemv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/filesystem-service/v1/resources"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/api/proto-common/v3/errors"

	"io"
	//iamService "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service"
	fileservice "gitlab.lainuoniao.cn/eden-quan/proto-service/api/filesystem-service/v1/services"
	"net/http"
)

type FileSystem struct {
	//IAM iam_service.IAMUserServiceClient
	FSClient   fileservice.FileSystemServiceV1Client
	FSIOClient fileservice.FileIOApiV1Client
	Log        *log.Helper
}

func NewFileSystem(
	client fileservice.FileSystemServiceV1Client,
	fsIOClient fileservice.FileIOApiV1Client,
	logger log.Logger) *FileSystem {
	return &FileSystem{
		FSClient:   client,
		FSIOClient: fsIOClient,
		Log:        log.NewHelper(logger),
	}
}

func (f *FileSystem) DeleteObj(ctx context.Context, path string) error {
	resp, err := f.FSClient.ObjectRemove(ctx, &filesystemv1.FSCommonRequest{Path: path})
	if err != nil {
		return errors.ERROR_INTERNAL.FromError(err)
	}
	if resp.Code == 0 {
		return nil
	}

	return errors.ERROR_INTERNAL.ToError("remove object with non zero code")
}

type DownloadUrlRequest struct {
	Md5      string `json:"md5"`
	Path     string `json:"path"`
	Filename string `json:"filename"`

	// DownloadUrl 为转换完成的结果
	DownloadUrl string `json:"download_url"`
	UploadUrl   string `json:"upload_url"`
}

func (f *FileSystem) GetDownloadUrl(ctx context.Context, request *DownloadUrlRequest) string {
	param := &filesystemv1.IOFSFileUrlRequest{
		Path:     request.Path,
		Md5:      request.Md5,
		Filename: request.Filename,
	}

	resp, err := f.FSIOClient.GetDownloadUrl(ctx, param)
	if err != nil {
		return ""
	}

	return resp.GetUrl()
}

func (f *FileSystem) GetUploadUrl(ctx context.Context, request *DownloadUrlRequest) string {
	param := &filesystemv1.IOFSFileUrlRequest{
		Path: request.Path,
		Md5:  request.Md5,
	}

	resp, err := f.FSIOClient.GetUploadUrl(ctx, param)
	if err != nil {
		return ""
	}

	return resp.GetUrl()
}

func (f *FileSystem) UploadData(_ context.Context, uploadUrl, data string) error {
	req, err := http.NewRequest(http.MethodPut, uploadUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		f.Log.Errorf("call upload url %s with %d bytes failed with status code %d", uploadUrl, len(body), resp.StatusCode)
		return errors.ERROR_INTERNAL.ToError("upload data failed")
	}

	return nil
}

// NewFileIOGRPCProvider 简化的 FileIO 注入
func NewFileIOGRPCProvider(creator interface{}) []interface{} {
	return fileservice.RegisterFileIOApiV1ClientGRPCProvider(creator)
}

// NewFileSystemGRPCProvider 简化的 FileSystem 注入
func NewFileSystemGRPCProvider(creator interface{}) []interface{} {
	return fileservice.RegisterFileSystemServiceV1ClientGRPCProvider(creator)
}

// NewFileRegistryGRPCProvider 简化的 FileRegistrySystem 注入
func NewFileRegistryGRPCProvider(creator interface{}) []interface{} {
	return fileservice.RegisterFileRegistrySystemServiceV1ClientGRPCProvider(creator)
}

// InjectFileSystemService 注入封装后的 FileSystem 等依赖到 injector 中
// 该封装同时包含了 FileSystem 及 FileIO
func InjectFileSystemService(inj *injection.Injector) {
	inj.InjectGRPCClient(NewFileIOGRPCProvider)
	inj.InjectGRPCClient(NewFileSystemGRPCProvider)
	inj.Inject(NewFileSystem)
}
