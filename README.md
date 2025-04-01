# 北理工校车抢票

## 快速使用

1. 使用抓包工具抓取 http://hqapp1.bit.edu.cn/vehicle 下的任意请求（触发方式为使用i北理打开班车服务，可以随便操作一些），找到请求参数中的`apitoken`、`apitime`、`userid`填入/config/config.go的对应位置
2. 运行`go build main.go`编译
3. 可选的运行时参数：`index`、`address`、`date`，index代表要抢列表中（这个列表可以去i北理上看到，i北理上展示多少，列表就是多少）的第index + 1趟班次，不提供时默认为0，代表列表中第一趟班车；address取值为0时代表中关村to良乡，为1时代表良乡to中关村，不提供时默认为0；date参数代表抢票的日期为今日之后第date天，不提供时默认为0，代表今日之后第0天，也即今天。例如：我要抢今天良乡到中关村的第2趟校车，就可以执行 `./main -index 1 -address 1 -date 0`
