package Action

import(
	"VideoAPI/Resources"
	"VideoAPI/Models"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// action : DescribeMediaInfos 
func DescribeMediaInfos(FileIds string) (ret map[string]interface{} ,err error){

	// 初始化sdk 接口调用的一些配置
	client,err := Resources.InitCommomApiConfig()
	if err != nil {
		return
	}

	// 构建请求结构体，并且将string类型参数转化成对应request结构体中对应的字段
	request := vod.NewDescribeMediaInfosRequest()

	// slice不初始化的话，默认len和cap都是0，所以用append自动扩容
	request.FileIds = append(request.FileIds,&FileIds)

	//调用修改类别方法
	response, err := client.DescribeMediaInfos(request)
	if err != nil {
		return
	}
	// 从response取得对应的信息
	vi := Models.VideoInfo{}
	for k,_ := range response.Response.MediaInfoSet {
		vi.FileId = *response.Response.MediaInfoSet[k].FileId
		vi.Name = *response.Response.MediaInfoSet[k].BasicInfo.Name
		vi.ClassId = *response.Response.MediaInfoSet[k].BasicInfo.ClassId
		vi.ClassName = *response.Response.MediaInfoSet[k].BasicInfo.ClassName
		vi.ClassPath = *response.Response.MediaInfoSet[k].BasicInfo.ClassPath
		vi.CoverUrl = *response.Response.MediaInfoSet[k].BasicInfo.CoverUrl
		vi.Description = *response.Response.MediaInfoSet[k].BasicInfo.Description
		vi.Duration = *response.Response.MediaInfoSet[k].MetaData.Duration
		vi.CreateTime = *response.Response.MediaInfoSet[k].BasicInfo.CreateTime
	}

	if ret,err = Resources.StructToMapJson(vi); err != nil {
		return
	}
	return
}