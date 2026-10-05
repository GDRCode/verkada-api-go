package client

type GetEventsOptions struct {
	Start_time    string   `json:"start_time"`
	End_time      string   `json:"end_time"`
	Product_types []string `json:"product_types"`
	Device_ids    []string `json:"device_ids"`
	Event_types   []string `json:"event_types"`
	Site_ids      []string `json:"site_ids"`
	Limit         *int     `json:"limit"`
	Cursor        string   `json:"cursor"`
}
