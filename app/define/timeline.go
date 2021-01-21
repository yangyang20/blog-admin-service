package define

type TimelineCreate struct {
	Title   string `v:"required#请输入标题" json:"title"`
	Content string `v:"required#请输入内容" json:"content"`
	Image   string `v:"required#请上传图片" json:"image"`
}

type TimelinePut struct {
	Id      int    ` json:"id"`
	Title   string ` json:"title"`
	Content string ` json:"content"`
	Image   string ` json:"image"`
}
