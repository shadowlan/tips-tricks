# 本地开发环境搭建  
好久没从零开始搭建新环境了，今天花了大部分时间在准备新机器，顺手把相关的配置和步骤记录下来，便于以后参考。机器是MacOS,版本是catalina 10.15.7。

## 基本工具

常用的工具列在这里，基本上都是去官网下载加双击安装，就不赘述，列一下工具集。
1. iterm2
2. visual studio code
3. docker desktop
4. firefox
5. vmware fusion

另有从app store安装的工具：
1. spark
2. slack
3. trello


## vscode配置

* 从命令行打开code： 在vscode UI界面利用快捷键control+shift+p打开命令行面板，然后输入 shell ,在出现的下拉列表选择 "Shell Command: Install 'code' command in PATH command."，然后就可用在terminal直接输入`$code`来打开vscode应用。
* 插件列表
    1. Go
    2. Markdown All in One


## Git配置  
初次运行git clone: `$git clone git@github.com:kubernetes/kubernetes.git` 试图拷贝项目到本地时出现了两个错误，记录如下

1. Permanently added the RSA host key for IP address '13.229.188.59' to the list of known hosts.  
Fix: 添加`13.229.188.59 github.com`到文件/etc/hosts

2. git@github.com: Permission denied (publickey). fatal: Could not read from remote repository.  
Fix: 需要follow[官方文档](https://docs.github.com/cn/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)生成key pair，然后将公钥添加到github账号。


## on-my-zsh

主要参考了👉https://my.oschina.net/u/2266513/blog/3103451

1. 安装： `ssh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"`
2. 修改主题
   1. 安装字体 
   ```bash
   git clone https://github.com/powerline/fonts.git powerline-fonts.git
   bash powerline-fonts.git/install.sh
   ```
   2. 修改~/.zshrc文件中ZSH_THEME配置为"agnoster"
   ```bash
   ZSH_THEME  = "agnoster"
   ```
重启iterm2后如果主题仍有乱码，可以参考[这里](https://segmentfault.com/a/1190000015962180#34-%E5%AE%89%E8%A3%85-oh-my-zsh)图示修改iterm2的profile。

默认agnoster主题的终端提示符是带主机名的，有点多余，可以通过修改theme文件来删除：
编辑 vim ~/.oh-my-zsh/themes/agnoster.zsh-theme
把prompt_context() 方法里prompt_segment最后的"@%m"删除掉,删掉后内容如下：
```
prompt_context() {
  if [[ "$USER" != "$DEFAULT_USER" || -n "$SSH_CLIENT" ]]; then
    prompt_segment black default "%(!.%{%F{yellow}%}.)%n"
  fi
}
```
3. 安装三个常用插件
* 语法高亮插件 zsh-syntax-highlighting： `git clone https://github.com/zsh-users/zsh-syntax-highlighting.git $ZSH_CUSTOM/plugins/zsh-syntax-highlighting`
* 自动补全插件 zsh-autosuggestions: `git clone https://github.com/zsh-users/zsh-autosuggestions $ZSH_CUSTOM/plugins/zsh-autosuggestions`
* 自动跳转插件 autojump: 
```bash
# clone 到本地
git clone https://github.com/wting/autojump.git
# 进入clone目录，接着执行安装文件
cd autojump
./install.py
# 接着根据安装完成后的提示，在~/.bashrc最后添加下面语句：
vim ~/.zshrc    
[[ -s /home/xxxx/.autojump/etc/profile.d/autojump.sh ]] && source /home/xxxx/.autojump/etc/profile.d/autojump.sh
autoload -U compinit && compinit -u
```
4. 启用插件
```bash
# 编辑~/.zshrc   
vim ~/.zshrc    
# 在plugins后括号里添加安装的插件名字
plugins=(   git 
            autojump 
            zsh-autosuggestions 
            zsh-syntax-highlighting
            )
# 最后刷新
source ~/.zshrc
```


## 安装homebrew 
从[官方网页](https://brew.sh/)看安装非常简单，一行搞定`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`.
然而，但是。。。官方源太慢了。。。（也可能是”墙“太强）...所以只能用国内的镜像来加速。主要参考了[here](https://www.jianshu.com/p/fdf7e316f096) and [there](https://www.jianshu.com/p/ff2ad9599a06)

1. 将安装脚本重定向为本地文件： `$curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh > brew_install.sh`
2. HOMEBREW_BREW_DEFAULT_GIT_REMOTE="https://github.com/Homebrew/brew" 替换为 HOMEBREW_BREW_DEFAULT_GIT_REMOTE="https://mirrors.ustc.edu.cn/brew.git"
3. chmod +x brew_install.sh;./brew_install.sh
4. 上一步执行到`Tapping homebrew/core`时会发现还是很慢，可以到`cd /usr/local/Homebrew/Library/Taps/homebrew`(如果没有homebrew,新建homebrew文件夹)，然后执行`git clone https://mirrors.ustc.edu.cn/homebrew-core.git;git clone https://mirrors.ustc.edu.cn/homebrew-cask.git`
5. 通过export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles 可以修改brew安装软件时使用的镜像
6. 上一步也可永久替换：`echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles' >> ~/.bash_profile;source ~/.bash_profile`

## 常用命令行工具  
因为都是用brew直接安装，比较简单，仅记录命令和github入口
```bash
 brew install fzf #https://github.com/junegunn/fzf
 brew install bat #https://github.com/sharkdp/bat
 brew install kubectl #https://github.com/kubernetes/kubectl
 brew install kube-ps1 #https://github.com/jonmosco/kube-ps1 
 brew install kubectx  #https://github.com/ahmetb/kubectx
```

## TBD

目测可能环境还不全，等后续有添加或者发现好的工具再加进来，