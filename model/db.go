package model

import (
	"github.com/patrickmn/go-cache"
	//"github.com/syndtr/goleveldb/leveldb/cache"
	"gorm.io/gorm"
)

var GoCache *cache.Cache //缓存
var DB *gorm.DB
