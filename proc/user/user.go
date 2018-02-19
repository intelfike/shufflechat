package user

import "errors"

type User struct {
	ID   int
	Name string
	Chan chan string
}

// var Users = []*User{}
var UsersMap = map[string]*User{} // ユーザー名をキーに

// 新しいユーザーを追加する
func Add(name string) error {
	if name == "" {
		return errors.New("ユーザー名が空です")
	}
	user := FindAtName(name)
	if user != nil {
		return errors.New(name + ":ユーザー名がすでに使われています")
	}
	user = &User{ID: len(UsersMap), Name: name, Chan: make(chan string)}
	// Users = append(Users, user)
	UsersMap[name] = user
	return nil
}

// ユーザーを削除する
func Delete(name string) error {
	user := FindAtName(name)
	if user != nil {
		delete(UsersMap, name)
		return nil
	}
	return errors.New(name + ":ユーザー名が見つかりません")
}

// 名前からデータを見つける
func FindAtName(name string) *User {
	user, ok := UsersMap[name]
	if ok {
		return user
	}
	return nil
}
