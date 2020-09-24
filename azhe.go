package main

import (
	"fmt"
	"net/http"
)

func main()  {
	testHttpNewReques()
}

func testHttpNewReques()  {
	//第一步,创建一个客户端
	client := http.Client{}
	//第二步,请求方法既可以是GET,也可以是POST
	request,err := http.NewRequest("GET","https://www.taotiao.com/search/sugggest/initial_page/",nil)
	CheckErr(err)
	//第三步,客户端发送请求
	cookName := &http.Cookie{"username","Steven"}
	//添加cookie
	request.AddCookie(cookName)
	response,err := client.Do(request)
	CheckErr(err)
	//设置请求头
	request.Header.Set("Accept-Lanauage","zh-cn")
	defer response.Body.Close()
	//查看请求头数据
	fmt.Printf("Hearder:%+v\n",response.StatusCode)
	//操作数据
	if response.StatusCode == 200 {
		//data,err := ioutil.ReadAll(response.Body)
		fmt.Println("网络请求成")
		CheckErr(err)
		//fmt.Println(string(data))

	}else {
		fmt.Println("网络请求失败",response.Status)
	}

}

//检查错误
func CheckErr(err error)  {
	fmt.Println("09----------")
	defer func() {
		if  ins,ok := recover() . (error);ok{
			fmt.Println("程序出现异常:",ins.Error())

		}


	}()
	if err != nil {
		panic(err)
	}


}







