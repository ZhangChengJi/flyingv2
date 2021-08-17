package api

type App struct {
}

/**
 /registry/app/user-server {appId:user-server,name: 用户服务,group: [dev,test,pro]}
/registry/app/user-server
/registry/app/user-server

/registry/group/dev  {name: dev,txt:"测试库"}

/registry/group/app/dev {[]}

*/

func (app *App) GetList() {

}
