package bucket

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"strings"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/minio/minio-go"
)

type PathExtra struct {
	EthAddr string
}

type Image interface {
	UploadContentImage(rawB64Image, folder, imageName string) (*pb_nft.ImageList, error)
}

// upload image to bucket return url
func (b *Bucket) uploadImageToBucket(img io.Reader, folder, imageName, contentType string) (string, error) {
	ext := fileExtensionFromContentType(contentType)
	fp := b.getImageFullPath(folder, imageName, ext)

	userMetaData := map[string]string{"x-amz-acl": "public-read"} // make it public
	cacheControl := "max-age=31536000"

	bs, _ := ioutil.ReadAll(img)

	r := bytes.NewReader(bs)

	_, err := b.Client.PutObject(b.S3BucketName, fp, r, int64(len(bs)), minio.PutObjectOptions{ContentType: contentType, CacheControl: cacheControl, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("PutObject:err [%v]", err.Error())
	}

	return b.GetCDNURL(fp), nil
}

func getB64ImageFromString(rawB64Image string) (*B64Image, error) {
	ss := strings.Split(rawB64Image, ";base64,")
	if len(ss) != 2 {
		return nil, fmt.Errorf("getB64ImageFromString:bad base64 image")
	}
	return &B64Image{
		Content:     []byte(ss[1]),
		ContentType: ss[0],
	}, nil

}

func (b64Img *B64Image) b64ToImage() (image.Image, error) {
	var img image.Image
	var err error
	switch b64Img.ContentType {
	case "data:image/jpeg":
		img, err = jpgFromB64(b64Img.Content)
		if err != nil {
			return nil, fmt.Errorf("b64ToImage:JPGFromB64: [%v]", err.Error())
		}
	case "data:image/png":
		img, err = pngFromB64(b64Img.Content)
		if err != nil {
			return nil, fmt.Errorf("b64ToImage:PNGFromB64: [%v]", err.Error())
		}
	default:
		return nil, fmt.Errorf("b64ToImage:PNGFromB64: File type is not supported [%s]", b64Img.ContentType)
	}
	return img, err
}

func imageFromString(rawB64Image string) (image.Image, error) {
	b64Img, err := getB64ImageFromString(rawB64Image)
	if err != nil {
		return nil, err
	}
	return b64Img.b64ToImage()
}

// upload single image with defined quality and	prefix to bucket
func (b *Bucket) uploadSingleImage(img image.Image, quality int, folder, imageName string) (string, error) {
	var buf bytes.Buffer
	imgWriter := bufio.NewWriter(&buf)

	err := encodeJPG(imgWriter, img, quality)
	if err != nil {
		return "", fmt.Errorf("Upload:EncodeJPG: [%v]", err.Error())
	}

	imgReader := bufio.NewReader(&buf)
	url, err := b.uploadImageToBucket(imgReader, folder, imageName, contentTypeJPEG)
	if err != nil {
		return "", fmt.Errorf("Upload:uploadImageToBucket: [%v]", err.Error())
	}
	return url, nil
}

// compose internal image object (with FullSize & Compressed formats) and upload it to S3
func (b *Bucket) uploadImageObj(img image.Image, folder, imageName string) (*pb_nft.ImageList, error) {
	imgObj := &pb_nft.ImageList{}
	var err error

	imgObj.FullSize, err = b.uploadSingleImage(img, 100, folder, fmt.Sprintf("%s_%s", imageName, "og"))
	if err != nil {
		return nil, fmt.Errorf("UploadProductImage:Upload:FullSize [%v]", err.Error())
	}

	imgObj.Compressed, err = b.uploadSingleImage(img, 60, folder, fmt.Sprintf("%s_%s", imageName, "compressed"))
	if err != nil {
		return nil, fmt.Errorf("UploadProductImage:Upload:Compressed [%v]", err.Error())
	}
	return imgObj, nil
}

// get raw image from b64 encoded string and upload full size and compressed images to s3
func (b *Bucket) UploadContentImage(rawB64Image, folder, imageName string) (*pb_nft.ImageList, error) {
	img, err := imageFromString(rawB64Image)
	if err != nil {
		return nil, err
	}
	return b.uploadImageObj(img, folder, imageName)
}
