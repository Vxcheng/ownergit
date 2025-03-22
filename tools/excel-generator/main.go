package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
)

var (
	words      = flag.String("words", "", "words of the repeat")
	count      = flag.Int("count", 1, "count of the repeat")
	sep        = flag.String("sep", "，", "separator of the repeat")
	configPath = flag.String("config.path", "./config.yaml", "config.path of app")
	Log        *log.Logger
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	flag.Parse()
	if *words != "" {
		fmt.Println(joinWords(*words, *count, *sep))
		return
	}
	// 初始化日志
	if err := initLogger(); err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		return
	}

	// 加载配置文件
	cfg, err := loadConfig(*configPath)
	if err != nil {
		LogPrintf("加载配置失败: %v\n", err)
		return
	}

	// 创建输出目录
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		LogPrintf("创建目录失败: %v\n", err)
		return
	}

	// 生成指定数量的文件
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= cfg.FileCount; i++ {
		// 可并发
		if err := generateFile(cfg, i); err != nil {
			LogPrintf("生成文件失败: %v\n", err)
			continue
		}
		LogPrintf("成功生成第 %d 个文件\n", i)
	}
}

// 配置文件结构
type Config struct {
	TemplatePath   string     `yaml:"template_path"` // 模板文件路径
	OutputDir      string     `yaml:"output_dir"`    // 输出目录
	FileCount      int        `yaml:"file_count"`    // 生成文件数量
	FileMaxCount   int        `yaml:"file_max_count"`
	FileNumPattern string     `yaml:"file_num_pattern"`
	SheetPoolSize  int        `yaml:"sheet_pool_size"`
	FilePoolSize   int        `yaml:"file_pool_size"`
	Tag            Tag        `yaml:"tag"`
	Templates      []Template `yaml:"templates"`
}

type Template struct {
	Pattern        string         `yaml:"pattern"` // 正则表达式（如 "\$num"）
	Type           RuleType       `yaml:"type"`    // 数据类型: int/date/text
	Min            float64        `yaml:"min,omitempty"`
	Max            float64        `yaml:"max,omitempty"`
	RandomFloating float64        `yaml:"random_floating,omitempty"`
	Options        []string       `yaml:"options,omitempty"`
	Decimal        int            `yaml:"decimal,omitempty"` // 小数位数
	Reg            *regexp.Regexp `yaml:"-"`

	Sep   string `yaml:"sep,omitempty"`
	Rule1 Rule   `yaml:"rule1"`
}
type Rule struct {
	Type           RuleType `yaml:"type"` // 数据类型: int/date/text
	Min            float64  `yaml:"min,omitempty"`
	Max            float64  `yaml:"max,omitempty"`
	RandomFloating float64  `yaml:"random_floating,omitempty"`
	// Options        []string `yaml:"options,omitempty"`
	Decimal int `yaml:"decimal,omitempty"` // 小数位数
}

func (t *Template) Pass(value Value) bool {
	v, _ := strconv.ParseFloat(value.v, 64)
	ok := v >= t.Min && v <= t.Max
	if value.v1 != "" {
		v1, _ := strconv.ParseFloat(value.v1, 64)
		ok = ok && (v1 >= t.Rule1.Min && v1 <= t.Rule1.Max)
	}

	return ok
}

func (t *Template) HasRules() bool {
	return t.Type != "" && t.Rule1.Type != ""
}

type RuleType string

const (
	RuleTypeInt   RuleType = "int"
	RuleTypeFloat RuleType = "float"
)

func (t Template) Number() bool {
	ok := t.Type == RuleTypeInt || t.Type == RuleTypeFloat
	if t.HasRules() {
		ok = ok && (t.Rule1.Type == RuleTypeInt || t.Rule1.Type == RuleTypeFloat)
	}
	return ok
}

func (temp Template) Check() (err error) {
	if temp.Type != "" && temp.Max <= temp.Min {
		return fmt.Errorf("template pattern(%s) rule max(%f) must be greater than min(%f)", temp.Pattern, temp.Max, temp.Min)
	}
	if temp.Rule1.Type != "" && temp.Rule1.Max <= temp.Rule1.Min {
		return fmt.Errorf("template pattern(%s) rule1 max(%f) must be greater than min(%f)", temp.Pattern, temp.Max, temp.Min)
	}
	return
}

