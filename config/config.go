package config

type (
	Config struct {
		SchemaImport   string `json:"schema_import"`
		WriteToPath    string `json:"write_to_path"`
		ModelName      string `json:"model_name"`
		LowerModelName string `json:"lower_model_name"`
		Fields         []Desc `json:"fields"`
	}
	Desc struct {
		LowerName     string `json:"lower_name"`
		PublicName    string `json:"public_name"`
		HasValidators bool   `json:"has_validators"`
		Comment       string `json:"comment"`
		Type          string `json:"type"`
	}
)
