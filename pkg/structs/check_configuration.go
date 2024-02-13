package structs

type CheckConfiguration struct {
	Config         map[string]interface{} `json:"config"`
	EditableFields []string               `json:"editable_fields"`
}
