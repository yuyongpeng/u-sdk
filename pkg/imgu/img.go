package imgutil

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
)

// CertificateCollection
// @Description: 集合的登记证书
type CertificateCollection struct {
	company                string // 公司
	nftPath                string // nft路径
	catetory               string // 作品类别
	firstPublish           string // 首次发表时间
	DigitalAssetsNO        string // 数字资产编号
	issuePrice             string // 发行价格
	hashiiValue            string // hashii 值
	owner                  string // 持有人
	registrationNO         string // 登记号
	productName            string // 作品名称
	author                 string // 作者
	creationCompletionTime string // 创作完成时间
	registrationDate       string // 登记日期
}

// CertificateNFT
// @Description: NFT的登记证书
type CertificateNFT struct {
	company                string // 公司
	nftPath                string // nft路径
	catetory               string // 作品类别
	firstPublish           string // 首次发表时间
	publishNum             string // 发布数量
	issuePrice             string // 发行价格
	registrationNO         string // 登记号
	productName            string // 作品名称
	author                 string // 作者
	creationCompletionTime string // 创作完成时间
	registrationDate       string // 登记日期
}

var backgroundImg string = "cert.jpg"
var backgroundImgX int
var backgroundImgY int
var sealImg string = "certificate.png"
var sealImgX int
var sealImgY int
var songtiFontFaceFile string = "STSONG.TTF"              // 宋体
var msYaHeiBoldFile string = "MicrosoftYaHei-Bold-01.ttf" // 微软雅黑粗体

func init() {
	file, err := os.Open(backgroundImg)
	if err != nil {
		log.Fatal(err)
	}
	fileConf, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}
	backgroundImgX = fileConf.Width
	backgroundImgY = fileConf.Height
	sealImgX = backgroundImgX - 533
	sealImgY = backgroundImgY - 413
}

