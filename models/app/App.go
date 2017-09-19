package app

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"MobileMonitor/models/user"
	"MobileMonitor/models/phone"
)

type App struct {
	Id int `orm:"pk;auto;index"`
	AppKey string `orm:"size(32)"`
	PackgeName string
	AppName string


	Version []*Version `orm:"reverse(many)"` // 设置一对多的反向关系
	Ip []* phone.IpAddr `orm:"reverse(many)"` // 设置一对多的反向关系
	User *user.Users `orm:"rel(fk)"`
	MobileType *MobileType `orm:"rel(fk)"`    //设置一对多关系

}

func init()  {
	orm.RegisterModel(new(App))
}

func SelectApp(appId int) (*App, error) {
	o:=orm.NewOrm()
	app := new(App)
	res := o.QueryTable(app).Filter("id",appId).One(app)
	return app,res
}

/**
 *	添加一个App
 */
func AddApp(userId int,AppKey string,AppName string,PackgeName string,MobileType string) (*App,error) {
	o:= orm.NewOrm()

	var User *user.Users

	//判断User是否存在
	User,_ = user.SelectUser(userId)

	if User != nil {
		return nil,errors.New("用户错误!")
	}


	//判断终端类型是否存在，如果不存在则创建
	var mType *MobileType
	mType,_ = SelectMobileName(MobileType)
	if mType == nil {
		mType,_ = AddMobileType(MobileType)
		if mType == nil{
			 return nil,errors.New("添加类型错误！")
		}
	}
	var app *App
	app.MobileType = mType
	app.User = User
	app.AppKey = AppKey
	app.AppName = AppName
	app.PackgeName = PackgeName

	//正式插入App数据
	result,err := o.Insert(app)

	if result <= 0 || err != nil {
		return nil,errors.New("添加App失败")
	}
	return app,nil
}