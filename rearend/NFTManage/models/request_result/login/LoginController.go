package loginResult

import (
	"time"
	"yqnft/NFTManage/models/model"
)

type LoginUser struct {
	ID           string            ` json:"id"`       //
	UserName     string            ` json:"userName"` // 姓名
	Account      string            ` json:"account"`  // 账号
	Password     string            ` json:"password"` // 密码,默认密码123
	Sex          string            ` json:"sex"`      // 性别
	Mobile       string            ` json:"mobile"`   // 手机号码
	Email        string            ` json:"email"`    // 邮箱
	OpenID       string            ` json:"openId"`   // 微信openId
	UnionID      string            ` json:"unionID"`  // 微信UnionID
	Avatar       string            ` json:"avatar"`   // 头像地址
	UserType     int64             ` json:"userType"` // 0会员,1员工,2店长收银,9系统管理员
	CreateTime   time.Time         ` json:"createTime"`
	CreateUserID string            ` json:"createUserId"`
	UpdateTime   time.Time         ` json:"updateTime"`
	UpdateUserID string            ` json:"updateUserId"`
	Active       int64             ` json:"active"` // 是否有效(0否,1是)
	Bak1         string            ` json:"bak1"`
	Bak2         string            ` json:"bak2"`
	Bak3         string            ` json:"bak3"`
	Bak4         string            ` json:"bak4"`
	Bak5         string            ` json:"bak5"`
	RoleName     string            ` json:"roleName"`
	RoleCode     string            ` json:"roleCode"`
	Codes        []string          ` json:"codes"`
	CodeMap      map[string]string ` json:"codeMap"`
	IndexPage    string            ` json:"indexPage"`
	DeptId       string            ` json:"deptId"`
	Roles        []model.TRole     ` json:"roles"`
	OrgList      []model.TOrg      ` json:"orgList"`
}

type LoginResult struct {
	Jwttoken    string    `json:"jwttoken"`
	User        LoginUser `json:"user"`
	ShowCaptcha bool      `json:"showCaptcha"`
}

type User struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	CaptchaKey string `json:"captchaKey"`
	Captcha    string `json:"captcha"`
}
