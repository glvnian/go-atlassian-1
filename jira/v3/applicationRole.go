package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

// ApplicationRoleService represents the Jira Cloud application roles
// Use it to get details of an application role or all application roles.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles
type ApplicationRoleService struct{ client *Client }

// Gets returns all application roles
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-all-application-roles
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-get
func (a *ApplicationRoleService) Gets(ctx context.Context) (result []*models.ApplicationRoleScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/applicationrole"

	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns an application role, this func needs the following parameters:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
func (a *ApplicationRoleService) Get(ctx context.Context, key string) (result *models.ApplicationRoleScheme, response *ResponseScheme, err error) {

	if key == "" {
		return nil, nil, models.ErrNoApplicationRoleError
	}

	var endpoint = fmt.Sprintf("rest/api/3/applicationrole/%v", key)

	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
