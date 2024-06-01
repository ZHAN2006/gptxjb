# CentOS7系统升级Python

# 下载并运行脚本以修改系统源
bash <(curl -sSL https://linuxmirrors.cn/main.sh) --source https://mirror.iscas.ac.cn/ --protocol https --install-epel true --close-firewall true --upgrade-software false --clean-cache true --ignore-backup-tips

# 安装dnf包管理器
yum install dnf -y
yum install dnf -y

# 更新系统
dnf upgrade -y

# 安装编译所需的软件包
dnf install make nano bison wget bzip2 gcc glibc-headers -y

# 安装Python3所需的依赖
dnf install zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel -y

# 下载并解压Python3安装包
cd ~
wget https://mirrors.huaweicloud.com/python/3.9.19/Python-3.9.19.tar.xz
tar xf Python-3.9.19.tar.xz
cd Python-3.9.19
mkdir build
cd build

# 配置Python3
../configure
make && make install

# 配置Python3软链接
mv /usr/bin/python /usr/bin/python.bak
ln -s /usr/local/bin/python3 /usr/bin/python
ln -s /usr/local/bin/pip3 /usr/bin/pip

# 修改yum和urlgrabber-ext-down的头部
sed -i '1s/$/2/' /usr/bin/yum
sed -i '1s/$/2/' /usr/libexec/urlgrabber-ext-down
cd ~
rm -rf ~/Python-3.9.19.tar.xz
rm -rf ~/Python-3.9.19

# 显示安装后信息
which python
which python3
which python2
which pip3
which pip
python -V
pip -V