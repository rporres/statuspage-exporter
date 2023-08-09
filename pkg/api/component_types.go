package api

type ComponentGroup struct {
	ID                 string `json:"id"`
	PageID             string `json:"page_id"`
	GroupID            string `json:"group_id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Group              bool   `json:"group"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Position           int    `json:"position"`
	Status             string `json:"status"`
	Showcase           bool   `json:"showcase"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded"`
	AutomationEmail    string `json:"automation_email"`
	StartDate          string `json:"start_date"`
}

type ComponentGroups []ComponentGroup

// Component was auto-generated by rea
type Component struct {
	AutomationEmail    string  `json:"automation_email"`
	CreatedAt          string  `json:"created_at"`
	Description        string  `json:"description"`
	Group              bool    `json:"group"`
	GroupId            string  `json:"group_id"`
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	OnlyShowIfDegraded bool    `json:"only_show_if_degraded"`
	PageId             string  `json:"page_id"`
	Position           float64 `json:"position"`
	Showcase           bool    `json:"showcase"`
	StartDate          string  `json:"start_date"`
	Status             string  `json:"status"`
	UpdatedAt          string  `json:"updated_at"`
}

// Components is multiple Component
type Components []Component
