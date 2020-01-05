package helper

import (
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// "encoding/json"
	// "fmt"
)

type Response struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	// Data gin.H `json:"data"`
}

func Responses(res Response, c *gin.Context) {
	msg := res.Message
	if msg == "" {
		res.Message = "Success!"
	}
	c.JSON(http.StatusOK, res)
}

// var key1 = "total"
// var key2 = "provinces"

// func (s Set) MarshalJSON() ([]byte, error) {
// 	// var s s
// 	// key1 := "total"
// 	// key2 := "provinces"
//     data := map[string]interface{}{
// 		key1: s.Records,
// 		key2: s.Results,
//     }
//     return json.Marshal(data)
// }
// type MyContext struct {
// 	*gin.Context
// }

func ErrorCustomStatus(code int, msg string, c *gin.Context) {
	var (
		// m MyContext
		// rw http.ResponseWriter
		res gin.H
	)
	newMsg := msg
	switch code {
	case 400:
		if newMsg == "" {
			newMsg = "Server tidak dapat mengenali permintaan anda. Silakan cek kembali URL yang ingin dituju!"
		}
		break
	case 401:
		if newMsg == "" {
			newMsg = "Anda tidak memiliki hak akses untuk membuka file/folder yang diminta!"
		}
		break
	case 403:
		if newMsg == "" {
			newMsg = "Anda tidak memiliki izin untuk mengakses halaman ini!"
		}
		break
	case 404:
		if newMsg == "" {
			newMsg = "Halaman/URL yang dituju tidak ditemukan!"
		}
		break
	case 405:
		if newMsg == "" {
			newMsg = "Terjadi kesalahan dalam mengakses URL ini!"
		}
		break
	case 408:
		if newMsg == "" {
			newMsg = "Waktu memuat server telah habis!"
		}
		break
	case 415:
		if newMsg == "" {
			newMsg = "Jenis media yang anda unggah tidak didukung/diizinkan oleh server!"
		}
		break
	case 422:
		if newMsg == "" {
			newMsg = "Terjadi kesalahan dalam menginput data. Silahkan anda cek kembali!"
		}
		break
	case 503:
		if newMsg == "" {
			newMsg = "Layanan server tidak tersedia untuk saat ini!"
		}
		break
	case 504:
		if newMsg == "" {
			newMsg = "Server sedang sibuk!"
		}
		break
	}

	res = gin.H{
		"code":    code,
		"message": newMsg,
	}
	c.JSON(code, res)
}

// func Resulting(code int, res gin.H, c *gin.Context) {
// 	c.JSON(code, res)
// 	// return nil
// }

func ErrorQuery(msg string, c *gin.Context) {
	var res gin.H
	// var n *gin.Context
	// n = c
	// fmt.Println(n)
	newMsg := msg
	if newMsg == "" {
		newMsg = "Internal Server Error!"
	}
	res = gin.H{
		"code":    500,
		"message": newMsg,
	}
	c.JSON(http.StatusInternalServerError, res)
}

// converter

func DateToTimestamp(date string) string {
	f1 := regexp.MustCompile(`\d{4}-\d{2}-\d{2}$`)
	f2 := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	if f1.MatchString(date) == true {
		unix, err := time.Parse("2006-01-02", date)
		if err != nil {
			return date
		}
		result := strconv.FormatInt(unix.Unix(), 10)
		return result
	} else if f2.MatchString(date) == true {
		unix, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			return date
		}
		result := strconv.FormatInt(unix.Unix(), 10)
		return result
	} else {
		return date
	}
}

func TimestampToDate(tmp int64) interface{} {
	// times, err := strconv.ParseInt(tmp, 10, 64)
	// if err != nil {
	// 	return tmp
	// }
	res := time.Unix(tmp, 0).Format("2006-01-02 15:04:05")
	// result, err := time.Parse("2006-01-02 15:04:05", res)
	// if err != nil {
	// 	return res
	// }
	return res
}

func Int64ToString(str int64) string {
	strings := strconv.FormatInt(int64(str), 10)
	return strings
}

func IsArray(array interface{}) bool {
	arr := reflect.ValueOf(array)

	if arr.Kind() != reflect.Array {
		return false
	}

	return true
}

func InArray(array interface{}, item interface{}) bool {
	arr := reflect.ValueOf(array)

	if arr.Kind() != reflect.Array {
		panic("Tipe data yang diinput harus array!")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
