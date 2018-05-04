package business

/*
 * 处理与频道标签有关的业务逻辑
 */

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
	entity "github.com/zcv8/YM.JinLiRead/entities"
)

// 获取所有频道标签
func GetChannels(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	channels, err := data.GetChannels()
	if err != nil {
		json.NewEncoder(w).Encode(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    nil,
			Message: common.ReadDataFailedError.SetOrginalErr(err).String(),
		})
		return
	}
	json.NewEncoder(w).Encode(&entity.ResponseStatus{
		Status:  entity.SUCCEED,
		Data:    channels,
		Message: "",
	})
}
