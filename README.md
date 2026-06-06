# 赛事赛程管理系统

## 项目简介

赛事赛程管理系统是一个基于 Go 语言开发的综合性体育赛事管理平台，提供运动员报名、分组编排、赛程排班、成绩录入、获奖档案和场地资源调度等完整功能。

## 技术栈

- **后端**: Go 1.22 + Gin Web Framework
- **数据库**: SQLite3
- **前端**: Go html/template 模板引擎 + CSS + JavaScript
- **认证**: Gin Sessions (Cookie Store) + bcrypt 密码加密

## 功能模块

### 1. 运动员报名
- 运动员信息录入（姓名、性别、年龄、代表队、参赛项目、联系方式、身份证号）
- 运动员列表查看、详情查看、编辑和删除
- 按项目筛选运动员

### 2. 分组编排
- 创建比赛分组（如预赛A组、决赛组等）
- 按项目将运动员编排到对应分组
- 分组状态管理（待编排、已编排、已完成）

### 3. 赛程排班
- 为分组分配比赛场地和时间
- 赛程状态管理（已排班、进行中、已完成、已取消）
- 赛程详情查看，包含成绩记录

### 4. 成绩录入
- 为指定赛程的运动员录入成绩
- 支持成绩、名次、备注等信息
- 成绩列表和详情查看

### 5. 获奖档案
- 记录运动员获奖信息（金牌、银牌、铜牌）
- 关联赛事名称
- 获奖记录查看和管理

### 6. 场地资源调度
- 场地信息管理（名称、位置、容量、状态）
- 场地状态管理（可用、维护中、占用中）
- 查看场地关联的赛程安排
- 场地资源调度总览

## 项目结构

```
repo/
├── src/
│   ├── config/          # 应用配置和数据库初始化
│   │   ├── app.go       # 应用常量配置
│   │   └── database.go  # 数据库连接和迁移
│   ├── controllers/     # HTTP 请求处理器
│   │   ├── auth.go      # 认证控制器
│   │   ├── athlete.go   # 运动员控制器
│   │   ├── group.go     # 分组控制器
│   │   ├── schedule.go  # 赛程控制器
│   │   ├── score.go     # 成绩控制器
│   │   ├── award.go     # 获奖控制器
│   │   ├── venue.go     # 场地控制器
│   │   └── dashboard.go # 仪表盘控制器
│   ├── middleware/       # 中间件
│   │   └── auth.go      # 认证中间件
│   ├── models/          # 数据模型定义
│   │   ├── user.go
│   │   ├── athlete.go
│   │   ├── group.go
│   │   ├── schedule.go
│   │   ├── score.go
│   │   ├── award.go
│   │   └── venue.go
│   ├── routes/          # 路由定义
│   │   └── routes.go
│   ├── services/        # 业务逻辑层
│   │   ├── user_service.go
│   │   ├── athlete_service.go
│   │   ├── group_service.go
│   │   ├── schedule_service.go
│   │   ├── score_service.go
│   │   ├── award_service.go
│   │   └── venue_service.go
│   ├── seed.go          # 种子数据
│   └── main.go          # 程序入口
├── templates/           # HTML 模板文件
│   ├── base.html
│   ├── index.html
│   ├── partials/        # 公共模板片段
│   ├── auth/            # 认证页面
│   ├── athletes/        # 运动员页面
│   ├── groups/          # 分组页面
│   ├── schedules/       # 赛程页面
│   ├── scores/          # 成绩页面
│   ├── awards/          # 获奖页面
│   ├── venues/          # 场地页面
│   └── dashboard/       # 仪表盘页面
├── static/              # 静态资源
│   ├── css/style.css
│   └── js/main.js
├── go.mod
└── go.sum
```

## 快速开始

### 本地运行

```bash
cd repo
go mod tidy
go run ./src/
```

访问 http://localhost:8080

### Docker 运行

```bash
cd event-racing
docker build -t event-racing .
docker run -p 8080:8080 event-racing
```

### 默认账号

| 用户名    | 密码         | 角色   |
|-----------|-------------|--------|
| admin     | admin123    | 管理员 |
| operator  | operator123 | 操作员 |

## 种子数据

系统首次启动时会自动插入种子数据，包括：

- 2 个用户（管理员和操作员）
- 15 名运动员（涵盖100米跑、200米跑、400米跑、800米跑、1500米跑、跳远、跳高、铅球、标枪、铁饼等项目）
- 6 个比赛分组
- 5 个比赛场地
- 5 条赛程安排
- 2 条成绩记录
- 5 条获奖记录

## 数据库

使用 SQLite3 数据库，数据文件为 `data.db`，首次运行自动创建。

### 数据表

- `users` - 用户表
- `athletes` - 运动员表
- `groups` - 分组表
- `group_athletes` - 分组-运动员关联表
- `venues` - 场地表
- `schedules` - 赛程表
- `scores` - 成绩表
- `awards` - 获奖表

## API 路由

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | / | 首页 |
| GET | /dashboard | 系统概览 |
| GET/POST | /auth/login | 登录 |
| GET/POST | /auth/register | 注册 |
| GET | /auth/logout | 退出 |
| GET | /athletes | 运动员列表 |
| GET/POST | /athletes/create | 运动员报名 |
| GET | /athletes/:id | 运动员详情 |
| GET | /athletes/:id/edit | 编辑运动员 |
| POST | /athletes/:id | 更新运动员 |
| POST | /athletes/:id/delete | 删除运动员 |
| GET | /groups | 分组列表 |
| POST | /groups | 创建分组 |
| GET | /groups/:id | 分组详情 |
| GET/POST | /groups/:id/arrange | 编排分组 |
| POST | /groups/:id/delete | 删除分组 |
| GET | /schedules | 赛程列表 |
| GET/POST | /schedules/create | 创建赛程 |
| GET | /schedules/:id | 赛程详情 |
| POST | /schedules/:id/status | 更新赛程状态 |
| POST | /schedules/:id/delete | 删除赛程 |
| GET | /scores | 成绩列表 |
| GET/POST | /scores/entry | 成绩录入 |
| GET | /scores/:id | 成绩详情 |
| POST | /scores/:id/delete | 删除成绩 |
| GET | /awards | 获奖列表 |
| POST | /awards | 新增获奖 |
| GET | /awards/:id | 获奖详情 |
| POST | /awards/:id/delete | 删除获奖 |
| GET | /venues | 场地列表 |
| POST | /venues | 新增场地 |
| GET | /venues/schedule | 场地调度 |
| GET | /venues/:id | 场地详情 |
| POST | /venues/:id | 更新场地 |
| POST | /venues/:id/delete | 删除场地 |

## 许可证

MIT License
