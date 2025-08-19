package client

import (
	"fmt"
	"strings"
)

// Retrieves a list of all access groups in a given organization.
//
// [Verkada API Docs - Get All Access Groups]
//
// [Verkada API Docs - Get All Access Groups]: https://apidocs.verkada.com/reference/getaccessgroupsviewv1
func (c *AccessClient) GetAllAccessGroups() (*GetAllAccessGroupsResponse, error) {
	var ret GetAllAccessGroupsResponse
	url := c.client.baseURL + "/access/v1/access_groups"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Delete an access group with the given group identifier within the given organization.
//
// [Verkada API Docs - Delete Access Group]
//
// [Verkada API Docs - Delete Access Group]: https://apidocs.verkada.com/reference/deleteaccessgroupviewv1
func (c *AccessClient) DeleteAccessGroups(group_id string) (*DeleteAccessGroupResponse, error) {
	options := &DeleteAccessGroupOptions{group_id: group_id}
	var ret DeleteAccessGroupResponse
	url := c.client.baseURL + "/access/v1/access_groups/group"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Retrieves an access group specified by its Verkada-defined unique identifier (Group ID).
//
// [Verkada API Docs - Get Access Group]
//
// [Verkada API Docs - Get Access Group]: https://apidocs.verkada.com/reference/getaccessgroupviewv1
func (c *AccessClient) GetAccessGroup(group_id string) (*AccessGroup, error) {
	options := &GetAccessGroupOptions{group_id: group_id}
	var ret AccessGroup
	url := c.client.baseURL + "/access/v1/access_groups/group"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Create an access group within the given organization using the given name.
// The name of the access group must be unique within the organization.
//
// [Verkada API Docs - Create Access Group]
//
// [Verkada API Docs - Create Access Group]: https://apidocs.verkada.com/reference/postaccessgroupviewv1
func (c *AccessClient) CreateAccessGroup(name string) (*AccessGroup, error) {
	body := &CreateAccessGroupBody{Name: name}
	var ret AccessGroup
	url := c.client.baseURL + "/access/v1/access_groups/group"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Remove an access user to an access group with the Verkada defined group ID and the user defined either by their User ID or their External ID.
//
// [Verkada API Docs - Remove User From Access Group]
//
// [Verkada API Docs - Remove User From Access Group]: https://apidocs.verkada.com/reference/deleteaccessgroupuserviewv1
func (c *AccessClient) RemoveUserFromAccessGroup(group_id string, options *RemoveUserFromAccessGroupOptions) (*RemoveUserFromAccessGroupResponse, error) {
	if options == nil {
		options = &RemoveUserFromAccessGroupOptions{}
	}
	options.group_id = group_id
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret RemoveUserFromAccessGroupResponse
	url := c.client.baseURL + "/access/v1/access_groups/group/user"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Add an access user to an access group with the Verkada defined group ID and either the user defined External ID or the Verkada defined user ID.
//
// [Verkada API Docs - Add User to Access Group]
//
// [Verkada API Docs - Add User tp Access Group]: https://apidocs.verkada.com/reference/putaccessgroupuserviewv1
func (c *AccessClient) AddUserToAccessGroup(group_id string, body *AddUserToAccessGroupBody) (*AddUserToAccessGroupResponse, error) {
	options := &AddUserToAccessGroupOptions{group_id: group_id}
	if body == nil {
		body = &AddUserToAccessGroupBody{}
	}
	// should not use both external_id and user_id, but need at least one
	if (body.External_id == "") == (body.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", body.External_id, body.User_id)
	}
	var ret AddUserToAccessGroupResponse
	url := c.client.baseURL + "/access/v1/access_groups/group/user"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, body, &ret, 0)
	return &ret, err
}

// This returns a list of All Access User Objects for all access members in an organization.
//
// [Verkada API Docs - Get All Access Users]
//
// [Verkada API Docs - Get All Access Users]: https://apidocs.verkada.com/reference/getaccessmembersviewv1
func (c *AccessClient) GetAllAccessUsers() (*GetAllAccessUsersResponse, error) {
	var ret GetAllAccessUsersResponse
	url := c.client.baseURL + "/access/v1/access_users"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Given the Verkada created User Id, user defined External Id, email address, or employee ID, retrieve the Access Object Information for the specified user.
//
// [Verkada API Docs - Get Access Information Object]
//
// [Verkada API Docs - Get Access Information Object]: https://apidocs.verkada.com/reference/getaccessuserviewv1
func (c *AccessClient) GetAccessInformationObject(options *GetAccessInformationObjectOptions) (*AccessInformationObject, error) {
	if options == nil {
		options = &GetAccessInformationObjectOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), activate bluetooth unlock capability for a user.
//
// [Verkada API Docs - Activate BLE for Access User]
//
// [Verkada API Docs - Activate BLE for Access User]: https://apidocs.verkada.com/reference/putactivateblemethodviewv1
func (c *AccessClient) ActivateUserBLE(options *ActivateUserBLEOptions) (*AccessInformationObject, error) {
	if options == nil {
		options = &ActivateUserBLEOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/ble/activate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), deactivate bluetooth unlock capability for a user.
//
// [Verkada API Docs - Deactivate BLE for Access User]
//
// [Verkada API Docs - Deactivate BLE for Access User]: https://apidocs.verkada.com/reference/putdeactivateblemethodviewv1
func (c *AccessClient) DeactivateUserBLE(options *DeactivateUserBLEOptions) (*AccessInformationObject, error) {
	if options == nil {
		options = &DeactivateUserBLEOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/ble/deactivate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), set the end date for an access users credentials to become invalid.
// After this time, all methods of access will be revoked.
//
// [Verkada API Docs - Set End Date for User]
//
// [Verkada API Docs - Set End Date for User]: https://apidocs.verkada.com/reference/putaccessenddateviewv1
func (c *AccessClient) SetUserEndDate(end_date string, options *SetUserEndDateOptions) (*AccessInformationObject, error) {
	body := struct {
		End_date string `json:"end_date"`
	}{
		End_date: end_date,
	}
	if options == nil {
		options = &SetUserEndDateOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/end_date"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, body, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), remove the entry code of a user.
//
// [Verkada API Docs - Remove Entry Code for User]
//
// [Verkada API Docs - Remove Entry Code for User]: https://apidocs.verkada.com/reference/deleteaccessuserpinviewv1
func (c *AccessClient) RemoveUserEntryCode(options *RemoveUserEntryCodeOptions) (*RemoveUserEntryCodeResponse, error) {
	if options == nil {
		options = &RemoveUserEntryCodeOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret RemoveUserEntryCodeResponse
	url := c.client.baseURL + "/access/v1/access_users/user/entry_code"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), set the entry code for a user.
//
// [Verkada API Docs - Set Entry Code for User]
//
// [Verkada API Docs - Set Entry Code for User]: https://apidocs.verkada.com/reference/putaccessuserpinviewv1
func (c *AccessClient) SetUserEntryCode(entry_code string, options *SetUserEntryCodeOptions) (*AccessInformationObject, error) {
	body := struct {
		Entry_code string `json:"entry_code"`
	}{
		Entry_code: entry_code,
	}
	if options == nil {
		options = &SetUserEntryCodeOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/entry_code"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, body, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), send out a Pass App invite.
//
// [Verkada API Docs - Send Pass App Invite for User]
//
// [Verkada API Docs - Send Pass App Invite for User]: https://apidocs.verkada.com/reference/postsendpassappinviteviewv1
func (c *AccessClient) SendPassInvite(options *SendPassInviteOptions) (*AccessInformationObject, error) {
	if options == nil {
		options = &SendPassInviteOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/pass/invite"
	err := c.client.MakeVerkadaRequest("POST", url, *options, nil, &ret, 0)
	return &ret, err
}

// Delete a profile photo for the specified user.
//
// [Verkada API Docs - Delete Profile Photo]
//
// [Verkada API Docs - Delete Profile Photo]: https://apidocs.verkada.com/reference/deleteprofilephotoviewv1
func (c *AccessClient) DeleteProfilePhoto(options *DeleteProfilePhotoOptions) (*DeleteProfilePhotoResponse, error) {
	if options == nil {
		options = &DeleteProfilePhotoOptions{}
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret DeleteProfilePhotoResponse
	url := c.client.baseURL + "/access/v1/access_users/user/profile_photo"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Retrieve a profile photo for the specified user.
//
// [Verkada API Docs - Get Profile Photo]
//
// [Verkada API Docs - Get Profile Photo]: https://apidocs.verkada.com/reference/getprofilephotoviewv1
func (c *AccessClient) GetProfilePhoto(options *GetProfilePhotoOptions, filename string) error {
	if options == nil {
		options = &GetProfilePhotoOptions{}
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	// filename validation and replacement if left blank
	if filename == "" {
		filename = "profilephoto.jpg"
	}
	if !strings.HasSuffix(filename, ".jpg") {
		return fmt.Errorf("included filename is not blank but does not end with \".jpg\" - received %s", filename)
	}
	url := c.client.baseURL + "/access/v1/access_users/user/profile_photo"
	err := c.client.MakeVerkadaRequestForFile("GET", url, *options, filename, 0)
	return err
}

// Upload a profile photo for the specified user.
//
// [Verkada API Docs - Upload Profile Photo]
//
// [Verkada API Docs - Upload Profile Photo]: https://apidocs.verkada.com/reference/putprofilephotoviewv1
func (c *AccessClient) UploadProfilePhoto(options *UploadProfilePhotoOptions, filename string) error {
	if options == nil {
		options = &UploadProfilePhotoOptions{}
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	// filename validation and replacement if left blank
	if !strings.HasSuffix(filename, ".jpg") {
		return fmt.Errorf("included filename does not end with \".jpg\" - received %s", filename)
	}
	var ret CreateProfilePhotoResponse
	url := c.client.baseURL + "/access/v1/access_users/user/profile_photo"
	err := c.client.MakeVerkadaRequestWithFile("PUT", url, *options, filename, "image/jpeg", &ret, 0)
	return err
}

// Given the user defined external ID or Verkada defined user ID (but not both), activate remote unlock capability for a user.
//
// [Verkada API Docs - Activate Remote Unlock for User]
//
// [Verkada API Docs - Activate Remote Unlock for User]: https://apidocs.verkada.com/reference/putactivateremoteunlockviewv1
func (c *AccessClient) ActivateUserRemoteUnlock(options *ActivateUserRemoteUnlockOptions) (*AccessInformationObject, error) {
	if options == nil {
		options = &ActivateUserRemoteUnlockOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/remote_unlock/activate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), deactivate remote unlock capability for a user.
//
// [Verkada API Docs - Deactivate Remote Unlock for User]
//
// [Verkada API Docs - Deactivate Remote Unlock for User]: https://apidocs.verkada.com/reference/putdeactivateremoteunlockviewv1
func (c *AccessClient) DeactivateUserRemoteUnlock(options *DeactivateUserRemoteUnlockOptions) (*AccessInformationObject, error) {
	if options == nil {
		options = &DeactivateUserRemoteUnlockOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/remote_unlock/deactivate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the user defined external ID or Verkada defined user ID (but not both), set the start date for an access users credentials to become valid.
// Before this time, all methods of access specified for this access user will invalid.
//
// [Verkada API Docs - Set Start Date for User]
//
// [Verkada API Docs - Set Start Date for User]: https://apidocs.verkada.com/reference/putaccessstartdateviewv1
func (c *AccessClient) SetStartDate(start_date string, options *SetStartDateOptions) (*AccessInformationObject, error) {
	body := struct {
		Start_date string `json:"start_date"`
	}{
		Start_date: start_date,
	}
	if options == nil {
		options = &SetStartDateOptions{}
	}
	keyCount := 0
	for _, key := range []string{options.Email, options.External_id, options.User_id, options.Employee_id} {
		if key == "" {
			keyCount++
		}
	}
	if keyCount != 1 {
		return nil, fmt.Errorf("need one out of email, external_id, user_id, and employee_id - received email: %s, external id: %s, user_id: %s, employee_id: %s", options.Email, options.External_id, options.User_id, options.Employee_id)
	}
	var ret AccessInformationObject
	url := c.client.baseURL + "/access/v1/access_users/user/start_date"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, body, &ret, 0)
	return &ret, err
}

// Deletes an access card of a specified access user given their user_id or external_id, the org_id, and the card_id.
//
// [Verkada API Docs - Delete Access Card]
//
// [Verkada API Docs - Delete Access Card]: https://apidocs.verkada.com/reference/deleteaccesscardviewv1
func (c *AccessClient) DeleteAccessCard(card_id string, options *DeleteAccessCardOptions) (*DeleteAccessCardResponse, error) {
	if options == nil {
		options = &DeleteAccessCardOptions{}
	}
	options.card_id = card_id
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret DeleteAccessCardResponse
	url := c.client.baseURL + "/access/v1/credentials/card"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Create and add an access card for a specified user_id or external_id and org_id.
// All cards require one of:
//   - card_number
//   - card_number_hex
//   - card_number_base36.
//
// [Verkada API Docs - Delete Access Card]
//
// [Verkada API Docs - Delete Access Card]: https://apidocs.verkada.com/reference/deleteaccesscardviewv1
func (c *AccessClient) AddAccessCard(format string, options *AddAccessCardOptions, body *AddAccessCardBody) (*Card, error) {
	if options == nil {
		options = &AddAccessCardOptions{}
	}
	fullBody := struct {
		Active             bool   `json:"active,omitempty"`
		Card_number        string `json:"card_number,omitempty"`
		Card_number_base36 string `json:"card_number_base36,omitempty"`
		Card_number_hex    string `json:"card_number_hex,omitempty"`
		Facility_code      string `json:"facility_code,omitempty"`
		Type               string `json:"type"`
	}{
		Active:             body.Active,
		Card_number:        body.Card_number,
		Card_number_base36: body.Card_number_base36,
		Card_number_hex:    body.Card_number_hex,
		Facility_code:      body.Facility_code,
		Type:               format,
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	// Card format must be one of the following:
	format_validation := map[string]bool{
		"Standard 26-bit Wiegand":     true,
		"HID":                         true,
		"HID 37-bit":                  true,
		"HID37wFacilityCode":          true,
		"HID 37-bit No Facility Code": true,
		"HID37woFacilityCode":         true,
		"HID 34-bit":                  true,
		"HID34":                       true,
		"Casi Rusco 40-Bit":           true,
		"CasiRusco":                   true,
		"HID Corporate 1000-35":       true,
		"Corporate1000_35":            true,
		"HID Corporate 1000-48":       true,
		"Corporate1000_48":            true,
		"HID iClass":                  true,
		"iClass":                      true,
		"DESFire CSN":                 true,
		"DESFire":                     true,
		"Verkada DESFire":             true,
		"VerkadaDESFire":              true,
		"DESFire 40X":                 true,
		"MiFareClassic1K_CSN":         true,
		"MiFareClassic4K_CSN":         true,
		"Apple Wallet Pass":           true,
		"MiFare 4-Byte (32 bit) CSN":  true,
		"MDC Custom 64-bit":           true,
		"MDCCustom_64":                true,
		"HID 36-bit Keyscan":          true,
		"HID36Keyscan":                true,
		"HID 33-bit DSX":              true,
		"HID33DSX":                    true,
		"HID 33-bit RS2":              true,
		"HID33RS2":                    true,
		"HID 36-bit Simplex":          true,
		"HID36Simplex":                true,
		"Cansec 37-bit":               true,
		"Cansec37":                    true,
		"Credit Card BIN Number":      true,
		"CreditCardBin":               true,
		"Kantech XSF":                 true,
		"KantechXSF":                  true,
		"Schlage 34-bit":              true,
		"Schlage34":                   true,
		"Schlage 37-bit":              true,
		"Schlage37x":                  true,
		"RBH 50-bit":                  true,
		"RBH50":                       true,
		"Guardall G-Prox II 36-bit":   true,
		"GProxII36":                   true,
		"AMAG 32-bit":                 true,
		"AMAG32":                      true,
		"Securitas 37-bit":            true,
		"Securitas37":                 true,
		"Kastle 32-bit":               true,
		"Kastle32":                    true,
		"PointGuard MDI 37-bit":       true,
		"PointGuardMDI37":             true,
		"Blackboard 64-bit":           true,
		"Blackboard64":                true,
		"IDm 64-bit":                  true,
		"IDm64bit":                    true,
		"Continental 36-bit":          true,
		"Continental36":               true,
		"AWID 34-bit":                 true,
		"AWID34":                      true,
		"License Plate":               true,
		"HID Infinity 37-bit":         true,
		"HIDInfinity37":               true,
		"HID Ceridian 26-bit":         true,
		"iClass 35-bit":               true,
		"Andover Controls 37-bit":     true,
	}
	if ok := format_validation[format]; !ok {
		return nil, fmt.Errorf("could not validate card format: %s", format)
	}
	var ret Card
	url := c.client.baseURL + "/access/v1/credentials/card"
	err := c.client.MakeVerkadaRequest("POST", url, *options, fullBody, &ret, 0)
	return &ret, err
}

// Given the card ID and Verkada defined user ID or user defined external ID, activate a specific access card for a user.
//
// [Verkada API Docs - Activate Access Card]
//
// [Verkada API Docs - Activate Access Card]: https://apidocs.verkada.com/reference/putaccesscardactivateviewv1
func (c *AccessClient) ActivateAccessCard(card_id string, options *ActivateAccessCardOptions) (*Card, error) {
	if options == nil {
		options = &ActivateAccessCardOptions{}
	}
	options.card_id = card_id
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret Card
	url := c.client.baseURL + "/access/v1/credentials/card/activate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the card ID and Verkada defined user ID or user defined external ID, deactivate a specific access card for a user.
//
// [Verkada API Docs - Deactivate Access Card]
//
// [Verkada API Docs - Deactivate Access Card]: https://apidocs.verkada.com/reference/putaccesscarddeactivateviewv1
func (c *AccessClient) DeactivateAccessCard(card_id string, options *DeactivateAccessCardOptions) (*Card, error) {
	if options == nil {
		options = &DeactivateAccessCardOptions{}
	}
	options.card_id = card_id
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret Card
	url := c.client.baseURL + "/access/v1/credentials/card/deactivate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Deletes a license plate credential from a specified user by providing the user_id or the external_id, the org_id, and the license_plate_number.
//
// [Verkada API Docs - Delete License Plate from User]
//
// [Verkada API Docs - Delete License Plate from User]: https://apidocs.verkada.com/reference/deletelicenseplateviewv1
func (c *AccessClient) DeleteUserLicensePlate(license_plate_number string, options *DeleteUserLicensePlateOptions) (*DeleteUserLicensePlateResponse, error) {
	if options == nil {
		options = &DeleteUserLicensePlateOptions{}
	}
	options.license_plate_number = license_plate_number
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret DeleteUserLicensePlateResponse
	url := c.client.baseURL + "/access/v1/credentials/license_plate"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Add a license plate credential to a user given a specified user_id or external_id and org_id.
// The plate number must be a string of 6 alphanumeric values.
//
// [Verkada API Docs - Add License Plate from User]
//
// [Verkada API Docs - Add License Plate from User]: https://apidocs.verkada.com/reference/postlicenseplateviewv1
func (c *AccessClient) AddUserLicensePlate(license_plate_number string, options *AddUserLicensePlateOptions, body *AddUserLicensePlatedBody) (*LicensePlate, error) {
	if options == nil {
		options = &AddUserLicensePlateOptions{}
	}
	fullBody := struct {
		Active               bool   `json:"active,omitempty"`
		License_plate_number string `json:"license_plate_number"`
		Name                 string `json:"name,omitempty"`
	}{
		Active:               body.Active,
		License_plate_number: license_plate_number,
		Name:                 body.Name,
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret LicensePlate
	url := c.client.baseURL + "/access/v1/credentials/license_plate"
	err := c.client.MakeVerkadaRequest("POST", url, *options, fullBody, &ret, 0)
	return &ret, err
}

// Given the Verkada defined user ID (OR user defined external ID) and Licence Plate Number, activate a users Licence Plate Credential.
//
// [Verkada API Docs - Activate License Plate]
//
// [Verkada API Docs - Activate License Plate]: https://apidocs.verkada.com/reference/putlicenseplateactivateviewv1
func (c *AccessClient) ActivateLicensePlate(license_plate_number string, options *ActivateLicensePlateOptions) (*LicensePlate, error) {
	if options == nil {
		options = &ActivateLicensePlateOptions{}
	}
	options.License_plate_number = license_plate_number
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret LicensePlate
	url := c.client.baseURL + "/access/v1/credentials/license_plate/activate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Given the Verkada defined user ID (OR user defined external ID) and Licence Plate Number, deactivate a users Licence Plate Credential.
//
// [Verkada API Docs - Deactivate License Plate]
//
// [Verkada API Docs - Deactivate License Plate]: https://apidocs.verkada.com/reference/putlicenseplatedeactivateviewv1
func (c *AccessClient) DeactivateLicensePlate(license_plate_number string, options *DeactivateLicensePlateOptions) (*LicensePlate, error) {
	if options == nil {
		options = &DeactivateLicensePlateOptions{}
	}
	options.License_plate_number = license_plate_number
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret LicensePlate
	url := c.client.baseURL + "/access/v1/credentials/license_plate/deactivate"
	err := c.client.MakeVerkadaRequest("PUT", url, *options, nil, &ret, 0)
	return &ret, err
}

// Deletes an MFA code credential from a specified user by providing the user_id or external_id, the org_id, and the mfa_code.
//
// [Verkada API Docs - Delete MFA Code from User]
//
// [Verkada API Docs - Delete MFA Code from User]: https://apidocs.verkada.com/reference/deletemfacodeviewv1
func (c *AccessClient) DeleteMFACode(code string, options *DeleteMFACodeOptions) (*DeleteMFACodeResponse, error) {
	if options == nil {
		options = &DeleteMFACodeOptions{}
	}
	options.code = code
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret DeleteMFACodeResponse
	url := c.client.baseURL + "/access/v1/credentials/mfa_code"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Add an MFA code credential to a user given a specified user_id or external_id and org_id.
//
// [Verkada API Docs - Add MFA Code to User]
//
// [Verkada API Docs - Add MFA Code to User]: https://apidocs.verkada.com/reference/postmfacodeviewv1
func (c *AccessClient) AddMFACode(code string, options *AddMFACodeOptions) (*AddMFACodeResponse, error) {
	if options == nil {
		options = &AddMFACodeOptions{}
	}
	body := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret AddMFACodeResponse
	url := c.client.baseURL + "/access/v1/credentials/mfa_code"
	err := c.client.MakeVerkadaRequest("POST", url, *options, body, &ret, 0)
	return &ret, err
}

// Retrieves a list of all Access Levels that the API Token is authorized to access.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for Access Levels that you want to retrieve.
// "Manage via API" must also be enabled for all doors listed in doors that you want to retrieve.
//
// [Verkada API Docs - Get All Available Access Levels]
//
// [Verkada API Docs - Get All Available Access Levels]: https://apidocs.verkada.com/reference/getaccesslevelview
func (c *AccessClient) GetAllAccessLevels() (*GetAllAccessLevelsResponse, error) {
	var ret GetAllAccessLevelsResponse
	url := c.client.baseURL + "/access/v1/door/access_level"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Create a new Access Level.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Create Access Level]
//
// [Verkada API Docs - Create Access Level]: https://apidocs.verkada.com/reference/postaccesslevelview
func (c *AccessClient) CreateAccessLevel(access_groups []string, access_schedule_events []AccessScheduleEvent, doors []string, name string, sites []string) (*AccessLevel, error) {
	body := AccessLevel{
		Access_groups:          access_groups,
		Access_schedule_events: access_schedule_events,
		Doors:                  doors,
		Name:                   name,
		Sites:                  sites,
	}
	var ret AccessLevel
	url := c.client.baseURL + "/access/v1/door/access_level"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Delete an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Delete Access Level]
//
// [Verkada API Docs - Delete Access Level]: https://apidocs.verkada.com/reference/deleteaccessleveldetailview
func (c *AccessClient) DeleteAccessLevel(access_level_id string) (*DeleteAccessLevelResponse, error) {
	var ret DeleteAccessLevelResponse
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

// Retrieves an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Get Access Level]
//
// [Verkada API Docs - Get Access Level]: https://apidocs.verkada.com/reference/getaccessleveldetailview
func (c *AccessClient) GetAccessLevel(access_level_id string) (*AccessLevel, error) {
	var ret AccessLevel
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Updates an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Update Access Level]
//
// [Verkada API Docs - Update Access Level]: https://apidocs.verkada.com/reference/putaccessleveldetailview
func (c *AccessClient) UpdateAccessLevel(access_level_id string, access_groups []string, access_schedule_events []AccessScheduleEvent, doors []string, name string, sites []string) (*AccessLevel, error) {
	body := AccessLevel{
		Access_groups:          access_groups,
		Access_schedule_events: access_schedule_events,
		Doors:                  doors,
		Name:                   name,
		Sites:                  sites,
	}
	var ret AccessLevel
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id
	err := c.client.MakeVerkadaRequest("PUT", url, nil, body, &ret, 0)
	return &ret, err
}

// Add a new Access Schedule Event to an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Add Access Schedule Event to Access Level]
//
// [Verkada API Docs - Add Access Schedule Event to Access Level]: https://apidocs.verkada.com/reference/postaccesslevelscheduleview
func (c *AccessClient) AddAccessScheduleEvent(access_level_id string, end_time string, start_time string, weekday string) (*AccessScheduleEvent, error) {
	body := AccessScheduleEvent{
		Door_status: "access_granted",
		End_time:    end_time,
		Start_time:  start_time,
		Weekday:     weekday,
	}
	// Card format must be one of the following:
	day_validation := map[string]bool{
		"SU":  true,
		"MO":  true,
		"TU":  true,
		"WE":  true,
		"TH":  true,
		"FR":  true,
		"SAT": true,
	}
	if ok := day_validation[body.Weekday]; !ok {
		return nil, fmt.Errorf("could not validate card format: %s", body.Weekday)
	}
	var ret AccessScheduleEvent
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id + "/access_schedule_event"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Delete an Access Schedule Event having access_schedule_event_id on an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Delete Access Schedule Event on Access Level]
//
// [Verkada API Docs - Delete Access Schedule Event on Access Level]: https://apidocs.verkada.com/reference/deleteaccesslevelscheduleview
func (c *AccessClient) DeleteAccessScheduleEvent(access_level_id string, event_id string) (*DeleteAccessScheduleEventResponse, error) {
	var ret DeleteAccessScheduleEventResponse
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id + "/access_schedule_event/" + event_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

// Retrieve an Access Schedule Event having access_schedule_event_id on an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Get Access Schedule Event Details]
//
// [Verkada API Docs - Get Access Schedule Event Details]: https://apidocs.verkada.com/reference/getaccesslevelscheduleview
func (c *AccessClient) GetAccessScheduleEvent(access_level_id string, event_id string) (*AccessScheduleEvent, error) {
	var ret AccessScheduleEvent
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id + "/access_schedule_event/" + event_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

//	Update an Access Schedule Event having access_schedule_event_id on an Access Level having the specified access_level_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites listed in sites for the Access Level.
// "Manage via API" must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Update Access Schedule Event on Access Level]
//
// [Verkada API Docs - Update Access Schedule Event on Access Level]: https://apidocs.verkada.com/reference/putaccesslevelscheduleview
func (c *AccessClient) UpdateAccessScheduleEvent(access_level_id string, event_id string, end_time string, start_time string, weekday string) (*AccessScheduleEvent, error) {
	body := AccessScheduleEvent{
		Door_status: "access_granted",
		End_time:    end_time,
		Start_time:  start_time,
		Weekday:     weekday,
	}
	// Card format must be one of the following:
	day_validation := map[string]bool{
		"SU":  true,
		"MO":  true,
		"TU":  true,
		"WE":  true,
		"TH":  true,
		"FR":  true,
		"SAT": true,
	}
	if ok := day_validation[body.Weekday]; !ok {
		return nil, fmt.Errorf("could not validate card format: %s", body.Weekday)
	}
	var ret AccessScheduleEvent
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id + "/access_schedule_event/" + event_id
	err := c.client.MakeVerkadaRequest("PUT", url, nil, body, &ret, 0)
	return &ret, err
}

// Sends a request to unlock a door with door_id irrespective of any user's door access privileges.
//
// [Verkada API Docs - Unlock Door as Admin]
//
// [Verkada API Docs - Unlock Door as Admin]: https://apidocs.verkada.com/reference/postaccessadminapiunlockviewv1
func (c *AccessClient) AdminUnlockDoor(door_id string) (*AdminUnlockDoorResponse, error) {
	body := struct {
		Door_id string `json:"door_id"`
	}{
		Door_id: door_id,
	}
	var ret AdminUnlockDoorResponse
	url := c.client.baseURL + "/access/v1/door/admin_unlock"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Sends a request to unlock a door with door_id as the user with the specified user_id or external_id.
// The unlock request will only be granted if the user with the specified user_id has door access permissions for that door_id through an Access Level in Verkada Command.
//
// [Verkada API Docs - Unlock Door as User]
//
// [Verkada API Docs - Unlock Door as User]: https://apidocs.verkada.com/reference/postaccessuserapiunlockviewv1
func (c *AccessClient) UserUnlockDoor(door_id string, options *UserUnlockDoorOptions) (*UserUnlockDoorResponse, error) {
	body := struct {
		Door_id     string `json:"door_id"`
		User_id     string `json:"user_id,omitempty"`
		External_id string `json:"external_id,omitempty"`
	}{
		Door_id:     door_id,
		User_id:     options.User_id,
		External_id: options.External_id,
	}
	// should not use both external_id and user_id, but need at least one
	if (options.External_id == "") == (options.User_id == "") {
		return nil, fmt.Errorf("should use one of external_id and user_id - received external_id: %s and user_id: %s", options.External_id, options.User_id)
	}
	var ret UserUnlockDoorResponse
	url := c.client.baseURL + "/access/v1/door/admin_unlock"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Retrieves a list of all doors in the organization.
// The response can optionally be filtered to doors within sites with the requested site_ids or to the specific doors represented by the specified door_ids.
// Only one of site_ids or door_ids can be used to filter the response for a single request (not both).
//
// [Verkada API Docs - Get Doors]
//
// [Verkada API Docs - Get Doors]: https://apidocs.verkada.com/reference/getaccessdoorinformationviewv1
func (c *AccessClient) GetDoors(options *GetDoorsOptions) (*GetDoorsResponse, error) {
	if options == nil {
		options = &GetDoorsOptions{}
	}
	// should not use both door_ids and site_ids
	if len(options.Door_ids) != 0 && len(options.Site_ids) != 0 {
		return nil, fmt.Errorf("should not use both door_ids and site_ids - received door_ids: %v and site_ids: %v", options.Door_ids, options.Site_ids)
	}
	var ret GetDoorsResponse
	url := c.client.baseURL + "/access/v1/doors"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Retrieves a list of all Door Exception Calendars that the API Token is authorized to access.
// The response can optionally be filtered to include Door Exception Calendars that were last updated after a provided date and time.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for Door Exception Calendars that you want to retrieve.
// "Manage via API" must also be enabled for all doors listed in doors that you want to retrieve.
//
// [Verkada API Docs - Get All Available Door Exception Calendars]
//
// [Verkada API Docs - Get All Available Door Exception Calendars]: https://apidocs.verkada.com/reference/getaccessdoorexceptioncalendarsviewv1
func (c *AccessClient) GetAllDoorExceptionCalendars(options *GetAllDoorExceptionCalendarsOptions) (*GetAllDoorExceptionCalendarsResponse, error) {
	if options == nil {
		options = &GetAllDoorExceptionCalendarsOptions{}
	}
	var ret GetAllDoorExceptionCalendarsResponse
	url := c.client.baseURL + "/access/v1/door/exception_calendar"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Create a new Door Exception Calendar.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Create Door Exception Calendar]
//
// [Verkada API Docs - Create Door Exception Calendar]: https://apidocs.verkada.com/reference/postaccessdoorexceptioncalendarsviewv1
func (c *AccessClient) CreateDoorExceptionCalendar(name string, body *CreateDoorExceptionCalendarBody) (*DoorExceptionCalendar, error) {
	fullBody := struct {
		Name       string          `json:"name"`
		Doors      []string        `json:"doors,omitempty"`
		Exceptions []DoorException `json:"exceptions,omitempty"`
	}{
		Name:       name,
		Doors:      body.Doors,
		Exceptions: body.Exceptions,
	}
	for _, exception := range body.Exceptions {
		if success, err := validateDoorException(exception); !success {
			return nil, err
		}
	}
	var ret DoorExceptionCalendar
	url := c.client.baseURL + "/access/v1/door/exception_calendar"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Delete a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Delete Door Exception Calendar]
//
// [Verkada API Docs - Delete Door Exception Calendar]: https://apidocs.verkada.com/reference/deleteaccessdoorexceptioncalendarviewv1
func (c *AccessClient) DeleteDoorExceptionCalendar(calendar_id string) (*DeleteDoorExceptionCalendarResponse, error) {
	var ret DeleteDoorExceptionCalendarResponse
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

// Retrieve a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Get Door Exception Calendar]
//
// [Verkada API Docs - Get Door Exception Calendar]: https://apidocs.verkada.com/reference/getaccessdoorexceptioncalendarviewv1
func (c *AccessClient) GetDoorExceptionCalendar(calendar_id string) (*DoorExceptionCalendar, error) {
	var ret DoorExceptionCalendar
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Update a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Update Door Exception Calendar]
//
// [Verkada API Docs - Update Door Exception Calendar]: https://apidocs.verkada.com/reference/putaccessdoorexceptioncalendarviewv1
func (c *AccessClient) UpdateDoorExceptionCalendar(calendar_id string, name string, body *UpdateDoorExceptionCalendarBody) (*DoorExceptionCalendar, error) {
	fullBody := struct {
		Name       string          `json:"name"`
		Doors      []string        `json:"doors,omitempty"`
		Exceptions []DoorException `json:"exceptions,omitempty"`
	}{
		Name:       name,
		Doors:      body.Doors,
		Exceptions: body.Exceptions,
	}
	for _, exception := range body.Exceptions {
		if success, err := validateDoorException(exception); !success {
			return nil, err
		}
	}
	var ret DoorExceptionCalendar
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id
	err := c.client.MakeVerkadaRequest("PUT", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Add a new Exception to a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Add Exception to Door Exception Calendar]
//
// [Verkada API Docs - Add Exception to Door Exception Calendar]: https://apidocs.verkada.com/reference/postaccessdoorexceptionsviewv1
func (c *AccessClient) AddExceptionToCalendar(calendar_id string, date string, start_time, end_time string, body *AddExceptionToCalendarBody) (*DoorException, error) {
	fullBody := DoorException{
		All_day_default:           body.All_day_default,
		Date:                      date,
		Door_status:               body.Door_status,
		Double_badge:              body.Double_badge,
		Double_badge_group_ids:    body.Double_badge_group_ids,
		End_time:                  end_time,
		First_person_in:           body.First_person_in,
		First_person_in_group_ids: body.First_person_in_group_ids,
		Recurrence_rule:           body.Recurrence_rule,
		Start_time:                start_time,
	}
	if success, err := validateDoorException(fullBody); !success {
		return nil, err
	}
	var ret DoorException
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id + "/exception"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Delete an Exception having the specified exception_id on a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Delete Exception on Door Exception Calendar]
//
// [Verkada API Docs - Delete Exception on Door Exception Calendar]: https://apidocs.verkada.com/reference/deleteaccessdoorexceptionviewv1
func (c *AccessClient) DeleteExceptionFromCalendar(calendar_id string, exception_id string) (*DeleteExceptionFromCalendarResponse, error) {
	var ret DeleteExceptionFromCalendarResponse
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id + "/exception/" + exception_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

// Retrieve an Exception having the specified exception_id on a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Get Exception on Door Exception Calendar]
//
// [Verkada API Docs - Get Exception on Door Exception Calendar]: https://apidocs.verkada.com/reference/getaccessdoorexceptionviewv1
func (c *AccessClient) GetExceptionFromCalendar(calendar_id string, exception_id string) (*DoorException, error) {
	var ret DoorException
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id + "/exception/" + exception_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Update an Exception having the specified exception_id on a Door Exception Calendar having the specified calendar_id.
//
// To use this API, you need an API Key having permission "Door Access Management" for all sites including doors listed in doors for the Door Exception Calendar.
// Manage via API must also be enabled for all doors listed in doors.
//
// [Verkada API Docs - Update Exception on Door Exception Calendar]
//
// [Verkada API Docs - Update Exception on Door Exception Calendar]: https://apidocs.verkada.com/reference/putaccessdoorexceptionviewv1
func (c *AccessClient) UpdateExceptionOnCalendar(calendar_id string, exception_id string, date string, start_time, end_time string, body *AddExceptionToCalendarBody) (*DoorException, error) {
	fullBody := DoorException{
		All_day_default:           body.All_day_default,
		Date:                      date,
		Door_status:               body.Door_status,
		Double_badge:              body.Double_badge,
		Double_badge_group_ids:    body.Double_badge_group_ids,
		End_time:                  end_time,
		First_person_in:           body.First_person_in,
		First_person_in_group_ids: body.First_person_in_group_ids,
		Recurrence_rule:           body.Recurrence_rule,
		Start_time:                start_time,
	}
	if success, err := validateDoorException(fullBody); !success {
		return nil, err
	}
	var ret DoorException
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id + "/exception/" + exception_id
	err := c.client.MakeVerkadaRequest("PUT", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Returns events for an organization within a specified time range.
//
// [Verkada API Docs - Get Access Events]
//
// [Verkada API Docs - Get Access Events]: https://apidocs.verkada.com/reference/geteventsviewv1
func (c *AccessClient) GetAccessEvents(options *GetAccessEventsOptions) (*GetAccessEventsResponse, error) {
	if options == nil {
		options = &GetAccessEventsOptions{}
	}
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	// notification_type validation
	event_type_validation := map[string]bool{
		"door_opened":                      true,
		"door_rejected":                    true,
		"door_granted":                     true,
		"door_forced_open":                 true,
		"door_held_open":                   true,
		"door_tailgating":                  true,
		"door_crowd_detection":             true,
		"door_tamper":                      true,
		"door_poi_detection":               true,
		"door_initialized":                 true,
		"door_armed":                       true,
		"door_armed_button_pressed":        true,
		"door_aux_unlock":                  true,
		"door_locked":                      true,
		"door_unlocked":                    true,
		"door_unarmed_event":               true,
		"door_code_entered_event":          true,
		"door_button_press_entered_event":  true,
		"door_lock_state_changed":          true,
		"door_lockdown":                    true,
		"door_auxinput_change_state":       true,
		"door_auxinput_held":               true,
		"door_low_battery":                 true,
		"door_critical_battery":            true,
		"door_mobile_nfc_scan_accepted":    true,
		"door_mobile_nfc_scan_rejected":    true,
		"door_user_database_corrupt":       true,
		"door_keycard_entered_accepted":    true,
		"door_keycard_entered_rejected":    true,
		"door_code_entered_accepted":       true,
		"door_code_entered_rejected":       true,
		"door_remote_unlock_accepted":      true,
		"door_remote_unlock_rejected":      true,
		"door_press_to_exit_accepted":      true,
		"door_ble_unlock_attempt_accepted": true,
		"door_ble_unlock_attempt_rejected": true,
		"door_acu_offline":                 true,
		"door_fire_alarm_triggered":        true,
		"door_fire_alarm_released":         true,
		"door_acu_fire_alarm_triggered":    true,
		"door_acu_fire_alarm_released":     true,
		"door_schedule_toggle":             true,
		"door_acu_dpi_cut":                 true,
		"door_acu_dpi_short":               true,
		"door_acu_rex_cut":                 true,
		"door_acu_rex_short":               true,
		"door_acu_rex2_cut":                true,
		"door_acu_rex2_short":              true,
		"door_acu_auxinput_cut":            true,
		"door_acu_auxinput_short":          true,
		"door_lockdown_debounced":          true,
		"door_lp_presented_accepted":       true,
		"door_lp_presented_rejected":       true,
		"door_apb_double_entry":            true,
		"door_apb_double_exit":             true,
		"all_access_granted":               true,
		"all_access_rejected":              true,
		"door_auxoutput_activated":         true,
		"door_auxoutput_deactivated":       true,
		"door_schedule_override_applied":   true,
		"door_schedule_override_removed":   true,
	}
	for _, param := range options.Event_type {
		if ok := event_type_validation[param]; !ok {
			return nil, fmt.Errorf("could not validate parameter in event_type: %s", param)
		}
	}
	var ret GetAccessEventsResponse
	url := c.client.baseURL + "/events/v1/access"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetAccessEventsResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Events = append(ret.Events, nextRet.Events...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

// Lists all access scenarios for the organization.
//
// To use this API, you need an API Key with "Read" permission for "Access Scenario Management" endpoints or an API Key with "Read" permission for "Access Control" endpoints.
//
// [Verkada API Docs - Get All Access Scenarios]
//
// [Verkada API Docs - Get All Access Scenarios]: https://apidocs.verkada.com/reference/getaccessscenariolistviewv1
func (c *AccessClient) GetAllAccessScenarios(options *GetAllAccessScenariosOptions) (*GetAllAccessScenariosResponse, error) {
	if options == nil {
		options = &GetAllAccessScenariosOptions{}
	}
	var ret GetAllAccessScenariosResponse
	url := c.client.baseURL + "/access/v1/scenarios"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Activates a scenario with the given scenario_id.
//
// To use this API, you need an API Key having "Read/Write" permission for "Access Control" endpoints.
// "Manage via API" must also be enabled for scenarios in your organization.
//
// [Verkada API Docs - Activate Access Scenario]
//
// [Verkada API Docs - Activate Access Scenario]: https://apidocs.verkada.com/reference/postaccessscenarioactivateviewv1
func (c *AccessClient) ActivateAccessScenario(scenario_id string) (*ActivateAccessScenarioResponse, error) {
	var ret ActivateAccessScenarioResponse
	url := c.client.baseURL + "/access/v1/scenarios/" + scenario_id + "/activate"
	err := c.client.MakeVerkadaRequest("POST", url, nil, nil, &ret, 0)
	return &ret, err
}

// Releases a scenario with the given scenario_id.
//
// To use this API, you need an API Key having "Read/Write" permission for "Access Control" endpoints.
// "Manage via API" must also be enabled for scenarios in your organization.
//
// [Verkada API Docs - Release Access Scenario]
//
// [Verkada API Docs - Release Access Scenario]: https://apidocs.verkada.com/reference/postaccessscenarioreleaseviewv1
func (c *AccessClient) DeactivateAccessScenario(scenario_id string) (*DeactivateAccessScenarioResponse, error) {
	var ret DeactivateAccessScenarioResponse
	url := c.client.baseURL + "/access/v1/scenarios/" + scenario_id + "/deactivate"
	err := c.client.MakeVerkadaRequest("POST", url, nil, nil, &ret, 0)
	return &ret, err
}

// Internally used for validation certain rules when creating/updating DoorException objects
func validateDoorException(exception DoorException) (bool, error) {
	// valdiating all_day_default rules
	if exception.All_day_default {
		if exception.Door_status != "access_controlled" {
			return false, fmt.Errorf("door_status must be \"access_controlled\" when all_day_default is true - received %s", exception.Door_status)
		}
		if !(exception.Start_time == "" && exception.End_time == "") {
			return false, fmt.Errorf("start_time and end_time both must be empty strings (not included) when all_day_default is true - received start_time%s and end_time:%s ", exception.Start_time, exception.End_time)
		}
		if exception.First_person_in || exception.Double_badge {
			return false, fmt.Errorf("first_person_in and double_badge must be false when all_day_default is true - received first_person_in: %v and double_badge: %v", exception.First_person_in, exception.Double_badge)
		}
	}
	// Door status must be one of the following:
	door_status_validation := map[string]bool{
		"":                  true,
		"locked":            true,
		"card_and_code":     true,
		"access_controlled": true,
		"unlocked":          true,
	}
	if ok := door_status_validation[exception.Door_status]; !ok {
		return false, fmt.Errorf("could not validate door_status: %s", exception.Door_status)
	}
	// validate double_badge rule
	if exception.Double_badge && exception.Door_status != "access_controlled" {
		return false, fmt.Errorf("door_status must be \"access_controlled\" when double_badge is true - received door_status: %s", exception.Door_status)
	}
	// validate double_badge_group_ids rule
	if len(exception.Double_badge_group_ids) > 0 && !exception.Double_badge {
		return false, fmt.Errorf("double_badge must be true if double_badge_group_ids is not empty")
	}
	// validate first_person_in rules
	if exception.First_person_in && !(exception.Door_status == "card_and_code" || exception.Door_status == "access_controlled" || exception.Door_status == "unlocked") {
		return false, fmt.Errorf("door_status must be \"card_and_code\", \"access_controlled\", or \"access_controlled\" when first_person_in is true - received %s", exception.Door_status)
	}
	// validate first_person_in_group_ids rule
	if len(exception.First_person_in_group_ids) > 0 && !exception.First_person_in {
		return false, fmt.Errorf("first_person_in must be true if first_person_in_group_ids is not empty")
	}
	return true, nil
}
