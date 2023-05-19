import os
import glob
import chardet
from tqdm import tqdm


# 获取文件夹下所有txt文件路径
def get_file_paths(folder_path):
    file_paths = glob.glob(os.path.join(folder_path, "*.txt"))
    return file_paths


# 合并所有txt文件到一个新文件中
def merge_files(file_paths, output_path, threads):
    with open(output_path, 'w', encoding='utf-8') as f_output:
        # 使用tqdm包装for循环，增加进度条
        for file_path in tqdm(file_paths, desc='合并中'):
            with open(file_path, 'rb') as f:
                file_data = f.read()
                encoding = chardet.detect(file_data)['encoding']
            with open(file_path, 'r', encoding=encoding) as f:
                file_content = f.read()
                f_output.write(file_content)
                f_output.write("\n")

    print('合并完成！')


if __name__ == '__main__':
    folder_path = input("请输入txt文件所在文件夹路径：")
    file_paths = get_file_paths(folder_path)
    output_path = os.path.join(folder_path, "p.txt")

    while True:
        threads = input("请输入使用的线程数（建议不要超过CPU核心数的2-4倍）：")
        try:
            threads = int(threads)
            break
        except ValueError:
            print("请输入有效的整数！")

    merge_files(file_paths, output_path, threads)
