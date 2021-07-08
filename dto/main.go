package dto

type Main struct {
	Id         int
	Title      string
	WebUrl     string
	UserName   string
	PassWord   string
	Remark     string
	Folder     string
	UpdateTime string
	CreateTime string
	Hash       string
}
type Fload struct {
	Id   int
	Name string
}

type Version struct {
	Id          int
	Version     string
	LastVersion string
	CreatTime   string
}

var MainList = make([]*Main, 0)
var VersionList = make([]*Version, 0)
var FloadList = make([]*Fload, 0)
