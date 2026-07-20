package client

type AccessInformationObject struct {
	Access_groups   []AccessGroup  `json:"access_groups"`
	Ble_unlock      bool           `json:"ble_unlock"`
	Cards           []Card         `json:"cards"`
	End_date        string         `json:"end_date"`
	Entry_code      string         `json:"entry_code"`
	External_id     string         `json:"external_id"`
	HasProfilePhoto bool           `json:"has_profile_photo"`
	License_plates  []LicensePlate `json:"license_plates"`
	Mfa_codes       []struct {
		Code string `json:"code"`
	} `json:"mfa_codes"`
	Profile_photo_last_updated string `json:"profile_photo_last_updated"`
	Remote_unlock              bool   `json:"remote_unlock"`
	Start_date                 string `json:"start_date"`
	User_id                    string `json:"user_id"`
}

type AccessGroupMetadata struct {
	Group_id string `json:"group_id"`
	Name     string `json:"name"`
}

type AccessGroup struct {
	Group_id string   `json:"group_id"`
	Name     string   `json:"name"`
	User_ids []string `json:"user_ids"`
}

type Card struct {
	Active             bool   `json:"active"`
	Card_id            string `json:"card_id"`
	Card_number        string `json:"card_number"`
	Card_number_base36 string `json:"card_number_base36"`
	Card_number_hex    string `json:"card_number_hex"`
	Facility_code      string `json:"facility_code"`
	Type               string `json:"type"`
}

type LicensePlate struct {
	Active               bool   `json:"active"`
	License_plate_number string `json:"license_plate_number"`
	Name                 string `json:"name"`
}

type AccessGroupAddition struct {
	Group_id          string   `json:"group_id"`
	Name              string   `json:"name"`
	Successful_adds   []string `json:"successful_adds"`
	Unsuccessful_adds []string `json:"unsuccessful_adds"`
}

type AccessUser struct {
	Company_name               string `json:"company_name"`
	Department                 string `json:"department"`
	Department_id              string `json:"department_id"`
	Email                      string `json:"email"`
	Employee_id                string `json:"employee_id"`
	Employee_title             string `json:"employee_title"`
	External_id                string `json:"external_id"`
	Full_name                  string `json:"full_name"`
	Has_profile_photo          bool   `json:"has_profile_photo"`
	Profile_photo_last_updated string `json:"profile_photo_last_updated"`
	User_id                    string `json:"user_id"`
}

type Door struct {
	Acu_name string `json:"acu_name"`
	Acu_id   string `json:"acu_id"`
	Name     string `json:"name"`
	Door_id  string `json:"door_id"`
	Site     Site   `json:"site"`
}

type Site struct {
	Name    string `json:"name"`
	Site_id string `json:"site_id"`
}

type AccessLevel struct {
	Access_groups          []string              `json:"access_groups"`
	Access_level_id        string                `json:"access_level_id,omitempty"`
	Access_schedule_events []AccessScheduleEvent `json:"access_schedule_events"`
	Doors                  []string              `json:"doors"`
	Last_updated_at        int64                 `json:"last_updated_at,omitempty"`
	Name                   string                `json:"name"`
	Sites                  []string              `json:"sites"`
}

type AccessScheduleEvent struct {
	Access_schedule_event_id string `json:"access_schedule_event_id"`
	Door_status              string `json:"door_status,omitempty"`
	End_time                 string `json:"end_time"`
	Start_time               string `json:"start_time"`
	Weekday                  string `json:"weekday"`
}

type DoorExceptionCalendar struct {
	Door_exception_calendar_id string          `json:"door_exception_calendar_id,omitempty"`
	Doors                      []string        `json:"doors"`
	Exceptions                 []DoorException `json:"exceptions"`
	Last_updated_at            int64           `json:"last_updated_at,omitempty"`
	Name                       string          `json:"name"`
}

type DoorException struct {
	All_day_default           bool            `json:"all_day_default,omitempty"`
	Calendar_id               string          `json:"calendar_id,omitempty"`
	Date                      string          `json:"date"`
	Door_exception_id         string          `json:"door_exception_id,omitempty"`
	Door_status               string          `json:"door_status"`
	Double_badge              bool            `json:"double_badge,omitempty"`
	Double_badge_group_ids    []string        `json:"double_badge_group_ids,omitempty"`
	End_time                  string          `json:"end_time"`
	First_person_in           bool            `json:"first_person_in,omitempty"`
	First_person_in_group_ids []string        `json:"first_person_in_group_ids,omitempty"`
	Recurrence_rule           *RecurrenceRule `json:"recurrence_rule,omitempty"`
	Start_time                string          `json:"start_time"`
}

