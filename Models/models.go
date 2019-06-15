package Models

import (
	"encoding/json"
	"time"
)

// APIGatewayProxyRequest contains data coming from the API Gateway proxy in integration way
type APIGatewayProxyRequest struct {
	Path                  string                        `json:"path"`        // The url path be called
	QueryString           map[string]string             `json:"queryString"` // Query string in request
	HTTPMethod            string                        `json:"httpMethod"`  // HTTP method
	Headers               map[string]string             `json:"headers"`
	QueryStringParameters map[string]string             `json:"queryStringParameters,omitempty"`
	PathParameters        map[string]string             `json:"pathParameters,omitempty"`
	HeaderParameters      map[string]string             `json:"headerParameters,omitempty"`
	StageVariables        map[string]string             `json:"stageVariables,omitempty"`
	RequestContext        APIGatewayProxyRequestContext `json:"requestContext"`
	Body                  string                        `json:"body,omitempty"`
	IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`
}

// APIGatewayProxyRequestContext contains the information of service and api config in api gateway
type APIGatewayProxyRequestContext struct {
	ServiceID       string                    `json:"serviceId"`
	Path            string                    `json:"path"`
	HTTPMethod      string                    `json:"httpMethod"`
	RequestID       string                    `json:"requestId"`
	Stage           string                    `json:"stage"`
	Identity        APIGatewayRequestIdentity `json:"identity"`
	SourceIP        string                    `json:"sourceIp"`
	WebsocketEnable bool                      `json:"websocketEnable,omitempty"`
}

// APIGatewayRequestIdentity contains identity information for the request caller.
type APIGatewayRequestIdentity struct {
	SecretID string `json:"secretId,omitempty"`
}


type CommonParamTemplate struct {
	UrlParam map[string]string
	BodyParams map[string]interface{}	
}

func (cpt *CommonParamTemplate) FromRequest(urlparams map[string]string, bodyparams string) (err error) {

	cpt.UrlParam = make(map[string]string)
	for k,v := range urlparams { //url后面的参数赋值给通用参数模板
		cpt.UrlParam[k] = v
	}
	if bodyparams != "" {
		if err = json.Unmarshal([]byte(bodyparams), &cpt.BodyParams); err != nil { //把body参数转换成固定结构
			return
		}
	}
	return nil
}

type VideoInfo struct {
	FileId     string `json:"FileId"`
	Name     string `json:"Name"`
	ClassId int64 `json:"ClassId"`
	ClassName      string `json:"ClassName"`
	ClassPath string `json:"ClassPath"`
	CoverUrl string `json:"CoverUrl"`
	Description string `json:"Description"`
	Duration float64 `json:"Duration"`
	CreateTime string `json:"CreateTime"`
}

type VideoList struct {
	Videolist []VideoInfo
}

type Comment struct {
	CommentId string 		`json:"commentid" bson:"COMMENTID"`
	UserId string 			`json:"userid" bson:"USERID"`
	UserName string			`json:"username" bson:"USERNAME"`
	CommentTime time.Time	`json:"commenttime" bson:"COMMENTTIME"`
	CommentText string		`json:"commenttext" bson:"COMMENTTEXT"`
	VideoId string			`json:"videoid" bson:"VIDEOID"`
	ParentCommentId string	`json:"parentcomentid" bson:"PARENTCOMENTID"`
	ReplyToUser string		`json:"replytouser" bson:"REPLYTOUSER"`
	Likes int64				`json:"likes" bson:"LIKES"`
}

type Comments struct {
	CommentList []Comment
}
 