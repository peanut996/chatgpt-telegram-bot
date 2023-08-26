
# chatgpt-telegram-bot

Telegram体验地址：[ChatGPTBot](https://t.me/simple8964bot)

## 启动

在项目的根目录下创建配置文件：
    
```bash
cp config.example.yaml config.yaml
```

引入ChatGPT的代理模块：
    
```bash
git submodule init
```

修改配置文件中的`secret`和`token`，然后运行：
    
```bash
docker-compose up --build -d
```


