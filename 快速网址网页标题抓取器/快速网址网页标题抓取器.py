import requests
import threading
from bs4 import BeautifulSoup
from concurrent.futures import ThreadPoolExecutor, as_completed
from tqdm import tqdm
import logging
from urllib3.exceptions import InsecureRequestWarning
import signal
import sys

# 禁用 SSL 证书验证警告
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

# 配置日志记录到文件
logging.basicConfig(filename='error.log', level=logging.ERROR, format='%(asctime)s [%(levelname)s]: %(message)s', datefmt='%Y-%m-%d %H:%M:%S')

# 读取包含网址的txt文件
with open('web.txt', 'r', encoding='utf-8') as file:
    urls = file.readlines()

# 创建一个锁，以确保多线程写文件时不会发生冲突
lock = threading.Lock()

# 用于捕获 Ctrl+C 信号的标志
exit_signal = False

# 定义一个信号处理函数
def signal_handler(signal, frame):
    global exit_signal
    exit_signal = True
    print('\nCtrl+C detected. Exiting gracefully.')

# 注册信号处理函数
signal.signal(signal.SIGINT, signal_handler)

# 定义一个函数，用于访问网页并写入文件
def process_url(url):
    global exit_signal  # 使用全局变量标志是否接收到 Ctrl+C 信号
    url = url.strip()  # 去除换行符和空格

    try:
        # 发送HTTP请求获取网页内容，禁用SSL证书验证
        response = requests.get(url, verify=False)

        # 检查是否接收到 Ctrl+C 信号
        if exit_signal:
            return

        # 检查请求是否成功
        if response.status_code == 200:
            # 获取网页内容的编码
            encoding = response.encoding

            # 使用BeautifulSoup解析HTML内容，指定编码
            soup = BeautifulSoup(response.content, 'html.parser', from_encoding=encoding)

            # 获取网页标题
            title = soup.title.string if soup.title else 'No Title'

            # 将网页标题写入txt文件
            with lock:
                with open('titles.txt', 'a', encoding='utf-8') as output_file:
                    output_file.write(f'{url}: {title}\n')
        else:
            logging.error(f'Failed to fetch {url}. Status code: {response.status_code}')

    except Exception as e:
        logging.error(f'Error processing {url}: {str(e)}')

# 创建线程池
max_threads = 100
with ThreadPoolExecutor(max_threads) as executor:
    # 创建一个字典，将每个线程与其对应的URL关联起来
    thread_url_mapping = {executor.submit(process_url, url): url for url in urls}

    # 启动所有线程，并使用tqdm创建进度条
    with tqdm(total=len(urls), desc="Processing URLs", unit="URL") as pbar:
        # 监控线程完成情况
        for future in as_completed(thread_url_mapping):
            url = thread_url_mapping[future]
            try:
                future.result()
            except Exception as exc:
                logging.error(f'Error processing {url}: {str(exc)}')
            
            pbar.update(1)
            
            # 检查是否接收到 Ctrl+C 信号
            if exit_signal:
                break

print('任务完成！')