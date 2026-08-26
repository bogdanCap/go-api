package validation

var ProductValidationMessages = map[string]map[string]string{
	"Age": {
		"gte": "age must be greater than 18",
	},
	"Name": {
		"min": "name must contain at least 3 characters",
	},
}
