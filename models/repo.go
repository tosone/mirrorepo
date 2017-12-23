package models

import (
	"time"

	"github.com/tosone/mirror-repo/common/defination"
)

type Repo struct {
	Id          int64                 `xorm:"pk autoincr"`    // 主键
	Address     string                `xorm:"notnull unique"` // 仓库地址
	Name        string                `xorm:"notnull unique"` // 本地存储所用的名字
	Schedule    uint                  `xorm:"notnull"`        // 仓库两次更新之间的时间间隔
	Status      defination.RepoStatus // 仓库状态
	CommitCount uint                  // commit 数量
	IsSendEmail bool                  // 是否发送邮件
	CreateAt    time.Time             `xorm:"created"` // 创建时间
	UpdateAt    time.Time             `xorm:"updated"` // 更新时间
	DeleteAt    *time.Time            `xorm:"deleted"` // 删除时间
}

// Create ..
func (repo *Repo) Create() (int64, error) {
	return engine.Insert(repo)
}

// Delete ..
func (repo *Repo) Delete() (int64, error) {
	return engine.Delete(&repo)
}

// Find ..
func (repo *Repo) Find() (*Repo, error) {
	var r = new(Repo)
	var err error
	_, err = engine.Id(repo.Id).Get(r)
	return r, err
}

func (repo *Repo) Update() (int64, error) {
	return engine.Id(repo.Id).Update(repo)
}

//func (repo *Repo) GetByAddress() (exist bool, err error) {
//	var keys []string
//	var retRepo *Repo
//	keys, err = redis.Strings(engine.Get().Do("keys", "*"))
//	if err != nil {
//		return
//	}
//	for _, key := range keys {
//		if strings.HasPrefix(key, fmt.Sprintf("hm:%s", funcs.BucketName(repo))) {
//			var values []interface{}
//			values, err = redis.Values(engine.Get().Do("HGETALL", key))
//			if err != nil {
//				return
//			}
//			err = redis.ScanStruct(values, &retRepo)
//			if err != nil {
//				return
//			}
//			if retRepo.Address == repo.Address && retRepo.DeleteAt == nil {
//				return
//			}
//		}
//	}
//	err = ErrNoSuchKey
//	return
//}
//
//func (repo *Repo) Create() (err error) {
//	_, err = repo.GetByID()
//	if err == ErrNoSuchKey {
//		var now = time.Now()
//		repo.ID = uuid.NewV4()
//		repo.CreateAt = now
//		repo.UpdateAt = now
//		_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(repo), repo.ID)).AddFlat(&repo)...)
//		if err != nil {
//			return
//		}
//		return
//	}
//	if err == nil {
//		err = ErrKeyAlreadyExist
//	}
//	return
//}
//
//func (repo *Repo) GetByID() (retRepo *Repo, err error) {
//	var keys []string
//	keys, err = redis.Strings(engine.Get().Do("keys", "*"))
//	if err != nil {
//		return
//	}
//	for _, key := range keys {
//		if strings.HasPrefix(key, fmt.Sprintf("hm:%s", funcs.BucketName(repo))) {
//			var values []interface{}
//			values, err = redis.Values(engine.Get().Do("HGETALL", key))
//			if err != nil {
//				return
//			}
//			err = redis.ScanStruct(values, &retRepo)
//			if err != nil {
//				return
//			}
//			if retRepo.ID == repo.ID && retRepo.DeleteAt == nil {
//				return
//			}
//		}
//	}
//	err = ErrNoSuchKey
//	return
//}
//
//func (repo *Repo) RemoveByID() (err error) {
//	_, err = repo.GetByID()
//	if err != nil {
//		return
//	}
//	now := time.Now()
//	repo.DeleteAt = &now
//	_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(repo), repo.ID)).AddFlat(&repo)...)
//	if err != nil {
//		return
//	}
//	return
//}
//
//func (repo *Repo) UpdateByID(newRepo Repo) (err error) {
//	_, err = repo.GetByID()
//	if err != nil {
//		return
//	}
//	now := time.Now()
//	repo.UpdateAt = now
//	_, err = engine.Get().Do("HMSET", redis.Args{}.Add(fmt.Sprintf("hm:%s:%s", funcs.BucketName(repo), repo.ID)).AddFlat(&repo)...)
//	if err != nil {
//		return
//	}
//	return
//}
