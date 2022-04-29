# 学单词后端服务
主要借助了heroku的自定义镜像的能力，将`node`的后端代码与`git`指令相结合。

# 获取修改words-db仓库的权限
使用github的token功能，在个人设置里找到token，创建个能操作所有repo的token，不要告诉别人，等会给他上传到heroku的容器仓库，heroku容器是私有的，可以把token设置到容器里，而不要直接出现在代码中。
![image](https://i.imgur.com/MJOAEF5.png)