# 更新日志

## v3.0.7

1. 聊天主界面：新增聊天引导页面，介绍产品功能
2. 功能重构：拆分项目，将函数插件以及微信机器人，MidJourney 机器人等功能拆分新项目独立部署。
3. 功能新增：新增 MidJourney AI 绘画支持，当识别到用户的绘画需求时，自动调用 MidJourney 绘画函数进行绘画。
4. 功能新增：支持导出聊天记录为 PDF 文件。
5. 功能优化：在后台 dashboard 页面新增统计今日众筹收入。
6. 功能优化：支持用户设置默认的 GPT 模型
7. Bug修复：修复若干已知的的 Bug

## v3.0.6

1. 管理后台：新增用户名和手机号码搜索功能
2. 管理后台：新增重置用户密码功能
3. 管理后台：支持关闭注册功能，新增添加用户功能，适用于内部使用场景
4. 管理后台：新增仪表盘页面，统计当天的新增用户，新增会话数据，以及 Token 消耗
5. Bug修复：修复注册页面验证码不显示 Bug
6. Bug修复：优化上下文 Token 计算算法，修复聊天上下文超出限制时循环发送消息的 Bug
7. 功能修正：允许用户使用手机号码登录
8. 功能优化：更新系统配置后同步更新服务端内存变量数据
9. 功能优化：优化打包脚本，减少容器镜像大小

## v3.0.5

重磅功能更新！！！ 新增函数插件支持，可以轻松地接入你的第三方插件服务，ChatGPT 自动帮您调用对应的函数完成任务。

1. 新增函数功能支持，全球早报，今日头条和微博热搜等插件服务，您也可以接入自己的第三方服务。
2. 集成微信机器人模块，可以通过微信个人收款码来完成充值，无需接入微信支付功能也可以完成收款功能。
3. 用户注册添加短信验证码功能，引入交互安全认证服务，有效防刷短信。
4. 支持配置聊天上下文深度，精确统计每轮对话所消耗的总 TOKEN 数量。
5. 修复已知的 Bug。

## v3.0.4

1. 调整项目目录结构，移除其他语言 API 目录
2. 修复 nodejs apple M1 跨平台打包，运行报错 exec format error
3. 增加用户 token 消耗统计功能

## v3.0.3

1. 优化启动参数接收处理，支持环境变量传参
2. 修复 PC 端聊天界面出现滚动条的 Bug
3. 修正前端 user_init_call 字段错误和用户注册初始化头像路径问题
4. 更改 docker 构建镜像的基础镜像，改用作者的阿里云镜像，这样打包更快一些。

## v3.0.2

1. Feat：新增移动端的聊天和用户设置功能
2. Fix: 修复 markdown 换行符解析的 Bug
3. Feat: 新增头像上传功能
4. Docs: 增加容器部署支持，支持 docker-compose 一键部署
5. Fix: 增加全局错误处理 handler，修复业务处理异常导致服务退出的 Bug

## v3.0.1

1. 紧急修复前端 Home 组件路由被后台管理 Home 组件路由覆盖的 Bug。
2. 增加 docker-compose 部署脚本

## v3.0.0

全新的重构版本！！！
新版的系统前后端都进行大改动的重构，后端还是用的 Gin Web 框架，但是作者整合了 fx 自动注入框架，整个后端应用结构非常简洁，特别适合二次开发。
另外，数据存储用 MySQL 替换了 leveldb, 因为要对 C 端，后期会涉及到很多业务数据查询统计，leveldb 已经完全不够用了。
前后台技术架构还是基于 `Vue3 + Element-Plus` ，但是页面风格已经全部变了，几乎所有页面样式代码都重写了，希望会你是希望的风格！

此次重构改版主要是为了后面功能的扩展准备了。

新版本已经实现的功能如下：

1. 引入用户体系，新增用户注册和登录功能。
2. 聊天页面改版，实现了跟 ChatGPT 官方版本一致的聊天体验。
3. 创建会话的时候可以选择聊天角色和模型。
4. 新增聊天设置功能，用户可以导入自己的 API KEY
5. 保存聊天记录，支持聊天上下文。
6. 重构后台管理模块，更友好，扩展性更好的后台管理系统。
7. 引入 ip2region 组件，记录用户的登录IP和地址。
8. 支持会话搜索过滤。