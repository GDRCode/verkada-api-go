package client

type GetAllAccessGroupsResponse struct {
	Access_groups []AccessGroupMetadata `json:"access_groups"`
}

type DeleteAccessGroupResponse struct {
}

type RemoveUserFromAccessGroupResponse struct {
}

type AddUserToAccessGroupResponse struct {
}

type GetAllAccessUsersResponse struct {
	Access_members []AccessUser `json:"access_members"`
}

type RemoveUserEntryCodeResponse struct {
}

type DeleteProfilePhotoResponse struct {
}

type CreateProfilePhotoResponse struct {
}

type DeleteAccessCardResponse struct {
}

type DeleteUserLicensePlateResponse struct {
}

type DeleteMFACodeResponse struct {
}

type DisableFaceUnlockExternalResponse struct {
}

type EnableFaceUnlockExternalProfilePhotoResponse struct {
	Identity_id string `json:"identity_id"`
	User_id     string `json:"user_id"`
}

type EnableFaceUnlockExternalMobileInviteResponse struct {
	Invitations_failed []string `json:"invitations_failed"`
	Invitations_sent   []string `json:"invitations_sent"`
	Live_link_id       string   `json:"live_link_id"`
	User_id            string   `json:"user_id"`
}

type EnableFaceUnlockExternalUploadPhotoResponse struct {
	Identity_id string `json:"identity_id"`
	User_id     string `json:"user_id"`
}

type DisableFaceUnlockInternalResponse struct {
}

type EnableFaceUnlockInternalProfilePhotoResponse struct {
	Identity_id string `json:"identity_id"`
	User_id     string `json:"user_id"`
}

type EnableFaceUnlockInternalMobileInviteResponse struct {
	Invitations_failed []string `json:"invitations_failed"`
	Invitations_sent   []string `json:"invitations_sent"`
	Live_link_id       string   `json:"live_link_id"`
	User_id            string   `json:"user_id"`
}

type EnableFaceUnlockInternalUploadPhotoResponse struct {
	Identity_id string `json:"identity_id"`
	User_id     string `json:"user_id"`
}

type AddMFACodeResponse struct {
	Code string `json:"code"`
}

type GetAllAccessLevelsResponse struct {
	Access_levels []AccessLevel `json:"access_levels"`
}

type DeleteAccessLevelResponse struct {
}

type DeleteAccessScheduleEventResponse struct {
}

type AdminUnlockDoorResponse struct {
	Door_id         string `json:"door_id"`
	Unlock_duration int    `json:"unlock_duration"`
}

type UserUnlockDoorResponse struct {
	Door_id         string `json:"door_id"`
	Unlock_duration int    `json:"unlock_duration"`
}

type GetDoorsResponse struct {
	Doors []Door `json:"doors"`
}

type GetAllDoorExceptionCalendarsResponse struct {
	Door_exception_calendars []DoorExceptionCalendar `json:"door_exception_calendars"`
}

type DeleteDoorExceptionCalendarResponse struct {
}

type DeleteExceptionFromCalendarResponse struct {
}

type GetAccessEventsResponse struct {
	Events          []Events `json:"events"`
	Next_page_token string   `json:"next_page_token"`
}

type GetAllAccessScenariosResponse struct {
	Scenarios []AccessScenario `json:"scenarios"`
}

type ActivateAccessScenarioResponse struct {
}

type DeactivateAccessScenarioResponse struct {
}
