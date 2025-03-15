### Excel 生成器操作使用说明

#### 1. 程序功能概述
该程序是一个基于模板生成多个 Excel 文件的自动化工具。主要功能包括：
- 命令行参数解析
- 配置文件加载
- 目录创建
- 文件生成
- 模板字符串替换
- 随机值生成
- Excel 文件操作
- 错误处理
- 日志输出

#### 2. 配置文件说明
配置文件 `config.yaml` 用于指定生成 Excel 文件的相关参数和模板规则。主要配置项如下：

- `template_path`: 模板文件路径
- `output_dir`: 输出目录
- `file_count`: 生成文件数量
- `file_num_pattern`: 文件序号替换模式
- `sheet_pool_size`: 工作表处理池大小
- `templates`: 模板规则列表，用于定义随机值的生成规则

#### 3. 模板规则说明
模板规则用于定义 Excel 文件中需要替换的字符串及其对应的随机值生成规则。每个模板规则包含以下字段：

- `pattern`: 规则标签名，需唯一值，例：@int<-30,30>.cell、@float.2<25,26>，其中 `@` 表示标签前缀，`int` 表示整数类型，`float` 表示浮点数类型，`<-30,30>` 表示整数范围，`<25,26>` 表示浮点数范围，`.2` 表示小数位数，`.cell`表示值占满单元格。
- `type`: 数据类型，支持 `int`、`float`、`text`
- `min`: 最小值（仅适用于 `int` 和 `float` 类型）
- `max`: 最大值（仅适用于 `int` 和 `float` 类型）
- `decimal`: 小数位数（仅适用于 `float` 类型）

#### 4. 命令行参数说明
程序支持以下命令行参数：

- `-words`: 指定重复的字符串
- `-count`: 指定重复的次数

示例：
```bash
./excel-generator -words "Hello" -count 3
```
输出：
```
Hello,Hello,Hello
```

#### 5. 运行程序
1. 确保配置文件 `config.yaml` 已正确配置。
2. 在命令行中运行程序：
   ```bash
   ./excel-generator
   ```
3. 程序将根据配置文件生成指定数量的 Excel 文件，并输出到指定目录。

#### 6. 注意事项
- 确保模板文件路径和输出目录路径正确。
- 配置文件中的模板规则需与模板文件中的字符串匹配。
- 程序支持并发处理工作表，可通过 `sheet_pool_size` 配置并发数。

#### 7. 示例配置文件
```yaml
template_path: ./templates/（2025HD10地块-@file_num田块）田块土方开挖单元工程施工质量验收评定表.xlsx
output_dir: ./output
file_count: 2
file_max_count: 100
file_num_pattern: "@file_num"
sheet_pool_size: 2

templates:
  - pattern: "@int<-30,30>.cell"
    type: int
    min: -32
    max: 30
  - pattern: "@float.2<25,26>"
    type: float
    min: 25
    max: 26
    decimal: 2
```

#### 8. 示例模板文件
模板文件中包含需要替换的字符串，如 `@file_num` 、`@int<-30,30>.cell`、`@float.2<25,26>`，程序将根据配置文件中的规则生成随机值并替换这些字符串。

#### 9. 输出文件
生成的 Excel 文件将保存在 `output_dir` 指定的目录中，文件名中的 `@file_num` 将被替换为文件序号。

#### 10. 错误处理
程序在运行过程中会捕获并输出错误信息，如配置文件加载失败、目录创建失败、文件生成失败等。

通过以上步骤，您可以成功使用该程序生成多个 Excel 文件。

