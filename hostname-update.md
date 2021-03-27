# 修改主机名

机器版本是Ubuntu20.04,主要参考[这里](https://www.cyberciti.biz/faq/ubuntu-18-04-lts-change-hostname-permanently/)。

## 重启方式修改主机名

- 输入`hostnamectl`能获得当前主机名及一些基本信息
- 修改主机名： `sudo hostnamectl set-hostname mynewhostname`
- 修改`/etc/hosts`文件，将任何含有旧主机名的地方修改为新主机名
- 重启机器：`reboot` 或者`shutdown -r now`

## 不重启机器修改主机名

参考的文章里提到不用重启机器来永久修改主机名的方式，没有实验，先记录在这里：

- `sudo hostnamectl set-hostname mynewname`
- 通过`ip`命令找到公网或私有ip：`ip a`,`ip a s ens33`，这里假设为`192.168.2.24`
- 修改`/etc/hosts`文件，添加或更新为`192.168.2.24 mynewhostname`,另外同时确保含有旧主机名的地方修改为新主机名。