type Tag struct {
	Cell     string `yaml:"cell"`
	PassNum  string `yaml:"pass_num"`
	PassRate string `yaml:"pass_rate"`
}

type ExcelizeFile struct {
	*excelize.File
	fileNum int
}

type Sheet struct {
	name     string
	stats    map[string]*RuleStat
	statAxis map[string]string
}
type RuleStat struct {
	values                   []Value
	PassNum, PassRate, Total int
	Template
}

type Value struct {
	v, v1 string
}

func initLogger() error {
	// 创建 logs 目录
	if err := os.MkdirAll("logs", 0755); err != nil {
		return fmt.Errorf("创建 logs 目录失败: %v", err)
	}

	// 生成带时间戳的日志文件名
	logFileName := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02_15-04-05"))
	logFile, err := os.Create(logFileName)
	if err != nil {
		return fmt.Errorf("创建日志文件失败: %v", err)
	}

	// 初始化 logger
	Log = log.New(logFile, "", log.LstdFlags)
	return nil
}

func LogPrintf(format string, v ...any) {
	fmt.Printf(format, v...)
	Log.Printf(format, v...)
}

func LogPrintln(v ...any) {
	fmt.Println(v...)
	Log.Println(v...)
}

func NewExcelizeFile(cfg *Config, fileNum int) (h *ExcelizeFile, err error) {
	f, err := excelize.OpenFile(cfg.TemplatePath)
	if err != nil {
		return
	}

	h = &ExcelizeFile{
		File:    f,
		fileNum: fileNum,
	}
	return
}

// joinWords words 指定count次数用,进行拼接
func joinWords(words string, count int, sep string) string {
	if count <= 0 {
		return ""
	}
	return strings.TrimSuffix(strings.Repeat(words+sep, count), sep)
}

// 加载配置文件
func loadConfig(path string) (cfg *Config, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	cfg = &Config{}
	if err = yaml.Unmarshal(data, cfg); err != nil {
		return
	}

	err = cfg.initTemplates()
	if err != nil {
		return
	}

	if err = cfg.Check(); err != nil {
		return
	}
	return
}

func (c *Config) initTemplates() (err error) {
	for i, temp := range c.Templates {
		expr := strings.TrimSpace(temp.Pattern)
		expr = regexp.QuoteMeta(expr)
		expr = fmt.Sprintf(`(%s)`, expr)
		c.Templates[i].Reg, err = regexp.Compile(expr)
		if err != nil {
			return fmt.Errorf("parse template pattern(%s) error: %v", temp.Pattern, err)
		}
		if err = temp.Check(); err != nil {
			return
		}
	}

	return
}

func (c *Config) Check() (err error) {
	if filepath.Ext(c.TemplatePath) != ".xlsx" {
		return errors.New("模板文件格式错误，请使用 .xlsx 格式")
	}

	if c.FileCount < 1 {
		c.FileCount = 1
	}
	if c.SheetPoolSize < 1 {
		c.SheetPoolSize = 1
	}
	if c.FilePoolSize < 1 {
		c.FilePoolSize = 1
	}
	if c.FileMaxCount < 1 {
		c.FileMaxCount = 1
	}

	// tag
	tag := c.Tag
	if tag.Cell == "" {
		tag.Cell = "cell"
	}
	if tag.PassNum == "" {
		tag.PassNum = "pnum"
	}
	if tag.PassRate == "" {
		tag.PassRate = "prate"
	}
	c.Tag = tag
	return
}

// 生成单个文件
func generateFile(cfg *Config, fileNum int) (err error) {
	// 打开模板文件
	f, err := NewExcelizeFile(cfg, fileNum)
	if err != nil {
		return
	}
	defer f.Close()

	// 计算序号位数（如 count=100 → 3位）
	numWidth := len(strconv.Itoa(cfg.FileMaxCount))
	fileNumStr := fmt.Sprintf("%0*d", numWidth, fileNum)

	// 处理模板中的 @file_num 标记
	// 新增：执行特殊字符串替换
	if err = f.replaceTemplateStrings(cfg, fileNumStr); err != nil {
		return err
	}

	// 保存新文件
	baseName := filepath.Base(cfg.TemplatePath)
	if strings.Contains(baseName, cfg.FileNumPattern) {
		// 替换后的新文本
		baseName = strings.ReplaceAll(baseName, cfg.FileNumPattern, fileNumStr)
	}
	outputPath := filepath.Join(cfg.OutputDir, baseName)
	return f.SaveAs(outputPath)
}

