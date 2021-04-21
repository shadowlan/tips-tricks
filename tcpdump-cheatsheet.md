# tcpdump cheatsheet

示例:

```
tcpdump -nnSX -i ens33 udp port 4500 dst host 172.20.1.1
获得http头信息： tcpdump -vvAls0 | grep 'User-Agent:'
根据网段查找： tcpdump net 1.2.3.0/24
显示尺寸小于32的包：tcpdump less 32
非ftp端口数据： tcpdump -vvAs0 port ! ftp
```

## 常用抓包参数

- -i any : 监听所有网络接口
- -i virbr0: 监听特定接口 virbr0
- -D: 显示所有可用网络接口
- -n: 不解析 IP 为主机名
- -nn: 不解析为主机或者端口名
- -c: 输出指定数量的包信息然后退出抓包
- -s: 定义包获取的字节大小.使用-s0 获取完整的包
- -S: 打印绝对顺序号
- -l: 基于行的输出，便于你保存查看，或者交给其它工具分析
- -E: 通过提供的解密密钥解密 IPSEC 数据流

## 常用打印参数

- -t: 不打印时间戳
- -tttt: 打印用户友好的时间戳
- -X: 以 HEX 和 ASCII 格式打印包内容
- -XX: 和-X 一样，但是打印包头
- -e: 获得包头信息
- -v, -vv, -vvv: 打印更多包信息
- -q: 安静输出，打印更少的协议信息
- -A: 以 ASCII 码打印包信息

## 表达式类型

| 表达式类型 | 可选项                  |
| ---------- | ----------------------- |
| 类型选项   | host, net, port         |
| 方向       | src, dst                |
| 协议       | udp, tcp, ipv6, icmp... |

## 逻辑操作符

| 操作符 | 符号    |
| ------ | ------- |
| 与     | and,&&  |
| 或     | or,\|\| |
| 或     | not,!   |
| 大于   | greater |
| 小于   | less    |

## 参考

- [tcpdump cheat sheet](https://www.comparitech.com/net-admin/tcpdump-cheat-sheet/)
- [tcpdump sample](https://gist.github.com/jforge/27962c52223ea9b8003b22b8189d93fb)
- [A tcpdump Tutorial with Examples](https://danielmiessler.com/study/tcpdump/)
- [A tcpdump Tutorial with Examples 中文版](https://colobu.com/2019/07/16/a-tcpdump-tutorial-with-examples/)
