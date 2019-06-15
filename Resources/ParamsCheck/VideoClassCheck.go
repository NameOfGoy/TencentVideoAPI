package ParamsCheck

import(
	"VideoAPI/Models"
	"errors"
)

func CreateVideoClassCheck(cpt Models.CommonParamTemplate) (err error) {
	if cpt.BodyParams == nil {
		err = errors.New("没有Params参数")
		return
	}
	if _,ok := cpt.BodyParams["ParentId"]; !ok {
		err = errors.New("没有ParentId参数")
		return
	}
	if _,ok := cpt.BodyParams["ClassName"]; !ok {
		err = errors.New("没有ClassName参数")
		return
	}
	return
}

func ModifyVideoClassCheck(cpt Models.CommonParamTemplate) (err error) {

	if cpt.BodyParams == nil {
		err = errors.New("没有Params参数")
		return
	}
	if _,ok := cpt.BodyParams["ClassId"]; !ok {
		err = errors.New("没有ClassId参数")
		return
	}
	if _,ok := cpt.BodyParams["ClassName"]; !ok {
		err = errors.New("没有ClassName参数")
		return
	}
	return
}

func DeleteVideoClassCheck(cpt Models.CommonParamTemplate) (err error) {

	if cpt.BodyParams == nil {
		err = errors.New("没有Params参数")
		return
	}
	if _,ok := cpt.BodyParams["ClassId"]; !ok {
		err = errors.New("没有ClassId参数")
		return
	}
	return
}