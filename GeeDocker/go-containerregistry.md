containerregistry

## crane

Crane 是一个管理容器镜像的工具

- [crane append](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_append.md) - Append contents of a tarball to a remote image
- [crane auth](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_auth.md) - Log in or access credentials
- [crane blob](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_blob.md) - Read a blob from the registry
- [crane catalog](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_catalog.md) - List the repos in a registry
- [crane config](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_config.md) - 获取镜像的配置
- [crane copy](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_copy.md) - Efficiently copy a remote image from src to dst while retaining the digest value
- [crane delete](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_delete.md) - 从registry仓库中删除镜像的引用
- [crane digest](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_digest.md) - 获取图像的摘要
- [crane export](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_export.md) - 将远程镜像的内容导出为 tarball
- [crane flatten](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_flatten.md) - Flatten an image's layers into a single layer
- [crane ls](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_ls.md) - 列出 repo 中的标签
- [crane manifest](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_manifest.md) - Get the manifest of an image
- [crane mutate](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_mutate.md) - Modify image labels and annotations. The container must be pushed to a registry, and the manifest is updated there.
- [crane pull](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_pull.md) - Pull remote images by reference and store their contents locally
- [crane push](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_push.md) - Push local image contents to a remote registry
- [crane rebase](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_rebase.md) - Rebase an image onto a new base image
- [crane tag](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_tag.md) - 有效地标记远程镜像
- [crane validate](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_validate.md) - 验证image镜像格式是否正确
- [crane version](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_version.md) - 打印版本

这里我最关心的是下载镜像，也就是crane pull命令