package service

import (
	"encoding/json"
	"pear-admin-golang/app/core/cache"
	"pear-admin-golang/app/dao"
	e2 "pear-admin-golang/app/global/e"
	"pear-admin-golang/app/model"
)

func GetPearConfig() (*model.PearConfigForm, error) {
	pear, err := dao.NewSiteConfigDaoImpl().FindOne(model.PearSiteConfig)
	if err != nil {
		return nil, err
	}
	var data model.PearConfigForm
	err = json.Unmarshal([]byte(pear.ConfigData), &data)
	if err != nil {
		return nil, err
	}
	cache.Instance().Set(e2.PearConfigCache, data, 0)
	return &data, nil
}
