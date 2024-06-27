<h1 align="center">Go Demo X</h1>
<p align="center">
<img alt="Static Badge" src="https://img.shields.io/badge/Go- 1.9-blue">
<img alt="Static Badge" src="https://img.shields.io/badge/Gin- 1.8-blue">

<img alt="Static Badge" src="https://img.shields.io/badge/Vue- 3.x -gren">
<img alt="Static Badge" src="https://img.shields.io/badge/Arco- 2.34.0 -gren">
<img alt="Static Badge" src="https://img.shields.io/badge/license- MIT-blue">

</p>

_面向开发者的API快速开发框架_
_更像一个DEMO 易于修改定制_

---
## 特点：
- 整个框架没有复杂的封装与接口设计，易于理解及定制。
- 模板项目包括基础api框架，admin框架(前后端)，带权限的api框架(前后端)。
- 自带可观测性解决方案，日志、追踪、指标。
- 自带文档及测试解决方案，方便的文档交流及测试。
- task自动化管理测试，文档生成，开发及部署。
- 完善的测试代码。
- 经过时间与项目的洗礼。
---
### 使用说明:
```
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/dangweiwu/goDemoX

```

---
### 其他
1. 本项目支持完整的可观测性，包括日志(log)，追踪(trace)，指标(metric)。方案基于[openobserve](https://openobserve.ai/),其中日志采集及指标采集都是基于go编写的采集工具，具有使用简单、可靠，与微服务的ELK相比，该方案使用资源极少。
2. 本项目快速开发基于代码生成理念，生成的代码不会是最终代码，但一定是最常用的代码部分。而最常用的代码并没有封装成组件提供使用，而是直接生成在项目里，这是为了修改方便，同时使整个项目逻辑简单，一看就懂。

---
### License
© Dangweiwu, 2023~time.Now

Released under the MIT License