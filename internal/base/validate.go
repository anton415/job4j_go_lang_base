package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

// Если какое-либо поле пустое, добавляется сообщение об ошибке.
func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil {
		res = append(res, "request is nil")
		return res
	}
	if req.UserID == "" {
		res = append(res, "user id is empty")
	}
	if req.Title == "" {
		res = append(res, "title is empty")
	}
	if req.Description == "" {
		res = append(res, "description is empty")
	}
	return res
}
