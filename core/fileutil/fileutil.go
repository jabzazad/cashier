package fileutil

import (
	"bufio"
	"cashier-api/core/fileutil/resize"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

// File file
type File struct {
	filename string
	basePath string
	md5      string
}

// New new file
func New(mf *multipart.FileHeader) (*File, error) {
	src, err := mf.Open()
	if err != nil {
		return nil, err
	}
	defer func() { _ = src.Close() }()

	td, err := os.MkdirTemp("", "cashier-")
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", td, mf.Filename))
	if err != nil {
		return nil, err
	}
	defer func() { _ = dst.Close() }()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	hash := md5.New()
	md5File, err := os.Open(fmt.Sprintf("%s/%s", td, mf.Filename))
	if err != nil {
		return nil, err
	}
	defer func() { _ = md5File.Close() }()
	if _, err := io.Copy(hash, md5File); err != nil {
		return nil, err
	}
	md5 := hex.EncodeToString(hash.Sum(nil))
	return &File{
		filename: mf.Filename,
		basePath: td,
		md5:      strings.ToUpper(md5),
	}, nil
}

// NewIOReader new file by io reader
func NewIOReader(src io.Reader) (*File, error) {
	td, err := os.MkdirTemp("", "cashier-")
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", td, "original"))
	if err != nil {
		return nil, err
	}
	defer func() { _ = dst.Close() }()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	hash := md5.New()
	md5File, err := os.Open(fmt.Sprintf("%s/%s", td, "original"))
	if err != nil {
		return nil, err
	}
	defer func() { _ = md5File.Close() }()
	if _, err := io.Copy(hash, md5File); err != nil {
		return nil, err
	}
	md5 := hex.EncodeToString(hash.Sum(nil))
	return &File{
		filename: "original",
		basePath: td,
		md5:      strings.ToUpper(md5),
	}, nil
}

// NewWDPath new file
func NewWDPath(mf *multipart.FileHeader) (*File, error) {
	src, err := mf.Open()
	if err != nil {
		return nil, err
	}
	defer func() { _ = src.Close() }()

	err = os.MkdirAll("file", os.ModePerm)
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("file/%s", mf.Filename))
	if err != nil {
		return nil, err
	}
	defer func() { _ = dst.Close() }()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	hash := md5.New()
	md5File, err := os.Open(fmt.Sprintf("file/%s", mf.Filename))
	if err != nil {
		return nil, err
	}
	defer func() { _ = md5File.Close() }()
	if _, err := io.Copy(hash, md5File); err != nil {
		return nil, err
	}
	md5 := hex.EncodeToString(hash.Sum(nil))
	return &File{
		filename: mf.Filename,
		basePath: "file",
		md5:      strings.ToUpper(md5),
	}, nil
}

// NewFromText file
func NewFromText(text, name string) (*File, error) {
	td, err := os.MkdirTemp("", "cashier-")
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", td, name))
	if err != nil {
		return nil, err
	}

	defer func() { _ = dst.Close() }()

	w := bufio.NewWriter(dst)
	_, err = w.Write([]byte(text))
	if err != nil {
		return nil, err
	}
	hash := md5.New()
	md5File, err := os.Open(fmt.Sprintf("%s/%s", td, name))
	if err != nil {
		return nil, err
	}
	defer func() { _ = md5File.Close() }()
	if _, err := io.Copy(hash, md5File); err != nil {
		return nil, err
	}
	md5 := hex.EncodeToString(hash.Sum(nil))
	return &File{
		filename: name,
		basePath: td,
		md5:      strings.ToUpper(md5),
	}, nil
}

// NewEmptyFile new file
func NewEmptyFile(name string) (*File, error) {
	td, err := os.MkdirTemp("", "cashier-")
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", td, name))
	if err != nil {
		return nil, err
	}
	defer func() { _ = dst.Close() }()
	return &File{
		filename: name,
		basePath: td,
	}, nil
}

// Path file path
func (f *File) Path() string {
	return fmt.Sprintf("%s/%s", f.basePath, f.filename)
}

// Name file name
func (f *File) Name() string {
	return f.filename
}

// MD5 md5
func (f *File) MD5() string {
	return strings.ToUpper(f.md5)
}

// GenMD5 gen md5
func (f *File) GenMD5() error {
	hash := md5.New()
	md5File, err := os.Open(f.Path())
	if err != nil {
		return err
	}
	defer func() { _ = md5File.Close() }()
	if _, err := io.Copy(hash, md5File); err != nil {
		return err
	}
	f.md5 = hex.EncodeToString(hash.Sum(nil))
	return nil
}

// Close close
func (f *File) Close() error {
	return os.RemoveAll(f.basePath)
}

// Image filename for resize images
type Image struct {
	OriginalImageFilename string
	ThumbImageFilename    string
	TempDir               string
	Size                  image.Point
}

// ResizeImage new model resize image
func (f *File) ResizeImage(imageSize, thumbImageSize int, originalFilePath string) (*Image, error) {
	originalImage, err := imaging.Open(originalFilePath)
	if err != nil {
		return nil, err
	}
	original, thumb := "original.jpg", "thumb.jpg"
	originalResizeImage := resize.Image(originalImage, imageSize)
	err = saveImage(originalResizeImage, f.basePath, original)
	if err != nil {
		return nil, err
	}
	thumbResizeImage := resize.Image(originalImage, thumbImageSize)
	err = saveImage(thumbResizeImage, f.basePath, thumb)
	if err != nil {
		return nil, err
	}
	return &Image{
		OriginalImageFilename: original,
		ThumbImageFilename:    thumb,
		TempDir:               f.basePath,
		Size:                  originalImage.Bounds().Size(),
	}, nil
}

func saveImage(src image.Image, tempDir, filename string) (err error) {
	f, err := os.Create(fmt.Sprintf("%s/%s", tempDir, filename))
	if err != nil {
		return err
	}
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()
	err = jpeg.Encode(f, src, &jpeg.Options{
		Quality: 100,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetThumbTempPath get temp file path for thumb
func (image *Image) GetThumbTempPath() string {
	return fmt.Sprintf("%s/%s", image.TempDir, image.ThumbImageFilename)
}

// GetOriginalTempPath get temp file path for original
func (image *Image) GetOriginalTempPath() string {
	return fmt.Sprintf("%s/%s", image.TempDir, image.OriginalImageFilename)
}

// Download for download file from url
func Download(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
