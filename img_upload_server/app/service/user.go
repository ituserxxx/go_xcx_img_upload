package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"img-xcx-tool/app/dao"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/model"
	"img-xcx-tool/constant"
)

var User *user

type user struct{}

func (u *user) Add(req *in_out.AddUserImage) error {
	var list g.List
	for _, i2 := range req.ImgList {
		list = append(list, g.Map{
			"img_url":     i2,
			"user_id":     req.Uid,
			"create_time": gtime.Now(),
			"status":      1,
		})
	}
	_, err := dao.UserImage.OmitEmpty().Data(list).Insert()
	if err != nil {
		return err
	}
	return nil
}
func (u *user) UserImageList(req *in_out.UserImageListReq) ([]in_out.UserImageListResp, error) {
	m := dao.UserImage.OmitEmpty().
		Order("id desc").
		Where("user_id = ?", req.Uid).
		Page(req.Page, constant.PageMinList)
	if req.Cid != 0 {
		m = m.Where("img_category_id =?", req.Cid)
	}
	l, err := m.FindAll()
	if err != nil {
		return nil, err
	}
	var list []model.UserImage
	if err := l.Structs(&list); err != nil {
		return nil, err
	}
	var data []in_out.UserImageListResp
	for _, i2 := range list {
		categroyNamme, err := dao.Category.OmitEmpty().Where("id=?", i2.ImgCategoryId).Value("name")
		if err != nil {
			return nil, err
		}
		data = append(data, in_out.UserImageListResp{
			ID:                i2.Id,
			ImgUrl:            i2.ImgUrl,
			CreateTime:        gconv.String(i2.CreateTime),
			ImgCategoryId:     i2.ImgCategoryId,
			ImgCategoryIdDesc: categroyNamme.String(),
		})
	}
	return data, nil
}
