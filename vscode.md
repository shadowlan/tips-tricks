# vscode tips

## title显示全路径

有时候用vscode打开两个不同文件夹下的同名文件时，容易混淆和误操作，这时可以配置vscode在标题栏显示全路径而不是仅文件名。
1. ctl+,（需要在英文输入状态）打开用户设置
2. 搜索title
3. 选择Window
4. 将${activeEditorShort}${separator}${rootName}替换为${activeEditorLong}${separator}${rootName}

## 命令行开启vscode

打开vscode，然后按Command + Shift + P， 然后在命令行面板输入Shell，会看到一个选项:`Shell Command : Install code in PATH`，选择即可。打开terminal，输入`code .`就会打开一个新的vscode界面。  
有时候会发现重启机器后code命令又丢了，大概率是因为MacOS对vscode添加了隔离属性。可以通过`xattr "/Applications/Visual Studio Code.app"`查看结果中是否有属性`com.apple.quarantine`，如果有则输入`sudo xattr -r -d com.apple.quarantine "/Applications/Visual Studio Code.app"`删除该属性。然后再次按照前面的步骤选择`Shell Command : Install code in PATH`即可。