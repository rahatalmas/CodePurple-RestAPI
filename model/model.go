package model 

type Comment struct{
	Id       int      `json:"id"`
    UserId   int      `json:"userId"`
	Comment  string   `json:"comment"`
}

type Post struct{
	Id        int        `json:"id"`
	UserId    int        `json:"userId"`
	Title     string     `json:"title"`
    Content   string     `json:"content`
	Comments  []Comment  `json:"comments"`
}
