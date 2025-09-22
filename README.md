# 北理工校车抢票

## 功能特性

- 🚌 自动抢票：支持班车座位自动抢票
- 📧 邮件通知：购票成功后自动发送邮件通知
- ⏰ 智能等待：自动计算并等待到合适的抢票时间
- 🎯 多座位并发：同时尝试多个座位，提高成功率

## 快速使用

### 1. 配置环境变量

创建 `.env` 文件并配置以下参数：

```bash
# API配置 - 需要通过抓包获取
API_HOST=hqapp1.bit.edu.cn
API_TOKEN=你的API令牌
API_TIME=你的API时间
USER_ID=你的用户ID

# 座位配置
TOTAL_SEATS=51

# 邮件配置
MAIL_USERNAME=registercode@yaphet.top
MAIL_PASSWORD=AYaphet677958
MAIL_DEFAULT_ENCODING=UTF-8
MAIL_HOST=smtpdm.aliyun.com
MAIL_PORT=465

# 用户邮箱 - 购票成功后发送通知邮件的目标邮箱
USER_EMAIL=你的邮箱@example.com
```

### 2. 获取API参数

使用抓包工具抓取 http://hqapp1.bit.edu.cn/vehicle 下的任意请求（触发方式为使用i北理打开班车服务，可以随便操作一些），找到请求参数中的`apitoken`、`apitime`、`userid`填入环境变量

### 3. 编译运行

```bash
go build main.go
```

### 4. 运行参数

可选的运行时参数：

- `index`：要抢列表中的第几趟班次（从0开始），默认为0
- `address`：校车路线（0：中关村→良乡，1：良乡→中关村），默认为0
- `date`：抢票日期（今日后第几天），默认为0（今天）

示例：抢今天良乡到中关村的第2趟校车
```bash
./main -index 1 -address 1 -date 0
```

## 邮件通知功能

购票成功后，系统会自动发送包含以下信息的邮件通知：

- 🚌 **班车信息**：路线、车次、发车时间
- 📅 **服务日期**：乘车日期
- 💺 **座位信息**：具体座位号
- 💰 **票价信息**：学生票价格
- ⏰ **下单时间**：成功购票的时间
- 📝 **重要提醒**：乘车注意事项

邮件采用精美的HTML模板，包含详细的购票信息和乘车提醒。
