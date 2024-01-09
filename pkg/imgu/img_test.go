package imgutil

import (
	"fmt"
	"github.com/disintegration/imaging"
	_ "image/jpeg"
	"regexp"
	"testing"
)

func TestDraw(t *testing.T) {
	Draw()
}

func TestDraw2(t *testing.T) {
	//Draw2()
	matched, err := regexp.MatchString(`^http(s{0,1})://(.+)`, "dhttps://wwww.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Print(matched)
}

func TestDrawCertNFTTest(t *testing.T) {
	nft := CertificateNFT{
		company:                "星矿科技（北京）有限公司",
		author:                 "topholder",
		catetory:               "数字文创",
		issuePrice:             "￥1~￥100",
		registrationNO:         "20231226161433985172082240120",
		firstPublish:           "2023/11/29",
		productName:            "数字文创",
		creationCompletionTime: "2023/11/12",
		publishNum:             "300",
		registrationDate:       "2023年11月30日",
		nftPath:                "nft.jpg",
	}
	outFile := "outfile.png"
	DrawCertNFTTest(nft, outFile)
}

func TestDrawCertNFT(t *testing.T) {
	nft := CertificateNFT{
		company:                "星矿科技（北京）有限公司",
		author:                 "topholder",
		catetory:               "数字文创",
		issuePrice:             "￥1~￥100",
		registrationNO:         "20231226161433985172082240120",
		firstPublish:           "2023/11/29",
		productName:            "数字文创",
		creationCompletionTime: "2023/11/12",
		publishNum:             "300",
		registrationDate:       "2023年11月30日",
		nftPath:                "nft.jpg",
	}
	outFile := "outfile2.png"
	DrawCertNFT(nft, outFile)
}
func TestDrawCertCollection(t *testing.T) {
	coll := CertificateCollection{
		company: "星矿科技（北京）有限公司",
		//nftPath:                "nft.jpg",
		nftPath:                "https://oss.raw-stones.com/uploads/2023/0930/JBYamMg1696081523719.png",
		catetory:               "数字文创",
		firstPublish:           "2023/11/29",
		DigitalAssetsNO:        "1",
		issuePrice:             "￥1~￥100",
		hashiiValue:            "0x625e888e83414386a76057528fee909d36e3437£3bb822511ec6fd68bd625a19",
		owner:                  "Oxcc79d384123BC57C0770fA884P3bD7198df600A7",
		registrationNO:         "20231226161433985172082240120",
		productName:            "数字文创",
		author:                 "topholder",
		creationCompletionTime: "2023/11/12",
		registrationDate:       "2023年11月30日",
	}
	outFile := "outfile3.png"
	DrawCertCollection(coll, outFile)
}

func TestDownload(t *testing.T) {
	url := "https://oss.raw-stones.com/uploads/2023/0930/JBYamMg1696081523719.png"
	local := "/Users/yuyongpeng/GolandProjects/test/"
	Download(url, local)
}

func TestExampleReadFrameAsJpeg(t *testing.T) {
	reader := ExampleReadFrameAsJpeg("./26_1680613282.mp4", 30)
	img, err := imaging.Decode(reader)
	if err != nil {
		t.Fatal(err)
	}
	err = imaging.Save(img, "outv12.jpeg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTest(t *testing.T) {
	//imgu := readImg("./out.png")
	//SavePic(imgu, "./out_x.gif")

	//Drawx()

	//drawCirclePic()
	//WriteFont()
	WriteFont2()
}
