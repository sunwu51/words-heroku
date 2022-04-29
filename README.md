# 学单词后端服务
主要借助了heroku的自定义镜像的能力，将`node`的后端代码与`git`指令相结合。详细见Dockerfile。

# heroku应用
使用heroku容器功能，保证本目录有Dockerfile之后
```
heroku create
heroku login -i
heroku container:login
heroku container:push web
heroku container:release web
```

# 获取修改words-db仓库的权限
使用github的token功能，在个人设置里找到token，创建个能操作所有repo的token，复制下来。

![image](https://i.imgur.com/MJOAEF5.png)

将token设置到heroku应用的环境变量中，在[代码](https://github.com/sunwu51/words-heroku/blob/master/index.js#L17)中通过环境变量获取该值。

![image](https://i.imgur.com/Erssrod.png)