# Go 格式化

今天遇到一个 Go 文件格式的问题，错误信息是"File is not `goimports`-ed with -local github.com/vmware-tanzu/antrea" 看了下 code 实在没看出来什么问题，于是在本地安装 goimports 后执行了`goimports -w -local github.com/vmware-tanzu/antrea pkg/apiserver/handlers/featuregates/handler.go` 才发现 import 的顺序从下面的顺序

```go
import (
	"encoding/json"
	"net/http"

	"github.com/vmware-tanzu/antrea/pkg/antctl/transform/common"
	"github.com/vmware-tanzu/antrea/pkg/features"
    "k8s.io/klog"
)
```

变成了

```go
import (
	"encoding/json"
	"net/http"

	"k8s.io/klog"

	"github.com/vmware-tanzu/antrea/pkg/antctl/transform/common"
	"github.com/vmware-tanzu/antrea/pkg/features"
)
```

查看了下 goimport 的帮助文档，

```shell
goimports -h
usage: goimports [flags] [path ...]
  -cpuprofile string
    	CPU profile output
  -d	display diffs instead of rewriting files
  -e	report all errors (not just the first 10 on different lines)
  -format-only
    	if true, don't fix imports and only format. In this mode, goimports is effectively gofmt, with the addition that imports are grouped into sections.
  -l	list files whose formatting differs from goimport's
  -local string
    	put imports beginning with this string after 3rd-party packages; comma-separated list
  -memprofile string
    	memory profile output
  -memrate int
    	if > 0, sets runtime.MemProfileRate
  -srcdir dir
    	choose imports as if source code is from dir. When operating on a single file, dir may instead be the complete file name.
  -trace string
    	trace profile output
  -v	verbose logging
  -w	write result to (source) file instead of stdout
```

`goimports -w -local github.com/vmware-tanzu/antrea pkg/apiserver/handlers/featuregates/handler.go` 这句话的意思是把字符串`github.com/vmware-tanzu/antrea`放到第三方包的后面，而我原始的 code 文件里顺序错误，导致`make golangci`报错。

我的代码编辑器 vscode 默认配置了 gofmt 来做 format，我本来理解文件格式肯定没问题，现在看来还要检查不同项目 golang ci 的配置，以 antrea 项目为例，它的[`.golangci.yml`](https://github.com/vmware-tanzu/antrea/blob/main/.golangci.yml)文件里明确指出了这个要求，所以本地需要做相应的更新来匹配要求。

```yml
# golangci-lint configuration used for CI
run:
  tests: true
  timeout: 10m
  skip-files:
    - ".*\\.pb\\.go"
  skip-dirs-use-default: true

linters-settings:
  goimports:
    local-prefixes: github.com/vmware-tanzu/antrea

linters:
  disable-all: true
  enable:
    - misspell
    - gofmt
    - deadcode
    - staticcheck
    - gosec
    - goimports
    - vet
```

本以为 goimport 只是负责 import 部分，继续深挖后发现原来 format 工具不止 gofmt/goimports，甚至还有 goreturns，而 goimports 也不仅仅是负责 import 部分，实际上是 gofmt+fixing imports，[这篇文章](https://alenkacz.medium.com/gofmt-goimports-goreturns-why-do-we-need-three-formatters-a3518ee6cc90)介绍了三者的区别联系，简单而言就是下面的关系,作者建议在 IDE 里使用 goreturns。

```
gofmt = golang formatter
goimports = gofmt + fixing imports
goreturns = goimports + return statement syntactic sugar
```

等有需要再研究 goreturns，初步想法是希望 vscode 在保存的时候能够运行 goimports，这个本来也很简单，快捷键 Command+,在 setting 里搜索 format，然后选择 Extensions 里面的 Go，右侧出现的“Go: Format Tool”里选择 goimports，但是这种方式只会执行默认的命令，无法解决我遇到的特殊顺序要求，再看发现配置中还有"Go: Format Flags"配置，于是尝试添加参数" -w -local github.com/vmware-tanzu/antrea ."，但是不工作，网上尝试几种途径都未能解决，google 一圈后发现有一个插件是"Run On Save"似乎能满足要求,安装该插件后，在 vscode setting 里添加如下配置，让指定目录下的文件保存时运行命令"goimports -w -local github.com/vmware-tanzu/antrea ${file}"， ${file}是插件支持的占位符，执行时将被替换为当前保存的文件全路径。要看设置的效果可以在修改文件保存后，在 OUTPUT tab 页的右上角的下拉框选择"Run On Save"来看命令行的输出结果。

```json
"emeraldwalk.runonsave": {
    "commands": [
        {
            "match": "/Users/codes/antrea/.*",
            "isAsync": true,
            "cmd": "goimports -w -local github.com/vmware-tanzu/antrea ${file}"
        }
    ]
}
```
