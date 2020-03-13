package appmanager

import (
	"context"
	"errors"
	"github.com/antihax/optional"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/internal/appmanager-api-client"
	"net/http"
)

var (
	// TODO: make custom types
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unathorized")
	ErrForbidden = errors.New("forbidden")
	ErrConflict     = errors.New("conflict")
	ErrNotFound     = errors.New("not found")
	ErrUnknown      = errors.New("unknown error")

	ErrAuthContext = errors.New("couldn't get authorized context")
)

func isResponseError(res *http.Response) bool {
	return res != nil && res.StatusCode >= 400
}

func formatResponseError(res *http.Response) error {
	switch res.StatusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 409:
		return ErrConflict
	default:
		return ErrUnknown
	}
}

type Client interface {
	DeploymentTargets() DeploymentTargetsService
	Events() EventsService
	Deployments() DeploymentsService
	Savepoints() SavepointsService
}

type client struct {
	apiClient                *appmanagerapi.APIClient
	authStore                *AuthStore // TODO: make an interface
	deploymentsService       DeploymentsService
	deploymentTargetsService DeploymentTargetsService
	eventsService            EventsService
	savepointsService        SavepointsService
}

func NewClient(apiClient *appmanagerapi.APIClient, authStore *AuthStore) Client {
	return &client{
		apiClient: apiClient,
		authStore: authStore,
	}
}

func (c *client) DeploymentTargets() DeploymentTargetsService {
	if c.deploymentTargetsService == nil {
		c.deploymentTargetsService = &deploymentTargetsService{client: c}
	}
	return c.deploymentTargetsService
}

func (c *client) Events() EventsService {
	if c.eventsService == nil {
		c.eventsService = &eventsService{client: c}
	}
	return c.eventsService
}

func (c *client) Deployments() DeploymentsService {
	if c.deploymentsService == nil {
		c.deploymentsService = &deploymentsService{client: c}
	}
	return c.deploymentsService
}

func (c *client) Savepoints() SavepointsService {
	if c.savepointsService == nil {
		c.savepointsService = &savepointsService{client: c}
	}
	return c.savepointsService
}

// Deployment targets

type DeploymentTargetsService interface {
	GetDeploymentTarget(ctx context.Context, namespaceName, name string) (*appmanagerapi.DeploymentTarget, error)
	CreateDeploymentTarget(ctx context.Context, namespaceName string, depTarget appmanagerapi.DeploymentTarget) (*appmanagerapi.DeploymentTarget, error)
	DeleteDeploymentTarget(ctx context.Context, namespaceName, name string) (*appmanagerapi.DeploymentTarget, error)
}

type deploymentTargetsService struct {
	client *client
}

func (s *deploymentTargetsService) GetDeploymentTarget(ctx context.Context, namespaceName, name string) (*appmanagerapi.DeploymentTarget, error) {
	depTarget, _, err := s.client.apiClient.DeploymentTargetResourceApi.GetDeploymentTargetUsingGET(ctx, name, namespaceName)
	return &depTarget, err
}

func (s *deploymentTargetsService) CreateDeploymentTarget(ctx context.Context, namespaceName string, depTarget appmanagerapi.DeploymentTarget) (*appmanagerapi.DeploymentTarget, error) {
	ctx, err := s.client.authStore.ContextForNamespace(context.Background(), namespaceName)
	if err != nil {
		return nil, ErrAuthContext
	}

	target, res, err := s.client.apiClient.DeploymentTargetResourceApi.
		CreateDeploymentTargetUsingPOST(ctx, depTarget, namespaceName)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &target, err
}

func (s *deploymentTargetsService) DeleteDeploymentTarget(ctx context.Context, namespaceName, name string) (*appmanagerapi.DeploymentTarget, error) {
	depTarget, res, err := s.client.apiClient.DeploymentTargetResourceApi.DeleteDeploymentTargetUsingDELETE(ctx, name, namespaceName)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &depTarget, err
}

// Events

type GetEventsOpts struct {
	DeploymentId optional.Interface
	JobId        optional.Interface
}

type EventsService interface {
	GetEvents(ctx context.Context, namespaceName string, opts *GetEventsOpts) (*appmanagerapi.ResourceListOfEvent, error)
}

type eventsService struct {
	client *client
}

func (s *eventsService) GetEvents(ctx context.Context, namespaceName string, opts *GetEventsOpts) (*appmanagerapi.ResourceListOfEvent, error) {
	ctx, err := s.client.authStore.ContextForNamespace(context.Background(), namespaceName)
	if err != nil {
		return nil, ErrAuthContext
	}
	eventsList, res, err := s.client.apiClient.EventResourceApi.GetEventsUsingGET(ctx, namespaceName, (*appmanagerapi.GetEventsUsingGETOpts)(opts))

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &eventsList, err
}

