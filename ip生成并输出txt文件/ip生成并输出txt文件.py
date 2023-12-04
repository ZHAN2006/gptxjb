import ipaddress
import progressbar

def generate_ip_range(start, end):
    start_ip = ipaddress.IPv4Address(start)
    end_ip = ipaddress.IPv4Address(end)

    current_ip = start_ip
    while current_ip <= end_ip:
        yield str(current_ip)
        current_ip += 1

def write_to_file(ip_range, output_file):
    with open(output_file, 'w') as file:
        for ip in ip_range:
            file.write(ip + '\n')

def main():
    #修改这里的ip生成范围
    start_ip = '10.0.0.0'
    end_ip = '10.9.225.225'
    output_file = 'ip_addresses.txt'

    ip_range = generate_ip_range(start_ip, end_ip)
    total_ips = sum(1 for _ in ip_range)

    ip_range = generate_ip_range(start_ip, end_ip)  # 重新生成IP范围

    # 创建进度条
    bar = progressbar.ProgressBar(maxval=total_ips, widgets=[progressbar.Bar('=', '[', ']'), ' ', progressbar.Percentage()])
    bar.start()

    ip_range = generate_ip_range(start_ip, end_ip)
    write_to_file(ip_range, output_file)

    bar.finish()

if __name__ == "__main__":
    main()
