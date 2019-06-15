package Resources

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
)

func InitCommomApiConfig() (apiClient *vod.Client,err error) {

	// 从环境变量中密钥对，用于访问接口，确定用户
	var SECRET_ID string = os.Getenv("SECRET_ID")
	var SECRET_KEY string = os.Getenv("SECRET_KEY")
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(SECRET_ID, SECRET_KEY)
	// 非必要步骤
    // 实例化一个客户端配置对象，可以指定超时时间等配置
	prof := profile.NewClientProfile()
	

	// 实例化要请求产品(以cvm为例)的client对象，此处是用vod/v20180717
    // 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
	apiClient, err = vod.NewClient(credential, regions.Guangzhou, prof)
	if err != nil {
		return 
	}
	return 
} 