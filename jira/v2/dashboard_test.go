package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestDashboardService_Gets_V2(t *testing.T) {

	testCases := []struct {
		name               string
		startAt            int
		maxResults         int
		filter             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetDashboardsWhenIsCorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetDashboardsWhenIsEndpointIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetDashboardsWhenTheStatusCodeIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheRequestMethodIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheFilterQueryIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "xxxx",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheQueryParametersAreIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=11111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheContextIsNil",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheResponseBodyLengthIsZero",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheResponseBodyHasADifferentFormat",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			service := &DashboardService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.startAt, testCase.maxResults, testCase.filter)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode()))
				assert.Equal(t, testCase.endpoint, fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode()))

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestDashboardService_Copy_V2(t *testing.T) {

	testCases := []struct {
		name               string
		dashboardID        string
		payload            *models.DashboardPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:        "CopyDashboardWhenTheParametersAreCorrect",
			dashboardID: "10001",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking #2 copy",
				Description: "Description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "CopyDashboardWhenTheDashboardIsNotProvided",
			dashboardID: "",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking #2 copy",
				Description: "Description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CopyDashboardWhenThePayloadIsNotProvided",
			dashboardID:        "10001",
			payload:            nil,
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CopyDashboardWhenTheRequestMethodIsIncorrect",
			dashboardID: "10001",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking #2 copy",
				Description: "Description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/dashboard/10001/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CopyDashboardWhenTheStatusCodeIsIncorrect",
			dashboardID: "10001",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking #2 copy",
				Description: "Description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "CopyDashboardWhenTheContextIsNotProvided",
			dashboardID: "10001",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking #2 copy",
				Description: "Description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001/copy",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &DashboardService{client: mockClient}
			gotResult, gotResponse, err := i.Copy(testCase.context, testCase.dashboardID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Log("--------------------------------")
				t.Logf("New copy Dashboard ID: %v", gotResult.ID)
				t.Logf("New copy Dashboard Name: %v", gotResult.Name)
				t.Logf("New copy Dashboard Self: %v", gotResult.Self)
				t.Log("-------------------------------- \n")

			}
		})

	}

}

func TestDashboardService_Create_V2(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.DashboardPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateDashboardWhenTheParametersAreCorrect",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking 2",
				Description: "description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/create-dashboard.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateDashboardWhenThePayloadIsNotProvided",
			payload:            nil,
			mockFile:           "../v3/mocks/create-dashboard.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CreateDashboardWhenTheContextIsNotProvided",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking 2",
				Description: "description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/create-dashboard.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CreateDashboardWhenTheRequestMethodIsIncorrect",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking 2",
				Description: "description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/create-dashboard.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/dashboard",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CreateDashboardWhenTheStatusCodeIsIncorrect",
			payload: &models.DashboardPayloadScheme{
				Name:        "Team Tracking 2",
				Description: "description sample",
				SharePermissions: []*models.SharePermissionScheme{
					{
						Type: "project",
						Project: &models.ProjectScheme{
							ID: "10000",
						},
						Role:  nil,
						Group: nil,
					},
					{
						Type:  "group",
						Group: &models.GroupScheme{Name: "jira-administrators"},
					},
				},
			},
			mockFile:           "../v3/mocks/create-dashboard.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &DashboardService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Log("--------------------------------")
				t.Logf("New copy Dashboard ID: %v", gotResult.ID)
				t.Logf("New copy Dashboard Name: %v", gotResult.Name)
				t.Logf("New copy Dashboard Self: %v", gotResult.Self)
				t.Log("-------------------------------- \n")

			}
		})

	}

}

func TestDashboardService_Delete_V2(t *testing.T) {

	testCases := []struct {
		name               string
		dashboardID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteDashboardWhenTheParametersAreCorrect",
			dashboardID:        "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteDashboardWhenTheDashboardIDIsNotSet",
			dashboardID:        "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteDashboardWhenTheRequestMethodIsIncorrect",
			dashboardID:        "10001",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteDashboardWhenTheStatusCodeIsIncorrect",
			dashboardID:        "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteDashboardWhenTheContextIsNil",
			dashboardID:        "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteDashboardWhenTheEndpointIsEmpty",
			dashboardID:        "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &DashboardService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.dashboardID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

			}
		})

	}

}

func TestDashboardService_Get_V2(t *testing.T) {

	testCases := []struct {
		name               string
		dashboardID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetDashboardWhenTheParametersAreCorrect",
			dashboardID:        "10001",
			mockFile:           "../v3/mocks/get-dashboard-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetDashboardWhenTheDashboardIsNotSet",
			dashboardID:        "",
			mockFile:           "../v3/mocks/get-dashboard-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetDashboardWhenTheRequestMethodIsIncorrect",
			dashboardID:        "10001",
			mockFile:           "../v3/mocks/get-dashboard-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetDashboardWhenTheStatusCodeIsIncorrect",
			dashboardID:        "10001",
			mockFile:           "../v3/mocks/get-dashboard-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetDashboardWhenTheContextIsNil",
			dashboardID:        "10001",
			mockFile:           "../v3/mocks/get-dashboard-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetDashboardWhenTheEndpointIsEmpty",
			dashboardID:        "10001",
			mockFile:           "../v3/mocks/get-dashboard-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardWhenTheResponseBodyHasADifferentFormat",
			dashboardID:        "10001",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &DashboardService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.dashboardID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Log("--------------------------------")
				t.Logf("New copy Dashboard ID: %v", gotResult.ID)
				t.Logf("New copy Dashboard Name: %v", gotResult.Name)
				t.Logf("New copy Dashboard Self: %v", gotResult.Self)
				t.Log("-------------------------------- \n")

			}
		})

	}

}

