package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	tagbiz "golang-training/modules/tag/biz"
	tagstorage "golang-training/modules/tag/storage"
	"net/http"
)

func FindTag(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := common.FromBase58(context.Param("id"))

		if err != nil {
			panic(err)
		}

		store := tagstorage.NewSQLStore(db)
		biz := tagbiz.NewFindTagBiz(store)

		result, err := biz.FindTag(context.Request.Context(), int(id.GetLocalID()))
		result.Mask()

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}
