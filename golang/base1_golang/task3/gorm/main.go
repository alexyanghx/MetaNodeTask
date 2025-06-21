package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SQL语句练习
// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	gorm.Model
	Name  string `gorm:"size:20;not null"`
	Age   int    `gorm:"size:3;check:age>0;not null"`
	Grade string `gorm:"size:20;not null"`
}

var db *gorm.DB

func init() {
	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
}

func task1_1() {
	db.AutoMigrate(&Student{})

	//清空数据
	db.Exec("DELETE FROM students")

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	var stus []Student
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	db.Where("age>?", 18).Find(&stus)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&Student{}).Where("name=?", "张三").Update("grade", "四年级")

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Delete(&Student{}, "age<?", 15)
}

// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	ID      int64
	Balance float64
}

type Transaction struct {
	ID            int64
	FromAccountId int64
	ToAccountId   int64
	Amount        float64
}

func task1_2() {
	db.AutoMigrate(&Account{}, &Transaction{})

	db.Exec("delete from accounts")
	db.Exec("delete from transactions")

	var accounts []Account = []Account{{
		ID:      1,
		Balance: 90.0,
	}, {
		ID:      2,
		Balance: 1000.0,
	}}

	db.Create(accounts)

	transfer := func(from_account_id int64, to_account_id int64, amount float64) {
		var fromAccount Account = Account{
			ID: from_account_id,
		}

		var toAccount Account = Account{
			ID: to_account_id,
		}

		db.Find(&fromAccount)
		db.Find(&toAccount)

		db.Transaction(func(tx *gorm.DB) error {
			fromAccount.Balance -= amount
			toAccount.Balance += amount

			if fromAccount.Balance < 0 {
				panic("余额不足")
			}
			tx.Save(&fromAccount)
			tx.Save(&toAccount)
			tx.Save(&Transaction{
				Amount:        amount,
				FromAccountId: from_account_id,
				ToAccountId:   to_account_id,
			})
			return nil
		})
	}

	transfer(1, 2, 100.0)
}

// Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Salary     float64 `db:"salary"`
	Department string  `db:"department"`
}

func task2_1() {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	var schema string = `
	CREATE TABLE if not exists employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(20) ,
		salary decimal(10,2) ,
		department varchar(20)
	);`

	db.MustExec(schema)

	db.MustExec("delete from employees")

	db.MustExec("insert into employees (name,Salary,department) values (?,?,?)", "Alex", 1000.0, "技术部")
	db.MustExec("insert into employees (name,Salary,department) values (?,?,?)", "Bob", 2000.0, "销售部")
	db.MustExec("insert into employees (name,Salary,department) values (?,?,?)", "Cindy", 3000.0, "财务部")
	db.MustExec("insert into employees (name,Salary,department) values (?,?,?)", "Judy", 3000.0, "财务部")
	//使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employees []Employee
	err = db.Select(&employees, "select * from employees where department=?", "技术部")
	if err != nil {
		panic(err)
	}
	for _, employee := range employees {
		fmt.Println(employee.Name, employee.Salary, employee.Department)
	}

	//使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employee Employee
	err = db.Get(&employee, "select * from employees order by salary desc limit 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(employee.Name, employee.Salary, employee.Department)

}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func task2_2() {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	var schema string = `
	CREATE TABLE if not exists books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title varchar(20) ,
		price decimal(10,2) ,
		author varchar(20)
	);`

	db.MustExec(schema)

	db.MustExec("delete from books")

	db.MustExec("insert into books(title,price,author)values('golang大全',10.0,'alex')")
	db.MustExec("insert into books(title,price,author)values('java大全',80.0,'alex')")

	var books []Book
	err = db.Select(&books, "select * from books where price>?", 50)
	if err != nil {
		panic(err)

	}
	fmt.Println(books)
}

// 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount int
}

type Post struct {
	ID           int `gorm:"primaryKey"`
	Title        string
	Content      string
	UserID       int
	Comments     []Comment `gorm:"foreignKey:PostID"`
	CommentCount int
	CommentState bool //false 表示“无评论”
}

type Comment struct {
	ID      int `gorm:"primaryKey"`
	Content string
	PostID  int
	UserID  int
	User    User `gorm:"foreignKey:UserID"`
}

