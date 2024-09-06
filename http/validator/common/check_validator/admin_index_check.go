package checkvalidator

import (
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/globalConst"
	"friends_ranking/http/controller"
	"friends_ranking/http/data_transfer"
	"friends_ranking/utils/response"

	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName string `form:"userName" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Pass1    string `form:"pass1" json:"pass1" binding:"required,min=6,max=20"`  //  密码为 必填，长度>=6
	Pass2    string `form:"pass2" json:"pass1" binding:"required,min=6,max=20"`  //  密码为 必填，长度>=6
}

type Regist struct {
	UserName string `form:"userName" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Mobile   string `form:"mobile" json:"mobile"  binding:"required,min=1"`      // 必填、对于文本,表示它的长度>=1
	Pass1    string `form:"pass" json:"pass" binding:"required,min=6,max=20"`    //  密码为 必填，长度>=6
	Sex      int32  `form:"sex" json:"sex" binding:"required,min=6,max=20"`      //  密码为 必填，长度>=6
}

// 验证器语法，参见 Register.go文件，有详细说明

func (l Login) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&l); err != nil {
		response.ReturnFail(context, errorMsg.ErrorsValidatorNoPass)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(l, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.IndexController{}).Login(extraAddBindDataContext)
	}

}

func (r Regist) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBind(r); err != nil {
		response.ReturnFail(context, errorMsg.ErrorsValidatorNoPass)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(r, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.IndexController{}).Register(extraAddBindDataContext)
	}

}