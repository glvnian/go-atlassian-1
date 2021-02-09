package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type IssueService struct {
	client     *Client
	Attachment *AttachmentService
	Comment    *CommentService
	Field      *FieldService
	Link       *IssueLinkService
	Priority   *PriorityService
	Resolution *ResolutionService
	Security   *IssueSecurityService
	Type       *IssueTypeService
	Votes      *VoteService
	Watchers   *WatcherService
	Label      *LabelService
	Search     *IssueSearchService
}

type IssueScheme struct {
	ID          string                        `json:"id,omitemtpy"`
	Key         string                        `json:"key,omitemtpy"`
	Self        string                        `json:"self,omitemtpy"`
	Transitions []IssueSearchTransitionScheme `json:"transitions,omitemtpy"`
	Changelog   IssueChangelogScheme          `json:"changelog"`
	Fields      struct {
		IssueType  IssueTypeScheme        `json:"issuetype,omitemtpy"`
		IssueLinks []IssueLinkScheme      `json:"issuelinks,omitemtpy"`
		Watcher    IssueWatcherScheme     `json:"watches,omitemtpy"`
		Votes      IssueVoteScheme        `json:"votes,omitemtpy"`
		Versions   []ProjectVersionScheme `json:"versions,omitemtpy"`
	} `json:"fields,omitemtpy"`
}

