package client

import "fmt"

// Returns audit logs for an organization within a specified time range.
//
// [Verkada API Docs - Verkada API Docs - Get Audit Logs]
//
// [Verkada API Docs - Verkada API Docs - Get Audit Logs]: https://apidocs.verkada.com/reference/getauditlogsviewv1
func (c *CoreClient) GetAuditLogs(options *GetAuditLogsOptions) (*GetAuditLogsResponse, error) {
	if options == nil {
		options = &GetAuditLogsOptions{}
	}
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	var ret GetAuditLogsResponse
	url := c.client.baseURL + "/core/v1/audit_log"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetAuditLogsResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Audit_logs = append(ret.Audit_logs, nextRet.Audit_logs...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

// Deletes a user for an organization based on either provided user ID or an external ID set during creation.
//
// [Verkada API Docs - Verkada API Docs - Delete User]
//
// [Verkada API Docs - Verkada API Docs - Delete User]: https://apidocs.verkada.com/reference/deleteuserviewv1
func (c *CoreClient) DeleteUser(options *DeleteUserOptions) (*DeleteUserResponse, error) {
	if options == nil {
		options = &DeleteUserOptions{}
	}
	// should not use both user_id and external_id, but need at least one
	if (options.User_id == "") == (options.External_id == "") {
		return nil, fmt.Errorf("should use one of user_id and external_id - received user_id: %s and external_id: %s", options.User_id, options.External_id)
	}
	var ret DeleteUserResponse
	url := c.client.baseURL + "/core/v1/user"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Returns a user for an organization based on either provided user ID or an external ID set during creation.
//
// [Verkada API Docs - Get User]
//
// [Verkada API Docs - Get User]: https://apidocs.verkada.com/reference/getuserviewv1
func (c *CoreClient) GetUser(options *GetUserOptions) (*GetUserResponse, error) {
	if options == nil {
		options = &GetUserOptions{}
	}
	// should not use both user_id and external_id, but need at least one
	if (options.User_id == "") == (options.External_id == "") {
		return nil, fmt.Errorf("should only use one of user_id and external_id - received user_id: %s and external_id: %s", options.User_id, options.External_id)
	}
	var ret GetUserResponse
	url := c.client.baseURL + "/core/v1/user"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Creates a user for in an organization. External ID required.
// Otherwise, the newly created user will contain a user ID which can be used for identification.
//
// Requires at least one of:
//   - First name
//   - Last name
//   - Email
//
// [Verkada API Docs - Create User]
//
// [Verkada API Docs - Create User]: https://apidocs.verkada.com/reference/postuserviewv1
func (c *CoreClient) CreateUser(body *CreateUserBody) (*CreateUserResponse, error) {
	var ret CreateUserResponse
	url := c.client.baseURL + "/core/v1/user"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Updates a user's metadata for an organization based on either provided user ID or an external ID set during creation.
//
// [Verkada API Docs - Update User]
//
// [Verkada API Docs - Update User]: https://apidocs.verkada.com/reference/putuserviewv1
func (c *CoreClient) UpdateUser(options *UpdateUserOptions, body *UpdateUserBody) (*UpdateUserResponse, error) {
	if options == nil {
		options = &UpdateUserOptions{}
	}
	// should not use both user_id and external_id, but need at least one
	if (options.User_id == "") == (options.External_id == "") {
		return nil, fmt.Errorf("should only use one of user_id and external_id - received user_id: %s and external_id: %s", options.User_id, options.External_id)
	}
	var ret UpdateUserResponse
	url := c.client.baseURL + "/core/v1/user"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, body, &ret, 0)
	return &ret, err
}
