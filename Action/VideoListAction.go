package Action

import(
	"VideoAPI/TencentVideoAPI/Resources"
	"VideoAPI/TencentVideoAPI/Models"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

func GetVideoListByClassId(ClassId int64) (ret map[string]interface{} ,err error){
	// 初始化sdk 接口调用的一些配置
	client,err := Resources.InitCommomApiConfig()
	if err != nil {
		return
	}

	// 构建请求结构体，并且将string类型参数转化成对应request结构体中对应的字段
	request := vod.NewSearchMediaRequest()

	// 赋值，指针类型，传引用
	// 当不为0时，返回对应的分类视频列表，为0则不入参，直接调用，返回所有视频列表
	if ClassId != 0 {
		request.ClassIds = append(request.ClassIds,&ClassId)
	}

	//调用修改类别方法
	response, err := client.SearchMedia(request)
	if err != nil {
		return
	}
	
	vi := Models.VideoInfo{}
	videoinfos := Models.VideoList{}
	for k,_ := range response.Response.MediaInfoSet {

		vi.FileId = *response.Response.MediaInfoSet[k].FileId
		vi.Name = *response.Response.MediaInfoSet[k].BasicInfo.Name
		vi.ClassId = *response.Response.MediaInfoSet[k].BasicInfo.ClassId
		vi.ClassName = *response.Response.MediaInfoSet[k].BasicInfo.ClassName
		vi.ClassPath = *response.Response.MediaInfoSet[k].BasicInfo.ClassPath
		vi.CoverUrl = *response.Response.MediaInfoSet[k].BasicInfo.CoverUrl
		vi.Description = *response.Response.MediaInfoSet[k].BasicInfo.Description
		vi.CreateTime = *response.Response.MediaInfoSet[k].BasicInfo.CreateTime

		videoinfos.Videolist = append(videoinfos.Videolist,vi)
	} 
	
	// 数组转不了map，故再加一层结构VideoInfos{}
	if ret,err = Resources.StructToMapJson(videoinfos); err != nil {
		return
	}
	return

}