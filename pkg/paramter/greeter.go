/**
 *@Description
 *@ClassName greeter
 *@Date 2020/11/22 2:09 下午
 *@Author ckhero
 */

package paramter

type BaseRsp struct {
	Code int64           `json:"code"`
	Msg  string          `json:"msg"`
}
type GreeterHelloReq struct {
	Name string `json:"name"`
}

type GreeterHelloRsp struct {
	BaseRsp
	Data GreeterHelloRspData `json:"data"`
}

type GreeterHelloRspData struct {
	Greeting string `json:"greeting"`
}