package main
import (
   "fmt"
   "github.com/gin-gonic/gin"
   "log"
   "net/http"
   "os/exec"
)

func main() {
   r := gin.Default()
   r.GET("/sync", func(c *gin.Context) {
      // 去同步我的文件
      go func() {
         cmd := exec.Command("git", "pull")
         cmd.Dir = "D:/dxm/yuncang/"
         out, err := cmd.CombinedOutput()
         fmt.Printf("sync result :\n%s\n", string(out))
         if err != nil {
            log.Printf("sync failed with %s\n", err)
         }
      }()

      c.JSON(http.StatusOK, gin.H{
         "message": "repo sync OK",
      })
   })
   r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

