package business

/*
 * 上传文件业务处理
 */

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
	entity "github.com/zcv8/YM.JinLiRead/entities"
)

//上传文章内的图片
func UploadArticleImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	defer file.Close()
	tempDir, err := common.GetTempDir()
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	//获取文件的后缀名
	extendName, err := common.GetFileExtendName(handler.Filename)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	//创建新的文件名
	newFileName := common.GetGuid() + "." + extendName
	newFilePath := tempDir + "/" + newFileName
	fi, err := common.OpenOrCreateFile(newFilePath)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	defer fi.Close()
	err = common.FileStreamCopy(fi, file)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&entity.ResponseStatus{
		Status: entity.SUCCEED,
		Data:   "/static/temps/" + newFileName,
	})
	fmt.Fprint(w, string(rtr))
}
