package comic

// ----------------------------------------------------------------------
// 漫画列表-控制器
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/app"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/e"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/setting"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/util"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/service/comic_service"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/logging"
)

func GetComicList(c *gin.Context) {
	appG := app.Gin{C: c}

	page := com.StrTo(c.Query("page")).MustInt()
	// logging.Info("page %s", page)
	valid := validation.Validation{}
	valid.Min(page, 1, "page")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	comicService := comic_service.ComicParam{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	comicList, err := comicService.GetList()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	total, err := comicService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = comicList
	data["total"] = total
	data["page"] = page
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
