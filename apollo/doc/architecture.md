# 技术方案设计

提供各类开关的配置管理，集成便于业务使用的用户侧友好接口，为业务的降级提供便捷的基础能力、及快速落地的抓手。
主要模块包括：配置系统(Apollo)、降级SDK、熔断器、灰度(ABTest)、toolkit(apollo-server apollo-cli apollo-agent)。
为解耦业务服务的依赖、引入复杂度，及程序健壮性，设计实现的是本机磁盘缓存 + 分布式存储 + 异步更新 的模型。
1. 业务服务读取基于本地 [ini](https://www.jianshu.com/p/7f60e3ee905b) 文件，通过SDK(监听本地文件变化，实时更新到内存)进行开关配置的取用
2. 配置管理基于分布式存储、或k3s ConfigMap 热更新等方式。管理员可通过修改分布式存储 或 ConfigMap 进行配置更新。

![apollo-instruction.png](apollo-instruction.png)

## 管理端（Manager）

管理端只需提供各类开关配置的维护更新，提交发布后即会变更到各个业务服务，无需重启服务。
每个业务服务独立维护各自的配置文件，系统会默认给所有服务生成一个带注释说明的空配置，如有需要各业务系统可直接基于默认配置进行更新修改。
目前提供如下可供参考的配置管理方式。


**单机模式 — Vim本地配置文件**

单机模式下， 管理员可直接基于本地文件的操作进行配置变更。


**集成Mars — k3s ConfigMap 热更新**

k3s/k8s集群模式下，可基于 ConfigMap 进行配置的更新管理。ConfigMap 既可以通过 watch 操作实现内容传播（默认形式），也可实现基于 TTL 的缓存，还可以直接经过所有请求重定向到 API 服务器。
[ConfigMap 热更新参考](https://jimmysong.io/kubernetes-handbook/concepts/configmap-hot-update.html)


**集成Mars - 老Mars mars_client 扩展**

扩展老 Mars 平台的 mars_client 组件，基于 Mysql + mars_client 进行配置的更新管理


**分布式存储 - ETCD**

以 ETCD 做为分布式存储，再分别提供 apollo-cli apollo-agent 两个工具用于配置的更新，以及配置的同步拉取。
apollo-cli：apollo 命令行工具，用于连接 ETCD（连接apollo-server），提供用户友好的命令行配置编辑接口
apollo-agent：默认集成到业务容器（或SDK），连接 ETCD （连接apollo-server），监听 Key 的变化，并实时同步到本地磁盘供业务使用
apollo-server(可选)：连接 ETCD，提供 HTTP API 接口用于配置的更新读取

## 业务端（Service-SDK）

业务服务通过引入 apollo-sdk 来获取配置，sdk 会默认加载本地配置文件，同时监听文件变化，文件一旦更新可实时更新到进程内存。

sdk提供如下接口：
```go
package apollo

type Agent struct{}

var agent *Agent // default agent

// InitAgent initializes the default apollo agent with the given file
func InitAgent(opts ...Option) error {
	return nil
}

// GetKey get apollo key with section
func GetKey(section, key string) *ini.Key {
	return agent.Get(section, key)
}


type Options struct{
	File string
	EndPoints string
	TTL time.Duration
}

type Option func(o *Options)

// WithFile sets the local file path
func WithFile(file string) Option {
	return func(o *Options) { o.File = file }
}

// WithEndpoints sets the remote endpoints, pull remote config to local with ttl
// endpoints schema support etcd://localhost, http://localhost
func WithEndpoints(endpoints string) Option {
	return func(o *Options) { o.EndPoints = endpoints }
}

// WithTTL sets the TTL of the remote endpoints synchronization
func WithTTL(ttl int64) Option {
	return func(o *Options) { o.TTL = time.Duration(ttl) *time.Second }
}
```


配置文件样例：
```ini
# apollo configuration file

# 降级
[degrade]
mysql = 0
redis = 1
tusd = 0

[degrade.fileSharing]
push_server = 2
key1 = 1

# 熔断
[circuit]
default_timeout = 1000 # DefaultTimeout is how long to wait for command to complete, in milliseconds
default_max_concurrent_requests = 10 # DefaultMaxConcurrent is how many commands of the same type can run at the same time
default_request_volume_threshold = 20 # DefaultVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
default_sleep_window = 5000 # DefaultSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
default_error_percent_threshold = 50 # DefaultErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests

[circuit.outgo.push]
timeout = 1000
max_concurrent_requests = 5
request_volume_threshold = 5
sleep_window = 2000
error_percent_threshold = 3

[circuit.flow.data_report]
timeout = 1000
max_concurrent_requests = 5
request_volume_threshold = 5
sleep_window = 2000
error_percent_threshold = 3


# others
[section_1]
key1 = 1
key2 = 2

[section_2]
key1 = 1
key2 = 2
```

### 业务降级开关

包装 Apollo 配置系统提供的 GetKey 接口， 使之更便于降级场景的使用。在需要使用手动降级的地方直接使用如下接口判断开关是否打开：
```go
package degrade

// Enabled determines whether the given section key is enabled
func Enabled(section, key string) bool {
	val := apollo.GetKey(section, key).MustInt(0)
	return val == Enable
}
```


### 依赖路径熔断/恢复

[熔断原理](https://www.jianshu.com/p/fc19f6ed6d0d)

使用 Netflix [hystrix-go](https://github.com/afex/hystrix-go) 实现熔断机制：
1. 服务初始化时通过 circuitbreaker.InitCommandConf 接口实现熔断器的配置初始化
2. 在需要熔断逻辑的地方使用 hystrix.Do("cmd", runHander, failHanler) 实现熔断逻辑

提供如下接口：
```go
package circuitbreaker

func InitCommandConf() {
	hystrix.Configure(nil)
}

func Do(key string, runHandler, failHanler func) error {
	return hystrix.Do(key, runHandler, failHandler)
}
```

### 业务功能灰度

灰度分布让一部分用户继续用产品特性A，一部分用户开始用产品特性B，如果用户对B没有什么反对意见，那么逐步扩大范围，把所有用户都迁移到B上面来。
灰度发布可以保证整体系统的稳定，在初始灰度的时候就可以发现、调整问题，以保证其影响度。
提供如下接口：
```go
package abtest

// Allow is the key allowed to be through
// ctx is the context.Context with common key-pairs e.g.: 
// phone: 123xxx 
// name: superAdmin 
// companyId: 0
func Allow(ctx context.Context, key string) bool {
	exprStr := apollo.GetKey("abtest", key).String()
	expr, err := ParseExpr(exprStr)
	if err != nil {
		return false
	}
	return expr.Execute(ctx)
}
```