# 一、Overlayfs介绍

## 二、Overlayfs实战

查看内核是否加载了overlay

lsmod | grep overlay



如果没加载，则采用如下命令加载：

modprobe overlay



可以看到merged文件夹中展示的文件是low和upper里面的文件集合。



同名覆盖测试

root@qa-control-pub-ci-build1:~/cxqdir/merged# vim 11.txt  

root@qa-control-pub-ci-build1:~/cxqdir/merged# cat 11.txt  333333333

分别查看low和upper目录下文件，发现low下面的同名文件内的内容没有变，而upper里面多了一个同名文件。 

root@qa-control-pub-ci-build1:~/cxqdir/merged# cat ../low/11.txt1111111111111111111111

root@qa-control-pub-ci-build1:~/cxqdir/merged# cat ../upper/11.txt333333333

值得注意的是upper中新增的11.txt文件和merge中的文件node是一样的。

 root@qa-control-pub-ci-build1:~/cxqdir/merged# ls -i ../upper/11.txt350031 ../upper/11.txt 

root@qa-control-pub-ci-build1:~/cxqdir/merged# ls -i 11.txt350031 11.txt 

但是跟底层low内的同名文件node不一样 

root@qa-control-pub-ci-build1:~/cxqdir/merged# ls -i ../low/11.txt  350033 ../low/11.txt



# 三、使用overlayfs搭建个镜像

How to Mount and Unmount File Systems in Linux

https://linuxize.com/post/how-to-mount-and-unmount-file-systems-in-linux/



参考链接：

公众号《Linux内核那些事》 容器三把斧之 | OverlayFS原理与实现

公众号《Linux内核那些事》Docker原理之 - OverlayFS设计与实现

容器如何工作:OverlayFS | Linux 中国