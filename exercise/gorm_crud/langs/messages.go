package langs

import "fmt"

func GenerateValidationMessage(field string, rule string) (message string) {
	switch rule {
	// required rule
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	//you can add more validator.v10 rule here
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}