// 处理特殊字符串替换
func (f *ExcelizeFile) replaceTemplateStrings(cfg *Config, fileNumStr string) (err error) {
	// 创建一个容量的 Goroutine 池
	p, err := ants.NewPool(cfg.SheetPoolSize)
	if err != nil {
		return
	}
	defer p.Release()

	sheetNum := len(f.GetSheetList())
	errC := make(chan error, sheetNum)
	for _, sheet := range f.GetSheetList() {
		// 遍历所有工作表，可并发
		_ = p.Submit(func() {
			err = f.replaceSheet(sheet, cfg, fileNumStr)
			errC <- err
		})
	}

	var errMsg []string
	for i := 0; i < sheetNum; i++ {
		if err1 := <-errC; err1 != nil {
			errMsg = append(errMsg, err1.Error())
		}
	}
	if len(errMsg) > 0 {
		err = errors.New(strings.Join(errMsg, "; "))
	}
	return
}

// replaceSheet 替换Excel文件中的指定工作表内容。
// 该函数接收一个工作表名称、一个配置对象和一个文件编号字符串作为参数。
// 它会根据配置中的规则替换工作表中的内容，并返回可能发生的错误。
func (f *ExcelizeFile) replaceSheet(sheet string, cfg *Config, fileNumStr string) (err error) {
	// 使用defer和recover来捕获并处理可能的panic，确保函数能够优雅地处理异常情况。
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()

	st := &Sheet{
		name:     sheet,
		stats:    make(map[string]*RuleStat),
		statAxis: make(map[string]string),
	}
	// 替换值，获取随机值数组暂存
	var match bool
	var newText string
	var values []Value

	// 获取指定工作表的所有行数据。
	rows, err := f.GetRows(sheet)
	if err != nil {
		return
	}

	// 遍历每个单元格
	for rowIdx, row := range rows {
		for colIdx, cellValue := range row {
			// 将单元格坐标转换为Excel中的单元格名称。
			axis, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)

			// 尝试用文件编号替换单元格内容。
			match, newText = replaceTitle(sheet, cfg.FileNumPattern, cellValue, fileNumStr)
			if match {
				// 替换成功后更新单元格内容，并继续处理下一个单元格。
				if err = f.SetCellValue(sheet, axis, newText); err != nil {
					return
				}
				continue
			}

			if cfg.Tag.Match(cellValue) {
				st.statAxis[axis] = cellValue
				continue
			}

			// 匹配每个单元格，匹配成功后，重新获取cellValue
			for _, tempRule := range cfg.Templates {
				// 根据配置中的模板规则替换单元格内容。
				match, newText, values = replaceContext(sheet, tempRule, cellValue)
				if match {
					// 替换成功后更新单元格内容，并使用新内容作为后续处理的基础。
					if err = f.SetCellValueByRule(sheet, tempRule, cfg.Tag, axis, newText); err != nil {
						return
					}
					cellValue = newText

					rs, ok := st.stats[tempRule.Pattern]
					if !ok {
						st.stats[tempRule.Pattern] = &RuleStat{
							Template: tempRule,
							values:   values,
						}
					} else {
						rs.values = append(rs.values, values...)
						st.stats[tempRule.Pattern] = rs
					}
				}
			}
		}
	}

	if !st.CheckCount() {
		return
	}
	st.Count() // 可校准合格
	if err = st.Replace(f, cfg.Tag); err != nil {
		return
	}
	return
}

func (f *ExcelizeFile) SetCellValueByRule(sheet string, tempRule Template, tag Tag, axis, newText string) (err error) {
	// 默认是文本内容
	var value interface{} = newText
	if strings.HasSuffix(tempRule.Pattern, "."+tag.Cell) {
		switch tempRule.Type {
		case RuleTypeInt:
			value, _ = strconv.Atoi(newText)
		case "float":
			value, _ = strconv.ParseFloat(newText, 64)
		}
	}

	return f.SetCellValue(sheet, axis, value)
}

