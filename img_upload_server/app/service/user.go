package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"img-xcx-tool/app/dao"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/model"
	"img-xcx-tool/constant"
	"strings"
)

var User *user

type user struct{}

func (u *user) DelImg(req *in_out.DelImg) error {
	_, err := dao.UserImage.OmitEmpty().
		Where("user_id = ?", req.Uid).
		Where("status = ?", 1).
		Where("id = ?", req.Id).
		Data(model.UserImage{
			Status: 2,
		}).
		Update()
	if err != nil {
		return err
	}
	return nil
}
func (u *user) LoveImg(req *in_out.LoveImg) error {
	_, err := dao.UserImage.OmitEmpty().
		Where("user_id = ?", req.Uid).
		Where("status = ?", 1).
		Where("id = ?", req.Id).
		Data(model.UserImage{
			Status: 3,
		}).
		Update()
	if err != nil {
		return err
	}
	return nil
}
func (u *user) Add(req *in_out.AddUserImage) error {
	var list g.List
	for _, i2 := range req.ImgList {
		list = append(list, g.Map{
			"img_url":     i2,
			"user_id":     req.Uid,
			"create_time": gtime.Now().Format("Y-m-d H:i:s"),
			"status":      1,
		})
	}
	_, err := dao.UserImage.OmitEmpty().Data(list).Insert()
	if err != nil {
		return err
	}
	return nil
}
func (u *user) UserImageList(req *in_out.UserImageListReq) (in_out.UserImgList, error) {

	var dd in_out.UserImgList
	d, err := dao.UserImage.OmitEmpty().
		Order("id desc").
		Where("user_id = ?", req.Uid).
		Group("DATE_FORMAT(`create_time`, '%Y-%m-%d')").
		Page(req.Page, constant.PageMinList).
		FindAll()
	if err != nil {
		return dd, err
	}
	var list []model.UserImage
	if err := d.Structs(&list); err != nil {
		return dd, err
	}
	var data []in_out.ImgItem
	for _, v := range list {
		v1 := v
		l89 := dao.UserImage.OmitEmpty().
			Where("user_id = ?", req.Uid)
		if req.StatusId > 0 {
			l89 = l89.Where("status = ?", req.StatusId)
		}
		//1-正常，2-删除,3-喜欢
		l1, err := l89.Where("create_time >= ?", gtime.New(gconv.String(v1.CreateTime)).FormatTo("Y-m-d")).
			Where("create_time < ?", gtime.New(gconv.String(v1.CreateTime)).AddDate(0, 0, 1).FormatTo("Y-m-d")).
			Order("id desc").
			FindAll()
		if err != nil {
			return dd, err
		}
		var l2 []model.UserImage
		if err := l1.Structs(&l2); err != nil {
			return dd, err
		}
		var da in_out.ImgItem
		da.CreateTime = gtime.New(v.CreateTime).Format("Y-m-d")
		for _, v2 := range l2 {
			da.PicList = append(da.PicList, in_out.ImgItemChild{
				ID:     v2.Id,
				ImgUrl: v2.ImgUrl,
				IsVideo: func() bool {
					index := strings.LastIndex(v2.ImgUrl, ".")
					s := []rune(v2.ImgUrl)
					switch string(s[index:]) {
					case ".png", ".jpg", ".gif", ".jpeg":
						return false
					default:
						return true
					}
				}(),
			})
		}
		if len(da.PicList) > 0 {
			data = append(data, da)
		}
	}
	dd.ImgList = data
	return dd, nil
}
