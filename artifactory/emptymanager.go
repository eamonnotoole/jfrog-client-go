package artifactory

import (
	"io"

	"github.com/jfrog/jfrog-client-go/artifactory/buildinfo"
	rthttpclient "github.com/jfrog/jfrog-client-go/artifactory/httpclient"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	_go "github.com/jfrog/jfrog-client-go/artifactory/services/go"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/config"
	"github.com/jfrog/jfrog-client-go/utils/io/content"
)

type ArtifactoryServicesManager interface {
	CreateLocalRepository() *services.LocalRepositoryService
	CreateRemoteRepository() *services.RemoteRepositoryService
	CreateVirtualRepository() *services.VirtualRepositoryService
	UpdateLocalRepository() *services.LocalRepositoryService
	UpdateRemoteRepository() *services.RemoteRepositoryService
	UpdateVirtualRepository() *services.VirtualRepositoryService
	DeleteRepository(repoKey string) error
	GetRepository(repoKey string) (*services.RepositoryDetails, error)
	CreatePermissionTarget(params services.PermissionTargetParams) error
	UpdatePermissionTarget(params services.PermissionTargetParams) error
	DeletePermissionTarget(permissionTargetName string) error
	PublishBuildInfo(build *buildinfo.BuildInfo) error
	DistributeBuild(params services.BuildDistributionParams) error
	PromoteBuild(params services.PromotionParams) error
	DiscardBuilds(params services.DiscardBuildsParams) error
	XrayScanBuild(params services.XrayScanParams) ([]byte, error)
	GetPathsToDelete(params services.DeleteParams) (*content.ContentReader, error)
	DeleteFiles(reader *content.ContentReader) (int, error)
	ReadRemoteFile(readPath string) (io.ReadCloser, error)
	initDownloadService() *services.DownloadService
	DownloadFiles(params ...services.DownloadParams) (totalDownloaded, totalExpected int, err error)
	DownloadFilesWithResultReader(params ...services.DownloadParams) (resultReader *content.ContentReader, totalDownloaded, totalExpected int, err error)
	GetUnreferencedGitLfsFiles(params services.GitLfsCleanParams) (*content.ContentReader, error)
	SearchFiles(params services.SearchParams) (*content.ContentReader, error)
	Aql(aql string) (io.ReadCloser, error)
	SetProps(params services.PropsParams) (int, error)
	DeleteProps(params services.PropsParams) (int, error)
	UploadFiles(params ...services.UploadParams) (artifactsFileInfo []utils.FileInfo, totalUploaded, totalFailed int, err error)
	Copy(params services.MoveCopyParams) (successCount, failedCount int, err error)
	Move(params services.MoveCopyParams) (successCount, failedCount int, err error)
	PublishGoProject(params _go.GoParams) error
	Ping() ([]byte, error)
	GetConfig() config.Config
	GetBuildInfo(params services.BuildInfoParams) (*buildinfo.BuildInfo, error)
	CreateToken(params services.CreateTokenParams) (services.CreateTokenResponseData, error)
	GetTokens() (services.GetTokensResponseData, error)
	RefreshToken(params services.RefreshTokenParams) (services.CreateTokenResponseData, error)
	RevokeToken(params services.RevokeTokenParams) (string, error)
	CreateReplication(params services.CreateReplicationParams) error
	UpdateReplication(params services.UpdateReplicationParams) error
	DeleteReplication(repoKey string) error
	GetReplication(repoKey string) ([]utils.ReplicationParams, error)
	GetVersion() (string, error)
	GetServiceId() (string, error)
	PromoteDocker(params services.DockerPromoteParams) error
	Client() *rthttpclient.ArtifactoryHttpClient
}

// By using this struct, you have the option of overriding only some of the ArtifactoryServicesManager
// interface's methods, but still implement this interface.
// This comes in very handy for tests.
type EmptyArtifactoryServicesManager struct {
}

func (esm *EmptyArtifactoryServicesManager) CreateLocalRepository() *services.LocalRepositoryService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) CreateRemoteRepository() *services.RemoteRepositoryService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) CreateVirtualRepository() *services.VirtualRepositoryService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) UpdateLocalRepository() *services.LocalRepositoryService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) UpdateRemoteRepository() *services.RemoteRepositoryService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) UpdateVirtualRepository() *services.VirtualRepositoryService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DeleteRepository(repoKey string) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetRepository(repoKey string) (*services.RepositoryDetails, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) CreatePermissionTarget(params services.PermissionTargetParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) UpdatePermissionTarget(params services.PermissionTargetParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DeletePermissionTarget(permissionTargetName string) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) PublishBuildInfo(build *buildinfo.BuildInfo) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DistributeBuild(params services.BuildDistributionParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) PromoteBuild(params services.PromotionParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DiscardBuilds(params services.DiscardBuildsParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) XrayScanBuild(params services.XrayScanParams) ([]byte, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetPathsToDelete(params services.DeleteParams) (*content.ContentReader, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DeleteFiles(reader *content.ContentReader) (int, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) ReadRemoteFile(readPath string) (io.ReadCloser, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) initDownloadService() *services.DownloadService {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DownloadFiles(params ...services.DownloadParams) (totalDownloaded, totalExpected int, err error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DownloadFilesWithResultReader(params ...services.DownloadParams) (resultReader *content.ContentReader, totalDownloaded, totalExpected int, err error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetUnreferencedGitLfsFiles(params services.GitLfsCleanParams) (*content.ContentReader, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) SearchFiles(params services.SearchParams) (*content.ContentReader, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) Aql(aql string) (io.ReadCloser, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) SetProps(params services.PropsParams) (int, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DeleteProps(params services.PropsParams) (int, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) UploadFiles(params ...services.UploadParams) (artifactsFileInfo []utils.FileInfo, totalUploaded, totalFailed int, err error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) Copy(params services.MoveCopyParams) (successCount, failedCount int, err error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) Move(params services.MoveCopyParams) (successCount, failedCount int, err error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) PublishGoProject(params _go.GoParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) Ping() ([]byte, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetConfig() config.Config {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetBuildInfo(params services.BuildInfoParams) (*buildinfo.BuildInfo, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) CreateToken(params services.CreateTokenParams) (services.CreateTokenResponseData, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetTokens() (services.GetTokensResponseData, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) RefreshToken(params services.RefreshTokenParams) (services.CreateTokenResponseData, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) RevokeToken(params services.RevokeTokenParams) (string, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) CreateReplication(params services.CreateReplicationParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) UpdateReplication(params services.UpdateReplicationParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) DeleteReplication(repoKey string) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetReplication(repoKey string) ([]utils.ReplicationParams, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetVersion() (string, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) GetServiceId() (string, error) {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) PromoteDocker(params services.DockerPromoteParams) error {
	panic("Failed: Method is not implemented")
}

func (esm *EmptyArtifactoryServicesManager) Client() *rthttpclient.ArtifactoryHttpClient {
	panic("Failed: Method is not implemented")
}
