# go-design-pattern
常用设计模式在 Golang 中的简单实现与示例

## 项目简介
该项目以简洁示例演示 Go 语言中的常见创建型设计模式：
- 简单工厂（Simple Factory）
- 工厂方法（Factory Method）
- 抽象工厂（Abstract Factory）
- 单例（Singleton）：饿汉式、懒汉式、加锁版、原子版、sync.Once 版

通过 `main.go` 一次性运行所有示例，便于快速理解不同模式的意图、结构与使用方式。

## 目录结构
```
go-design-pattern
└── create
    ├── factory
    │   ├── simple_factory_patterns.go           # 简单工厂
    │   ├── factory_method_patterns.go           # 工厂方法
    │   └── abstract_factory_patterns.go         # 抽象工厂
    └── singleton
        ├── singleton.go                         # 单例模式统一入口
        ├── eager_initialization.go              # 饿汉式
        ├── lazy_initialization.go               # 懒汉式（非线程安全基础版）
        ├── lazy_initialization_with_lock.go     # 懒汉式（互斥锁版）
        ├── lazy_initialization_with_atomic.go   # 懒汉式（原子操作版）
        └── lazy_initialization_with_once.go     # 懒汉式（sync.Once 高性能版）
main.go
```

## 运行环境
- Go 版本：1.20+（建议）
- 操作系统：跨平台（Mac/Linux/Windows）

## 快速开始
1. 克隆本仓库后，在项目根目录执行：
   ```
   go run ./main.go
   ```
   若你自定义了模块名，请确保 `import` 路径与当前 `go.mod` 模块名一致（示例代码使用 `go-design-pattern/...`），必要时：
   ```
   go mod init go-design-pattern
   go run ./main.go
   ```

## 设计模式简述
- 简单工厂模式
  - 通过一个工厂类 `SimpleFruitFactory.CreateFruit(name)` 根据入参决定实例化哪个具体产品，客户端无需关心构造细节。
  - 适用：产品种类较少且变化不频繁。
  - 优点：封装创建逻辑、调用方简单、统一入口便于管理。
  - 缺点：对新增产品不友好（需修改工厂，违反开闭原则）；易出现大量分支判断。

- 工厂方法模式
  - 为每种产品定义独立工厂（如 `AppleFactory`、`OrangeFactory`），通过工厂的多态创建产品。
  - 适用：产品类型较多、变化频繁，需要遵守开闭原则的场景。
  - 优点：新增产品仅需新增工厂与产品类，符合开闭原则；创建与使用解耦，职责更清晰。
  - 缺点：类与工厂数量增多，增加系统复杂度；客户端需选择合适工厂（可能引入额外抽象或工厂注册表）。

- 抽象工厂模式
  - 定义一组相关或依赖的产品族创建接口（如 `CreateApple`、`CreateOrange`），由具体工厂（如 `CNFactory`、`USFactory`）生成同一产品族对象。
  - 适用：需要在多个产品族之间切换，且保证族内产品协同与兼容。
  - 优点：保证产品族一致性与兼容性；切换产品族（如地区/平台）成本低；集中管理多产品的创建。
  - 缺点：对“新增产品类型”不友好（需改抽象接口与所有具体工厂）；结构更复杂，学习与维护成本更高。

- 单例模式
  - 饿汉式：启动时直接实例化。
    - 优点：实现简单、天然线程安全、无竞争。
    - 缺点：非延迟加载，可能浪费资源；初始化顺序不当可能造成依赖问题。
  - 懒汉式：首次使用时创建。
    - 优点：延迟加载、按需初始化。
    - 缺点：非线程安全，存在竞态条件；并发环境不可用。
  - 加锁版：获取实例时加互斥锁。
    - 优点：线程安全、实现直观。
    - 缺点：每次获取都加锁，性能一般。
  - 原子版：利用原子操作减少锁开销。
    - 优点：在高并发下较少锁竞争，性能优于全局加锁。
    - 缺点：实现复杂，边界仍可能需要锁；代码可读性与维护性下降。
  - sync.Once 版：使用标准库一次性初始化。
    - 优点：高性能、语义清晰、避免双检锁陷阱。
    - 缺点：只支持一次初始化，复位困难；在测试中需特别处理全局状态。