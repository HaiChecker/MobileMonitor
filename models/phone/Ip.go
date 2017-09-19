package phone

import (
	"MobileMonitor/models/app"
	"github.com/astaxie/beego/orm"
	"errors"
)

type IpAddr struct {
	Id int `orm:"pk;auto"`
	IdAddr string
	Location string
	
	App *app.App `orm:"rel(fk)"`    //设置一对多关系
	Version *app.Version `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(IpAddr))
}

/**
 *	创建IP地址
 */
func AddIp(IpAddr string,Location string,appId int,versionCode int) (*IpAddr,error) {
	o:=orm.NewOrm()
	var appObj *app.App
	var versionObj *app.Version
	var ipObj *IpAddr
	var err error
	var result int64
	appObj,err = app.SelectApp(appId)
	if appObj == nil || err != nil {
		return nil,errors.New("App不存在，请创建App")
	}

	versionObj,err=app.SelectVersion(versionCode,appObj.Id)

	if versionObj == nil || err != nil {
		return nil,errors.New("版本不存在，请创建版本")
	}

	ipObj.Version = versionObj
	ipObj.App = appObj
	ipObj.IdAddr = IpAddr
	ipObj.Location = Location

	result,err = o.Insert(ipObj)
	if result <= 0 || err != nil{
		return nil,errors.New("创建IP失败")
	}

	return ipObj,nil
}