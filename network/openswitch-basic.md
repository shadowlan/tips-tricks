# Open vSwitch

## 安装

这里介绍从源码安装openvswitch-2.14.1，下载地址是https://www.openvswitch.org/releases/openvswitch-2.14.1.tar.gz。操作系统是Ubuntu20.04

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
   - 加载openvswitch和需要的相关模块：`modprobe openvswitch;modprobe vport-stt;modprobe vport-lisp;modprobe vport-vxlan;modprobe vport-geneve;modprobe vport-gre`
   - 检查模块是否加载： `lsmod | grep -E 'openvswitch|vport_lisp|vport_stt|vport_gre|vport_vxlan|vport_geneve|vport_geneve'`
5. 显示模块信息和依赖：`modinfo openvswitch`,`modprobe -D openvswitch`

## 概述

Open vSwitch架构图  
![Open vSwitch架构图](./imgs/ovs-arch.png)

![OVS internal](./imgs/ovs-internal.png)

- ovs-vswitchd, 实现交换机的守护进程，与Linux内核模块共同实现基于流的报文交换。
- ovsdb-server, 一个轻量级的数据库，ovs-vswitchd 查询以获取其配置信息。
- ovs-vsctl: 用于查询和更新 ovs-vswitchd 的配置的实用程序。
- ovs-ofctl: 用于查询和控制OpenFlow交换机和控制器的应用程序。
- ovs-dpctl: 用于配置交换机内核模块的工具。
- ovs-appctl: 一个向Open vSwitch守护进程发送命令的程序。

## 常用命令

### [ovs-vsctl](http://www.openvswitch.org/support/dist-docs/ovs-vsctl.8.txt)

```shell
# 查看bridge信息
ovs-vsctl list bridge br-int
# 查看接口信息
ovs-vsctl list interface
# 以表格方式查看指定接口信息
ovs-vsctl --columns=ofport,name --format=table list Interface
```

### [ovs-ofctl](http://www.openvswitch.org/support/dist-docs/ovs-ofctl.8.txt)

```shell
# 查看流表
ovs-ofctl dump-flows br-int -O Openflow13
ovs-ofctl dump-flows br-int table=40
# 添加流表
ovs-ofctl add-flow br-int priority=301,ip,nw_dst=226.94.1.1,actions=output:5,output:6
# 删除流表
ovs-ofctl del-flows br-int table=0,ip,nw_dst=226.94.1.1
```
### [ovs-dpctl](http://www.openvswitch.org/support/dist-docs/ovs-dpctl.8.txt)

```shell
ovs-dpctl dump-conntrack
```
### [ovs-appctl](http://www.openvswitch.org/support/dist-docs/ovs-appctl.8.txt)

```shell
# 查看datapath端口
ovs-appctl dpif/show
# 查看缓存的路由：
ovs-appctl ovs/route/show
# 追踪包
ovs-appctl ofproto/trace br-int in_port="poda-net-ea06dd",udp,nw_src=192.168.224.3,nw_dst=226.94.1.1,udp_dst=5001
```


### 其他

* 事务命令
可以将多个命令在一行执行以实现原子事务操作, 例如将下面两个命令转成事务命令: `ovs-vsctl add-br br0 -- add-port br0 eth0`，事务命令意味着这两条命令必须都执行成功，否则回退。
```
ovs-vsctl add-br br0
ovs-vsctl add-port br0 eth0
```

* 查看OVS版本: `ovs-vswitchd --version`
* 启动OVS daemon: `/usr/share/openvswitch/scripts/ovs-ctl start`

## 组播相关

启用组播嗅探：
```shell
ovs-vsctl set Bridge br-int mcast_snooping_enable=true
ovs-vsctl set Bridge br-int other_config:mcast-snooping-disable-flood-unregistered=true
```

禁用组播嗅探：
```shell
ovs-vsctl set Bridge br-int mcast_snooping_enable=false
ovs-vsctl remove Bridge br-int other_config mcast-snooping-disable-flood-unregistered=true
```

验证组播通信：
```shell
#启动客户端
iperf -c 226.94.1.1 -u -t 3600
#启动服务端
iperf -u -s -B 226.94.1.1
```

## 参考

* [ovs概述](http://www.nfvschool.cn/?p=561)
* [ovs-fields](http://www.openvswitch.org/support/dist-docs/ovs-fields.7.txt)
* [ovs advance](https://docs.openvswitch.org/en/latest/tutorials/ovs-advanced/)
* [ovs 流表规则](https://segmentfault.com/a/1190000038767587)
* [ovs cheatsheet](https://gist.github.com/djoreilly/c5ea44663c133b246dd9d42b921f7646)
* [从源码安装Open vSwitch](https://github.com/ebiken/doc-network/wiki/How-To:-Install-OVS-(Kernel-Module)-from-Source-Code)
* [ovs-ofctl ubuntu man page](http://manpages.ubuntu.com/manpages/trusty/man8/ovs-ofctl.8.html)
* [OpenvSwitch/OpenFlow 架构解析与实践案例](https://www.cnblogs.com/jmilkfan-fanguiju/p/10589725.html#_0)
* [Introduction to Open vSwitch](https://www.youtube.com/watch?v=rYW7kQRyUvA)