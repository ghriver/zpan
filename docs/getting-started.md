## Install

### Linux
```bash
# 安装服务
curl -sSLf https://dl.saltbo.cn/install.sh | sh -s zpan

# 启动服务
systemctl start zpan

# 查看服务状态
systemctl status zpan

# 设置开机启动
systemctl enable zpan

# 查看日志
journalctl -xe -u zpan -f
```

### Docker
```bash
docker run -it -p 8222:8222 -v /etc/zpan:/etc/zpan --name zpan saltbo/zpan
```

### StartWithMinIO
```bash
mkdir localzpan && cd localzpan
curl -L https://raw.githubusercontent.com/saltbo/zpan/master/quickstart/docker-compose.yaml -o docker-compose.yaml
docker-compose up -d
```

## Usage

visit http://localhost:8222
