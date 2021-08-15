# Ubuntu 内核版本问题追踪

这两天在调试一个内核版本的问题，准确来说是为了解决给kubernetes的CNI启用IPSec+Geneve模式遇到的一个坑。
一开始我在本地实验IPSec+Geneve模式时很顺利，但是放到e2e测试环境时，基本的ping连接测试也会失败，因为CNI本身只用到了内核的OVS,
于是怀疑是内核的版本不一致导致，通过`uname -a`查看了两个环境后发现的确有小版本的差异，可工作的版本是`Linux 4.15.0-143-generic`,
不工作的版本是`Linux 4.15.0-66-generic`,然后又比较了下内核自带的OVS内核模块文件夹，发现路径`ls -alt /lib/modules/$version/kernel/net/openvswitch/`内的文件大小都有差异，这就更让我肯定了是内核版本的问题。
但是小版本66和143之间还是存在很多个中间版本的，因为需要给用户准确的文档说明，这中间到底是什么时候引入的fix是需要解决的问题。
这里记录下为了定位到具体版本经历的过程，给以后一些参考。

## 灵活选择内核版本

为了测试版本差异和效果，需要安装不同的内核版本，想用旧的内核版本，首先需要安装上旧版本，但如果不做特别配置的话，系统默认都是用当前最新的内核版本，并不会选择老版本来启动机器。参考了[这篇文章](https://docs.digitalocean.com/products/droplets/how-to/kernel/use-non-default/)后，我在本地配置了允许用户灵活选择内核版本。测试的系统版本是Ubuntu 18.04，具体步骤如下：

编辑文件`/etc/default/grub`,注释掉`GRUB_DEFAULT=0`这一行，然后加上如下三行：
```conf
GRUB_DEFAULT=saved
GRUB_SAVEDEFAULT=true
GRUB_DISABLE_SUBMENU=y
```


GRUB_DEFAULT 设置允许 Grub 使用我们保存的任何值为默认内核，而不是硬编码默认值。 GRUB_SAVEDEFAULT 告诉 Grub 将默认内核设置为我们在菜单中明确选择的条目。 GRUB_DISABLE_SUBMENU 选项使菜单结构扁平化，以便我们可以更轻松地解析它。 

另外需要检查文件夹`/etc/default/grub.d` 下名为`50-cloudimg-settings.cfg`的文件. 它可能包含覆盖`/etc/default/grub`文件内配置的默认选择。如果有也在`GRUB_DEFAULT=0`所在行添加注释，改为`#GRUB_DEFAULT=0`。

然后配置变量GRUB_CONFIG：`export GRUB_CONFIG='/boot/grub/grub.cfg'`,并运行`update-grub`来重新构建grub.cfg文件以使前面的更改生效。现在配置已重建且菜单已展平，我们可以解析文件中的可用条目。以下命令显示条目索引号和标题。我们可以使用其中任何一个来引用特定条目。
`grep 'menuentry ' $GRUB_CONFIG | cut -f 2 -d "'" | nl -v 0`
这将返回所有可用的引导选项。记下要引导的条目的索引号或标题。
```
0	Ubuntu, with Linux 4.15.0-151-generic
1	Ubuntu, with Linux 4.15.0-151-generic (recovery mode)
2	Ubuntu, with Linux 4.15.0-141-generic
3	Ubuntu, with Linux 4.15.0-141-generic (recovery mode)
4	Ubuntu, with Linux 4.15.0-140-generic
...
```
Grub 包括从命令行设置新的默认内核的命令。我们可以使用索引号或条目标题来指定引导选项。例如使用索引号：`grub-set-default 2`
还可以选择仅适用于下次启动的临时启动选项:`grub-reboot 2`,配置结束后就可以重启以便于使用配之后的内核版本。


## 查找Ubuntu版本变更

既然是Ubuntu内核版本差异导致的，那就要查找内核变更，先搜到了[Ubuntu的Git源码位置](https://kernel.ubuntu.com/git/), 而我本地用的是18.04，对应的是bionic，所以Git路径是[ubuntu-bionic](https://kernel.ubuntu.com/git/ubuntu/ubuntu-bionic.git/)所在地址。为了方便定位具体变更，我把整个源码克隆到了本地: `git clone git://kernel.ubuntu.com/ubuntu/ubuntu-bionic.git`.

我先在[log页面](https://kernel.ubuntu.com/git/ubuntu/ubuntu-bionic.git/log/)通过`log msg`方式查找了geneve关键字，发现了不少相关的commit，另外在[tag页面](https://kernel.ubuntu.com/git/ubuntu/ubuntu-bionic.git/refs/tags)能够看到具体的tag以及他们创建的日期时间。这中间定位到正确的fix commit饶了点弯路，切换不同的内核做了一些不必要的测试，最后才定位到是在[这个commit](https://kernel.ubuntu.com/git/ubuntu/ubuntu-bionic.git/commit/?id=be556894e8eaefd5d21690d56614d76e45786ecb)里提交的Fix。然后在本地通过`git tag --contains $commit`找到了包含这个commit的所有tag。如果没有克隆源码，也可以在页面上比较commit提交时间和tag的创建时间，在commit提交时间之后创建的tag肯定都会包含相关的代码。

其实看tag的历史可以发现`Linux 4.15.0-66-generic`和`Linux 4.15.0-143-generic`之间的时间跨度还是非常大的，如果不是Fix commit在搜索出来的commit列表里的头几个，估计我也没有耐心或者找到更好的方式来定位问题，希望下次能想到更快更好的方式定位问题。

## Ubuntu版本号

`4.15.0-66`这个版本号里每个字段包含不同含义，更多内容可以参考[这篇博客](https://jasonhzy.github.io/2019/02/05/linux-kernel-version/)，介绍的很详细。我把其中介绍Ubuntu版本号的部分拷贝到了这里：
```
Linux localhost 3.2.0-67-generic #101-Ubuntu SMP Tue Jul 15 17:46:11 UTC 2014 x86_64 x86_64 x86_64 GNU/Linux
第一个组数字：3, 主版本号
第二个组数字：2, 次版本号，当前为稳定版本
第三个组数字：0, 修订版本号
第四个组数字：67，当前内核版本（3.2.0）的第67次微调patch
generic：当前内核版本为通用版本，另有表示不同含义的server（针对服务器）、i386（针对老式英特尔处理器）
pae（Physical Address Extension）：物理地址扩展，为了弥补32位地址在PC服务器应用上的不足而推出，表示此32位系统可以支持超过4G的内存
x86_64：采用的是64位的CPU
SMP：对称多处理机，表示内核支持多核、多处理器
Tue Jul 15 17:46:11 UTC 2014：内核的编译时间（build date）为 2014/07/15 17:46:11  
```

## 其他

在调试中间用到一些命令，记录在这里：
```bash
# 用日期查找git log
git log --date=iso --pretty=format:'%ad%x08%aN %s' | awk '$0 >= "2020-02-28" && $0 <= "2020-04-01"'
# 用日期和作者查找git log
git log --pretty=format:"%ad - %an: %s" --after="2020-02-28" --until="2020-04-01" --author="John Doe"
# 查看tag间差异
git diff Ubuntu-4.15.0-91.92 Ubuntu-4.15.0-96.97 --stat
# 安装内核版本
apt install linux-image-4.15.0-91-generic linux-headers-4.15.0-91-generic
# 删除内核版本
apt remove linux-image-4.15.0-91-generic linux-headers-4.15.0-91-generic
# 搜索内核版本
apt list | grep linux-image-4.15.0-
# 标记软件包状态，用来禁止更新等
apt-mark hold linux-image-generic linux-headers-generic
```
