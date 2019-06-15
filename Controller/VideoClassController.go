package Controller

import(
	"context"
	"VideoAPI/TencentVideoAPI/Models"
	"errors"
	check "VideoAPI/TencentVideoAPI/Resources/ParamsCheck"
	"VideoAPI/TencentVideoAPI/Action"
)

func VideoClassController(ctx context.Context, request Models.APIGatewayProxyRequest) (ret map[string]interface{}, err error) {
	
	var cpt Models.CommonParamTemplate
	err = cpt.FromRequest(request.QueryString,request.Body) //把url参数和body参数转成对应结构
	if err != nil {
		return
	}

	switch request.HTTPMethod {
	case "POST" :
		//校验参数合法性
		err = check.CreateVideoClassCheck(cpt)
		if err != nil {
			return
		}
		ParentId := cpt.BodyParams["ParentId"]
		ClassName := cpt.BodyParams["ClassName"]
		//interface转其他类型,由于传进来的是float64，所以要强转int64
		ret,err = Action.CreateVideoClass(int64(ParentId.(float64)),ClassName.(string))
		return
	case "GET" :
		ret,err = Action.GetAllVideoClass()
		return
	case "PUT" :
		err = check.ModifyVideoClassCheck(cpt)
		if err != nil {
			return
		}
		ClassId := cpt.BodyParams["ClassId"]
		ClassName := cpt.BodyParams["ClassName"]
		//interface转其他类型,由于传进来的是float64，所以要强转uint64
		ret,err = Action.ModifyVideoClass(uint64(ClassId.(float64)),ClassName.(string))
		return
	case "DELETE" :
		err = check.DeleteVideoClassCheck(cpt)
		if err != nil {
			return
		}
		ClassId := cpt.BodyParams["ClassId"]
		//interface转其他类型,由于传进来的是float64，所以要强转int64
		ret,err = Action.DeleteVideoClass(int64(ClassId.(float64)))
		return
	default :
		err = errors.New("no this http mothod ! ")
		return
	}
}