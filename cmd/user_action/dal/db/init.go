// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"dy/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			//预编译sql语句，提高效率（缓存）
			PrepareStmt: true,
			//跳过默认事务
			//如果使用gorm的hook或者关联创建时，false
			SkipDefaultTransaction: true,
		},
	)
	//初始化失败
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	//m := DB.Migrator()
	////迁移数据库
	//if m.HasTable(&Note{}) {
	//	return
	//}
	//if err = m.CreateTable(&Note{}); err != nil {
	//	panic(err)
	//}
}
