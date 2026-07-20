package client

type GetAuditLogsOptions struct {
	Start_time              *int   `name:"start_time"`
	End_time                *int   `name:"end_time"`
	Use_processed_timestamp *bool  `name:"use_processed_timestamp"`
	Page_token              string `name:"page_token"`
	Page_size               *int   `name:"page_size"`
}

type DeleteUserOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type GetUserOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type CreateUserBody struct {
	Company_name   string `json:"company_name,omitempty"`
	Department     string `json:"department,omitempty"`
	Department_id  string `json:"department_id,omitempty"`
	Email          string `json:"email,omitempty"`
	Employee_id    string `json:"employee_id,omitempty"`
	Employee_title string `json:"employee_title,omitempty"`
	Employee_type  string `json:"employee_type,omitempty"`
	External_id    string `json:"external_id,omitempty"`
	First_name     string `json:"first_name,omitempty"`
	Last_name      string `json:"last_name,omitempty"`
	Middle_name    string `json:"middle_name,omitempty"`
	Phone          string `json:"phone,omitempty"`
}

type UpdateUserOptions struct {
	User_id     string `name:"user_id"`
	External_id string `name:"external_id"`
}

type UpdateUserBody struct {
	Company_name   string `json:"company_name,omitempty"`
	Department     string `json:"department,omitempty"`
	Department_id  string `json:"department_id,omitempty"`
	Email          string `json:"email,omitempty"`
	Employee_id    string `json:"employee_id,omitempty"`
	Employee_title string `json:"employee_title,omitempty"`
	Employee_type  string `json:"employee_type,omitempty"`
	External_id    string `json:"external_id,omitempty"`
	First_name     string `json:"first_name,omitempty"`
	Last_name      string `json:"last_name,omitempty"`
	Middle_name    string `json:"middle_name,omitempty"`
	Phone          string `json:"phone,omitempty"`
}
