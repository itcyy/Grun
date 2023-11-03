package Grun

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

/*
* InitGrun
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @Description:
*  @return *Engine
 */
func (grun *GrunServer) InitGrun() *Engine {
	return gin.Default()
}

/*
* POSTController
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) POSTController(url string, handler gin.HandlerFunc, controller gin.HandlerFunc) GinRouter {
	return router.POST(url, controller, handler)
}

/*
* POST
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) POST(url string, handler gin.HandlerFunc) GinRouter {
	return router.POST(url, handler)
}

/*
* AdminPOST
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) AdminPOST(url string, handler gin.HandlerFunc, controller gin.HandlerFunc) GinRouter {
	return router.POST("/admin_api/"+url, controller, handler)
}

/*
* ApiPOST
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) ApiPOST(url string, handler gin.HandlerFunc) GinRouter {
	return router.POST("/api/"+url, handler)
}

/*
* VitalPOST
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) VitalPOST(url string, handler gin.HandlerFunc) GinRouter {
	return router.POST("/vital_api/"+url, handler)
}

/*
* ClientPOST
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) ClientPOST(url string, controller gin.HandlerFunc, handler gin.HandlerFunc) GinRouter {
	return router.POST("/client_api/"+url, controller, handler)
}

/*
* GETController
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) GETController(url string, handler gin.HandlerFunc, controller gin.HandlerFunc) GinRouter {
	return router.GET(url, controller, handler)
}

/*
* GET
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) GET(url string, handler gin.HandlerFunc) GinRouter {
	return router.GET(url, handler)
}

/*
* AdminGET
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) AdminGET(url string, handler gin.HandlerFunc, controller gin.HandlerFunc) GinRouter {
	return router.GET("/admin_api/"+url, controller, handler)
}

/*
* ApiGET
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) ApiGET(url string, handler gin.HandlerFunc) GinRouter {
	return router.GET("/api/"+url, handler)
}

/*
* VitalGET
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) VitalGET(url string, handler gin.HandlerFunc) GinRouter {
	return router.GET("/vital_api/"+url, handler)
}

/*
* ClientGET
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) ClientGET(url string, controller gin.HandlerFunc, handler gin.HandlerFunc) GinRouter {
	return router.GET("/client_api/"+url, controller, handler)
}

/*
* WebSocket
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) WebSocket(url string, handler gin.HandlerFunc) GinRouter {
	return router.GET("/ws/"+url, handler)
}

/*
* ApiPOSTSearch
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param handler
*  @param controller
*  @Description:
*  @return GinRouter
 */
func (grun *GrunServer) ApiPOSTSearch(url string, handler gin.HandlerFunc, controller gin.HandlerFunc) GinRouter {
	return router.POST("/api/"+url, controller, handler)
}

/*
* HttpRun
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @Description:
*  @return error
 */
func (grun *GrunServer) HttpRun(url string) error {
	return router.Run(url)
}

/*
* HttpsRun
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param url
*  @param crt
*  @param key
*  @Description:
*  @return error
 */
func (grun *GrunServer) HttpsRun(url string, crt string, key string) error {
	return router.RunTLS(url, crt, key)
}

/*
* WebSocketInit
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-11-3]
*  @receiver grun
*  @param c
*  @Description:
*  @return *websocket.Conn
*  @return error
 */
func (grun *GrunServer) WebSocketInit(c gin.Context) (*websocket.Conn, error) {
	return (&websocket.Upgrader{
		// 决解跨域问题
		CheckOrigin: func(r *http.Request) bool { return true },
	}).Upgrade(c.Writer, c.Request, nil)
}
