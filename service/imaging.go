package service

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/disintegration/imaging"
)

func getNextRect(pre image.Rectangle, maxPoint image.Point, length int) image.Rectangle {

	p := image.Point{}
	prePoint := pre.Min

	//如果不是原点第一张图片
	if pre.Empty() == false {
		p.Y = prePoint.Y
		p.X = prePoint.X + length
		if p.X >= maxPoint.X {
			p.X = 0
			p.Y += length
		}
		if p.Y > maxPoint.Y {
			p.Y = maxPoint.Y
		}

	}
	mp := p.Add(image.Point{X: length, Y: length})
	if mp.X > maxPoint.X {
		mp.X = maxPoint.X
	}
	if mp.Y > maxPoint.Y {
		mp.Y = maxPoint.Y
	}

	return image.Rectangle{
		Min: p,
		Max: mp,
	}

}

func SplitToSquare(srcName string, length int, savePath string) (x int, y int, err error) {

	if length <= 0 {
		err = errors.New("切片图片大小定义非法")
		return
	}

	if _, err = os.Stat(srcName); os.IsNotExist(err) {
		err = errors.New("文件不存在")
		return
	}

	var img image.Image
	img, err = imaging.Open(srcName)
	if err != nil {
		return
	}
	// mkdir if save dir is not exist.
	if _, err = os.Stat(savePath); os.IsNotExist(err) {
		if err = os.Mkdir(savePath, 0777); err != nil {
			return
		}
	}

	//默认分割成此格式图片
	ext := ".JPG"

	_, name := filepath.Split(srcName)
	fileName := path.Join(savePath, strings.Split(name, ".")[0]+ext)
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	//保存缩略图,缩略图最大值
	if thubMaxLen, _ := beego.AppConfig.Float("thubMaxLen"); thubMaxLen > 0 {
		if float64(w) > thubMaxLen || float64(h) > thubMaxLen {
			x, y := float64(w), float64(h)
			bili := x / y

			if x > thubMaxLen {
				x = thubMaxLen
			}
			if y > thubMaxLen {
				y = thubMaxLen
			}

			if x >= y {
				y = x / bili
			} else {
				x = y * bili
			}
			w = int(x)
			h = int(y)
		}
	}
	thub := imaging.Resize(img, w, h, imaging.Lanczos)
	if err = imaging.Save(thub, fileName); err != nil {
		return
	}

	// get image real type for sub.
	rgbImg := img.(*image.NRGBA)
	cur := image.Rectangle{}
	for {
		cur = getNextRect(cur, rgbImg.Rect.Max, length)
		subImg := rgbImg.SubImage(cur)
		saveName := path.Join(savePath, fmt.Sprintf("(%d,%d)%s", cur.Min.Y/length, cur.Min.X/length, ext))
		err = imaging.Save(subImg, saveName)
		if err != nil {
			return
		}

		if cur.Max == rgbImg.Rect.Max {
			break
		}

	}
	return cur.Min.X / length, cur.Min.Y / length, nil
}
