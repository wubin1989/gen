// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/tests/.gen/dal_5/model"
	"github.com/wubin1989/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.User{}) fail: %s", err)
	}
}

func Test_userQuery(t *testing.T) {
	user := newUser(_gen_test_db)
	user = *user.As(user.TableName())
	_do := user.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(user.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <users> fail:", err)
		return
	}

	_, ok := user.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from user success")
	}

	err = _do.Create(&model.User{})
	if err != nil {
		t.Error("create item in table <users> fail:", err)
	}

	err = _do.Save(&model.User{})
	if err != nil {
		t.Error("create item in table <users> fail:", err)
	}

	err = _do.CreateInBatches([]*model.User{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <users> fail:", err)
	}

	_, err = _do.Select(user.ALL).Take()
	if err != nil {
		t.Error("Take() on table <users> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <users> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <users> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <users> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.User{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <users> fail:", err)
	}

	_, err = _do.Select(user.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <users> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <users> fail:", err)
	}

	_, err = _do.Select(user.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <users> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <users> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <users> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <users> fail:", err)
	}

	_, err = _do.ScanByPage(&model.User{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <users> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <users> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <users> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <users> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <users> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <users> fail:", err)
	}
}
