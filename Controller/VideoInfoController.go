package Controller

import(
	"context"
	"VideoAPI/TencentVideoAPI/Models"
	"errors"
	_"VideoAPI/TencentVideoAPI/Resources/ParamsCheck"
	"VideoAPI/TencentVideoAPI/Action"
)

func VideoInfoController(ctx context.Context, request Models.APIGatewayProxyRequest) (ret map[string]interface{}, err error) {

	var cpt Models.CommonParamTemplate
	err = cpt.FromRequest(request.QueryString,request.Body) //把url参数和body参数转成对应结构
	if err != nil {
		return
	}

	switch request.HTTPMethod {
	case "POST" :
		
		return
	case "GET" :
		if _,ok := cpt.UrlParam["VideoId"]; ok {
			ret,err = Action.DescribeMediaInfos(cpt.UrlParam["VideoId"])
			return
		}
		err = errors.New("missing param VideoId")
		return
	case "PUT" :
		
		return
	case "DELETE" :
		
		return
	default :
		err = errors.New("no this http mothod ! ")
		return
	}
}