func TestDashboardService_Search_V2(t *testing.T) {

	testCases := []struct {
		name                string
		opts                *models.DashboardSearchOptionsScheme
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name: "SearchDashboardsWhenTheParametersAreCorrect",
			opts: &models.DashboardSearchOptionsScheme{
				OwnerAccountID:      "as48ashashash4hsahashdahsd",
				DashboardName:       "Bug",
				GroupPermissionName: "administrators",
				OrderBy:             "description",
				Expand:              []string{"description", "favourite", "sharePermissions"},
			},
			startAt:        0,
			maxResults:     50,
			mockFile:       "../v3/mocks/search-dashboards.json",
			wantHTTPMethod: http.MethodGet,
			endpoint: "/rest/api/2/dashboard/search?accountId=as48ashashash4hsahashdahsd&dashboardName=as48as" +
				"hashash4hsahashdahsd&expand=description%2Cfavourite%2CsharePermissions&groupname=as48ashashash4hsahashdahsd&maxResults=50&orderBy=as48ashashash4hsahashdahsd&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "SearchDashboardsWhenTheOptionsAreNil",
			opts:               nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/search-dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/search?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "SearchDashboardsWhenTheRequestMethodIsIncorrect",
			opts: &models.DashboardSearchOptionsScheme{
				DashboardName:       "Bug",
				GroupPermissionName: "administrators",
				OrderBy:             "description",
				Expand:              []string{"description", "favourite", "sharePermissions"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/search-dashboards.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/dashboard/search?dashboardName=&expand=description%2Cfavourite%2CsharePermissions&groupname=&maxResults=50&orderBy=&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchDashboardsWhenTheStatusCodeIsIncorrect",
			opts: &models.DashboardSearchOptionsScheme{
				DashboardName:       "Bug",
				GroupPermissionName: "administrators",
				OrderBy:             "description",
				Expand:              []string{"description", "favourite", "sharePermissions"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/search-dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/search?dashboardName=&expand=description%2Cfavourite%2CsharePermissions&groupname=&maxResults=50&orderBy=&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "SearchDashboardsWhenTheContextIsNil",
			opts: &models.DashboardSearchOptionsScheme{
				DashboardName:       "Bug",
				GroupPermissionName: "administrators",
				OrderBy:             "description",
				Expand:              []string{"description", "favourite", "sharePermissions"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/search-dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/search?dashboardName=&expand=description%2Cfavourite%2CsharePermissions&groupname=&maxResults=50&orderBy=&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchDashboardsWhenTheEndpointIsEmpty",
			opts: &models.DashboardSearchOptionsScheme{
				DashboardName:       "Bug",
				GroupPermissionName: "administrators",
				OrderBy:             "description",
				Expand:              []string{"description", "favourite", "sharePermissions"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/search-dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchDashboardsWhenTheResponseBodyHasADifferentFormat",
			opts: &models.DashboardSearchOptionsScheme{
				DashboardName:       "Bug",
				GroupPermissionName: "administrators",
				OrderBy:             "description",
				Expand:              []string{"description", "favourite", "sharePermissions"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/dashboard/search?dashboardName=&expand=description%2Cfavourite%2CsharePermissions&groupname=&maxResults=50&orderBy=&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &DashboardService{client: mockClient}

			gotResult, gotResponse, err := i.Search(testCase.context, testCase.opts, testCase.startAt, testCase.maxResults)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				for _, dashboard := range gotResult.Values {

					t.Log("--------------------------------")
					t.Logf("New copy Dashboard ID: %v", dashboard.ID)
					t.Logf("New copy Dashboard Name: %v", dashboard.Name)
					t.Logf("New copy Dashboard Self: %v", dashboard.Self)
					t.Log("-------------------------------- \n")
				}

			}
		})

	}

}

func TestDashboardService_Update_V2(t *testing.T) {

	testCases := []struct {
		name               string
		dashboardID        string
		payload            *models.DashboardPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:        "UpdateDashboardWhenTheParametersAreCorrect",
			dashboardID: "1001",
			payload: &models.DashboardPayloadScheme{
				Name: "new dashboard update name",
			},
			mockFile:           "../v3/mocks/update-dashboard.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/dashboard/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "UpdateDashboardWhenThePayloadIsNotProvided",
			dashboardID:        "1001",
			payload:            nil,
			mockFile:           "../v3/mocks/update-dashboard.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/dashboard/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateDashboardWhenTheDashboardIDIsNotProvided",
			dashboardID: "",
			payload: &models.DashboardPayloadScheme{
				Name: "new dashboard update name",
			},
			mockFile:           "../v3/mocks/update-dashboard.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/dashboard/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateDashboardWhenTheContextIsNotProvided",
			dashboardID: "1001",
			payload: &models.DashboardPayloadScheme{
				Name: "new dashboard update name",
			},
			mockFile:           "../v3/mocks/update-dashboard.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/dashboard/1001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateDashboardWhenTheRequestMethodIsIncorrect",
			dashboardID: "1001",
			payload: &models.DashboardPayloadScheme{
				Name: "new dashboard update name",
			},
			mockFile:           "../v3/mocks/update-dashboard.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/dashboard/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateDashboardWhenTheStatusCodeIsIncorrect",
			dashboardID: "1001",
			payload: &models.DashboardPayloadScheme{
				Name: "new dashboard update name",
			},
			mockFile:           "../v3/mocks/update-dashboard.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/dashboard/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &DashboardService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.dashboardID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Log("--------------------------------")
				t.Logf("Dashboard Updated ID: %v", gotResult.ID)
				t.Logf("Dashboard Updated Name: %v", gotResult.Name)
				t.Logf("Dashboard Updated Self: %v", gotResult.Self)
				t.Log("-------------------------------- \n")

			}
		})

	}

}