// DrawCert
//
//	@Description: 绘制证书  1803×2543  229×227
//	@param ttf: ttf字体文件路径,使用的是x/image/gofont 字体库
//	@param certImg: 证书图片路径
//	@param cert: 证书信息
func DrawCertNFTTest(cert CertificateNFT, outFile string) {
	// 字体颜色
	fontR := 164
	fontG := 128
	fontB := 94

	// Draw 底图
	dc := gg.NewContext(backgroundImgX, backgroundImgY)
	src, openErr := imaging.Open(backgroundImg)
	if openErr != nil {
		panic(openErr)
	}
	dc.DrawImage(src, 0, 0)

	songtiFontFace1, err := getOpenTypeFontFace(songtiFontFaceFile, 30, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace1)
	dc.SetRGB255(fontR, fontG, fontB)
	dc.DrawStringWrapped("以下事项，由“"+cert.company+"”申请", float64(backgroundImgX/2), 430, 0.5, 0.5, 800, 1.0, gg.AlignCenter)
	dc.DrawStringWrapped("经由“北京国际大数据交易所”审核，根据《数字资产平台登记暂行条例》规定，予以登记", float64(backgroundImgX/2), 500, 0.5, 0.5, 700, 1.0, gg.AlignCenter)

	// Draw NFT
	nftImg, openErr := imaging.Open(cert.nftPath)
	if openErr != nil {
		panic(openErr)
	}
	nft := imaging.Resize(nftImg, 0, 710, imaging.Lanczos)
	dc.DrawImageAnchored(nft, backgroundImgX/2, 1000, 0.5, 0.5)

	// Draw title
	fontFace, err := getOpenTypeFontFace(msYaHeiBoldFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*fontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	var initX float64 = 503
	var initY float64 = 1510
	var incr float64 = 80
	var title []string = []string{"作  品  类  别：", "首次发表时间：", "发  行  数  量：", "发  行  价  格：", "登     记     号：", "作   品  名  称：", "作              者：", "创作完成时间："}
	for i := 0; i < len(title); i++ {
		//dc.DrawStringAnchored(title[i], initX, initY, 0.5, 0.5)
		dc.DrawStringWrapped(title[i], initX, initY, 0.5, 0.5, 200, 2.5, gg.AlignRight)
		initY += incr
	}

	//
	var qianzhanX float64 = 1163
	var qianzhanY float64 = 2303
	dc.DrawString("登记机构签章", qianzhanX, qianzhanY)
	dc.DrawString("登记日期："+cert.registrationDate, qianzhanX, qianzhanY+40)

	// Draw content
	//dc.Clear()
	songtiFontFace, err := getOpenTypeFontFace(songtiFontFaceFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	initX = 615
	initY = 1528
	incr = 80
	//title = []string{"数字文创", "2023/11/29", "30", "￥1~￥100", "20231226161433985172082240120", "1209-1344测试", "daop-12091344wangmj", "2023/11/29"}
	title = []string{cert.catetory, cert.firstPublish, cert.publishNum, cert.issuePrice, cert.registrationNO, cert.productName, cert.author, cert.creationCompletionTime}
	for i := 0; i < len(title); i++ {
		dc.DrawString(title[i], initX, initY)
		initY += incr
	}

	// Draw Cert
	certificate, openErr := imaging.Open(sealImg)
	if openErr != nil {
		panic(openErr)
	}
	dc.DrawImage(certificate, sealImgX, sealImgY)
	dc.SavePNG(outFile)
}

func DrawCertNFT(cert CertificateNFT, outFile string) {
	// 字体颜色
	fontR := 164
	fontG := 128
	fontB := 94

	dc := InitCertificate()
	DrawTitle(dc, float64(backgroundImgX/2), 430, cert.company, fontR, fontG, fontB)
	DrawNFTImg(dc, backgroundImgX/2, 1000, cert.nftPath, 710)
	DrawContentNFT(dc, cert, fontR, fontG, fontB)
	var qianzhanX float64 = 1163
	var qianzhanY float64 = 2303
	DrawEnd(dc, qianzhanX, qianzhanY, cert.registrationDate, fontR, fontG, fontB)
	DrawCert(dc, sealImgX, sealImgY, sealImg)
	dc.SavePNG(outFile)
}

// DrawCertCollection
//
//	@Description: 	画 集合 证书
//	@param cert		数据
//	@param outFile	输出文件路径
func DrawCertCollection(coll CertificateCollection, outFile string) {
	// 字体颜色
	fontR := 164
	fontG := 128
	fontB := 94

	dc := InitCertificate()
	DrawTitle(dc, float64(backgroundImgX/2), 430, coll.company, fontR, fontG, fontB)
	DrawNFTImg(dc, backgroundImgX/2, 1000, coll.nftPath, 710)
	DrawContentCollection(dc, coll, fontR, fontG, fontB)
	var qianzhanX float64 = 1163
	var qianzhanY float64 = 2303
	DrawEnd(dc, qianzhanX, qianzhanY, coll.registrationDate, fontR, fontG, fontB)
	DrawCert(dc, sealImgX, sealImgY, sealImg)
	dc.SavePNG(outFile)
}

func Draw() {
	// Draw 底图
	dc := gg.NewContext(1803, 2543)
	src, openErr := imaging.Open("cert.jpg")
	if openErr != nil {
		panic(openErr)
	}
	dc.DrawImage(src, 0, 0)

	songtiTTF1 := "STSONG.TTF"
	songtiFontFace1, err := getOpenTypeFontFace(songtiTTF1, 30, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace1)
	dc.SetRGB255(164, 128, 94)
	companyName := "星矿科技（北京）有限公司"
	dc.DrawStringWrapped("以下事项，由“"+companyName+"”申请", 1803/2, 430, 0.5, 0.5, 800, 1.0, gg.AlignCenter)
	dc.DrawStringWrapped("经由“北京国际大数据交易所”审核，根据《数字资产平台登记暂行条例》规定，予以登记", 1803/2, 500, 0.5, 0.5, 700, 1.0, gg.AlignCenter)

	// Draw NFT
	nftImg, openErr := imaging.Open("nft.jpg")
	if openErr != nil {
		panic(openErr)
	}
	nft := imaging.Resize(nftImg, 0, 710, imaging.Lanczos)
	dc.DrawImageAnchored(nft, 1803/2, 1000, 0.5, 0.5)

	// Draw title
	fontFilePath := "MicrosoftYaHei-Bold-01.ttf"
	fontFace, err := getOpenTypeFontFace(fontFilePath, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*fontFace)
	dc.SetRGB255(164, 128, 94)
	var initX float64 = 503
	var initY float64 = 1510
	var incr float64 = 80
	var title []string = []string{"作  品  类  别：", "首次发表时间：", "发  行  数  量：", "发  行  价  格：", "登     记     号：", "作   品  名  称：", "作              者：", "创作完成时间："}
	for i := 0; i < len(title); i++ {
		//dc.DrawStringAnchored(title[i], initX, initY, 0.5, 0.5)
		dc.DrawStringWrapped(title[i], initX, initY, 0.5, 0.5, 200, 2.5, gg.AlignRight)
		initY += incr
	}

	//
	var qianzhanX float64 = 1163
	var qianzhanY float64 = 2303
	dc.DrawString("登记机构签章", qianzhanX, qianzhanY)
	dc.DrawString("登记日期：2023年12月26日", qianzhanX, qianzhanY+40)

	// Draw content
	//dc.Clear()
	songtiTTF := "STSONG.TTF"
	//songtiTTF := "fzchuheisong.TTF"
	songtiFontFace, err := getOpenTypeFontFace(songtiTTF, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace)
	dc.SetRGB255(164, 128, 94)
	initX = 615
	initY = 1528
	incr = 80
	title = []string{"数字文创", "2023/11/29", "30", "￥1~￥100", "20231226161433985172082240120", "1209-1344测试", "daop-12091344wangmj", "2023/11/29"}
	for i := 0; i < len(title); i++ {
		dc.DrawString(title[i], initX, initY)
		initY += incr
	}

	// Draw Cert
	certificate, openErr := imaging.Open("certificate.png")
	if openErr != nil {
		panic(openErr)
	}
	dc.DrawImage(certificate, 1270, 2130)
	dc.SavePNG("out2.png")
}

// InitCertificate
//
//	@Description: 初始化证书底图
//	@return *gg.Context
func InitCertificate() *gg.Context {
	// Draw 底图
	dc := gg.NewContext(backgroundImgX, backgroundImgY)
	src, openErr := imaging.Open(backgroundImg)
	if openErr != nil {
		panic(openErr)
	}
	dc.DrawImage(src, 0, 0)
	return dc
}

// DrawTitle
//
//	@Description: 画 证书的title
//	@param dc
//	@param x	title的X坐标
//	@param y	title的Y坐标
//	@param companyName	title的公司名称
//	@param fontR	字体的颜色R值
//	@param fontG	字体的颜色G值
//	@param fontB	字体的颜色B值
//	@return *gg.Context
func DrawTitle(dc *gg.Context, x, y float64, companyName string, fontR, fontG, fontB int) *gg.Context {
	x = float64(backgroundImgX / 2)
	y = 430
	songtiFontFace1, err := getOpenTypeFontFace(songtiFontFaceFile, 30, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace1)
	dc.SetRGB255(fontR, fontG, fontB)
	dc.DrawStringWrapped("以下事项，由“"+companyName+"”申请", x, y, 0.5, 0.5, 800, 1.0, gg.AlignCenter)
	dc.DrawStringWrapped("经由“北京国际大数据交易所”审核，根据《数字资产平台登记暂行条例》规定，予以登记", x, y+70, 0.5, 0.5, 700, 1.0, gg.AlignCenter)

	return dc
}

// DrawNFTImg
//
//	@Description: 画证书的NFT图片
//	@param dc
//	@param x	图片的X坐标
//	@param y	图片的Y坐标
//	@param nftImg	NFT的文件路径（目前只支持本地文件）
//	@param height	图片限定最高值（图片太高会和底部的文字重叠）
//	@return *gg.Context
func DrawNFTImg(dc *gg.Context, x, y int, nftImg string, height int) *gg.Context {
	matched, err := regexp.MatchString(`^http(s{0,1})://(.+)`, nftImg)
	if err != nil {
		panic(err)
	}
	if matched {
		// TODO: 下载文件到本地
		nftImg = Download(nftImg, "/Users/yuyongpeng/GolandProjects/test/")
		//nftImg = "nft.jpg"
	}
	nftImg2, openErr := imaging.Open(nftImg)
	if openErr != nil {
		panic(openErr)
	}
	// 限定高度
	nft := imaging.Resize(nftImg2, 0, height, imaging.Lanczos)
	dc.DrawImageAnchored(nft, x, y, 0.5, 0.5)
	return dc
}

// DrawContentCollection
//
//	@Description:  画 集合 证书的文字内容
//	@param dc
//	@param coll		文字信息
//	@param fontR	文字的颜色R值
//	@param fontG	文字的颜色G值
//	@param fontB	文字的颜色B值
func DrawContentCollection(dc *gg.Context, coll CertificateCollection, fontR, fontG, fontB int) *gg.Context {
	// Draw key
	fontFace, err := getOpenTypeFontFace(msYaHeiBoldFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*fontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	var initX float64 = 503
	var initY float64 = 1510
	var incr float64 = 80
	var title []string = []string{"作  品  类  别：",
		"首次发表时间：",
		"数字资产编号：",
		"发  行  价  格：",
		"Hashii      值：",
		"持     有     人：",
		"登     记     号：",
		"作   品  名  称：",
		"作              者：",
		"创作完成时间："}
	for i := 0; i < len(title); i++ {
		//dc.DrawStringAnchored(title[i], initX, initY, 0.5, 0.5)
		dc.DrawStringWrapped(title[i], initX, initY, 0.5, 0.5, 200, 2.5, gg.AlignRight)
		initY += incr
	}

	// Draw value
	songtiFontFace, err := getOpenTypeFontFace(songtiFontFaceFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	initX = 615
	initY = 1528
	incr = 80
	title = []string{coll.catetory, // 作品类别
		coll.firstPublish,           // 首次发表时间
		coll.DigitalAssetsNO,        // 数字资产编号
		coll.issuePrice,             // 发行价格
		coll.hashiiValue,            // Hashii 值
		coll.owner,                  // 持有人
		coll.registrationNO,         // 登记号
		coll.productName,            // 作品名称
		coll.author,                 // 作者
		coll.creationCompletionTime, // 创作完成时间
	}
	for i := 0; i < len(title); i++ {
		dc.DrawString(title[i], initX, initY)
		initY += incr
	}
	return dc
}

// DrawContentNFT
//
//	@Description: 画 nft 证书的 文字内容
//	@param dc
//	@param cert		文字信息
//	@param fontR	文字的颜色R值
//	@param fontG	文字的颜色G值
//	@param fontB	文字的颜色B值
//	@return *gg.Context
func DrawContentNFT(dc *gg.Context, cert CertificateNFT, fontR, fontG, fontB int) *gg.Context {
	// Draw key
	fontFace, err := getOpenTypeFontFace(msYaHeiBoldFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*fontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	var initX float64 = 503
	var initY float64 = 1510
	var incr float64 = 80
	var title []string = []string{"作  品  类  别：",
		"首次发表时间：",
		"发  行  数  量：",
		"发  行  价  格：",
		"登     记     号：",
		"作   品  名  称：",
		"作              者：",
		"创作完成时间："}
	for i := 0; i < len(title); i++ {
		//dc.DrawStringAnchored(title[i], initX, initY, 0.5, 0.5)
		dc.DrawStringWrapped(title[i], initX, initY, 0.5, 0.5, 200, 2.5, gg.AlignRight)
		initY += incr
	}

	// Draw value
	songtiFontFace, err := getOpenTypeFontFace(songtiFontFaceFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*songtiFontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	initX = 615
	initY = 1528
	incr = 80
	//title = []string{"数字文创", "2023/11/29", "30", "￥1~￥100", "20231226161433985172082240120", "1209-1344测试", "daop-12091344wangmj", "2023/11/29"}
	title = []string{cert.catetory, cert.firstPublish, cert.publishNum, cert.issuePrice, cert.registrationNO, cert.productName, cert.author, cert.creationCompletionTime}
	for i := 0; i < len(title); i++ {
		dc.DrawString(title[i], initX, initY)
		initY += incr
	}
	return dc
}

// DrawEnd
//
//	@Description: 画证书的底部文字
//	@param dc
//	@param x	X坐标
//	@param y	Y坐标
//	@param RegistDate	登记日期
//	@param fontR	字体颜色R值
//	@param fontG	字体颜色G值
//	@param fontB	字体颜色B值
//	@return *gg.Context
func DrawEnd(dc *gg.Context, x, y float64, RegistDate string, fontR, fontG, fontB int) *gg.Context {
	fontFace, err := getOpenTypeFontFace(msYaHeiBoldFile, 27, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*fontFace)
	dc.SetRGB255(fontR, fontG, fontB)
	var qianzhanX float64 = 1163
	var qianzhanY float64 = 2303
	dc.DrawString("登记机构签章", qianzhanX, qianzhanY)
	dc.DrawString("登记日期："+RegistDate, qianzhanX, qianzhanY+40)
	return dc
}

// DrawCert
//
//	@Description:  画 北数所 印章
//	@param dc
//	@param x 坐标
//	@param y 坐标
//	@param certFile  印章文件(本地)
func DrawCert(dc *gg.Context, x, y int, certFile string) *gg.Context {
	// Draw Cert
	certificate, openErr := imaging.Open(certFile)
	if openErr != nil {
		panic(openErr)
	}
	dc.DrawImage(certificate, x, y)
	return dc
}

func generateWxImage(title string, savingFileName string) {
	width := 1050
	height := 442
	dc := gg.NewContext(width, height)
	dc.SetRGB255(47, 54, 66)
	dc.Clear()

	fontFilePath := "FangZhengKaiTiJianTi.TTF"
	fontFace, err := getOpenTypeFontFace(fontFilePath, 76, 72)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(*fontFace)
	dc.SetRGB255(238, 241, 247)

	// 文本框最大宽度
	maxWordsWidth := 660.0
	x := 665.0
	y := 224.0
	dc.DrawStringWrapped(title, x, y, 0.5, 0.6, maxWordsWidth, 1.1, gg.AlignCenter)

	//开始压缩图片
	src, openErr := imaging.Open("logo.png")
	if openErr != nil {
		panic(openErr)
	}
	src = imaging.Resize(src, 360, 360, imaging.Lanczos)

	dc.DrawImageAnchored(src, 182, 220, 0.5, 0.5)

	dc.SavePNG(savingFileName + ".png")
}

// getOpenTypeFontFace
//
//	@Description: 	获得字体文件对应的对象
//	@param fontFilePath 	字体文件路径，只支持TTF文件格式
//	@param fontSize			字体大小
//	@param dpi				字体dpi值
//	@return *font.Face
//	@return error
func getOpenTypeFontFace(fontFilePath string, fontSize, dpi float64) (*font.Face, error) {
	fontData, fontFileReadErr := ioutil.ReadFile(fontFilePath)
	if fontFileReadErr != nil {
		return nil, fontFileReadErr
	}
	otfFont, parseErr := opentype.Parse(fontData)
	if parseErr != nil {
		return nil, parseErr
	}
	otfFace, newFaceErr := opentype.NewFace(otfFont, &opentype.FaceOptions{
		Size: fontSize,
		DPI:  dpi,
	})
	if newFaceErr != nil {
		return nil, newFaceErr
	}
	return &otfFace, nil
}

func Download(fileUrl, localPath string) (localFile string) {
	// 测试的 URL
	//strURL := "https://www.ucg.ac.me/skladiste/blog_44233/objava_64433/fajlovi/Computer%20Networking%20_%20A%20Top%20Down%20Approach,%207th,%20converted.pdf"
	strURL := fileUrl

	resp, err := http.Head(strURL)
	if err != nil {
		fmt.Println("resp, err := http.Head(strURL)  报错: strURL = ", strURL)
		log.Fatalln(err)
	}

	// fmt.Printf("%#v\n", resp)
	fileLength := int(resp.ContentLength)

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", 0, fileLength))
	// fmt.Printf("%#v", req)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient.Do(req)", "error")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 创建文件
	filename := path.Base(strURL)
	flags := os.O_CREATE | os.O_WRONLY
	f, err := os.OpenFile(localPath+"/"+filename, flags, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		log.Fatal("err")
	}
	defer f.Close()

	// 写入数据
	buf := make([]byte, 16*1024)
	_, err = io.CopyBuffer(f, resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println("io.EOF")
			return
		}
		fmt.Println(err)
		log.Fatal(err)
	}
	return localPath + "/" + filename
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
