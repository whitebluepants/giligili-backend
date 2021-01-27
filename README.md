前端项目: https://github.com/whitebluepants/giligili-vue



**Tips**

1. 代码中使用到的数据库等各种配置, 通过.env环境变量文件获取. 避免代码泄露的时候同时泄露数据库配置等重要信息. .env文件一般不上传
2. 前端访问不了后端api可能会是**跨域问题**, 需要在middleware/cors中添加允许请求的域名.
3. 使用docker生成镜像部署到服务器上
   1. 登录docker, 可以使用阿里云的容器镜像服务, 把镜像推送到阿里云的Docker Registry 例子: docker login --username=xxx registry.cn-hongkong.aliyuncs.com
   2. docker build -t "镜像名:版本号" "docker配置文件的目录" (不带双引号)
   3. docker push "镜像名:版本号" (同上)
4. docker 使用portainer图形化管理容器.
   1. 绑定阿里云的registry
   2. 在Stack界面创建前端、后端、mysql、redis容器.

