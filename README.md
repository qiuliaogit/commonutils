# commonutils
## 使用办法
```bash
# https://github.com/qiuliaogit/commonutils.git
go get github.com/qiuliaogit/commonutils
```
## 主要内容
- int.go 定义泛型的整数类型和字符串转整数的函数和随机范围的函数
- isin.go 判断指定元素是否再数组中的函数
- page.go mysql用于分页的类
- ret.go 通用返回值的类
- set.go 基于map实现的集合功能
- string.go 主要是字符串链接和pad的函数
- time.go 用于时间处理的功能函数
- array.go 数组相关工具函数
- map.go map相关工具函数
- json.go json相关工具函数
- md5.go md5相关工具函数
- utils.go 其他工具函数
- version.go 版本信息
- datenum.go 一个数字日期的工具函数
- tdatetime.go 一个时间日期TDateTime的类

## 更新日志
### 1.0.11
- 增加ParamDateTime是否为空字符的函数
### 1.0.10
- 增加NewGormWhere函数，用于简化Gorm的条件查询
### 1.0.9
- 增加GormWhere类，用于简化Gorm的条件查询
- 增加ParamDateTime类，用于处理日期参数
