# bilibili 动态抽奖工具

## 简介

使用 GO 语言编写的动态抽奖工具 (non-GUI), 支持多种抽奖方式以及自定义权重

使用 SQLite 存储数据

### 功能
1. 扫码登录，无需用户手动获取 Cookies
2. 可获取转发、评论、点赞列表
3. 可自定义权重
4. 可区分用户等级以及是否大会员
5. 若抽奖用户过多，可选择持续抓取列表
6. 支持 HTTP Proxy
7. 支持自定义配置文件路径
8. **支 持 黑 幕 功 能** (选择特定用户中奖)

### 食用方法
1. 下载 Release 中的文件，哪个平台就下哪个
2. 解压文件，将 `example.toml` 复制或改名为 `config.toml`
3. 填写你的动态ID, 并可选择修改权重及可选项
4. 在终端中运行 `./bili-lottery-tool` (Windows 用户运行 `./bili-lottery-tool.exe`)
5. 可使用 -c 选择配置文件 `./bili-lottery-tool -c /etc/bili/config.toml`

## 配置文件

**config.toml**

```
#版本号
Version="1.0"

# 权重
[Weights]
Forward = 1  # 转发
Comment = 1  # 评论
Like = 1     # 点赞
Fan = 5      # 粉丝
Friend = 10  # 互关
VIP = 2      # 大会员

# 可选项
[Select]
Level = 0    # 排除几级以下的用户, 0 为不排除
Member = 0   # 黑幕, 填写 UID, 直接抽这位用户
Time = 120   # 抓取间隔, 单位 s
EndTime = 1656648000  # 结束时间, 为 0 或已超时则直接出结果

# 动态
[Dynamic]
Id = ""      # 动态ID, 需要填写
Oid = ""     # 自动获取，无需填写
Type = 0     # 自动获取，无需填写

# Cookie，无需填写
[Cookies]
SESSDATA = ""
bili_jct = ""
DedeUserID = ""
DedeUserID__ckMd5 = ""

# 代理
[ProxyURL]
api = ""
api_vc = ""
api_live = ""
```
