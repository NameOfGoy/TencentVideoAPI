package Action

import(
	"VideoAPI/Resources"
	_"VideoAPI/Models"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// action : CreateVideoClass 
func CreateVideoClass(ParentId int64,ClassName string) (ret map[string]interface{} ,err error){

	// 初始化sdk 接口调用的一些配置
	client,err := Resources.InitCommomApiConfig()
	if err != nil {
		return
	}

	// 构建请求结构体，并且将string类型参数转化成request结构体中对应的字段
	request := vod.NewCreateClassRequest()

	// 赋值，指针类型，传引用
	request.ParentId = &ParentId
	request.ClassName = &ClassName

	//调用创建类别方法
	response, err := client.CreateClass(request)
	if err != nil {
		return
	}

	// response结构转map
	if ret,err = Resources.StructToMapJson(response); err != nil {
		return
	}

	return
}

// action ：GetAllVedioClass 获取全部视频分类信息
func GetAllVideoClass() (ret map[string]interface{},err error) {
	
	// 初始化sdk 接口调用的一些配置
	client,err := Resources.InitCommomApiConfig()
	if err != nil {
		return
	}

	// 调用获取全部视频分类的方法，传入一个初始化的对应方法的request结构体
	Response, err := client.DescribeAllClass(vod.NewDescribeAllClassRequest()) 
	if err != nil {
		return 
	}

	// response结构转map
	if ret,err = Resources.StructToMapJson(Response); err != nil {
		return
	}

	return 
}

// action : ModifyVideoClass 
func ModifyVideoClass(ClassId uint64,ClassName string) (ret map[string]interface{} ,err error){

	// 初始化sdk 接口调用的一些配置
	client,err := Resources.InitCommomApiConfig()
	if err != nil {
		return
	}

	// 构建请求结构体，并且将string类型参数转化成对应request结构体中对应的字段
	request := vod.NewModifyClassRequest()

	// 赋值，指针类型，传引用
	request.ClassId = &ClassId
	request.ClassName = &ClassName

	//调用修改类别方法
	response, err := client.ModifyClass(request)
	if err != nil {
		return
	}

	// response结构转map
	if ret,err = Resources.StructToMapJson(response); err != nil {
		return
	}
	return
}

// action : DeleteVideoClass 
func DeleteVideoClass(ClassId int64) (ret map[string]interface{} ,err error){

	// 初始化sdk 接口调用的一些配置
	client,err := Resources.InitCommomApiConfig()
	if err != nil {
		return
	}

	// 构建请求结构体，并且将string类型参数转化成对应request结构体中对应的字段
	request := vod.NewDeleteClassRequest()

	// 赋值，指针类型，传引用
	request.ClassId = &ClassId

	//调用修改类别方法
	response, err := client.DeleteClass(request)
	if err != nil {
		return
	}

	// response结构转map
	if ret,err = Resources.StructToMapJson(response); err != nil {
		return
	}
	return
}