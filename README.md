# 学单词后端服务


# render
Heroku从2022.11开始不再提供免费服务，能免费部署DIY的docker镜像的平台目前找到了[render](https://render.com/)。
![image](https://i.imgur.com/Qen3soq.png)

render使用方式非常简单，直接注册账号，右上角new，创建web service，绑定当前的git repo，自动即可识别为docker类型的应用，根据Dockerfile开始构建。

![image](https://i.imgur.com/4m9k5oG.png)

github token是这里生成的

![image](https://i.imgur.com/MJOAEF5.png)

render的免费版本有这些致命缺点，尤其是30s的启动时间，使得不得不通过[cron-job](https://console.cron-job.org/jobs)平台来定时调用服务，保持活跃。

![image](https://i.imgur.com/S83ZVc7.png)
# railway
[railway](https://railway.app)和render类似，计费方式稍有不同，railway的免费版本是给了5$/mon的使用额度，主要是内存和cpu的计费，但是一般小应用花不了这么多钱。

但时间限制是500h，500h就不到1月了，这有点可惜，好在可以升级为付费用户就没有时间限制了，并且付费用户是按需收费的，即并不需要每个月付固定的钱数，而是用了多少cpu内存资源才会进行付费，这样就对小项目比较合适。付费用户前5$同样是免费，如果没有超过5$的话，其实不需要花任何钱，但是获得了无限的时间的使用。

指的注意的是目前使用nodejs写的代码，虽然只是简单的express web应用，但是有150+M的常驻内存占用，是比较高的。使用rust写的另一个项目，同样是web应用，常驻内存就只有10M。因而考虑可以将当前项目用rust或者golang重构，golang应该内存占用也很小，而且相比rust更简单。

![image](https://user-images.githubusercontent.com/15844103/207232865-b085180b-dc00-40a6-9a6c-39b42517d28c.png)

![image](https://user-images.githubusercontent.com/15844103/207233078-b9b9834f-2687-4237-823a-3f5af90dc26e.png)




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
