
## Golang之家

> 技术栈：dotweb + mgo + vue

> 预览地址： [http://59.110.212.239/gohome](http://59.110.212.239/gohome)

### Dev环境, 项目运行

```
1.环境安装
  配置好GOROOT和GOPATH
  cd src
  mkdir work-codes
  git clone https://code.aliyun.com/wukc/gohome.git
  cd wuyue/app
  glide install

2. 运行项目
  启动mongodb:  mongod
  启动项目： bee run  或者 go build 

  运行测试： go test -v

3. 访问项目：
  浏览器输入： localhost:8080

4. 前端编译：
  cd view
  cnpm i
  npm run build 后生成dist文件夹

5. 前后端结合： 
  将编译后的dist文件， 复制到app下， 并将名称改为public
  cd view 
  cp -R dist ../app && cd ../app && mv dist public

```

### Prod环境, 项目部署

```
docker + caddy

```