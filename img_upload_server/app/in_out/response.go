package in_out

type UserImgList struct {
	ImgList []ImgItem `json:"img_list"`
}
type ImgItem struct {
	PicList []ImgItemChild `json:"pic_list"`
	CreateTime string `json:"create_time"`
}

type ImgItemChild struct {
	ID int `json:"id"`
	ImgUrl string `json:"img_url"`
}


// TokenResponse .
type QiniuResponse struct {
	Token  string `json:"token"`
	Domain string `json:"domain"`
	Region string `json:"region"`
}
