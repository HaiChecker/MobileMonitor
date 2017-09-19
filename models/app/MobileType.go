package app

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type MobileType struct {
	Id int `orm:"pk;aoto"`
	Name string

	//一对多
	App []*App `orm:"reverse(many)"`
}

func init()  {
	orm.RegisterModel(new(App))
}

/**
 *	查询移动端类型(通过名字)
 */
func SelectMobileName(MobileName string) (*MobileType,error) {
	o:=orm.NewOrm()
	mType := new(MobileType)
	res := o.QueryTable(mType).Filter("name",MobileName).One(mType)
	return mType,res
}

/**
 *	添加类型
 */
func AddMobileType(Name string) (*MobileType,error) {
	o:=orm.NewOrm()

	var mType *MobileType

	mType.Name = Name
	result,err := o.Insert(mType)
	if result <= 0 || err != nil {
		return nil,errors.New("添加Type失败")
	}

	return mType,nil
}