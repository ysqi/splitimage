package service

import (
	"image"
	"io/ioutil"
	"path"
	"testing"
)

func TestPotint(t *testing.T) {

	maxPoint := image.Point{}
	maxPoint.X = 1728
	maxPoint.Y = 2592
	length := 500

	cur := image.Rectangle{}
	for {
		cur = getNextRect(cur, maxPoint, length)
		t.Log(cur)
		if cur.Max == maxPoint {
			break
		}

	}

}

func TestSpiteOneImg(t *testing.T) {

	srcPath := `E:\PIC\光盘\`
	savePath := `E:\WORK\09.Dev\06.外包项目\图片切片\new\`

	x, y, err := SplitToSquare(path.Join(srcPath, "7K0A6694.JPG"), 670, path.Join(savePath, "4"))
	if err != nil {
		t.Error(err)
	}
	t.Log(x, y)

}

func TestSpiteImg(t *testing.T) {

	srcPath := `E:\PIC\60张精修\`
	// srcPath := `E:\WORK\09.Dev\06.外包项目\图片切片\src\`
	savePath := `E:\WORK\09.Dev\06.外包项目\图片切片\new\`
	files, _ := ioutil.ReadDir(srcPath)
	for _, fi := range files {
		if fi.IsDir() {
			continue
		}
		t.Log(fi.Name())
		x, y, err := SplitToSquare(path.Join(srcPath, fi.Name()), 600, path.Join(savePath, fi.Name()))
		if err != nil {
			t.Error(err)
		}
		t.Log(x, y)

	}

}