// Deployments

type DeploymentsService interface {
	GetDeployment(ctx context.Context, namespaceName, id string) (*appmanagerapi.Deployment, error)
	ListDeployments(ctx context.Context, namespaceName string) (*appmanagerapi.ResourceListOfDeployment, error)
	GetDeploymentByName(ctx context.Context, namespaceName, name string) (*appmanagerapi.Deployment, error)
	CreateDeployment(ctx context.Context, namespaceName string, dep appmanagerapi.Deployment) (*appmanagerapi.Deployment, error)
	UpdateDeployment(ctx context.Context, namespaceName, id string, dep appmanagerapi.Deployment) (*appmanagerapi.Deployment, error)
	DeleteDeployment(ctx context.Context, namespaceName, id string) (*appmanagerapi.Deployment, error)
}

type deploymentsService struct {
	client *client
}

func (s *deploymentsService) GetDeployment(ctx context.Context, namespaceName, id string) (*appmanagerapi.Deployment, error) {
	ctx, err := s.client.authStore.ContextForNamespace(context.Background(), namespaceName)
	if err != nil {
		return nil, ErrAuthContext
	}
	deployment, res, err := s.client.apiClient.DeploymentResourceApi.GetDeploymentUsingGET(ctx, id, namespaceName)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &deployment, err
}

func (s *deploymentsService) ListDeployments(ctx context.Context, namespaceName string) (*appmanagerapi.ResourceListOfDeployment, error) {
	depList, res, err := s.client.apiClient.DeploymentResourceApi.GetDeploymentsUsingGET(ctx, namespaceName, nil)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &depList, err
}

func (s *deploymentsService) GetDeploymentByName(ctx context.Context, namespaceName, name string) (*appmanagerapi.Deployment, error) {
	if len(namespaceName) == 0 || len(name) == 0 {
		return nil, errors.New("namespace and name must not be empty")
	}

	deploymentsList, err := s.ListDeployments(ctx, namespaceName)

	if err != nil {
		return nil, err
	}

	for _, deployment := range deploymentsList.Items {
		if deployment.Metadata.Name == name {
			return &deployment, nil
		}
	}

	return nil, ErrNotFound
}

func (s *deploymentsService) CreateDeployment(ctx context.Context, namespaceName string, dep appmanagerapi.Deployment) (*appmanagerapi.Deployment, error) {
	deployment, res, err := s.client.apiClient.DeploymentResourceApi.CreateDeploymentUsingPOST(ctx, dep, namespaceName)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &deployment, err
}

func (s *deploymentsService) UpdateDeployment(ctx context.Context, namespaceName, id string, dep appmanagerapi.Deployment) (*appmanagerapi.Deployment, error) {
	deployment, res, err := s.client.apiClient.DeploymentResourceApi.UpdateDeploymentUsingPATCH(ctx, dep, id, namespaceName)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &deployment, err
}

func (s *deploymentsService) DeleteDeployment(ctx context.Context, namespaceName, id string) (*appmanagerapi.Deployment, error) {
	deployment, res, err := s.client.apiClient.DeploymentResourceApi.DeleteDeploymentUsingDELETE(ctx, id, namespaceName)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &deployment, err
}

// Savepoints

type SavepointsService interface {
	GetSavepoint(ctx context.Context, namespaceName, id string) (*appmanagerapi.Savepoint, error)
	CreateSavepoint(ctx context.Context, namespaceName string, savepoint appmanagerapi.Savepoint) (*appmanagerapi.Savepoint, error)
}

type savepointsService struct {
	client *client
}

func (s savepointsService) GetSavepoint(ctx context.Context, namespaceName, id string) (*appmanagerapi.Savepoint, error) {
	savepoint, res, err := s.client.apiClient.SavepointResourceApi.GetSavepointUsingGET(ctx, namespaceName, id)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &savepoint, err
}

func (s savepointsService) CreateSavepoint(ctx context.Context, namespaceName string, savepoint appmanagerapi.Savepoint) (*appmanagerapi.Savepoint, error) {
	savepoint, res, err := s.client.apiClient.SavepointResourceApi.CreateSavepointUsingPOST(ctx, namespaceName, savepoint)

	if isResponseError(res) {
		return nil, formatResponseError(res)
	}

	return &savepoint, err
}
