package controllers

import (
	"github.com/astaxie/beego"

	"github.com/ysqi/splitimage/service"
)

type ImageController struct {
	beego.Controller
}

//将图片分割成指定大小的正方形小方块
func (c *ImageController) Split() {
	var result = struct {
		Success bool        `json:"success"`
		Msg     string      `json:"msg"`
		Data    interface{} `json:"data"`
	}{
		Success: false,
	}

	var x, y int
	defer func() {
		result.Success = result.Msg == ""
		if result.Success {
			result.Data = map[string]int{
				"x": x,
				"y": y,
			}
		}
		c.Data["json"] = result
		c.ServeJSON()
	}()

	src := c.GetString("src")
	if src == "" {
		result.Msg = "入参不完整，缺失 src"
		return
	}

	length, err := c.GetInt("length")
	if err != nil {
		result.Msg = "入参 length 非法，必须为整数"
		return
	}
	if length <= 0 {
		result.Msg = "入参 length 非法，必须大于0"
		return
	}

	savePath := c.GetString("save")
	if savePath == "" {
		result.Msg = "入参不完整，缺失 save"
		return
	}

	x, y, err = service.SplitToSquare(src, int(length), savePath)
	if err != nil {
		result.Msg = err.Error()
	}

}