func task3_1() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	db.Exec("delete from users")
	db.Exec("delete from posts")
	db.Exec("delete from comments")

	var alex User = User{
		Name:     "alex",
		Email:    "alex@email.com",
		Password: "password",
		Posts: []Post{
			{
				Title:   "alex post1",
				Content: "content1",
			},
			{
				Title:   "alex post2",
				Content: "content2",
			},
		},
	}
	db.Create(&alex)

	var zhangsan User = User{
		Name:     "张三",
		Email:    "zhangsan@email.com",
		Password: "password",
		Posts: []Post{
			{
				Title:   "张三 post1",
				Content: "content1",
			},
			{
				Title:   "张三 post2",
				Content: "content2",
			},
		},
	}

	db.Create(&zhangsan)

	db.Create([]Comment{
		{
			Content: "张三对alex post1的评论1",
			UserID:  zhangsan.ID,
			PostID:  alex.Posts[0].ID,
		},
	})
	// alex.Posts[0].CommentCount = 1
	var lisi User = User{
		Name:  "lisi",
		Email: "lisi@163.com",
	}
	db.Create(&lisi)

	db.Create([]Comment{
		{
			Content: "lisi对张三 post1的评论1",
			UserID:  lisi.ID,
			PostID:  zhangsan.Posts[0].ID,
		},
		{
			Content: "lisi对张三 post1的评论1",
			UserID:  lisi.ID,
			PostID:  zhangsan.Posts[0].ID,
		},
	})

}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func task3_2() {
	var user User
	db.Model(&User{}).Preload("Posts").Preload("Posts.Comments").Where("name=?", "alex").Find(&user)
	//使用joins关联查询user对象,并填充Posts和Posts.Comments

	// rows, err := db.Table("users").Joins("left join posts on posts.user_id=users.id").Joins("left join comments on comments.post_id=posts.id").Where("users.name=?", "alex").Rows()
	// if err != nil {
	// 	panic(err)
	// }

	// postMap := make(map[int]Post)
	// for rows.Next() {
	// 	var post Post
	// 	var comment Comment

	// 	if user.ID == 0 {
	// 		err = db.ScanRows(rows, &user)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}

	// 	err = db.ScanRows(rows, &post)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	err = db.ScanRows(rows, &comment)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if p, exists := postMap[post.ID]; exists {
	// 		p.Comments = append(p.Comments, comment)
	// 	} else {
	// 		post.Comments = []Comment{comment}
	// 		postMap[post.ID] = post
	// 		user.Posts = append(user.Posts, post)
	// 	}

	// }
	fmt.Println(user)

	var post Post
	//使用Gorm查询评论数量最多的文章信息。
	db.Order("comment_count DESC").First(&post)

	fmt.Println(post)
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (post *Post) AfterCreate(db *gorm.DB) error {
	fmt.Println("post AfterCreate ")
	db.Model(&User{}).Where("id = ?", post.UserID).Update("post_count", gorm.Expr("post_count + ?", 1))

	return nil
}

func (comment *Comment) AfterCreate(db *gorm.DB) error {
	var post Post
	db.Where("id = ?", comment.PostID).Find(&post)

	post.CommentCount++

	if !post.CommentState && post.CommentCount > 0 {
		post.CommentState = true
	}
	db.Updates(&post)

	return nil
}

func (comment *Comment) AfterDelete(db *gorm.DB) error {
	var post Post
	db.Find(&post, "id  = ?", comment.PostID)

	post.CommentCount--

	if post.CommentCount == 0 {
		post.CommentState = false
	}
	db.Select("comment_count", "comment_state").Updates(&post)

	return nil
}

func task3_3() {
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	var user User
	db.Where("name=?", "alex").Find(&user)
	db.Create(&Post{
		Title:   "alex post3",
		Content: "content3",
		UserID:  user.ID,
	})

	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	var post Post
	db.Find(&post, "id=?", 57)
	fmt.Println(post.CommentCount, post.CommentState)

	var comment Comment
	db.Find(&comment, "id=?", 38)
	db.Delete(&comment)

	db.Find(&post, "id=?", 57)
	fmt.Println(post.CommentCount, post.CommentState)

}

func main() {
	// task1_1()
	// task1_2()
	// task2_1()
	// task2_2()
	// task3_1()
	// task3_2()
	task3_3()
}
