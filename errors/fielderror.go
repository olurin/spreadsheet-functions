package fielderror

//ErrorType struct contains the error type, field error
type ErrorType struct {
	Error   string
	Type    string
	Formula string
}

// ErrorCode function
// 1 = #NULL!
// 2 = #DIV/0!
// 3 = #VALUE!
// 4 = #REF!
// 5 = #NAME?
// 6 = #NUM!
// default = #N/A
func ErrorCode(code int) string {
	switch code {
	case 1:
		return "#NULL!"
	case 2:
		return "#DIV/0!"
	case 3:
		return "#REF!"
	case 4:
		return "#NAME?"
	case 5:
		return "#NUM!"
	case 6:
		return "#NUM!"
	default:
		return "#N/A"
	}
}
