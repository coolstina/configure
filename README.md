![configure](assets/banner/configure.jpg)


## 安装

```shell script
go get -u github.com/coolstina/configure
```

## Configure是什么?

Configure 是 viper 增强包（用于快速添加项目配置）。支持功能如下：

- 通过选项可具体配置，配置文件名称，默认是 `config`。
- 通过选项可具体配置，配置文件格式如（toml/json/yaml/yml等），默认是 `toml`。
- 通过选项可具体配置，配置文件所在目录，默认是当前目录 `.`。
- 通过选项可具体配置，配置文件内容要映射的结构体，如 &Configure{}。

## 为什么要使用Configure?

每开启一个新的项目都离不开要配置一堆的配置信息，每次都去实现配置文件到配置结构体的映射，实属是浪费时间，导致不能提高开发效率，为了将这一部分重复的工作工具化，也就产生了该工具包，只需要几分钟的时间你就可以搞定项目配置信息到结构化的映射，非常方便和高效~

## 如何使用？

### Step1: 添加配置文件

例如，你的项目需要使用到MYSQL数据库和Redis(NoSQL)数据库，你可能需要添加如下配置文件config.toml：

```toml
[Application]
Name = "coolstina"
Version = "v1.0.0"

[Database]
[Database.MYSQL]
Host = "localhost:3306"
Username = "coolstina"
Password = ""
Database = "coolstina"

[Database.Redis]
Host = "localhost:6379"
Password = ""
Database = 5
```

### Step2: 添加映射结构体

要将 config.toml 配置文件中的配置信息映射到结构体，我们需要定义对应的结构体。根据上面的配置文件信息，定义的结构体如下：

```go
// Configure Configure the structure type.
type Configure struct {
	Application struct {
		Name    string 
		Version string 
	}

	Database struct {
		MYSQL struct {
			Host     string
			Username string 
			Password string
			Database string
		}

		Redis struct {
			Host     string
			Password string
			Database int
		}
	}
}
```

### Step3: 映射配置信息到结构体

configure.NewConfigure 构造函数参数为 Option 动态参数，可用的动态参数有：

- configure.WithSpecificConfigure: 指定要将配置文件信息映射到的结构体，该参数值具体是一个接口类型，必传参数。
- configure.WithSpecificConfigPath: 指定配置文件所在目录，如果不指定则在当前目录下查找配置文件。
- configure.WithSpecificConfigType: 指定配置文件类型，例如 `toml/yaml/yml/json`等viper支持的类型. 如果不指定默认是 `toml`。
- configure.WithSpecificConfigName: 指定配置文件的名称，如果不指定默认是 `config`。

```go
conf := configure.NewConfigure(
    configure.WithSpecificConfigure(&config.Configure{}),
    configure.WithSpecificConfigPath("example"),
).(*config.Configure)

marshal, err := json.MarshalIndent(conf, "", "	")
if err != nil {
    log.Panicf("JSON marshal has error: %s\n", err.Error())
}
```

具体映射也就是所期望的值，具体JSON后的数据如下所示：

```json
{
	"Application": {
		"Name": "coolstina",
		"Version": "v1.0.0"
	},
	"Database": {
		"MYSQL": {
			"Host": "localhost:3306",
			"Username": "coolstina",
			"Password": "",
			"Database": "coolstina"
		},
		"Redis": {
			"Host": "localhost:6379",
			"Password": "",
			"Database": 5
		}
	}
}
```
