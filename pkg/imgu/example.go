package imgutil

import (
	"bufio"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/ccitt"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	_ "golang.org/x/image/riff"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/vp8"
	_ "golang.org/x/image/vp8l"
	_ "golang.org/x/image/webp"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
)

func readImg(picFile string) image.Image {
	f, err := os.Open(picFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, fmtName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %v, Bounds: %+v, Color: %+v", fmtName, img.Bounds(), img.ColorModel())
	return img
}

func SavePic(img image.Image, picFile string) {
	outFile, err := os.Create(picFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	bio := bufio.NewWriter(outFile)
	if err := jpeg.Encode(bio, img, &jpeg.Options{Quality: 60}); err != nil {
		panic(err)
	}
	bio.Flush()
}

func Drawx() {
	f, err := os.Open("./out.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	// 200 * 200 的 RGBA 画布
	dst := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(dst, image.Rect(0, 0, 100, 100), img, image.Point{300, 300}, draw.Src)

	fDst, err := os.Create("xxx.png")
	if err != nil {
		log.Fatal(err)
	}
	defer fDst.Close()
	err = png.Encode(fDst, dst)
	if err != nil {
		log.Fatal(err)
	}
}

type circle struct { // 这里需要自己实现一个圆形遮罩，实现接口里的三个方法
	p image.Point // 圆心位置
	r int         // 半径
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

// 对每个像素点进行色值设置，在半径以内的图案设成完全不透明
func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{A: 255}
	}
	return color.Alpha{}
}
func drawCirclePic() {
	f, err := os.Open("./outv12.jpeg")
	if err != nil {
		panic(err)
	}
	gopherImg, _, err := image.Decode(f)

	d := gopherImg.Bounds().Dx()

	//将一个cicle作为蒙层遮罩，圆心为图案中点，半径为边长的一半
	c := circle{p: image.Point{X: d / 2, Y: d / 2}, r: d / 2}
	circleImg := image.NewRGBA(image.Rect(0, 0, d, d))
	draw.DrawMask(circleImg, circleImg.Bounds(), gopherImg, image.Point{}, &c, image.Point{}, draw.Over)

	//SavePng(circleImg)

	fDst, err := os.Create("xxx2.png")
	if err != nil {
		log.Fatal(err)
	}
	defer fDst.Close()
	err = png.Encode(fDst, circleImg)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFont() {
	ttfBytes, err := ioutil.ReadFile("fzchuheisong.TTF") // 读取 ttf 文件
	if err != nil {
		panic(err)
	}
	ft, err := freetype.ParseFont(ttfBytes)
	if err != nil {
		panic(err)
	}
	// 绘制底图 400*400
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)

	fc := freetype.NewContext()
	fc.SetDPI(72)  // 每英寸的分辨率,默认72
	fc.SetFont(ft) // 使用的字体
	fc.SetHinting(font.HintingNone)
	fc.SetFontSize(28)
	fc.SetClip(img.Bounds()) // 绘制的矩形
	fc.SetDst(img)           // 绘制在哪个图片上
	fc.SetSrc(image.White)   // 字体颜色,使用纯色图片 image.Uniform
	//image.NewUniform(color.RGBA{0, 0, 255, 255})

	fc.PointToFixed(28)
	//fontHeight := int(fc.PointToFixed(28) >> 6)
	// 文字的源点坐标是文字的左下角，所以对应到图片上就是文字的高度
	// 字体高度=字体大小*72/96
	//_, err = fc.DrawString("hello world", fc.PointToFixed(28))
	_, err = fc.DrawString("hello world", freetype.Pt(0, int(fc.PointToFixed(28)>>6)))
	if err != nil {
		panic(err)
	}
	file, err := os.Create("font_test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		panic(err)
	}
}

func WriteFont2() {
	fontBytes, err := ioutil.ReadFile("fzchuheisong.TTF") // 读取 ttf 文件
	if err != nil {
		panic(err)
	}
	ft, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	// 绘制底图 400*400
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)

	drawer := font.Drawer{
		Dst: img,         // 绘制在哪个图片上
		Src: image.White, // 字体颜色
		Face: truetype.NewFace(ft, // 使用的字体文件
			&truetype.Options{
				Size:    28, // 字体大小
				DPI:     72, // 字体 dpi
				Hinting: font.HintingFull,
			}),
		//fontHeight := int(fc.PointToFixed(28) >> 6)
		// 文字的源点坐标是文字的左下角，所以对应到图片上就是文字的高度
		// 字体高度=字体大小*72/96
		Dot: freetype.Pt(0, 128), // 字体的位置 文字的源点坐标是文字的左下角，所以对应到图片上就是文字的高度
	}
	drawer.DrawString("hello world")

	file, err := os.Create("font_test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		panic(err)
	}
}
