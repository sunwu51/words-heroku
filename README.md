# 学单词后端服务
Heroku从2022.11开始不再提供免费服务，能免费部署DIY的docker镜像的平台目前找到了[render](https://render.com/)。
![image](https://i.imgur.com/Qen3soq.png)

render使用方式非常简单，直接注册账号，右上角new，创建web service，绑定当前的git repo，自动即可识别为docker类型的应用，根据Dockerfile开始构建。

![image](https://i.imgur.com/4m9k5oG.png)

github token是这里生成的

![image](https://i.imgur.com/MJOAEF5.png)

~~主要借助了heroku的自定义镜像的能力，将`node`的后端代码与`git`指令相结合。详细见Dockerfile。~~

# ~~heroku应用~~
~~使用heroku容器功能，保证本目录有Dockerfile之后~~
```
heroku create
heroku login -i
heroku container:login
heroku container:push web
heroku container:release web
```
~~如果已经创建好了，则需要关联即可~~
```
heroku login -i
heroku container:login
git remote add heroku [你的heroku git地址]
heroku container:push web
heroku container:release web
```

# ~~获取修改words-db仓库的权限~~
~~使用github的token功能，在个人设置里找到token，创建个能操作所有repo的token，复制下来。~~


~~将token设置到heroku应用的环境变量中，在[代码](https://github.com/sunwu51/words-heroku/blob/master/index.js#L17)中通过环境变量获取该值。~~

![image](https://i.imgur.com/tlz5URQ.png)
