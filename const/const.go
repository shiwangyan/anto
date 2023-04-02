package _const

const (
	Author     = "speauty"
	Version    = "v2.0.0"
	EnvDebug   = "debug"
	EnvRelease = "release"

	GoUidLen = 8

	UITitle       = "Anto"
	UIIcon        = "favicon.ico"
	UIResourceDir = "./assets"

	CfgYmlDefaultContent = `
huawei_cloud_nlp: # 华为云机器翻译
  ak_id: ""             # Access Key
  sk_key: ""            # Secret Access Key
  region: "cn-north-4"  # 当前接口开发的区域, 目前仅支持华北-北京四终端节点
  project_id: ""        # 项目ID

ling_va:
  data_id: ""           # 数据ID

baidu:
  app_id: ""            # 应用ID
  app_key: ""           # 应用密钥

tencent_cloud_mt:       # 腾讯云机器翻译
  secret_id: ""         # 密钥ID 用于标识接口调用者身份
  secret_key: ""        # 密钥关键字 用于验证接口调用者的身份
`
)
