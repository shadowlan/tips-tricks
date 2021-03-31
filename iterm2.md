# iterm2 tips
工欲善其事，必先利其器， 把iterm2快捷键利用起来，利用分屏避免在各个窗口间切换来去，另外学习必要的提高效率的快捷键。
- Command + D 垂直分屏（使用默认profile）
- Command + Shift + D 垂直分屏（使用默认profile）
- Command + Option + Shift + H 水平分屏（弹窗让用户选择profile）
- Command + Option + Shift + V 垂直分屏（弹窗让用户选择profile）
- Control + Command + Up/Down/Left/Right 调整选中的分屏窗口的大小
- Option + Command + Left/Right 切换分屏窗口
- Command + Shift + Enter 将当前选中的窗口最大化
- Command + [ / ] 选择前一个或下一个窗口
- Commond + T 新建tab窗口
- Command + Shift + H 查看命令行历史
- Command + Shift + E 在最右侧显示时间戳
- Command + Shift + O 快速选择并打开一个窗口
- Commond + / 高亮当前鼠标
- Command + Shift + I 将命令广播到所有tab
- Command + Option + I 将命令广播到当前tab的所有面板
- Command + Option + F 打开密码管理器
- Command + Shift + S 保存当前面板布局
- Command + Shift + W 关闭当前终端
- Command + Shift + Space 打开表情输入
- Command + Option + B 显示快照

* **将分屏合并为tab**  
  按下Command + Option + Shift，鼠标会变成小手，此时拉拽该窗口到标签页即可。
* **预定义profile**  
  在做水平或者垂直分屏的时候，会弹出窗口提示选择profile，可以预先建立一些profile来自动执行常见的命令，比如登陆机器。通过Command + O 打开Profiles，选择"Edit Profiles",然后点击加号按钮，在General面板的Command下拉列表里选择Command，然后输入要跑的命令比如"/bin/bash"，在Send Text at Start框输入'ssh user@192.168.1.2'。
* **无鼠标选择文本**  
  通过Command + F进入搜索状态，输入搜索关键字，定位到要搜索的行之后按tab键会继续选中，此时粘贴面板就是当前选中的内容，到其他窗口或者激活当前命令行按下Control + V即可复制。
