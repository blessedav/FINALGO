package reqresp

type CreateCourseRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
}

type CreateCourseResponse struct {
	ID string `json:"id"`
}

type GetCourseResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
}

type ListCoursesResponse struct {
	Courses []GetCourseResponse `json:"courses"`
}

type SearchCoursesRequest struct {
	Query string `json:"query"`
}

type SearchCoursesResponse struct {
	Courses []GetCourseResponse `json:"courses"`
} 