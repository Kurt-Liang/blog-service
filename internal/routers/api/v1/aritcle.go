package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 取得單個文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {}

// @Summary 取得多個文章
// @Produce json
// @Param name query string false "文章名稱"
// @Param tag_id query int false "標籤ID"
// @Param state query int false "狀態"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary 新增文章
// @Produce json
// @Param tag_id body string true "標籤ID"
// @Param title body string true "文章標題"
// @Param desc body string false "文章簡述"
// @Param cover_image_url body string true "封面圖片位址"
// @Param content body string true "文章內容"
// @Param created_by body int true "建立人"
// @Param state body int false "狀態"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce json
// @Param tag_id body string false "標籤ID"
// @Param title body string false "文章標題"
// @Param desc body string false "文章簡述"
// @Param cover_image_url body string false "封面圖片位址"
// @Param content body string false "文章內容"
// @Param modified_by body string true "修改人"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
