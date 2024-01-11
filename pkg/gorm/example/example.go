package example

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
	"u-sdk/pkg/gorm/query"
)

// //////////////////////////////////////////////////
// 用于测试 给模型添加一个方法
type CommonMethod struct {
	ID   int32
	Name *string
}

func (m *CommonMethod) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

func (m *CommonMethod) GetName() string {
	if m == nil || m.Name == nil {
		return ""
	}
	return *m.Name
}

// //////////////////////////////////////////////////
// 用于测试 给模型添加方法（动态SQL）
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)

	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error) // GetByID query data by id and return it as *struct*

	// SELECT * FROM @@table WHERE id IN @ids
	MGet(ids ...string) ([]*gen.T, error)

	// QueryWith
	//SELECT * FROM @@table
	//  {{if p != nil}}
	//      {{if p.ID > 0}}
	//          WHERE id=@p.ID
	//      {{else if p.ShopID != ""}}
	//          WHERE created_time=@p.ShopID
	//      {{end}}
	//  {{end}}
	//QueryWith(p *gen.T) (gen.T, error)
}

// //////////////////////////////////////////////////
func GenerateCode() {
	// 官方文档：https://gorm.io/zh_CN/gen/dao.html
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../query", // 生成代码的目录
		OutFile:      "gen.go",   // 生成文件名
		ModelPkgPath: "model",    //  表对应的结构体单独存放的目录
		//WithUnitTest: true,	// 生成测试用例

		// 生成模型全局配置
		FieldNullable:     true, // 数据库中的字段可为空，则生成struct字段为指针类型
		FieldCoverable:    true, // 如果数据库中字段有默认值，则生成指针类型的字段，以避免零值（zero-value）问题，见：https://gorm.io/docs/create.html#Default-Values
		FieldSignable:     true, // 根据数据库中列的数据类型，使用可签名类型作为字段的类型
		FieldWithIndexTag: true, // 为结构体生成gorm index tag，如gorm:"index:idx_name"，默认：false
		FieldWithTypeTag:  true, // 为结构体生成gorm type tag，如：gorm:"type:varchar(12)"，默认：false

		//gen.WithDefaultQuery	是否生成全局变量Q作为DAO接口，如果开启，你可以通过这样的方式查询数据 dal.Q.User.First()
		//gen.WithQueryInterface	生成查询API代码，而不是struct结构体。通常用来MOCK测试
		//gen.WithoutContext	生成无需传入context参数的代码。如果无需传入context，则代码调用方式如：dal.User.First()，否则，调用方式要像这样：dal.User.WithContext(ctx).First()
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/dar?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb)
	/////////////////////////////////////
	///////////// 全局设置  //////////////
	/////////////////////////////////////
	g.WithOpts(
		gen.FieldType("deleted_at", "gorm.DeletedAt"),
	)
	// 表名的命名策略
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if strings.HasPrefix(tableName, "_") { //忽略下划线开头的表
			return ""
		}
		return tableName
	})
	// 指定数据类型的映射关系
	g.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		// tinyint(1) -> int8
		// tinyint -> int16
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "int8"
			}
			return "int16"
		},
	})

	/////////////////////////////////////
	///////////// 单表设置  //////////////
	/////////////////////////////////////

	// 只会生成对应的表结构，不会生成查询方法。
	role := g.GenerateModel("admin_role")
	// 生成基础的CRUD方法
	g.ApplyBasic(
		role,
		// 生成 admin_account 表的结构体 （单表设置）
		g.GenerateModelAs(
			"admin_account",
			"AdminAccount",
			//通用的选项，允许你对整个字段进行修改（名字、类型、标签、关联关系等）
			//下面其他的选项都是对这个封装/简化
			//比如针对id字段，删除default标签
			gen.FieldModify(func(f gen.Field) gen.Field {
				if f.ColumnName == "id" {
					f.GORMTag.Remove(field.TagKeyGormDefault)
				}
				return f
			}),
			//
			gen.FieldNew("name", "*string", field.Tag{"json": "name", "gorm": "gormName"}), // 新增一个数据库不存在，但是model上需要的字段
			gen.FieldIgnore("password"),                   // 排除需要忽略的字段
			gen.FieldIgnoreReg("^password$"),              // 基于正则匹配，排除需要忽略的字段
			gen.FieldRename("email", "Email_yyp"),         // 自定义字段的model属性名
			gen.FieldComment("email", "这个是自定义的email注释信息"), // 自定义字段的注释信息，默认会同步数据库的备注信息
			gen.FieldType("email", "*string"),             // email字段在struct的类型
			gen.FieldTypeReg("^email$", "*string"),        // 根据正则表达式匹配字段，指定结构体的类型
			//gen.FieldGenType("status", "TSS"),             // 给 status 字段 使用一个新的字段类型。
			//gen.FieldGenTypeReg("^status$", "TSS"),        // 给 status 字段 使用一个新的字段类型。
			gen.FieldTag("mobile", func(tag field.Tag) field.Tag { // 给字段添加自定义标签 json2:id,string,omitempty
				return tag.Set("json2", "id,string,omitempty")
			}),
			gen.FieldJSONTag("mobile", "id,string,omitempty"), // 给 mobile字段添加json标签
			gen.FieldNewTagWithNS("avatar", func(columnName string) (tagContent string) { // 所有字段都加上 avatar:"列名称" 的标签
				return columnName
			}),
			//gen.FieldTrimPrefix("account_"), // 去掉字段前缀
			//gen.FieldTrimSuffix("_id"),      // 去掉字段后缀
			//gen.FieldAddPrefix("pre_"),      // 字段添加前缀
			//gen.FieldAddSuffix("_suf"),      // 字段添加后缀
			// 生成关联关系，并不会删除role_id字段，只是新增了一个数组
			gen.FieldRelate(
				field.HasOne,
				"Roles",
				g.GenerateModel("admin_role"),
				&field.RelateConfig{
					RelatePointer:      true,                     // *Roles
					RelateSlice:        false,                    //[]Roles
					RelateSlicePointer: false,                    //[]*Roles
					JSONTag:            "roles,string,omitempty", // 创建 json tag
					//GORMTag:            field.GormTag{"foreignKey": []string{"RoleID"}}.Append("references", "RoleID"), // 创建 GORM tag
					GORMTag: field.GormTag{}.Append("foreignKey", "RoleID").Append("references", "RoleID"),
					Tag:     field.Tag{"json_x_tag": "abcdef"}, // 新家一个tag
					//OverwriteTag:       field.Tag{"json_x_tag_2": "abcdef_over"},         // 会删除 json 和 json_x_tag 两个tag
				}),
			// 给 AdminAccount 模型添加一个方法
			//gen.WithMethod(CommonMethod{}),         // CommonMethod 的所有方法都添加到 model 中
			//gen.WithMethod(CommonMethod{}.IsEmpty), // 只添加 IsEmpty 方法
		),
	)

	// 在已有的表结构上生成自定义的接口(会覆盖上一个)
	//g.ApplyInterface(func(Querier) {},
	//	g.GenerateModel(
	//		"admin_account",
	//		gen.FieldModify(func(f gen.Field) gen.Field {
	//			if f.ColumnName == "id" {
	//				f.GORMTag.Remove(field.TagKeyGormDefault)
	//			}
	//			return f
	//		}),
	//		gen.FieldTag("phone", func(tag field.Tag) field.Tag {
	//			//tag.Set("kms","xxx")
	//			return tag.Set("encrypt", "xxx")
	//		}),
	//
	//		gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//			if strings.Contains(columnName, "time") {
	//				return columnName + ",omitempty"
	//			}
	//			return columnName
	//		}),
	//		gen.FieldJSONTag("id", "id,string,omitempty"),
	//		gen.FieldIgnore("create_time"),
	//		gen.FieldNewTag("phone", field.Tag{"encrypt": "xxx"}),
	//		gen.FieldNewTagWithNS("form", func(columnName string) (tagContent string) {
	//			return columnName
	//		}),
	//	),
	//	//model.AdminAccount{},
	//)
	// 生成代码
	g.Execute()
}

