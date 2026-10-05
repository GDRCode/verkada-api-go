package client

import "encoding/json"

type GetEventsResponse struct {
	Cursor string  `json:"cursor"`
	Items  []Event `json:"items"`
}

type Event struct {
	Device_id   string          `json:"device_id"`
	Device_type string          `json:"device_type"`
	Event_id    string          `json:"event_id"`
	Event_info  json.RawMessage `json:"event_info"`
	Event_type  string          `json:"event_type"`
	Timestamp   string          `json:"timestamp"`
}

type CameraEvent struct {
	Crowd_threshold                  int      `json:"crowd_threshold"`
	Detected_objects                 []string `json:"detected_objects"`
	Event_description                string   `json:"event_description"`
	License_plate_number             string   `json:"license_plate_number"`
	License_plate_state              string   `json:"license_plate_state"`
	Location                         string   `json:"location"`
	Location_lat                     string   `json:"location_lat"`
	Location_lon                     string   `json:"location_lon"`
	Person_label                     string   `json:"person_label"`
	Rounded_inactivity_duration_mins string   `json:"rounded_inactivity_duration_mins"`
}

type SensorEvent struct {
	End_time           string  `json:"end_time"`
	Most_extreme_value float64 `json:"most_extreme_value"`
	Reading            string  `json:"reading"`
	Threshold          float64 `json:"threshold"`
}

type AccessEventInfo struct {
	Accepted               bool           `json:"accepted"`
	Access_controller_id   string         `json:"access_controller_id"`
	Access_controller_name string         `json:"access_controller_name"`
	Aux_input_id           string         `json:"aux_input_id"`
	Aux_input_name         string         `json:"aux_input_name"`
	Building_id            string         `json:"building_id"`
	Building_name          string         `json:"building_name"`
	Direction              string         `json:"direction"`
	Door_id                string         `json:"door_id"`
	Door_name              string         `json:"door_name"`
	Entity_id              string         `json:"entity_id"`
	Entity_name            string         `json:"entity_name"`
	Entity_type            string         `json:"entity_type"`
	Floor_id               string         `json:"floor_id"`
	Floor_name             string         `json:"floor_name"`
	Floors                 []Floor        `json:"floors"`
	Granted                bool           `json:"granted"`
	Input_value            string         `json:"input_value"`
	Message                string         `json:"message"`
	Raw_card               string         `json:"raw_card"`
	Scenario_info          AccessScenario `json:"scenario_info"`
	User_info              UserInfo       `json:"user_info"`
}

type Floor struct {
	Door_id  string `json:"door_id"`
	Floor_id string `json:"floor_id"`
	Uuid     string `json:"uuid"`
}

type UserInfo struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	User_id    string `json:"user_id"`
}

type AccessScenario struct {
	Action   string `json:"action"`
	Scenario struct {
		Enabled     bool   `json:"enabled"`
		Message     string `json:"message"`
		Name        string `json:"name"`
		Scenario_id string `json:"scenario_id"`
	} `json:"scenario"`
}

type IntercomEventInfo struct {
	Answered_by_name   string `json:"answered_by_name"`
	Answered_timestamp string `json:"answered_timestamp"`
	Device_name        string `json:"device_name"`
	End_timestamp      string `json:"end_timestamp"`
	Start_timestamp    string `json:"start_timestamp"`
}
