package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 全局数据库 db
var db *gorm.DB

// User 用户表
type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        *string
	Age          uint8 `gorm:"default:18"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// type User2 struct {
//	gorm.Model
//	Name string
// }

func (u User) TableName() string {
	return "user"
}

// 包初始化函数，可以用来初始化 gorm
func init() {

	var err error
	// 连接 mysql 获取 db 实例
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "huang:root123456@tcp(182.61.6.85:3306)/demo_test?charset=utf8mb4&parseTime=True", // data source name
		DefaultStringSize:         256,                                                                               // default size for string fields
		DisableDatetimePrecision:  true,                                                                              // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                              // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                              // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                             // auto configure based on currently MySQL version
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)
}

// 获取 gorm db，其他包调用此方法即可拿到 db
// 无需担心不同协程并发时使用这个 db 对象会公用一个连接，因为 db 在调用其方法时候会从数据库连接池获取新的连接
func GetDB() *gorm.DB {
	return db
}

func autoBuildTable() {
	// https://gorm.io/docs/migration.html
	// 迁移
	_ = db.AutoMigrate(&User{})

	// _ = db.AutoMigrate(&User{}, &Product{}, &Order{})

	// Add table suffix when creating tables
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	// TODO Migrator Interface

	// Returns current using database name
	// db.Migrator().CurrentDatabase()

	// table
	// Create table for `User`
	// db.Migrator().CreateTable(&User{})

	// Append "ENGINE=InnoDB" to the creating table SQL for `User`
	// db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})

	// Check table for `User` exists or not
	// db.Migrator().HasTable(&User{})
	// db.Migrator().HasTable("users")

	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	// db.Migrator().DropTable(&User{})
	// db.Migrator().DropTable("users")

	// Rename old table to new table
	// db.Migrator().RenameTable(&User{}, &UserInfo{})
	// db.Migrator().RenameTable("users", "user_infos")

	// Columns

	// Constraints

	// Indexes
}

func curdOperate() {
	//
	// c()
	//
	// u()
	//
	r()
	//
	// d()

}

func c() {
	// https://gorm.io/docs/create.html

	// create
	// user := User{Name: "huang", Age: 18}
	// result := db.Create(&user) // pass pointer of data to Create
	// fmt.Println(user.ID, result.Error, result.RowsAffected)

	// Create a record and assign a value to the fields specified.
	// db.Select("Name", "Age", "CreatedAt").Create(&User{Name: "hhhhh", Age: 18})

	// Create a record and ignore the values for fields passed to omit.
	// db.Omit("Name", "Age", "CreatedAt").Create(&User{Name: "zzzz", Age: 19})

	// Batch Insert
	var users = []User{{Name: "hhh111"}, {Name: "hhh222"}, {Name: "hhh333"}}
	db.Create(&users)
	// batch size 100
	// db.CreateInBatches(users, 100)

	// Create Hooks

	// Create From Map

	// Create From SQL Expression/Context Valuer

	// Create With Associations
}

func u() {
	// https://gorm.io/docs/update.html

	var user User
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	// Save is a combination function. If save value does not contain primary key, it will execute Create, otherwise it will execute Update (with all fields)
	// Don’t use Save with Model, it’s an Undefined Behavior.
	db.Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

	// Update single column
	// Update with conditions
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User's ID is `111`:
	db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// Update with conditions and model value
	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// Updates multiple columns
	// Update attributes with `struct`, will only update non-zero fields
	db.Model(&user).Updates(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// Update attributes with `map`
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// Update Selected Fields

	// Update Hooks

	// Batch Updates

	// Update with SQL Expression
}

func r() {
	// https://gorm.io/docs/query.html
	var (
		user  User
		user2 = User{ID: 2}
		users []User
	)
	// Get the first record ordered by primary key
	result := db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	result2 := db.First(&user2)
	fmt.Println(result.Row(), result2.Row())

	// Get one record, no specified order
	// db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// Get last record, ordered by primary key desc
	// db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// result := db.First(&user)
	// fmt.Println(result.RowsAffected, result.Error)
	// result.RowsAffected // returns count of records found
	// result.Error        // returns error or nil

	// check error ErrRecordNotFound
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	// !!!note: the three methods below will return the error ErrRecordNotFound if no record is found
	// If you want to avoid the ErrRecordNotFound error, you could use Find like db.Limit(1).Find(&user), the Find method accepts both struct and slice data

	db.Find(&users)

	// Conditions
	// Get first matched record
	db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// Get all matched records
	// db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	// db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	// db.Where("updated_at > ?", "lastWeek").Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", "lastWeek", "today").Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// Struct & Map Conditions
	// Struct
	// db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	// db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// Slice of primary keys
	// db.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);
}

func d() {
	// db.Delete(&User{}, 1)

}

type Result struct {
	ID   int
	Name string
	Age  int
}

func rawSqlAndSQLBuilder() {
	// Query Raw SQL with Scan

	var (
		result Result
		user   User
	)
	db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

	db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

	var age int
	db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)

	var users []User
	db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)

	// Exec with Raw SQL
	db.Exec("DROP TABLE users")
	db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

	// Exec with SQL Expression
	db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")

	// Named Argument
	db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)
	// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu"

	db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "jinzhu2"}).First(&result)
	// SELECT * FROM `users` WHERE name1 = "jinzhu2" OR name2 = "jinzhu2" ORDER BY `users`.`id` LIMIT 1

	// Named Argument with Raw SQL
	db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
		sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)
	// SELECT * FROM users WHERE name1 = "jinzhu1" OR name2 = "jinzhu2" OR name3 = "jinzhu1"

	db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
		sql.Named("name", "jinzhunew"), sql.Named("name2", "jinzhunew2"))
	// UPDATE users SET name1 = "jinzhunew", name2 = "jinzhunew2", name3 = "jinzhunew"

	db.Raw("SELECT * FROM users WHERE (name1 = @name AND name3 = @name) AND name2 = @name2",
		map[string]interface{}{"name": "jinzhu", "name2": "jinzhu2"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"

	type NamedArgument struct {
		Name  string
		Name2 string
	}

	db.Raw("SELECT * FROM users WHERE (name1 = @Name AND name3 = @Name) AND name2 = @Name2",
		NamedArgument{Name: "jinzhu", Name2: "jinzhu2"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"
}

func main() {
	// 迁移
	// autoBuildTable()

	// 执行数据库查询操作
	curdOperate()
}

// TODO Goanno插件
