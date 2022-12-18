package course

type CreateCourseSchema struct {
	Name    string `json:"name" binding:"required"`
	Teacher *int   `json:"teacher"`
}
type UpdateCourseSchema struct {
	Name    *string `json:"name"`
	Teacher *int    `json:"teacher"`
}

type ManageCourseMemberSchema struct {
	Students []int `json:"students"`
}

type ManageCourseMemberResult struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}
