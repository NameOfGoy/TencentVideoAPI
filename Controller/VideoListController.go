package Controller

import(
	"context"
	"VideoAPI/TencentVideoAPI/Models"
	"errors"
	_"VideoAPI/TencentVideoAPI/Resources/ParamsCheck"
	"VideoAPI/TencentVideoAPI/Action"
	"strconv"
)

func VideoListController(ctx context.Context, request Models.APIGatewayProxyRequest) (ret map[string]interface{}, err error) {

	var cpt Models.CommonParamTemplate
	err = cpt.FromRequest(request.QueryString,request.Body) //把url参数和body参数转成对应结构
	if err != nil {
		return
	}

	switch request.HTTPMethod {
	case "POST" :
		
		return
	case "GET" :
		if _,ok := cpt.UrlParam["ClassId"]; ok {
			if cpt.UrlParam["ClassId"] == "ALL" {
				ret,err = Action.GetVideoListByClassId(0) //参数为all，传0.使之返回全部数据
				return
			}
			ClassId,err1 := strconv.ParseInt(cpt.UrlParam["ClassId"], 10, 64) //string转int64
			if err1 != nil {
				err = err1
				return 
			}
			ret,err = Action.GetVideoListByClassId(ClassId)
			return
		}
		err = errors.New("missing param ClassId")
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