package v1
import (
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
)

func GetPushUrl(c *gin.Context)  {
	var url string
	url = "rtmp://192.168.43.250/live/1"
	response.OkWithData(gin.H{"url": url}, c)
}