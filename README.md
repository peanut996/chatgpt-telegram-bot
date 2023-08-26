
# chatgpt-telegram-bot

Telegram体验地址：[ChatGPTBot](https://t.me/simple8964bot)

## 准备

+ [chatgpt-engine](https://github.com/peanut996/chatgpt-engine)

## 运行

```bash
git clone https://github.com/peanut996/chatgpt-telegram-bot.git

cd chatgpt-telegram-bot

# 需要写入配置
vim config.yaml

go mod download

go run .
```

### 配置文件

这里只是telegram bot的控制程序，实际的chatgpt通信仍然要交给[chatgpt-engine](https://github.com/peanut996/chatgpt-engine),
所以需要运行起engine服务然后再启动  

下面是配置文件：
```yaml
bot:
  type: telegram
  token: <your tg bot token> #你的telegram bot token
engine:
  port: <engine-port> # engine服务host
  host: <engine-host> # engine服务端口号
```



