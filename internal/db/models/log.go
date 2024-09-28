package models

import "gorm.io/gorm"

/*
Создает модель `Log`, которая будет связана с таблицей `logs`, которая выглядит так:

create table logs (

	-- наше поле string
	text nvarchar not null,

	-- поля, которые добавились автоматически
	id uint primary key autoincrement,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,

)
*/
type Log struct {
	gorm.Model
	Text string
}
