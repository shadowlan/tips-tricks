# Makefile

要全面了解Makefile的使用，可以移步参考里的链接。本文仅记录平时常用的一些功能或者曾经遇到的坑。

## 基础知识

* Makefile 文件由一系列的 `规则 (rules)` 组成，一个规则类似下面的结构：
```makefile
targets: prerequisites
    command
    command
    command
```

targets 指的是文件名称，多个以空格分隔。通常，一个规则只对应一个文件。
commands 通常是一系列用于制作（make）一个或多个目标的步骤。它们 需要以一个制表符开头，而不是空格。
prerequisites 也是文件名称，多个以空格分隔。在运行目标的 commands 之前，要确保这些文件是存在的。它们也被称为 依赖。

* makefile默认的target是文件中第一个target。
* .PHONY: 向一个目标中添加 .PHONY 会避免把目标识别为一个文件名。在下面这个例子中，即便文件 clean 被创建了，make clean 仍会运行。
```makefile
.PHONY: clean
clean:
    rm -f some_file
    rm -f clean
```


## 多平台构建应用

这里以Golang项目为例：

```makefile
SHELL              := /bin/bash
# go options
GO                 ?= go
LDFLAGS            :=
GOFLAGS            :=
BINDIR             ?= $(CURDIR)/bin

.PHONY: build
build: demo

DEMO_BINARIES := demo-darwin demo-linux demo-windows
$(DEMO_BINARIES): demo-%:
	@GOOS=$* $(GO) build -o $(BINDIR)/$@ $(GOFLAGS) -ldflags '$(LDFLAGS)' $(CURDIR)/main.go
	@if [[ $@ != *windows ]]; then \
	  chmod 0755 $(BINDIR)/$@; \
	else \
	  mv $(BINDIR)/$@ $(BINDIR)/$@.exe; \
	fi

.PHONY: demo
demo: $(DEMO_BINARIES)

.PHONY: clean
clean:
	@rm -rf $(BINDIR)
```

DEMO_BINARIES定义了一个变量  
`%`的意思是匹配零或若干字符  
`$*`指代匹配符 % 匹配的部分， 比如%匹配demo-darwin中的darwin ，`$*` 就表示 darwin。  
`$@`指代当前目标，就是Make命令当前构建的那个目标。比如，make build的 `$@` 就指代build.  
`@` 正常情况下，make会打印每条命令，然后再执行，这就叫做回声（echoing），在命令前加`@`会关闭回声。

## 常见错误

* `Makefile:25: *** missing separator.  Stop.`，这个提示很令人迷惑，但通常是因为makefile里用了空格而不是tab，需要将空格替换为tab制表符。
  
## 参考

* [GNU make](https://www.gnu.org/software/make/manual/make.html#toc-Overview-of-make)
* [GNU make中文手册](http://hacker-yhj.github.io/resources/gun_make.pdf)
* [Learn Makefiles](https://makefiletutorial.com/)
* [Learn Makefiles中文版](https://makefiletutorial.vercel.app)
* [跟我一起写Makefile](https://seisman.github.io/how-to-write-makefile/index.html)
* [Make命令教程](https://www.ruanyifeng.com/blog/2015/02/make.html)