/**
 * @auther:  zu1k
 * @date:    2020/1/24
 */
package bot

import "github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"

func SendMsg(qqnum int64, msg string, isGroup bool) {
	if isGroup {
		cqp.SendGroupMsg(qqnum, msg)
	} else {
		cqp.SendPrivateMsg(qqnum, msg)
	}
}
