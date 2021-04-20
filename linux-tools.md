# Linux Tools

## rsync

```bash
rsync -avh --delete --exclude="bin/*" --exclude=".git" /Users/workspaces/codes/src/github.com/sample/ 192.168.1.10:/home/ubuntu/sample/
# -a 是一个参数集合，等于-rlptgoD
# -r 是递归同步文件夹
# -l --links 拷贝软链接
# -p --perms 保留权限信息
# -t --times 保留时间信息
# -o --owner 保留用户信息
# -g --group 保留组信息
# -D 等于--devices --specials 保留设备文件和特殊文件
# -v 打印详细信息
# --delete 删除target中多余的文件
# --exclude 不需要同步的文件
```

注意： 当source路径中包含*通配符时，--delete参数不生效，因为\*会被shell解释扩展，实际上rsync是在针对每个文件进行同步，而不是针对文件夹。
稍详尽的介绍可以参考ruanyifeng的[博客](https://www.ruanyifeng.com/blog/2020/08/rsync.html)