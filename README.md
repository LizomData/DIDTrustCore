# DIDTrustCore

> 🧾 SBOM + 🛡️ 漏洞扫描 + 🪪 DID 上链 + 🔑 鉴权的一站式服务。默认端口 `:8000`，Swagger 在 `/swagger/index.html`。

## SBOM 背景（为什么重要）
- SBOM = Software Bill of Materials，像硬件 BOM 一样列出软件的组件、版本、来源。
- 作用：快速定位依赖漏洞（如 Log4j）、做许可证/合规审计、满足供应链安全要求，为可信溯源打基础。
- 常见格式：SPDX、CycloneDX。本项目用 Syft 生成 SPDX/CycloneDX JSON，后续漏洞扫描直接复用。

## 亮点速览
- ✅ 一条链路跑完：上传包 → 生成 SBOM → 漏洞扫描 → DID 上链。
- 🔒 可信标识：DID + Fabric，标识可追溯、不可抵赖。
- 🧭 快速接入：Swagger 交互文档即开即用。
- 🧩 易扩展：SBOM/报告都用标准格式，便于对接现有安全平台。

## 你能做什么
- 📦 上传软件包（zip/tar.gz）并拿到下载链接。
- 🧾 生成 SBOM（SPDX 或 CycloneDX），保存并提供下载。
- 🛡️ 基于 SBOM 跑 Grype 漏洞扫描，生成报告并提供下载。
- 🪪 基于 Fabric 管理 DID：生成、查询、更新、吊销 DID 文档。
- 🗂️ 维护个人漏洞记录，查询列表、统计严重度分布。
- 🔑 账号注册/登录，接口用 JWT 鉴权。

## 工作原理
- Gin 作为 Web 框架，`routers` 里注册所有路由，`util/Middlewares.go` 负责 CORS。
- 鉴权：`util/JwtUtil.go` 生成/校验 JWT；`AuthMiddleware` 保护接口，`AuthMiddlewareV2` 校验权限等级。
- 存储：MySQL + GORM（连接配置在 `util/dataBase/DataBase.go`），启动时自动建用户/软件包/SBOM/扫描报告/漏洞表。
- DID：`did_connnection` 用 Fabric Gateway（默认证书在 `./org1.example.com`，链码 `basic`，通道 `mychannel`），`didController` 生成 ECDSA 密钥、签名 DID 文档并上链。
- SBOM：`controller/sbomController` 先解压包，再用 `util/sbom`（封装 Syft）生成 SBOM，存 `./tmp/sbomStorage`，静态暴露 `/sbom_list/`。
- 漏洞扫描：`util/grype/grypeService` 封装 Grype，输入 `sbom:<path>`，输出 JSON 报告到 `./tmp/scanResult`，静态暴露 `/scanResult/`。
- 上传：`pkgUploadController` 保存文件到 `./tmp/uploads`，静态暴露 `/uploads/`。

## 流程示意
```
上传包 → 解压 → 生成 SBOM → 保存/暴露 → 基于 SBOM 扫描 → 报告保存/暴露
        ↘ 生成 DID 文档 → 上链（Fabric） → 查询/更新/吊销
```

## 目录速览
- `controller/`：用户、DID、软件包、SBOM、漏洞库、漏洞扫描。
- `util/`：JWT、CORS、解压、Syft/Grype 封装、Swagger。
- `did_connnection/`：Fabric 连接。
- `model/`：数据模型与请求/响应体。
- `docs/`：Swagger 源文件。
- `tmp/`：运行期文件（上传、SBOM、扫描报告）。

## 启动前准备
1) Go 1.20+（`go.mod` 要求 1.24.1）。  
2) 数据库：修改 `util/dataBase/DataBase.go` 的 `dsn`、`dsn_vulnerability`，确保能连 MySQL。  
3) Fabric：根据实际网络调整 `did_connnection/did_connnection.go` 的证书路径、`peerEndpoint`、链码名、通道名。  
4) 目录权限：`./tmp/uploads`、`./tmp/sbomStorage`、`./tmp/scanResult` 需可写（启动会自动创建）。  
5) 安全：替换 `util/JwtUtil.go` 的 `jwtSecret` 及硬编码账号。

## 启动方式
```bash
# 开发模式
go run main.go

# 构建 Linux 版本
sh build.sh
```
启动后打开 `http://localhost:8000/swagger/index.html` 调试接口。

## 常用接口速查（更多见 Swagger）

| 模块 | 方法 | 路径 | 说明 | 鉴权 |
| --- | --- | --- | --- | --- |
| 账号 | POST | `/api/v1/account/register` | 注册 | 无 |
| 账号 | POST | `/api/v1/account/login` | 登录获取 JWT | 无 |
| 账号 | GET | `/api/v1/account/getUserInfo` | 获取当前用户信息 | JWT |
| DID | POST | `/api/v1/did/create_identity` | 传软件名生成 DID 与文档 | JWT（通过后续接口调用时需） |
| DID | GET | `/api/v1/did/query_identity?did=xxx|all` | 查询单个或全部 DID | JWT |
| DID | PUT | `/api/v1/did/update_identity` | 更新 DID 文档 | JWT |
| DID | DELETE | `/api/v1/did/remove_identity` | 吊销 DID（表单 `didID`） | JWT |
| 软件包 | POST | `/api/v1/pkg/upload` | 上传压缩包，返回访问地址 | JWT |
| 软件包 | POST | `/api/v1/pkg/query` | 分页查询上传记录 | JWT |
| 软件包 | POST | `/api/v1/pkg/getDetail` | 根据文件名获取 DID/SBOM/报告 | JWT |
| SBOM | POST | `/api/v1/sbom/generate` | 基于上传包生成 SBOM（SPDX/CycloneDX） | JWT |
| 漏洞扫描 | POST | `/api/v1/vulnerability/scan` | 基于 SBOM 生成漏洞报告 | JWT |
| 漏洞库 | POST | `/api/v1/vulnerability/query` | 查询公开漏洞 | JWT |
| 漏洞库 | POST | `/api/v1/vulnerability/queryByUser` | 查询个人创建的漏洞 | JWT |
| 漏洞库 | POST | `/api/v1/vulnerability/create` | 创建个人漏洞记录 | JWT |
| 漏洞库 | GET | `/api/v1/vulnerability/getRankDistribution` | 严重度分布统计 | JWT |
| 健康检查 | POST | `/api/v1/service/checkSoft` | 健康探针 | JWT |
| 健康检查 | POST | `/api/v1/service/checkSoftV2` | 健康探针（权限校验版） | JWT（高权限） |

## 静态访问路径
- 上传包：`http://<host>:8000/uploads/<filename>`
- SBOM：`http://<host>:8000/sbom_list/<filename>`
- 漏洞报告：`http://<host>:8000/scanResult/<filename>`