type RecurrenceRule struct {
	By_day         []string `json:"by_day"`
	By_month       int      `json:"by_month"`
	By_month_day   int      `json:"by_month_day"`
	By_set_pos     int      `json:"by_set_pos"`
	Count          int      `json:"count"`
	Excluded_Dates []string `json:"excluded_dates"`
	Frequency      string   `json:"frequency"`
	Interval       *int     `json:"interval"`
	Until          string   `json:"until"`
}

type AccessScenario struct {
	Doors_to_lock                   []string `json:"doors_to_lock"`
	Message                         string   `json:"message"`
	Name                            string   `json:"name"`
	Scenario_id                     string   `json:"scenario_id"`
	Site_count                      int      `json:"site_count"`
	Sites                           []string `json:"sites"`
	State                           string   `json:"state"`
	Type                            string   `json:"type"`
	User_groups_with_disable_access []string `json:"user_groups_with_disable_access"`
	User_groups_with_door_access    []string `json:"user_groups_with_door_access"`
	User_groups_with_enable_access  []string `json:"user_groups_with_enable_access"`
}

type DoorInfo struct {
	Access_controller_id   string `json:"accessControllerId"`
	Access_controller_name string `json:"accessControllerName"`
	Name                   string `json:"name"`
}

type Floors struct {
	Door_id  string `json:"doorId"`
	Floor_id string `json:"floorId"`
	Uuid     string `json:"uuid"`
}

type Lockdown struct {
	Enabled     bool   `json:"enabled"`
	Lockdown_id string `json:"lockdownId"`
	Message     string `json:"message"`
	Name        string `json:"name"`
}

type LockdownInfo struct {
	Action   string   `json:"action"`
	Lockdown Lockdown `json:"lockdown"`
}

type UserInfo struct {
	Email           string `json:"email"`
	First_name      string `json:"firstName"`
	Last_name       string `json:"lastName"`
	Name            string `json:"name"`
	Organization_id string `json:"organizationId"`
	Phone           string `json:"phone"`
	User_id         string `json:"userId"`
}

type EventInfo struct {
	Accepted        bool         `json:"accepted"`
	Aux_input_id    string       `json:"auxInputId"`
	Aux_input_name  string       `json:"auxInputName"`
	Building_id     string       `json:"buildingId"`
	Building_name   string       `json:"buildingName"`
	Direction       string       `json:"direction"`
	Door_id         string       `json:"doorId"`
	Door_info       DoorInfo     `json:"doorInfo"`
	Entity_id       string       `json:"entityId"`
	Entity_name     string       `json:"entityName"`
	Entity_type     string       `json:"entityType"`
	Event_type      string       `json:"eventType"`
	Floor_id        string       `json:"floorId"`
	Floor_name      string       `json:"floorName"`
	Floors          []Floors     `json:"floors"`
	Input_value     string       `json:"inputValue"`
	Lockdown_info   LockdownInfo `json:"lockdownInfo"`
	Message         string       `json:"message"`
	Organization_id string       `json:"organizationId"`
	Raw_card        string       `json:"rawCard"`
	Site_id         string       `json:"siteId"`
	Site_name       string       `json:"siteName"`
	Type            string       `json:"type"`
	User_id         string       `json:"userId"`
	User_info       UserInfo     `json:"userInfo"`
	User_name       string       `json:"userName"`
	Uuid            string       `json:"uuid"`
}

type Events struct {
	Device_id       string    `json:"device_id"`
	Device_type     string    `json:"device_type"`
	End_timestamp   string    `json:"end_timestamp"`
	Event_id        string    `json:"event_id"`
	Event_info      EventInfo `json:"event_info"`
	Event_type      string    `json:"event_type"`
	Organization_id string    `json:"organization_id"`
	Site_id         string    `json:"site_id"`
	Timestamp       string    `json:"timestamp"`
}
