import os
import chardet
import concurrent.futures

# 需要合并的所有txt文件所在的文件夹路径
folder_path = "./"

# 新文件的名称和路径
new_file_name = "p.txt"
new_file_path = os.path.join(folder_path, new_file_name)

# 逐行读取文件内容并逐行写入新文件
def merge_file(file_path):
    with open(file_path, 'rb') as binary_file:
        encoding = chardet.detect(binary_file.read())['encoding']
    with open(file_path, 'r', encoding=encoding) as single_file:
        with open(new_file_path, 'a', encoding='utf-8') as merged_file:
            for line in single_file:
                merged_file.write(line)

# 使用多线程遍历文件夹中的所有txt文件（包括子文件夹中的txt文件），并将它们的内容添加到新文件中
def merge_files_with_threads(thread_num):
    with concurrent.futures.ThreadPoolExecutor(max_workers=thread_num) as executor:
        future_to_file = {executor.submit(merge_file, os.path.join(root, file_name)): os.path.join(root, file_name)
                          for root, dirs, files in os.walk(folder_path)
                          for file_name in files if file_name.endswith('.txt')}

        # 检查任务并处理异常
        for future in concurrent.futures.as_completed(future_to_file):
            file_path = future_to_file[future]
            try:
                _ = future.result()
            except Exception as exc:
                print('%r 处理文件 %s 时出现错误: %s' % (exc, file_path, exc))
    print("合并完成！")

#修改这里的参数可改变使用多少线程
merge_files_with_threads(thread_num=120)
