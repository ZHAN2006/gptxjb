import threading
import os

def remove_duplicates(filepath, thread_count):
    with open(filepath, 'r') as f:
        lines = f.readlines()

    lines = [line.strip() for line in lines if line.strip()]
    lines_set = set(lines)

    # 创建一个新列表，用于存储去重后的数据
    unique_lines = []

    # 使用多个线程来并行地去重
    def remove_duplicates_thread(lines_set, unique_lines, thread_index):
        start = int(thread_index * len(lines_set) / thread_count)
        end = int((thread_index + 1) * len(lines_set) / thread_count)
        for line in list(lines_set)[start:end]:
            unique_lines.append(line)

    threads = []
    for i in range(thread_count):
        thread = threading.Thread(target=remove_duplicates_thread, args=(lines_set, unique_lines, i))
        threads.append(thread)
        thread.start()

    # 等待所有线程结束
    for thread in threads:
        thread.join()

    # 将去重后的数据写入文件
    output_file = os.path.join(os.path.dirname(filepath), 'dedup_' + os.path.basename(filepath))
    with open(output_file, 'w') as f:
        f.writelines(line + '\n' for line in unique_lines)
    print("文件已输出至：", output_file)

# 交互式输入
filepath = input("请输入需要去重的文件路径：")
while not os.path.exists(filepath):
    filepath = input("文件路径无效，请重新输入：")

thread_count = input("请输入线程数：")
while not thread_count.isdigit():
    thread_count = input("线程数必须是一个正整数，请重新输入：")
thread_count = int(thread_count)

remove_duplicates(filepath, thread_count)
