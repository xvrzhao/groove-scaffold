package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"

	gormGen "gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Generator interface {
	GenModel(db *gorm.DB, tableName, asModelName string) (dstFile string, err error)
	GenService(modelName string) (dstFile string, err error)
	GenController(modelName string) (dstFile string, err error)
	GenRouter(modelName string) (dstFile string, err error)
	GenAll(db *gorm.DB, tableName, asModelName string) error
}

type generator struct {
	moduleName string

	modelTplPath  string
	svcTplPath    string
	ctrlTplPath   string
	routerTplPath string

	dstModelPathFmt string
	dstSvcPathFmt   string
	dstCtrlPathFmt  string
}

func NewGenerator() Generator {
	f, err := os.Open("./go.mod")
	if err != nil {
		log.Printf("err: failed to open go.mod: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	firstLine := scanner.Text()
	moduleName := strings.TrimPrefix(firstLine, "module")
	moduleName = strings.TrimSpace(moduleName)

	return &generator{
		moduleName:      moduleName,
		modelTplPath:    "./cmd/gencode/templates/model.go.tpl",
		svcTplPath:      "./cmd/gencode/templates/svc.go.tpl",
		ctrlTplPath:     "./cmd/gencode/templates/ctrl.go.tpl",
		routerTplPath:   "./cmd/gencode/templates/router.out.tpl",
		dstModelPathFmt: "./db/model/%s_model.go",
		dstSvcPathFmt:   "./handler/%s_svc.go",
		dstCtrlPathFmt:  "./handler/%s_ctrl.go",
	}
}

func (generator) lowerCamelCase(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	for i := 0; i < len(s); i++ {
		d := s[i]
		if d == '_' {
			j = true
			continue
		}
		if j && d >= 'a' && d <= 'z' {
			d = d - 32
			j = false
		} else if i == 0 && d >= 'A' && d <= 'Z' {
			d = d + 32
		}
		data = append(data, d)
	}
	return string(data)
}

func (g generator) upperCamelCase(s string) string {
	s = g.lowerCamelCase(s)
	return strings.ToUpper(s[:1]) + s[1:]
}

func (g generator) GenModel(db *gorm.DB, tableName, asModelName string) (dstFile string, err error) {
	conf := gormGen.Config{}
	conf.WithJSONTagNameStrategy(g.lowerCamelCase)

	gg := gormGen.NewGenerator(conf)
	gg.UseDB(db)

	meta := gg.GenerateModelAs(
		tableName,
		asModelName,
		gormGen.FieldIgnore("id", "created_at", "updated_at", "deleted_at"),
		gormGen.FieldGORMTagReg(".*", func(tag field.GormTag) field.GormTag { return field.GormTag{} }),
	)

	tpl, err := template.ParseFiles(g.modelTplPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, map[string]any{
		"module":    g.moduleName,
		"struct":    meta.ModelStructName,
		"tableName": meta.TableName,
		"fields":    meta.Fields,
	}); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	dstFile = fmt.Sprintf(g.dstModelPathFmt, g.lowerCamelCase(asModelName))
	if _, err := os.Stat(dstFile); !os.IsNotExist(err) {
		return "", fmt.Errorf("file %s is already exists", dstFile)
	}

	f, err := os.Create(dstFile)
	if err != nil {
		return "", fmt.Errorf("failed to create dst file: %w", err)
	}
	defer f.Close()

	// format file
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", buf.String(), parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("failed to parse file ast: %w", err)
	}
	if err := format.Node(f, fset, astFile); err != nil {
		return "", fmt.Errorf("failed to format file: %w", err)
	}

	return dstFile, nil
}

func (g generator) GenService(modelName string) (dstFile string, err error) {
	tpl, err := template.ParseFiles(g.svcTplPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	dstFile = fmt.Sprintf(g.dstSvcPathFmt, g.lowerCamelCase(modelName))
	if _, err := os.Stat(dstFile); !os.IsNotExist(err) {
		return "", fmt.Errorf("file %s is already exists", dstFile)
	}

	f, err := os.Create(dstFile)
	if err != nil {
		return "", fmt.Errorf("failed to create dst file: %w", err)
	}
	defer f.Close()

	if err := tpl.Execute(f, map[string]any{
		"module":          g.moduleName,
		"lowerCamelModel": g.lowerCamelCase(modelName),
		"upperCamelModel": g.upperCamelCase(modelName),
	}); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return dstFile, nil
}

func (g generator) GenController(modelName string) (dstFile string, err error) {
	tpl, err := template.ParseFiles(g.ctrlTplPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	dstFile = fmt.Sprintf(g.dstCtrlPathFmt, g.lowerCamelCase(modelName))
	if _, err := os.Stat(dstFile); !os.IsNotExist(err) {
		return "", fmt.Errorf("file %s is already exists", dstFile)
	}

	f, err := os.Create(dstFile)
	if err != nil {
		return "", fmt.Errorf("failed to create dst file: %w", err)
	}
	defer f.Close()

	if err := tpl.Execute(f, map[string]any{
		"module":          g.moduleName,
		"lowerCamelModel": g.lowerCamelCase(modelName),
		"upperCamelModel": g.upperCamelCase(modelName),
	}); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return dstFile, nil
}

func (g generator) GenRouter(modelName string) (code string, err error) {
	tpl, err := template.ParseFiles(g.routerTplPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, map[string]any{
		"lowerCamelModel": g.lowerCamelCase(modelName),
		"upperCamelModel": g.upperCamelCase(modelName),
	}); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return buf.String(), nil
}

func (g generator) GenAll(db *gorm.DB, tableName, asModelName string) error {
	modelFile, err := g.GenModel(db, tableName, asModelName)
	if err != nil {
		return fmt.Errorf("failed to gennerate model file: %w", err)
	}
	log.Printf("model file: 👉 \033[3;4;32m%s\033[0m", modelFile)

	svcFile, err := g.GenService(asModelName)
	if err != nil {
		return fmt.Errorf("failed to gennerate service file: %w", err)
	}
	log.Printf("service file: 👉 \033[3;4;32m%s\033[0m", svcFile)

	ctrlFile, err := g.GenController(asModelName)
	if err != nil {
		return fmt.Errorf("failed to gennerate controller file: %w", err)
	}
	log.Printf("controller file: 👉 \033[3;4;32m%s\033[0m", ctrlFile)

	routerCode, err := g.GenRouter(asModelName)
	if err != nil {
		return fmt.Errorf("failed to gennerate router code: %w", err)
	}
	log.Printf("router code here, copy to the right place: \n\033[3;32m--------------------------------------------------------------\n\n%s\n\n--------------------------------------------------------------\033[0m\n", routerCode)

	log.Print("🎉 Complete!")
	return nil
}
