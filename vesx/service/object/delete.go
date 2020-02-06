package objectservice

import (
	"github.com/Myriad-Dreamin/go-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/vesx/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func (svc *Service) deleteHook(c controller.MContext, object *model.Object) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
		Code: types.CodeDeleteError,
		Err:  "generated delete api has not been implemented yet",
	})
	return false
}