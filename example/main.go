package main

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/kiraxie/dbgorm"
	"github.com/kiraxie/logzap"
	"github.com/kiraxie/rrule-go"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

//go:embed migrate/*.sql
var sqlMigration embed.FS

func main() {
	fmt.Println("Hello, playground")
	logzap.Reload(zapcore.DebugLevel, nil)
	db, err := dbgorm.New(
		context.Background(),
		dbgorm.Config{
			PrimaryDSN:   "postgres://postgres:password@localhost:35432/rrule?sslmode=disable&TimeZone=Asia/Tokyo",
			MaxOpenConns: 10,
		},
		dbgorm.WithLogger(log),
		dbgorm.WithGormLogger(log),
		dbgorm.WithMigrate(sqlMigration),
	)
	if err != nil {
		panic(err)
	}
	// if err := db.Rx(context.Background(), func(tx *gorm.DB) error {
	// 	user := User{}
	// 	if err := tx.Where("id = ?", 1).First(&user).Error; err != nil {
	// 		return err
	// 	}
	// 	// fmt.Println(fmt.Sprintf("%+v", user))
	// 	fmt.Println(user.Rule.String())

	// 	return nil
	// }); err != nil {
	// 	panic(err)
	// }
	if err := db.Rx(context.Background(), func(tx *gorm.DB) error {
		z := ZZZ{}
		if err := tx.First(&z).Error; err != nil {
			return err
		}
		// fmt.Println(fmt.Sprintf("%+v", user))
		// fmt.Println(user.Rule.String())
		fmt.Println(z.R.String())

		return nil
	}); err != nil {
		panic(err)
	}

	// r2, err := rrule.NewRRule(rrule.ROption{
	// 	Freq:     rrule.Daily,
	// 	Dtstart:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Until:    time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC),
	// 	Byhour:   []int{13, 17},
	// 	Byminute: []int{0, 30},
	// 	Bysecond: []int{0},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r2.Value())
	// if err := db.Wx(context.Background(), func(tx *gorm.DB) error {
	// 	return tx.Create(&Example{
	// 		R: *r2,
	// 	}).Error
	// }); err != nil {
	// 	panic(err)
	// }
	// if err := db.Wx(context.Background(), func(tx *gorm.DB) error {
	// 	return tx.Create(&User{
	// 		ID:        54,
	// 		Rule:      *r2,
	// 		UpdatedAt: time.Now(),
	// 		CreatedAt: time.Now(),
	// 	}).Error
	// }); err != nil {
	// 	panic(err)
	// }
}

type User struct {
	ID        int64       `json:"id" gorm:"column:id"`
	Rule      rrule.RRule `json:"rule" gorm:"column:rule"`
	Set       rrule.Set   `json:"set" gorm:"column:set"`
	UpdatedAt time.Time   `json:"-" gorm:"column:updated_at"`
	CreatedAt time.Time   `json:"-" gorm:"column:created_at"`
}

type Example struct {
	R rrule.RRule `json:"r" gorm:"column:r"`
}

func (Example) TableName() string {
	return "example"
}

type ZZZ struct {
	R rrule.Set `json:"r" gorm:"column:r"`
}

func (ZZZ) TableName() string {
	return "zzz"
}

var log = logzap.Get("main")
