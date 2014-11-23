// @SubApi helloworld [/helloworld]
package controllers

import "github.com/revel/revel"

type App struct {
  *revel.Controller
}

type AjaxResult struct {
  Success bool   `json:"success"`
  Error   string `json:"error"`
  Data    string `json:"data"`
}

var mHelloWorld = map[int]string{}

func (c App) Index() revel.Result {
  return c.Render()
}

// @Title AjaxGetHelloWorld
// @Description 获取指定 key 数据
// @Param   key     path   int    true "id 必须 > 0"
// @Success 200 {object}  AjaxResult
// @Failure 400 {object}  AjaxResult    key 必须 > 0
// @Resource /helloworld
// @Router /j/helloworld [get]
func (c *App) AjaxGetHelloWorld(key int) revel.Result {
  if key < 0 {
    c.Response.Status = 400
    return c.RenderJson(AjaxResult{
      Success: false,
      Error:   "key 必须 > 0",
    })
  }

  if _, ok := mHelloWorld[key]; !ok {
    c.Response.Status = 400
    return c.RenderJson(AjaxResult{
      Success: false,
      Error:   "key 不存在",
    })
  }

  return c.RenderJson(AjaxResult{
    Success: true,
    Data:    mHelloWorld[key],
  })
}

// @Title AjaxPostHelloWorld
// @Description 设置指定 key 数据
// @Param   key     form   int    true "id 必须 > 0"
// @Param   data    form   string true "数据"
// @Success 200 {object}  AjaxResult
// @Failure 400 {object}  AjaxResult    key 必须 > 0
// @Resource /helloworld
// @Router /j/helloworld [post]
func (c *App) AjaxPostHelloWorld(key int, data string) revel.Result {

  c.Validation.Required(data)

  if c.Validation.HasErrors() {
    c.Response.Status = 400
    return c.RenderJson(AjaxResult{
      Success: false,
      Error:   "data 不能为空",
    })
  }

  if key < 0 {
    c.Response.Status = 400
    return c.RenderJson(AjaxResult{
      Success: false,
      Error:   "key 必须 > 0",
    })
  }

  mHelloWorld[key] = data

  return c.RenderJson(AjaxResult{
    Success: true,
  })
}
