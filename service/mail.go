package service

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
	"time"

	"awesomeProject/config"
	"awesomeProject/models"
)

// MailService 邮件服务结构体
type MailService struct {
	auth smtp.Auth
}

// NewMailService 创建新的邮件服务实例
func NewMailService() *MailService {
	mailConfig := config.AppConfig.Mail
	auth := smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.Host)
	return &MailService{auth: auth}
}

// SendBookingSuccessEmail 发送购票成功邮件
func (ms *MailService) SendBookingSuccessEmail(shuttle models.ShuttleRoute, date string, seatNumber int, userEmail string) error {
	mailConfig := config.AppConfig.Mail

	// 验证邮箱格式
	if !validateEmail(userEmail) {
		return fmt.Errorf("无效的邮箱地址: %s", userEmail)
	}

	// 构建邮件内容
	from := mail.Address{Name: "校车购票系统", Address: mailConfig.Username}
	to := mail.Address{Name: "", Address: userEmail}

	subject := "校车购票成功通知"
	body := ms.generateBookingEmailBody(shuttle, date, seatNumber)

	// 构建邮件头部
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"
	headers["Content-Transfer-Encoding"] = "quoted-printable"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 使用SSL连接发送邮件
	return ms.sendMailWithTLS(mailConfig.Host, mailConfig.Port, mailConfig.Username, mailConfig.Password,
		mailConfig.Username, []string{userEmail}, []byte(message))
}

// sendMailWithTLS 使用TLS/SSL发送邮件
func (ms *MailService) sendMailWithTLS(host string, port int, username, password, from string, to []string, msg []byte) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	// 创建TLS连接
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return fmt.Errorf("TLS连接失败: %v", err)
	}
	defer conn.Close()

	// 创建SMTP客户端
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("创建SMTP客户端失败: %v", err)
	}
	defer client.Quit()

	// 身份验证
	auth := smtp.PlainAuth("", username, password, host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP身份验证失败: %v", err)
	}

	// 设置发件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("设置发件人失败: %v", err)
	}

	// 设置收件人
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return fmt.Errorf("设置收件人失败: %v", err)
		}
	}

	// 发送邮件内容
	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("获取邮件写入器失败: %v", err)
	}
	defer writer.Close()

	_, err = writer.Write(msg)
	if err != nil {
		return fmt.Errorf("写入邮件内容失败: %v", err)
	}

	return nil
}

// generateBookingEmailBody 生成购票成功邮件内容
func (ms *MailService) generateBookingEmailBody(shuttle models.ShuttleRoute, date string, seatNumber int) string {
	orderTime := time.Now().Format("2006-01-02 15:04:05")

	html := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>校车购票成功通知</title>
    <style>
        body { font-family: "Microsoft YaHei", Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; background: #f9f9f9; }
        .header { background: #4CAF50; color: white; padding: 20px; text-align: center; border-radius: 10px 10px 0 0; }
        .content { background: white; padding: 30px; border-radius: 0 0 10px 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .info-row { display: flex; justify-content: space-between; margin: 15px 0; padding: 10px; background: #f5f5f5; border-radius: 5px; }
        .label { font-weight: bold; color: #2196F3; }
        .value { color: #333; }
        .highlight { background: #E8F5E8; border-left: 4px solid #4CAF50; padding: 15px; margin: 20px 0; }
        .footer { text-align: center; color: #666; margin-top: 20px; font-size: 14px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🚌 校车购票成功通知</h1>
            <p>恭喜您！您的校车票已成功预订</p>
        </div>
        <div class="content">
            <div class="highlight">
                <h3>📋 订票信息</h3>
            </div>
            
            <div class="info-row">
                <span class="label">🚌 班车路线：</span>
                <span class="value">%s → %s</span>
            </div>
            
            <div class="info-row">
                <span class="label">🗓️ 服务日期：</span>
                <span class="value">%s</span>
            </div>
            
            <div class="info-row">
                <span class="label">🕐 发车时间：</span>
                <span class="value">%s</span>
            </div>
            
            <div class="info-row">
                <span class="label">🕐 到达时间：</span>
                <span class="value">%s</span>
            </div>
            
            <div class="info-row">
                <span class="label">🚌 班车车次：</span>
                <span class="value">%s</span>
            </div>
            
            <div class="info-row">
                <span class="label">💺 座位号：</span>
                <span class="value">%d 号</span>
            </div>
            
            <div class="info-row">
                <span class="label">💰 票价：</span>
                <span class="value">学生票 %s 元</span>
            </div>
            
            <div class="info-row">
                <span class="label">📅 下单时间：</span>
                <span class="value">%s</span>
            </div>
            
            <div class="highlight">
                <h3>📝 重要提醒</h3>
                <ul>
                    <li>请提前10分钟到达乘车地点</li>
                    <li>请携带有效证件上车</li>
                    <li>如需取消，请在发车前30分钟操作</li>
                    <li>请保存此邮件作为乘车凭证</li>
                </ul>
            </div>
        </div>
        <div class="footer">
            <p>此邮件由校车购票系统自动发送，请勿回复</p>
            <p>如有疑问，请联系相关工作人员</p>
        </div>
    </div>
</body>
</html>`

	return fmt.Sprintf(html,
		shuttle.OriginAddress,      // 起点
		shuttle.EndAddress,         // 终点
		date,                       // 服务日期
		shuttle.OriginTime,         // 发车时间
		shuttle.EndTime,            // 到达时间
		shuttle.TrainNumber,        // 班车车次
		seatNumber,                 // 座位号
		shuttle.StudentTicketPrice, // 票价
		orderTime,                  // 下单时间
	)
}

// validateEmail 验证邮箱格式
func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
