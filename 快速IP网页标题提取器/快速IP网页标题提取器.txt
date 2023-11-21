import requests
from bs4 import BeautifulSoup
import chardet
from concurrent.futures import ThreadPoolExecutor
from tqdm import tqdm
import sys

def get_title(ip):
    url = f'http://{ip}'

    try:
        # 发送HEAD请求
        response = requests.head(url, headers={'Accept-Encoding': 'utf-8'}, timeout=5)
        response.raise_for_status()

        # 获取编码
        encoding = chardet.detect(response.content)['encoding']

        # 发送GET请求，只获取标题
        response = requests.get(url, headers={'Accept-Encoding': 'utf-8'}, timeout=5)
        response.raise_for_status()

        # 使用BeautifulSoup解析网页内容
        soup = BeautifulSoup(response.content, 'html.parser', from_encoding=encoding)

        # 获取网页标题
        title = soup.title.text.strip()

        # 将结果写入txt文件
        with open('titles.txt', 'a', encoding='utf-8') as output_file:
            output_file.write(f'{url}: {title}\n')

    except requests.exceptions.RequestException as e:
        # 将错误信息写入error.log文件
        with open('error.log', 'a', encoding='utf-8') as error_file:
            error_file.write(f'Error fetching {url}: {str(e)}\n')
    except Exception as e:
        # 忽略其他异常，不打印 "An error occurred:"
        pass

def process_ips(ip_addresses):
    # 使用ThreadPoolExecutor并发处理
    with ThreadPoolExecutor(max_workers=50) as executor:
        # 使用tqdm创建进度条
        with tqdm(total=len(ip_addresses), desc="Processing IPs", unit="IP", file=sys.stdout, position=0, leave=True) as pbar:
            # 使用submit而不是map以便在循环内更新进度条
            futures = [executor.submit(get_title, ip) for ip in ip_addresses]
            for future in futures:
                try:
                    future.result(timeout=10)  # 调整timeout的值
                except KeyboardInterrupt:
                    # 如果收到Ctrl+C，中断程序并退出
                    print('\n用户中断，正在退出...')
                    sys.exit(1)
                except Exception as e:
                    print(f'An error occurred: {str(e)}')
                finally:
                    pbar.update(1)  # 更新进度条

# 读取IP地址文件
with open('ip.txt', 'r') as file:
    ip_addresses = file.read().splitlines()

try:
    process_ips(ip_addresses)
except KeyboardInterrupt:
    print('\n用户中断，正在退出...')
    sys.exit(1)

print('任务完成，结果已写入titles.txt文件。')
