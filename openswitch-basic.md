# Open vSwitch

## 安装

这里介绍从源码安装openvswitch-2.14.1，下载地址是https://www.openvswitch.org/releases/openvswitch-2.14.1.tar.gz。

1. 先安装一些依赖包： `sudo apt-get install autoconf make libtool`
2. 解压文件并编译：
```bash
tar -zxf openvswitch-2.14.1.tar.gz
cd openvswitch-2.14.1
./configure --with-linux=/lib/modules/`uname -r`/build 
make 
```
3. 安装执行文件和man页:  `make install` . 默认路径在/usr/local,执行文件会放在`/usr/local/bin/`，同时会创建一个空文件夹`/usr/local/etc/openvswitch/`,库文件被装入`/usr/local/lib`.
4. 编译安装内核模块：
   - `make modules_install`将内核模块装至`/lib/modules/5.4.0-72-generic/extra`,这一步会有关于签名的Err报错，但实际是warning，提示模块未签名，可以忽略。
   - 备份系统的原始模块文件： `sudo cp /lib/modules/5.4.0-67-generic/kernel/net/openvswitch/* backup`
   - 删除原始模块文件：`sudo rm /lib/modules/5.4.0-67-generic/kernel/net/openvswitch/*`
   - 为modprobe命令运行更新依赖列表：`sudo depmod`
   - 加载openvswitch和需要的相关模块：`modprobe openvswitch;modprobe vport-stt;modprobe vport-lisp`
   - 检查模块是否加载： `lsmod | grep -E 'openvswitch|vport_lisp|vport_stt'`
5. 显示模块信息和依赖：`modinfo openvswitch`,`modprobe -D openvswitch`

## 参考

[从源码安装Open vSwitch](https://github.com/ebiken/doc-network/wiki/How-To:-Install-OVS-(Kernel-Module)-from-Source-Code)