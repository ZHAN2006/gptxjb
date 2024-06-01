# CentOS7系统升级gcc

# 下载并运行脚本以升级系统软件
bash <(curl -sSL https://linuxmirrors.cn/main.sh) --source https://mirror.iscas.ac.cn/ --protocol https --install-epel true --close-firewall true --upgrade-software true --clean-cache true

# 安装dnf包管理器
yum install dnf -y
yum install dnf -y

# 安装编译所需的软件包
dnf install make nano bison wget bzip2 gcc glibc-headers -y

# 安装Python3所需的依赖
dnf install zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel -y

# 升级make
cd ~/
wget https://mirror.iscas.ac.cn/gnu/make/make-4.4.1.tar.gz
tar xf make-4.4.1.tar.gz
cd make-4.4.1
mkdir build
cd build

# 配置并安装make
../configure --prefix=/usr/
make && make install
make -v
cd ~
rm -rf ~/make-4.4.1.tar.gz
rm -rf ~/make-4.4.1

mkdir ~/g
cd ~/g


# 下载并解压工具包
wget https://mirror.iscas.ac.cn/gnu/gmp/gmp-6.3.0.tar.xz
wget https://mirror.iscas.ac.cn/gnu/mpfr/mpfr-4.2.1.tar.xz
wget https://mirror.iscas.ac.cn/gnu/mpc/mpc-1.3.1.tar.gz
wget https://gcc.gnu.org/pub/gcc/infrastructure/isl-0.24.tar.bz2
wget https://mirror.iscas.ac.cn/gnu/gettext/gettext-0.22.5.tar.xz
tar xf gmp-6.3.0.tar.xz
tar xf mpfr-4.2.1.tar.xz
tar xf mpc-1.3.1.tar.gz
tar xf isl-0.24.tar.bz2
tar xf gettext-0.22.5.tar.xz

# 安装编译工具包
cd ~/g/gettext-0.22.5
./configure
make && make install 
cd ~/g/gmp-6.3.0
./configure
make && make install 
cd ~/g
cd ~/g/mpfr-4.2.1
./configure
make && make install 
cd ~/g
cd ~/g/mpc-1.3.1
./configure
make && make install 
cd ~/g
cd ~/g/isl-0.24
./configure
make && make install 
cd ~

rm -rf ~/g

# 升级gcc
wget https://mirror.iscas.ac.cn/gnu/gcc/gcc-14.1.0/gcc-14.1.0.tar.xz
tar xf gcc-14.1.0.tar.xz
cd gcc-14.1.0
mkdir build
cd build

# 配置并安装gcc
dnf install bison wget bzip2 gcc gcc-c++ glibc-headers -y
export LD_LIBRARY_PATH=/usr/local/lib
../configure --enable-bootstrap --enable-checking=release --enable-languages=c,c++ --disable-multilib --with-gmp=/usr/local/lib --with-mpfr=/usr/local/lib --with-mpc=/usr/local/lib
make
dnf remove gcc g++ -y
make install
cd ~
rm -rf gcc-14.1.0.tar.xz
rm -rf gcc-14.1.0

# 显示安装后信息
which make
which gcc
make -v
gcc -v
