package models

import "github.com/jinzhu/gorm"

type Repo struct {
	gorm.Model
	Address     string // 仓库地址
	Name        string // 本地存储所用的名字
	Schedule    uint   // 仓库两次更新之间的时间间隔
	Status      string // 仓库状态
	Progress    int    // 进度
	CommitCount uint   // commit 数量
	IsSendEmail bool   // 是否发送邮件
}

func (repo Repo) IsExist() (err error) {
	var num int
	if err := Major.Where(Repo{Address: repo.Address}).Count(&num).Error; err != nil {
		return err
	}
	return nil
}

func (repo Repo) Add() error {
	if err := Major.Create(&repo).Error; err != nil {
		return err
	}
	return nil
}

func (repo Repo) Remove() error {
	if err := Major.Delete(new(Repo), repo).Error; err != nil {
		return err
	}
	return nil
}

func (repo Repo) Update(newRepo Repo) error {
	if err := Major.Where(repo).Updates(newRepo).Error; err != nil {
		return err
	}
	return nil
}

func (repo Repo) StopAble() (bool, error) {
	if err := Major.First(new(Repo), "(status = ? OR status = ?) AND id = ?", "Connecting", "Receiving", repo.ID).Error; err != nil {
		return false, err
	}
	return true, nil
}
