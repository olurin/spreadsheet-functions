package formulas

//FunctionError struct contains the error type, field error
type FunctionError struct {
	Error   string
	Type    string
	Formula string
}

//FormulaErrors is the type of map[string]*FunctionError
type FormulaErrors map[string]*FunctionError

// ErrorExport function
func ErrorExport(code int) string {
	switch code {
	case 1:
		return "#NULL!"
	case 2:
		return "#DIV/0!"
	case 3:
		return "#VALUE?"
	case 4:
		return "#REF!"
	case 5:
		return "#NAME?"
	case 6:
		return "#NUM!"
	case 7:
		return "#ERROR!"
	case 8:
		return "GETTING_DATA"
	default:
		return "#N/A"
	}
}
