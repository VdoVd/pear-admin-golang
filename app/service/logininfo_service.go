package service

import (
	"pear-admin-golang/app/core/log"
	dao2 "pear-admin-golang/app/dao"

	"pear-admin-golang/app/global/request"
	"pear-admin-golang/app/model"
)

func LoginInfoListJsonService(f request.LayerListForm) (count int, list []model.LoginInfo, err error) {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}
	list, count, err = dao2.NewLoginInfoImpl().FindByPage(f.Page, f.Limit)
	if err != nil {
		log.Instance().Error("LoginInfoListJsonService.FindByPage:" + err.Error())
		return 0, nil, err
	}
	return count, list, nil
}
