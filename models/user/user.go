package user

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
)

type Users struct {
	Id int `orm:"pk;auto"`
	UserName string
	UserPwd string `orm:"size(32)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

/**
 *	注册用户
 */
func AddUser(UserName string,UserPwd string) (*Users,error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.UserName = UserName

	//生成32位 md5
	h := md5.New()
	h.Write([]byte(UserPwd)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	user.UserPwd = hex.EncodeToString(cipherStr)
	//开始插入
	res,err := o.Insert(user)
	//判断是否注册成功
	if err == nil && res > 0 {
		return user,nil
	}else {
		return user,err
	}
}