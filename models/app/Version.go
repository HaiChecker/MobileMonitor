package app

import (
	"MobileMonitor/models/phone"
	"github.com/astaxie/beego/orm"
)

type Version struct {
	Id int `orm:"pk;auto"`
	App *App `orm:"rel(fk)"`    //设置一对多关系
	VersionCode int

	IpAddr *phone.IpAddr `orm:"rel(one)"`
}

func init() {
	orm.RegisterModel(new(Version))
}

/**
 *	查询版本，通过版本号和AppId
 */
func SelectVersion(versionCode int,appId int) (*Version,error) {
	o := orm.NewOrm()
	var version *Version
	err := o.QueryTable(version).Filter("version_code",versionCode).Filter("app_Id",appId).RelatedSel().One(version)
	return version,err
}