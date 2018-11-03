### 安装
```bash
# 安装
go get github.com/Masterminds/glide

# 确认安装
glide -v

```
### 介绍
```bash
$ glide create                            # 初始化一个新项目，并创建一个glide.yaml文件
$ open glide.yaml                         # 编辑glide.yaml文件
$ glide get github.com/Masterminds/cookoo # 获取包，并将其添加至glide.yaml
$ glide install                           # 安装包和依赖，从glide.lock文件安装特定的版本
$ go build                                # Go tools work normally
$ glide up                                # 下载或更新glide.yaml文件中列出的所有库，并将它们放在vendor目录中。
```