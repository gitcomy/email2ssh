# email2ssh
###1.终端登录，创建目录

```
mkdir 755 -p /sendmail/bin
```
###2.下载文件
```
cd /sendmail
git clone https://github.com/gitcomy/email2ssh.git
cd email2ssh
```
###3.修改配置
```
 vi /sendmail/email2ssh/sendmail.json
```
建议163邮箱SMTP发信:
```
  "host": "smtp.163.com", //SMTP地址
  "port": 25, //端口号
  "user": "***@163.com",//用户名为发信邮箱地址
  "pass": "***",//授权码，邮箱管理获取
  "from": "name <***@163.com>", //发件人填主机名方便查看，发信邮箱地址
  "to": ["***@qq.com"],  //邮箱地址
  "subject": "ssh login code" //邮件标题
```
以上需均为英文填写。
###4.准备发信
####安装go
```
wget -qO- https://raw.githubusercontent.com/skiy/golang-install/master/install.sh | sh
go env -w GO111MODULE=on
go mod init sendmail
```
####下载go发送邮件的库
```
go get gopkg.in/gomail.v2
```
####编译发送邮件的可执行程序
```
go build -o /sendmail/bin/sendmail sendmail.go
```
####复制配置文件
```
cp sendmail.json /sendmail/bin/sendmail.json
```
####复制发信脚本
```
cp sendmail.sh /bin/sendmail.sh
chmod 755 /bin/sendmail.sh
```
###测试发信
```
/bin/sendmail.sh
```
###修改登录ssh发信
```
vi /etc/passwd
```
修改第一行
```
root:x:0:0:root:/root:/bin/bash
```
为
```
root:x:0:0:root:/root:/bin/sendmail.sh
```
###重新登录ssh即可使用