type rst struct {
	AccountID   uint32
	AccountName string
	RoleID      uint32
	GroupNameEn string
}

func Query() {
	gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/dar?charset=utf8mb4&parseTime=True&loc=Local"))
	gormdb.Debug()
	gormdb.Logger.LogMode(logger.Info)
	q := query.Use(gormdb)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// 查询关联数据，必须使用 Preload()方法 将关联数据预加载，会将预加载关联表的所有字段的数据都查询出来。
	// 两张表联查，2个sql(有一个预加载sql，一个查询sql)，合并到一个对象中
	// SELECT * FROM `admin_role` WHERE `admin_role`.`role_id` = 1
	// SELECT * FROM `admin_account` WHERE `admin_account`.`account_id` = 1 ORDER BY `admin_account`.`account_id` LIMIT 1
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	adm, err := q.AdminAccount.Preload(q.AdminAccount.Roles).Where(q.AdminAccount.AccountID.Eq(1)).Debug().First()
	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Printf("%+v\n", adm)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// 因为有 预加载 Preload() ,所以会将预加载关联表的所有字段的数据都查询出来。
	// 两张表联查，一个查询sql，一个预加载sql，合并到一个对象中。预加载表是全部数据都查询出来
	// SELECT * FROM `admin_role` WHERE `admin_role`.`role_id` IN (8,1,9,10)
	// SELECT `admin_account`.`account_id`,`admin_account`.`account_name`,`admin_role`.`role_id`,`admin_role`.`group_name_en`
	// FROM `admin_account` LEFT JOIN `admin_role` ON `admin_account`.`role_id` = `admin_role`.`role_id` WHERE `admin_account`.`account_id` IN (10057,10058,10064,10065)
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	adm2, err := q.AdminAccount.Preload(q.AdminAccount.Roles).
		Select(q.AdminAccount.AccountID, q.AdminAccount.AccountName, q.AdminRole.RoleID, q.AdminRole.GroupNameEn).
		LeftJoin(q.AdminRole, q.AdminAccount.RoleID.EqCol(q.AdminRole.RoleID)).
		Where(q.AdminAccount.AccountID.In(10057, 10058, 10064, 10065)).
		Debug().Find()
	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Printf("%+v\n", adm2)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// model模型中字段是 string 类型，无法倒出到result。必须是 *string 类型才能倒出
	// 一个sql，只查询需要的字段
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	var scanResult []rst
	// 两张表联查，一个sql，合并到一个对象中
	err3 := q.AdminAccount.Select(q.AdminAccount.AccountID, q.AdminAccount.AccountName, q.AdminRole.RoleID, q.AdminRole.GroupNameEn).
		LeftJoin(q.AdminRole, q.AdminAccount.RoleID.EqCol(q.AdminRole.RoleID)).
		Where(q.AdminAccount.AccountID.In(10057, 10058, 10064, 10065)).
		Debug().Limit(4).Scan(&scanResult)
	if err != nil {
		log.Printf("%s", err3)
	}

	// Dynamic SQL API
	//adms, err := query.AdminAccount.FilterWithNameAndRole("modi", "admin")
	//if err != nil {
	//	log.Printf("", err)
	//}
	//fmt.Printf("%+v\n", adms)
}
