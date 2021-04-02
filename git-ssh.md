# Github多账号登陆

因为想把公司和个人账号分开，在Github上建了两个账号，结果发现在一台机器上要同时使用两个账号还需要一些额外配置，这里记录下具体配置步骤。

首先两个账号的ssh key登陆要配置好，这里就不记录，具体可以参考[官方文档](https://docs.github.com/cn/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)或者这个issue文章里最前面的介绍[配置多个git账号](https://github.com/jawil/notes/issues/2)。

[配置多个git账号](https://github.com/jawil/notes/issues/2)这里面提供的多账号登陆的配置其实基本上全了，但我配完后遇到的一个问题是账号并没有根据指定域名来切换认证的key信息，导致账号权限错误。搜索一番后才发现少了“IdentitiesOnly yes”这个配置。最终使用的完整~/.ssh/config配置文件内容示例如下：
```
Host s.github.com
  hostname github.com
  User user1
  IdentityFile ~/.ssh/id_ed25519user1
  IdentitiesOnly yes
Host github.com
  hostname github.com
  User user2
  IdentityFile ~/.ssh/id_ed25519user2
  IdentitiesOnly yes
```

根据[StackOverflow上一个问题的解答](https://stackoverflow.com/questions/7927750/specify-an-ssh-key-for-git-push-for-a-given-domain)了解到IdentityFile实际上只是将用户指定的key文件添加到SSH agent里（默认的key比如版本1使用的~/.ssh/identity和版本2使用的~/.ssh/id_dsa, ~/.ssh/id_ecdsa, ~/.ssh/id_rsa）。而`IdentitiesOnly yes`阻止使用默认的id，否则如果有id文件匹配默认的名字，他们将会在ssh认证过程中先被尝试使用。

配置好后可以通过`ssh -T git@s.github.com`来验证输出的账户信息是否和User中指定的用户名一致。
另外由于user1使用的Host变成了`s.github.com`,需要修改该用户项目的remote配置:
```bash
$git remote rm origin
$git remote add origin git@s.github.com:user1/demo.git 
```