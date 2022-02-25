package todo

import "github.com/Convenience-Tools/convenience-server/pkg/models/base"

type Todo struct {
	base.BaseModel `bson:",inline" json:",inline"`

	Title    string `json:"title" bson:"title"`
	Content  string `json:"content" bson:"content"`
	Status   string `json:"status" bson:"status"`
	UserName string `json:"username" bson:"username"`
}
