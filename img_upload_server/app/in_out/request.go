package in_out


type LoginReq struct {
	Code string `json:"code" v:"required"`
	Str  string `json:"str"`
}

type UidPageListReq struct {
	Uid int `json:"uid" v:"required"`
	Page int `json:"page" d:"1"`
}

type AddCategoryReq struct {
	Uid int `json:"uid" v:"required"`
	Name string `json:"name" v:"required"`
}
type UpdateCategoryReq struct {
	Uid int `json:"uid" v:"required"`
	Cid int `json:"cid" v:"required"`
	Name string `json:"name" v:"required"`
}

type UserImageListReq struct {
	Uid int `json:"uid" v:"required"`
	StatusId int `json:"status_id" d:"1"`//1-正常，2-删除,3-喜欢的
	Page int `json:"page" d:"1"`
}

type AddUserImage struct {
	Uid int `json:"uid" v:"required"`
	ImgList []string `json:"img_list"`
}

type DelImg struct {
	Uid int `json:"uid" v:"required"`
	Id int `json:"id" v:"required"`
}