func replaceContext(sheet string, tempRule Template, cellValue string) (match bool, newText string, values []Value) {
	re := tempRule.Reg
	//将cellValue匹配re的每个值都调用 generateTemplateValue获取随机值，并替换当前匹配的值，可用string.Count+values
	if re.MatchString(cellValue) {
		match = true
		newText = re.ReplaceAllStringFunc(cellValue, func(s string) string {
			rule := Rule{
				Type:           tempRule.Type,
				Decimal:        tempRule.Decimal,
				Max:            tempRule.Max,
				Min:            tempRule.Min,
				RandomFloating: tempRule.RandomFloating,
			}
			value := Value{}
			value.v = generateTemplateValue(rule)
			disp := value.v
			if tempRule.HasRules() {
				value.v1 = generateTemplateValue(tempRule.Rule1)
				disp = strings.Join([]string{value.v, value.v1}, tempRule.Sep)
			}
			values = append(values, value)
			return disp
		})
	}
	return
}

func replaceTitle(sheet, pattern string, cellValue, numStr string) (match bool, newText string) {
	if strings.Contains(cellValue, pattern) {
		match = true
		// 替换后的新文本
		newText = strings.ReplaceAll(cellValue, pattern, numStr)
	}
	return
}

// 生成替换值
func generateTemplateValue(rule Rule) (value string) {
	max := rule.Max
	min := rule.Min
	if rule.RandomFloating > 0 {
		min -= rule.RandomFloating
		max += rule.RandomFloating
	}
	switch rule.Type {
	case RuleTypeInt:
		v := rand.Intn(int(max-min)) + int(min)
		value = strconv.Itoa(v)
	case RuleTypeFloat:
		v := rand.Float64()*(max-min) + min
		value = strconv.FormatFloat(v, 'f', rule.Decimal, 64)
	// case "text":
	// 	if len(rule.Options) > 0 {
	// 		value = rule.Options[rand.Intn(len(rule.Options))]
	// 	}
	default:
	}
	return
}

func (st *Sheet) Count() {
	for _, stat := range st.stats {
		stat.Count()
	}
}

func (st *Sheet) CheckCount() bool {
	return len(st.statAxis) > 0
}

// RuleStat Count，计算规则的统计信息
func (rs *RuleStat) Count() {
	rs.Total = len(rs.values)
	if !rs.Template.Number() {
		return
	}

	for _, vs := range rs.values {
		if rs.Template.Pass(vs) {
			rs.PassNum++
		}
	}

	rs.PassRate = int(float64(rs.PassNum) / float64(rs.Total) * 100)
}
func (st *Sheet) Replace(f *ExcelizeFile, tag Tag) (err error) {
	var (
		match   bool
		newText string
	)
	// 根据暂存，替换统计字段
	for axis, cellValue := range st.statAxis {
		// 匹配tag，进行替换
		match, newText = st.replace(tag, cellValue)
		if match {
			// 替换成功后更新单元格内容，并继续处理下一个单元格。
			if err = f.SetCellValue(st.name, axis, newText); err != nil {
				return
			}
			continue
		}
	}
	return
}
func (st *Sheet) replace(tag Tag, cellValue string) (match bool, newText string) {
	for _, stat := range st.stats {
		key := stat.Pattern + "." + tag.PassNum
		if strings.Contains(cellValue, key) {
			match = true
			cellValue = strings.ReplaceAll(cellValue, key, strconv.Itoa(stat.PassNum))
		}
		key = stat.Pattern + "." + tag.PassRate
		if strings.Contains(cellValue, key) {
			match = true
			cellValue = strings.ReplaceAll(cellValue, key, strconv.Itoa(stat.PassRate))
		}
	}
	newText = cellValue
	return
}

func (tag Tag) Match(cellValue string) (match bool) {
	key := "." + tag.PassNum
	if strings.Contains(cellValue, key) {
		match = true
		return
	}

	key = "." + tag.PassRate
	if strings.Contains(cellValue, key) {
		match = true
	}
	return
}
