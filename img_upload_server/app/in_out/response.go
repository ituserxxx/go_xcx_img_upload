package in_out

type UserImageListResp struct {

	ID int `json:"id"`
	ImgUrl string `json:"img_url"`
	CreateTime string `json:"create_time"`
	ImgCategoryId int `json:"img_category_id"`
	ImgCategoryIdDesc string `json:"img_category_id_desc"`
}

// TokenResponse .
type QiniuResponse struct {
	Token  string `json:"token"`
	Domain string `json:"domain"`
	Region string `json:"region"`
}
