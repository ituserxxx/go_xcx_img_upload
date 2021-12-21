package service

import (
	"img-xcx-tool/app/dao"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/model"
)

var ImageCategory *imageCategory

type imageCategory struct{}

func (ic *imageCategory) Update(req *in_out.UpdateCategoryReq) error {
	_, err := dao.Category.OmitEmpty().Update(model.Category{
		UserId: req.Uid,
		Id: req.Cid,
		Name:   req.Name,
	})
	if err != nil {
		return err
	}
	return nil
}
func (ic *imageCategory) Add(req *in_out.AddCategoryReq) error {
	_, err := dao.Category.OmitEmpty().Save(model.Category{
		UserId: req.Uid,
		Name:   req.Name,
		Status: 1,
	})
	if err != nil {
		return err
	}
	return nil
}
