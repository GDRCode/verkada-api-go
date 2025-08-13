package client

import (
	"fmt"
	"strings"
)

func (c *AccessClient) GetAllAccessGroups() (*GetAllAccessGroupsResponse, error) {
	var ret GetAllAccessGroupsResponse
	url := c.client.baseURL + "/access/v1/access_groups"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) DeleteAccessGroups(group_id string) (*DeleteAccessGroupResponse, error) {
	options := &DeleteAccessGroupOptions{group_id: group_id}
	var ret DeleteAccessGroupResponse
	url := c.client.baseURL + "/access/v1/access_groups/group"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) GetAccessGroup(group_id string) (*AccessGroup, error) {
	options := &GetAccessGroupOptions{group_id: group_id}
	var ret AccessGroup
	url := c.client.baseURL + "/access/v1/access_groups/group"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) CreateAccessGroup(name string) (*AccessGroup, error) {
	body := &CreateAccessGroupBody{Name: name}
	var ret AccessGroup
	url := c.client.baseURL + "/access/v1/access_groups/group"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) GetAllAccessUsers() (*GetAllAccessUsersResponse, error) {
	var ret GetAllAccessUsersResponse
	url := c.client.baseURL + "/access/v1/access_users"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) ActivateLicensePlated(license_plate_number string, options *ActivateLicensePlateOptions) (*LicensePlate, error) {
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

func (c *AccessClient) DeactivateLicensePlated(license_plate_number string, options *DeactivateLicensePlateOptions) (*LicensePlate, error) {
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

func (c *AccessClient) GetAllAccessLevels() (*GetAllAccessLevelsResponse, error) {
	var ret GetAllAccessLevelsResponse
	url := c.client.baseURL + "/access/v1/door/access_level"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) DeleteAccessLevel(access_level_id string) (*DeleteAccessLevelResponse, error) {
	var ret DeleteAccessLevelResponse
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) GetAccessLevel(access_level_id string) (*AccessLevel, error) {
	var ret AccessLevel
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) DeleteAccessScheduleEvent(access_level_id string, event_id string) (*DeleteAccessScheduleEventResponse, error) {
	var ret DeleteAccessScheduleEventResponse
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id + "/access_schedule_event/" + event_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) GetAccessScheduleEvent(access_level_id string, event_id string) (*AccessScheduleEvent, error) {
	var ret AccessScheduleEvent
	url := c.client.baseURL + "/access/v1/door/access_level/" + access_level_id + "/access_schedule_event/" + event_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) GetAllDoorExceptionCalendars(options *GetAllDoorExceptionCalendarsOptions) (*GetAllDoorExceptionCalendarsResponse, error) {
	if options == nil {
		options = &GetAllDoorExceptionCalendarsOptions{}
	}
	var ret GetAllDoorExceptionCalendarsResponse
	url := c.client.baseURL + "/access/v1/door/exception_calendar"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) DeleteDoorExceptionCalendar(calendar_id string) (*DeleteDoorExceptionCalendarResponse, error) {
	var ret DeleteDoorExceptionCalendarResponse
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) GetDoorExceptionCalendar(calendar_id string) (*DoorExceptionCalendar, error) {
	var ret DoorExceptionCalendar
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) DeleteExceptionFromCalendar(calendar_id string, exception_id string) (*DeleteExceptionFromCalendarResponse, error) {
	var ret DeleteExceptionFromCalendarResponse
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id + "/exception/" + exception_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) GetExceptionFromCalendar(calendar_id string, exception_id string) (*DoorException, error) {
	var ret DoorException
	url := c.client.baseURL + "/access/v1/door/exception_calendar/" + calendar_id + "/exception/" + exception_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

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

func (c *AccessClient) GetAllAccessScenarios(options *GetAllAccessScenariosOptions) (*GetAllAccessScenariosResponse, error) {
	if options == nil {
		options = &GetAllAccessScenariosOptions{}
	}
	var ret GetAllAccessScenariosResponse
	url := c.client.baseURL + "/access/v1/scenarios"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) ActivateAccessScenario(scenario_id string) (*ActivateAccessScenarioResponse, error) {
	var ret ActivateAccessScenarioResponse
	url := c.client.baseURL + "/access/v1/scenarios/" + scenario_id + "/activate"
	err := c.client.MakeVerkadaRequest("POST", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *AccessClient) DeactivateAccessScenario(scenario_id string) (*DeactivateAccessScenarioResponse, error) {
	var ret DeactivateAccessScenarioResponse
	url := c.client.baseURL + "/access/v1/scenarios/" + scenario_id + "/deactivate"
	err := c.client.MakeVerkadaRequest("POST", url, nil, nil, &ret, 0)
	return &ret, err
}

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
