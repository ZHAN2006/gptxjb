import subprocess
import requests
import re
from datetime import datetime

# 企业微信机器人 webhook 地址
webhook_url = 'https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key='

# 中文月份到英文月份的映射
month_map = {
    '1月': 'Jan', '2月': 'Feb', '3月': 'Mar', '4月': 'Apr', '5月': 'May', '6月': 'Jun',
    '7月': 'Jul', '8月': 'Aug', '9月': 'Sep', '10月': 'Oct', '11月': 'Nov', '12月': 'Dec'
}

def send_wechat_message(date_time, username, ip):
    headers = {
        "Content-Type": "application/json"
    }
    payload = {
        "msgtype": "text",
        "text": {
            "content": (
                f"**SSH 登录警告**\n"
                f"**日期时间**: {date_time}\n"
                f"**用户名**: {username}\n"
                f"**IP 地址**: {ip}"
            )
        }
    }
    response = requests.post(webhook_url, headers=headers, json=payload)
    if response.status_code == 200:
        print(f"成功发送警告消息：{ip}")
    else:
        print(f"发送警告消息失败：{response.text}")

def convert_to_standard_time(date_time):
    # 提取月份和日期时间
    match = re.match(r'(\d+月) (\d+ \d+:\d+:\d+)', date_time)
    if match:
        month_cn = match.group(1)
        rest_date = match.group(2)
        # 将中文月份转换为英文月份
        month_en = month_map.get(month_cn, 'Jan')
        # 使用英文月份组装新的日期时间字符串
        new_date_time = f"{month_en} {rest_date}"
        # 使用当前年份组装完整的日期时间字符串
        current_year = datetime.now().year
        new_date_time = f"{current_year} {new_date_time}"
        # 将日期时间解析为标准格式
        return datetime.strptime(new_date_time, "%Y %b %d %H:%M:%S").strftime("%Y-%m-%d %H:%M:%S")
    return date_time

def monitor_journalctl():
    process = subprocess.Popen(['journalctl', '-f', '-u', 'sshd'], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    while True:
        line = process.stdout.readline().decode('utf-8')
        if line:
            # 匹配 SSH 登录成功的日期时间、用户名和 IP 地址
            match = re.search(r'(\d+月 \d+ \d+:\d+:\d+).*sshd.*Accepted .* for (\w+) from (\d+\.\d+\.\d+\.\d+)', line)
            if match:
                date_time = match.group(1)
                username = match.group(2)
                ip_address = match.group(3)
                # 转换日期时间为标准格式
                formatted_date_time = convert_to_standard_time(date_time)
                send_wechat_message(formatted_date_time, username, ip_address)

if __name__ == '__main__':
    monitor_journalctl()
