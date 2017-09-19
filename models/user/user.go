package user

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"MobileMonitor/models/app"
)

type Users struct {
	Id int `orm:"pk;auto;index"`
	UserName string `orm:"index"`
	UserPwd string `orm:"size(32)"`
	Email string

	App []*app.App `orm:"reverse(many)"`
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

/**
 *	登录，需要已经加密成功过后
 */
func Login(UserName string,UserPwd string) (*Users,error)  {
	o:= orm.NewOrm()
	user := new(Users)
	res := o.QueryTable(user).Filter("user_name",UserName).Filter("user_pwd",UserPwd).One(user)
	return user,res
}


/**
 *	查询用户，通过ID
 */
func SelectUser(userId int) (*Users,error) {
	o:= orm.NewOrm()
	user := new(Users)
	res := o.QueryTable(user).Filter("id",userId).One(user)
	return user,res
}

/**
 *	检查用户名是否存在，存在则返回true
 */
func CheckUserName(UserName string) bool {
	o:= orm.NewOrm()
	user := new(Users)
	res := o.QueryTable(user).Filter("user_name",UserName).Exist()
	return res
}

/**
 *	检查邮箱是否存在，存在则返回true
 */
func CheckEmail(Email string) bool {
	o:= orm.NewOrm()
	user := new(Users)
	res := o.QueryTable(user).Filter("email",Email).Exist()
	return res
}