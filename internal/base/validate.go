package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

// Если какое-либо поле пустое, добавляется сообщение об ошибке.
func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil || req.UserID == "" || req.Title == "" || req.Description == "" {
		res = append(res, "validate message")
	}
	return res
}
