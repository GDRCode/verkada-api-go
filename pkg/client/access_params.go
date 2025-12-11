package client

type DeleteAccessGroupOptions struct {
	group_id string `name:"group_id"`
}

type GetAccessGroupOptions struct {
	group_id string `name:"group_id"`
}

type CreateAccessGroupBody struct {
	Name string `json:"name"`
}

type RemoveUserFromAccessGroupOptions struct {
	group_id    string `name:"group_id"`
	External_id string `name:"external_id"`
	User_id     string `name:"user_id"`
}

type AddUserToAccessGroupOptions struct {
	group_id string `name:"group_id"`
}

type AddUserToAccessGroupBody struct {
	External_id string `json:"external_id,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}

type GetAccessInformationObjectOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type ActivateUserBLEOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type DeactivateUserBLEOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type SetUserEndDateOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type RemoveUserEntryCodeOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type SetUserEntryCodeOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
	Override    *bool  `name:"override"`
}

type SendPassInviteOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type DeleteProfilePhotoOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type GetProfilePhotoOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Original    *bool  `name:"original"`
}

type UploadProfilePhotoOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Overwrite   *bool  `name:"overwrite"`
}

type ActivateUserRemoteUnlockOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type DeactivateUserRemoteUnlockOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type SetStartDateOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	Email       string `name:"email"`
	Employee_id string `name:"employee_id"`
}

type DeleteAccessCardOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	card_id     string `name:"card_id"`
}

type AddAccessCardOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type AddAccessCardBody struct {
	Active             bool
	Card_number        string
	Card_number_base36 string
	Card_number_hex    string
	Facility_code      string
}

type ActivateAccessCardOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	card_id     string `name:"card_id"`
}

type DeactivateAccessCardOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	card_id     string `name:"card_id"`
}

type DeleteUserLicensePlateOptions struct {
	User_id              string `name:"user_id"`
	External_id          string `name:"external_id"`
	license_plate_number string `name:"license_plate_number"`
}

type AddUserLicensePlateOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type AddUserLicensePlatedBody struct {
	Active bool
	Name   string
}

type ActivateLicensePlateOptions struct {
	User_id              string `name:"user_id"`
	External_id          string `name:"external_id"`
	License_plate_number string `name:"license_plate_number"`
}

type DeactivateLicensePlateOptions struct {
	User_id              string `name:"user_id"`
	External_id          string `name:"external_id"`
	License_plate_number string `name:"license_plate_number"`
}

type DeleteMFACodeOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
	code        string `name:"code"`
}

type AddMFACodeOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type UserUnlockDoorOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type GetDoorsOptions struct {
	Door_ids []string `name:"door_ids"`
	Site_ids []string `name:"site_ids"`
}

type GetAllDoorExceptionCalendarsOptions struct {
	Last_updated_after *int `name:"last_updated_after"`
}

type CreateDoorExceptionCalendarBody struct {
	Doors      []string
	Exceptions []DoorException
}

type UpdateDoorExceptionCalendarBody struct {
	Doors      []string
	Exceptions []DoorException
}

type AddExceptionToCalendarBody struct {
	All_day_default           bool
	Door_status               string
	Double_badge              bool
	Double_badge_group_ids    []string
	First_person_in           bool
	First_person_in_group_ids []string
	Recurrence_rule           *RecurrenceRule
}

type GetAccessEventsOptions struct {
	Start_time *int     `name:"start_time"`
	End_time   *int     `name:"end_time"`
	Page_token string   `name:"page_token"`
	Page_size  *int     `name:"page_size"`
	Event_type []string `name:"event_type"`
	Site_id    string   `name:"site_id"`
	Device_id  string   `name:"device_id"`
	User_id    string   `name:"user_id"`
}

type GetAllAccessScenariosOptions struct {
	Scenario_ids []string `name:"scenario_ids"`
	Site_ids     []string `name:"site_ids"`
	Types        []string `name:"types"`
}
