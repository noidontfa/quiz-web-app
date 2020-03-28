package main

//import (
//	"./controller"
//	router "./http"
//	"fmt"
//	"net/http"
//)
//
//var (
//	postController = controller.NewPostController()
//	httpRouter = router.NewMuxRouter()
//)
//
//
//func main() {
//	const port string = ":8000"
//	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
//		fmt.Fprint(resp, "hello world")
//	})
//	//
//	httpRouter.GET("/posts", postController.AddPost )
//	httpRouter.POST("/posts", postController.GetPosts)
//
//	httpRouter.SERVE(port)
//}
import (
	"./config"
	"./repository"
)

var (
	cf = config.NewConfiguration()
)

func main() {
	config, err := cf.GetConfig()
	if err != nil {
		panic(err.Error())
	}
	repo := repository.NewMySqlRepository(config)
	repo.AutoMigration()


}