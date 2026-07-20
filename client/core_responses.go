package client

type GetAuditLogsResponse struct {
	Audit_logs []struct {
		Details             any               `json:"details"`
		Devices             []AuditLogsDevice `json:"devices"`
		Event_description   string            `json:"event_description"`
		Event_name          string            `json:"event_name"`
		Ip_address          string            `json:"ip_address"`
		Organization_id     string            `json:"organization_id"`
		Processed_timestamp string            `json:"processed_timestamp"`
		Timestamp           string            `json:"timestamp"`
		User_email          string            `json:"user_email"`
		User_id             string            `json:"user_id"`
		User_name           string            `json:"user_name"`
		Verkada_support_id  string            `json:"verkada_support_id"`
	} `json:"audit_logs"`
	Next_page_token string `json:"next_page_token"`
}

type AuditLogsDevice struct {
	Details          any    `json:"details"`
	Device_id        string `json:"device_id"`
	Device_name      string `json:"device_name"`
	Device_site_name string `json:"device_site_name"`
	Device_type      string `json:"device_type"`
}

type DeleteUserResponse struct {
}

type GetUserResponse struct {
	Active         *bool  `json:"active"`
	Company_name   string `json:"company_name"`
	Department     string `json:"department"`
	Department_id  string `json:"department_id"`
	Email          string `json:"email"`
	Employee_id    string `json:"employee_id"`
	Employee_title string `json:"employee_title"`
	Employee_type  string `json:"employee_type"`
	External_id    string `json:"external_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Middle_name    string `json:"middle_name"`
	Phone          string `json:"phone"`
	User_id        string `json:"user_id"`
}

type CreateUserResponse struct {
	Active         *bool  `json:"active"`
	Company_name   string `json:"company_name"`
	Department     string `json:"department"`
	Department_id  string `json:"department_id"`
	Email          string `json:"email"`
	Employee_id    string `json:"employee_id"`
	Employee_title string `json:"employee_title"`
	Employee_type  string `json:"employee_type"`
	External_id    string `json:"external_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Middle_name    string `json:"middle_name"`
	Phone          string `json:"phone"`
	User_id        string `json:"user_id"`
}

type UpdateUserResponse struct {
	Active         *bool  `json:"active"`
	Company_name   string `json:"company_name"`
	Department     string `json:"department"`
	Department_id  string `json:"department_id"`
	Email          string `json:"email"`
	Employee_id    string `json:"employee_id"`
	Employee_title string `json:"employee_title"`
	Employee_type  string `json:"employee_type"`
	External_id    string `json:"external_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Middle_name    string `json:"middle_name"`
	Phone          string `json:"phone"`
	User_id        string `json:"user_id"`
}
