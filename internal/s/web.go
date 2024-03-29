package s

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var tmpl embed.FS

type CliInfo struct {
	Vip    string
	IP     string
	VpnIPs template.HTML
	LanIPs template.HTML

	CreateTime     string
	LastAccessTime string
	TransBytes     string
}

type ClientView interface {
	GetCliInfos() []*CliInfo
}

func RunWeb(listen string, cv ClientView) {
	r := gin.New()
	t, _ := template.ParseFS(tmpl, "templates/*.tpl")
	r.SetHTMLTemplate(t)
	r.Use(Cors())
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web.tpl", gin.H{
			"clis": cv.GetCliInfos(),
		})
	})

	srv := &http.Server{
		Addr:    listen,
		Handler: r,
	}

	log.Println("listen on:", listen)

	log.Fatal(srv.ListenAndServe())
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// nolint: lll
			c.Header("Access-Control-Allow-Headers",
				"Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//  允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers",
				"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
