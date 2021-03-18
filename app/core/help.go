package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	sanitizer "github.com/go-sanitize/sanitize"
	"gopkg.in/mgo.v2/bson"
)

// getIDParamFromURL returns the ID bson parameter from the URL and makes sure is propperly coded
func GetIDParamFromURL(c *gin.Context, paramName string) bson.ObjectId {
	var id bson.ObjectId
	if bson.IsObjectIdHex(c.Param(paramName)) {
		id = bson.ObjectIdHex(c.Param(paramName))
		return id
	}
	panic("ID from URL not an Object ID!")
}

// validateSanitaze validates and sanitizes the inputs to avoid code injection
func ValidateSanitaze(c *gin.Context, st interface{}, v *validator.Validate) {
	// Validation (with Gin)
	if err := c.Bind(st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Customs validations
	err := v.Struct(st)
	if err != nil {
		fmt.Print("Error: ", err)
		panic("It has beeen an error with the custom validation")
	}
	// Sanitization (with Sanitizer)
	s, err := sanitizer.New()
	if err != nil {
		fmt.Print("Error: ", err)
		panic("It has beeen an error sanitizing the data")
	}
	err = s.Sanitize(st)
	if err != nil {
		fmt.Print("Error: ", err)
		panic("It has beeen an error sanitizing the data")
	}
}

// getPhotoFromHTML gets the Photo information of the HTML in a POST request with server validation
// and then saves it in the directory given. If not directory given, it creates it
func GetPhotoFromHTML(c *gin.Context, name string, IDelement string, directory string) string {
	file, err := c.FormFile(name)
	if err != nil {
		fmt.Print("Error: ", err)
		panic("Error retrieving the file from HTML")
	}
	if file.Filename != "" {
		file.Filename = IDelement + path.Ext(file.Filename)
		if err := c.SaveUploadedFile(file, directory+file.Filename); err != nil {
			os.MkdirAll(directory, os.ModePerm)
			if err = c.SaveUploadedFile(file, directory+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("Upload photo err: %s", err.Error()))
				panic("Upload photo err")
			}
		}
	} else {
		panic("File not retrieved from HTML")
	}
	return file.Filename
}

// getMultiplePhotosFromHTML gets the multiple photos from the HTML
func GetMultiplePhotosFromHTML(c *gin.Context, IDelement string, directory string, photos ...string) []string {
	var files []string
	for i, fphoto := range photos {
		files = append(files, GetPhotoFromHTML(c, fphoto, IDelement+strconv.Itoa(i), directory))
	}
	return files
}

// copyFile copies a single file from src to dst, taking care of creating the destiny directory if does not exist
// The src and the dst should be the FULL ROUTE TO THE FILE, INCLUDING THE EXTENSION
func CopyFile(src, dst string) error {
	// First create the directory just in case is not there
	last := strings.LastIndex(dst, "/")
	os.MkdirAll(dst[0:last+1], os.ModePerm)

	fmt.Println("Copy " + src)
	fmt.Println("Into " + dst[0:last+1])

	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	fmt.Println("Copy file finished executing")

	return os.Chmod(dst, srcinfo.Mode())
}

// copyDirectory copies a whole directory recursively, taking care of creating the destiny directory if does not exist
func CopyDirectory(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDirectory(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