// Creates an issue or, where the option to create subtasks is enabled in Jira, a subtask.
// A transition may be applied, to move the issue or subtask to a workflow step other than the default start step, and issue properties set.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-post
func (i *IssueService) Create(ctx context.Context, payload interface{}) (result *IssueScheme, response *Response, err error) {

	var endpoint = "rest/api/3/issue"
	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type BulkIssueScheme struct {
	Issues []IssueScheme `json:"issues"`
}

// Creates issues and, where the option to create subtasks is enabled in Jira, subtasks.
// Transitions may be applied, to move the issues or subtasks to a workflow step other than the default start step, and issue properties set.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-bulk-post
func (i *IssueService) Creates(ctx context.Context, payload interface{}) (result *BulkIssueScheme, response *Response, err error) {

	var endpoint = "rest/api/3/issue/bulk"
	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(BulkIssueScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueMetadataScheme struct {
	Expand string `json:"expand"`

	Projects []struct {
		Expand     string `json:"expand"`
		Self       string `json:"self"`
		ID         string `json:"id"`
		Key        string `json:"key"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`

		IssueTypes []struct {
			Self             string `json:"self"`
			ID               string `json:"id"`
			Description      string `json:"description"`
			IconURL          string `json:"iconUrl"`
			Name             string `json:"name"`
			UntranslatedName string `json:"untranslatedName"`
			Subtask          bool   `json:"subtask"`
			Expand           string `json:"expand"`

			Fields struct {
				Summary struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"summary"`
				IssueType struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string        `json:"name"`
					Key             string        `json:"key"`
					HasDefaultValue bool          `json:"hasDefaultValue"`
					Operations      []interface{} `json:"operations"`
					AllowedValues   []struct {
						Self        string `json:"self"`
						ID          string `json:"id"`
						Description string `json:"description"`
						IconURL     string `json:"iconUrl"`
						Name        string `json:"name"`
						Subtask     bool   `json:"subtask"`
						AvatarID    int    `json:"avatarId"`
					} `json:"allowedValues"`
				} `json:"issuetype"`
				Components struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						Items  string `json:"items"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
					AllowedValues   []struct {
						Self        string `json:"self"`
						ID          string `json:"id"`
						Name        string `json:"name"`
						Description string `json:"description"`
					} `json:"allowedValues"`
				} `json:"components"`
				Description struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"description"`
				Project struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
					AllowedValues   []struct {
						Self           string `json:"self"`
						ID             string `json:"id"`
						Key            string `json:"key"`
						Name           string `json:"name"`
						ProjectTypeKey string `json:"projectTypeKey"`
						Simplified     bool   `json:"simplified"`
						AvatarUrls     struct {
							Four8X48  string `json:"48x48"`
							Two4X24   string `json:"24x24"`
							One6X16   string `json:"16x16"`
							Three2X32 string `json:"32x32"`
						} `json:"avatarUrls"`
					} `json:"allowedValues"`
				} `json:"project"`
				Reporter struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"reporter"`
				Priority struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
					AllowedValues   []struct {
						Self    string `json:"self"`
						IconURL string `json:"iconUrl"`
						Name    string `json:"name"`
						ID      string `json:"id"`
					} `json:"allowedValues"`
					DefaultValue struct {
						Self    string `json:"self"`
						IconURL string `json:"iconUrl"`
						Name    string `json:"name"`
						ID      string `json:"id"`
					} `json:"defaultValue"`
				} `json:"priority"`
				Customfield10002 struct {
					Required bool `json:"required"`
					Schema   struct {
						Type     string `json:"type"`
						Items    string `json:"items"`
						Custom   string `json:"custom"`
						CustomID int    `json:"customId"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"customfield_10002"`
				Customfield10003 struct {
					Required bool `json:"required"`
					Schema   struct {
						Type     string `json:"type"`
						Items    string `json:"items"`
						Custom   string `json:"custom"`
						CustomID int    `json:"customId"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"customfield_10003"`
				Labels struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						Items  string `json:"items"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"labels"`
				Customfield10026 struct {
					Required bool `json:"required"`
					Schema   struct {
						Type     string `json:"type"`
						Items    string `json:"items"`
						Custom   string `json:"custom"`
						CustomID int    `json:"customId"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"customfield_10026"`
				Attachment struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						Items  string `json:"items"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string        `json:"name"`
					Key             string        `json:"key"`
					HasDefaultValue bool          `json:"hasDefaultValue"`
					Operations      []interface{} `json:"operations"`
				} `json:"attachment"`
				Duedate struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"duedate"`
				Issuelinks struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						Items  string `json:"items"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"issuelinks"`
				Assignee struct {
					Required bool `json:"required"`
					Schema   struct {
						Type   string `json:"type"`
						System string `json:"system"`
					} `json:"schema"`
					Name            string   `json:"name"`
					Key             string   `json:"key"`
					AutoCompleteURL string   `json:"autoCompleteUrl"`
					HasDefaultValue bool     `json:"hasDefaultValue"`
					Operations      []string `json:"operations"`
				} `json:"assignee"`
			} `json:"fields"`
		} `json:"issuetypes"`
	} `json:"projects"`
}

type IssueMetadataOptions struct {
	ProjectIDs     []string
	ProjectKeys    []string
	IssueTypeIDs   []string
	IssueTypeNames []string
	Expand         []string
}

// Returns details of projects, issue types within projects, and, when requested,
// the create screen fields for each issue type for the user.
// Use the information to populate the requests in Create and Creates methods
func (i *IssueService) CreateMetadata(ctx context.Context, opts *IssueMetadataOptions) (result *IssueMetadataScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range opts.Expand {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var projectIDs string
	for index, value := range opts.ProjectIDs {

		if index == 0 {
			projectIDs = value
			continue
		}

		projectIDs += "," + value
	}

	if len(expand) != 0 {
		params.Add("projectIds", projectIDs)
	}

	var projectKeys string
	for index, value := range opts.ProjectKeys {

		if index == 0 {
			projectKeys = value
			continue
		}

		projectKeys += "," + value
	}

	if len(expand) != 0 {
		params.Add("projectKeys", projectKeys)
	}

	var issueTypeIDs string
	for index, value := range opts.IssueTypeIDs {

		if index == 0 {
			issueTypeIDs = value
			continue
		}

		issueTypeIDs += "," + value
	}

	if len(expand) != 0 {
		params.Add("issuetypeIds", issueTypeIDs)
	}

	var issueTypeNames string
	for index, value := range opts.IssueTypeNames {

		if index == 0 {
			issueTypeNames = value
			continue
		}

		issueTypeNames += "," + value
	}

	if len(expand) != 0 {
		params.Add("issuetypeNames", issueTypeNames)
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/issue/createmeta?%v", params.Encode())
	} else {
		endpoint = "rest/api/3/issue/createmeta"
	}

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueMetadataScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns the details for an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-issueidorkey-get
func (i *IssueService) Get(ctx context.Context, issueKeyOrID string, fields []string, expands []string) (result *IssueScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var fieldsNames string
	for index, value := range fields {

		if index == 0 {
			fieldsNames = value
			continue
		}

		fieldsNames += "," + value
	}

	if len(expand) != 0 {
		params.Add("fields", fieldsNames)
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v?%v", issueKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v", issueKeyOrID)
	}

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Edits an issue. A transition may be applied and issue properties updated as part of the edit.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-issueidorkey-put
func (i *IssueService) Update(ctx context.Context, issueKeyOrID string, payload interface{}) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v", issueKeyOrID)
	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-issueidorkey-delete
func (i *IssueService) Delete(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v", issueKeyOrID)
	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Assigns an issue to a user.
// Use this operation when the calling user does not have the Edit Issues permission but has the
// Assign issue permission for the project that the issue is in.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-issueidorkey-assignee-put
func (i *IssueService) Assign(ctx context.Context, issueKeyOrID, accountID string) (response *Response, err error) {

	payload := struct {
		AccountID string `json:"accountId"`
	}{AccountID: accountID}

	var endpoint = fmt.Sprintf("/rest/api/3/issue/%v/assignee", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}
