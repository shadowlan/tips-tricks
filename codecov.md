# Codecov Lab

这里面主要放和Codecov相关的测试代码，用来验证与Codecov相关的一些功能。

## [Codecov uploader](https://docs.codecov.com/docs/codecov-uploader)

Codecov uploader是Codecov提供的独立工具，能够用来上传相关report到Codecov。基本命令如下：
```bash
# -c[--clean]  删除寻找到的coverage reports
# -t[--token]  Codecov上传的token串
# -F[--flags]  coverage report的分类标签名，更多flag的使用介绍可查看https://docs.codecov.com/docs/flags
# -s[--dir] 用来搜索coverage reports的文件夹路径
# -C[--sha] 指定的commit SHA
# -r[--slug] owner/repo信息
# sample: ./codecov -c -t ***** -F integration -f .coverage/coverage-integration.txt -s . -C abcea2fac89 -r antrea-io/antrea
# 更多参数可参考 👉 https://docs.codecov.com/docs/codecov-uploader#uploader-command-line-arguments
./codecov -c -t ${CODECOV_TOKEN} -F ${FLAG} -f ${COVERAGE_FILE} -s ${DIR} -C ${SHA} -r ${OWNER}/${REPO}
```
