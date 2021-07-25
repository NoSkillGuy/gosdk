package errors

/* Is - tells whether actual error is targer error
where, actual error can be either Error/withError
if actual error is wrapped error then if any internal error
matches the target error then function results in true
*/
func Is(actual error, target *Error) bool {
	actualError := isError(actual)
	if actualError != nil {
		if actualError.Code == "" && target.Code == "" {
			return actualError.Msg == target.Msg
		} else {
			return actualError.Code == target.Code
		}
	} else {
		actualWithError := isWithError(actual)
		if actualWithError != nil {
			return Is(actualWithError.current, target) || Is(actualWithError.previous, target)
		} else {
			return false
		}
	}
}
