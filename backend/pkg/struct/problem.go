package structure

type CreateProblemRequest struct {
	UserId      string `json:"user_id" binding:"required"`
	Problem     string `json:"problem" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateProblemResponse struct {
	ProblemId   string `json:"problem_id"`
	UserId      string `json:"user_id"`
	Problem     string `json:"problem"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type GetProblemDetailByIdResponse struct {
	ProblemId     string `json:"problem_id"`
	AdminUsername string `json:"admin_username"`
	Problem       string `json:"problem"`
	Description   string `json:"description"`
	Reply         string `json:"reply"`
	Status        string `json:"status"`
}

type ProblemList struct {
	ProblemId     string `json:"problem_id"`
	AdminUsername string `json:"admin_username"`
	Problem       string `json:"problem"`
	Description   string `json:"description"`
	Reply         string `json:"reply"`
	Status        string `json:"status"`
}

type GetProblemListsRequest struct {
	UserId string `json:"user_id"`
	Status string `json:"status"`
}

type GetProblemListsResponse struct {
	ProblemLists []ProblemList `json:"problem_lists"`
}

type UpdateProblemRequest struct {
	ProblemId     string  `json:"problem_id" binding:"required"`
	AdminUsername *string `json:"admin_username"`
	Problem       string  `json:"problem"`
	Description   string  `json:"description"`
	Reply         *string `json:"reply"`
	Status        string  `json:"status"`
}
