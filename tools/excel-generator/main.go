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
	TemplatePath string `yaml:"template_path"` // 模板文件路径
	OutputDir    string `yaml:"output_dir"`    // 输出目录
	FileCount    int    `yaml:"file_count"`    // 生成文件数量
	// FileMaxCount   int        `yaml:"file_max_count"`
	FileNumPattern string     `yaml:"file_num_pattern"`
	SheetPoolSize  int        `yaml:"sheet_pool_size"`
	FilePoolSize   int        `yaml:"file_pool_size"`
	Templates      []Template `yaml:"templates"`
}

type Template struct {
	Pattern string         `yaml:"pattern"` // 正则表达式（如 "\$num"）
	Type    string         `yaml:"type"`    // 数据类型: int/date/text
	Min     float64        `yaml:"min,omitempty"`
	Max     float64        `yaml:"max,omitempty"`
	Options []string       `yaml:"options,omitempty"`
	Decimal int            `yaml:"decimal,omitempty"` // 小数位数
	Reg     *regexp.Regexp `yaml:"-"`
}

type ExcelizeFile struct {
	*excelize.File
	fileNum int
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
	for i := range c.Templates {
		expr := strings.TrimSpace(c.Templates[i].Pattern)
		expr = regexp.QuoteMeta(expr)
		expr = fmt.Sprintf(`(%s)`, expr)
		c.Templates[i].Reg, err = regexp.Compile(expr)
		if err != nil {
			return fmt.Errorf("parse template pattern(%s) error: %v", c.Templates[i].Pattern, err)
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
	numWidth := len(strconv.Itoa(cfg.FileCount))
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

func (f *ExcelizeFile) replaceSheet(sheet string, cfg *Config, fileNumStr string) (err error) {
	var match bool
	var newText string
	rows, err := f.GetRows(sheet)
	if err != nil {
		return
	}
	//提高性能， 正则初始化缓存
	// 遍历每个单元格
	for rowIdx, row := range rows {
		for colIdx, cellValue := range row {
			axis, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)
			match, newText = f.replaceTitle(sheet, cfg.FileNumPattern, cellValue, fileNumStr)
			if match {
				if err = f.SetCellValue(sheet, axis, newText); err != nil {
					return
				}
				continue
			}

			// 匹配每个单元格，匹配成功后，重新获取cellValue
			for _, tempRule := range cfg.Templates {
				match, newText = f.replaceContext(sheet, tempRule, cellValue)
				if match {
					if err = f.SetCellValueByRule(sheet, tempRule, axis, newText); err != nil {
						return
					}
					cellValue = newText
				}
			}
		}
	}
	return
}

func (f *ExcelizeFile) SetCellValueByRule(sheet string, tempRule Template, axis, newText string) (err error) {
	// 默认是文本内容
	var value interface{} = newText
	if strings.Contains(tempRule.Pattern, ".cell") {
		switch tempRule.Type {
		case "int":
			value, _ = strconv.Atoi(newText)
		case "float":
			value, _ = strconv.ParseFloat(newText, 64)
		}
	}

	return f.SetCellValue(sheet, axis, value)
}

func (f *ExcelizeFile) replaceContext(sheet string, tempRule Template, cellValue string) (match bool, newText string) {
	re := tempRule.Reg
	//将cellValue匹配re的每个值都调用 generateTemplateValue获取随机值，并替换当前匹配的值
	if re.MatchString(cellValue) {
		match = true
		newText = re.ReplaceAllStringFunc(cellValue, func(s string) string {
			return generateTemplateValue(tempRule)
		})
	}
	return
}

func (f *ExcelizeFile) replaceTitle(sheet, pattern string, cellValue, numStr string) (match bool, newText string) {
	if strings.Contains(cellValue, pattern) {
		match = true
		// 替换后的新文本
		newText = strings.ReplaceAll(cellValue, pattern, numStr)
	}
	return
}

// 生成替换值
func generateTemplateValue(rule Template) (value string) {
	switch rule.Type {
	case "int":
		value = strconv.Itoa(rand.Intn(int(rule.Max-rule.Min)) + int(rule.Min))
	case "float":
		v := rand.Float64()*(rule.Max-rule.Min) + rule.Min
		value = strconv.FormatFloat(v, 'f', rule.Decimal, 64)
	case "text":
		if len(rule.Options) > 0 {
			value = rule.Options[rand.Intn(len(rule.Options))]
		}
	default:
		value = rule.Pattern
	}
	return
}
