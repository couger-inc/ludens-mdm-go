package userStore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/couger-inc/ludens-mdm/crud/db"
)

type TableBasics struct{
	PrismaClient *db.PrismaClient
}

func CreateClient() (*TableBasics, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	return &TableBasics{client}, nil
}


type Manager struct {
	Name string `json:"name" validate:"required,max=64"`
	Email string `json:"email" validate:"required,email"`
}

func (basics *TableBasics) GetStores(ctx context.Context, offset int, limit int, storeId string, storeName string, managerEmail string, managerName string) ([]db.StoreModel, int, error) {
	var res []map[string]string
	err := basics.PrismaClient.Prisma.QueryRaw(fmt.Sprintf("SELECT COUNT(*) FROM `UserStore` where email LIKE '%s%%' AND name LIKE '%s%%' AND `storeId` IN (SELECT id FROM `Store` where id LIKE '%s%%' AND name LIKE '%s%%')", managerEmail, managerName, storeId, storeName)).Exec(ctx, &res)
	if (err != nil) {
		return nil, 0, err
	}
	totalCount, err := strconv.Atoi(res[0]["COUNT(*)"])
	if (err != nil) {
		return nil, 0, err
	}
	stores, err := basics.PrismaClient.Store.FindMany(
		db.Store.ID.StartsWith(storeId),
		db.Store.Name.StartsWith(storeName),
	).With(db.Store.UserStore.Fetch(
		db.UserStore.Email.StartsWith(managerEmail),
		db.UserStore.Name.StartsWith(managerName),
	)).Skip(offset).Take(limit).Exec(ctx)
	return stores, totalCount, err
}

func (basics *TableBasics) GetUserStores(ctx context.Context, offset int, limit int, storeId string) ([]db.StoreModel, int, error) {
	var res []map[string]string
	err := basics.PrismaClient.Prisma.QueryRaw(fmt.Sprintf("SELECT COUNT(*) FROM `UserStore` where `storeId` IN (SELECT id FROM `Store` where id='%s')", storeId)).Exec(ctx, &res)
	if (err != nil) {
		return nil, 0, err
	}
	totalCount, err := strconv.Atoi(res[0]["COUNT(*)"])
	if (err != nil) {
		return nil, 0, err
	}
	stores, err := basics.PrismaClient.Store.FindMany(
		db.Store.ID.StartsWith(storeId),
	).With(db.Store.UserStore.Fetch()).Skip(offset).Take(limit).Exec(ctx)
	return stores, totalCount, err
}

func (basics *TableBasics) AddUserStore(ctx context.Context, storeId string, manager Manager) (*db.UserStoreModel, error) {
	_, err := basics.PrismaClient.Store.FindUnique(db.Store.ID.Equals(storeId)).Exec(ctx)
	if err != nil && err.Error() == "ErrNotFound" {
		return nil, fmt.Errorf("store %v not found", storeId)
	}
	_, err = basics.PrismaClient.UserStore.FindUnique(db.UserStore.EmailStoreID(db.UserStore.Email.Equals(manager.Email), db.UserStore.StoreID.Equals(storeId))).Exec(ctx)
	if err == nil || err.Error() != "ErrNotFound" {
		return nil, fmt.Errorf("manager %v:%v already exists", manager.Email, manager.Name)
	}
	createdManager, err :=basics.PrismaClient.UserStore.CreateOne(
		db.UserStore.Email.Set(manager.Email),
		db.UserStore.Name.Set(manager.Name),
		db.UserStore.Store.Link(db.Store.ID.Equals(storeId)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return createdManager, nil
}

func (basics *TableBasics) DeleteUserStore(ctx context.Context, storeId string, managerEmail string) (*db.UserStoreModel, error) {
	_, err := basics.PrismaClient.Store.FindUnique(db.Store.ID.Equals(storeId)).Exec(ctx)
	if err != nil && err.Error() == "ErrNotFound" {
		return nil, fmt.Errorf("store not found: %v", storeId)
	}
	_, err = basics.PrismaClient.UserStore.FindUnique(db.UserStore.EmailStoreID(db.UserStore.Email.Equals(managerEmail), db.UserStore.StoreID.Equals(storeId))).Exec(ctx)
	if err != nil && err.Error() == "ErrNotFound" {
		return nil, fmt.Errorf("manager, %v, not found for store: %v", managerEmail, storeId)
	}
	deletedManager, err := basics.PrismaClient.UserStore.FindUnique(db.UserStore.EmailStoreID(db.UserStore.Email.Equals(managerEmail), db.UserStore.StoreID.Equals(storeId))).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return deletedManager, nil
}

func (basics *TableBasics) Disconnect() {
	if err := basics.PrismaClient.Disconnect(); err != nil {
		panic(err)
	}
}
