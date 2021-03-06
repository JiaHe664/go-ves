
package control

import (
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"time"
)

var _ controller.MContext


type UserService interface {
    UserServiceSignatureXXX() interface{}
    ListUsers(c controller.MContext)
    Login(c controller.MContext)
    Register(c controller.MContext)
    ChangePassword(c controller.MContext)
    InspectUser(c controller.MContext)
    GetUser(c controller.MContext)
    PutUser(c controller.MContext)
    Delete(c controller.MContext)

}
type ListUsersRequest = gorm_crud_dao.Filter

type ListUsersReply struct {
    Code int `json:"code" form:"code"`
    Users []ListUserReply `json:"users" form:"users"`
}

type ListUserReply struct {
    Id uint `json:"id" form:"id"`
    LastLogin time.Time `json:"last_login" form:"last_login"`
}

type LoginRequest struct {
    Id uint `form:"id" json:"id"`
    Name string `form:"name" json:"name"`
    Password string `json:"password" form:"password" binding:"required"`
}

type LoginReply struct {
    Code int `form:"code" json:"code"`
    Id uint `json:"id" form:"id"`
    Name string `json:"name" form:"name"`
    Identity []string `form:"identity" json:"identity"`
    Token string `json:"token" form:"token"`
    RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

type RegisterRequest struct {
    Name string `binding:"required" json:"name" form:"name"`
    Password string `form:"password" binding:"required" json:"password"`
}

type RegisterReply struct {
    Code int `json:"code" form:"code"`
    Id uint `json:"id" form:"id"`
}

type ChangePasswordRequest struct {
    OldPassword string `json:"old_password" form:"old_password" binding:"required"`
    NewPassword string `form:"new_password" binding:"required" json:"new_password"`
}

type InspectUserReply struct {
    Code int           `json:"code" form:"code"`
    User *model.User `json:"user" form:"user"`
}

type GetUserReply struct {
    Code int `json:"code" form:"code"`
    Id uint `json:"id" form:"id"`
    LastLogin time.Time `json:"last_login" form:"last_login"`
}

type PutUserRequest struct {

}
func PSerializeListUsersReply(_code int, _users []ListUserReply) *ListUsersReply {

    return &ListUsersReply{
        Code: _code,
        Users: _users,
    }
}
func SerializeListUsersReply(_code int, _users []ListUserReply) ListUsersReply {

    return ListUsersReply{
        Code: _code,
        Users: _users,
    }
}
func _packSerializeListUsersReply(_code int, _users []ListUserReply) ListUsersReply {

    return ListUsersReply{
        Code: _code,
        Users: _users,
    }
}
func PackSerializeListUsersReply(_code []int, _users [][]ListUserReply) (pack []ListUsersReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListUsersReply(_code[i], _users[i]))
	}
	return
}
func PSerializeListUserReply(vUser model.User) *ListUserReply {

    return &ListUserReply{
        Id: vUser.ID,
        LastLogin: vUser.LastLogin,
    }
}
func SerializeListUserReply(vUser model.User) ListUserReply {

    return ListUserReply{
        Id: vUser.ID,
        LastLogin: vUser.LastLogin,
    }
}
func _packSerializeListUserReply(vUser model.User) ListUserReply {

    return ListUserReply{
        Id: vUser.ID,
        LastLogin: vUser.LastLogin,
    }
}
func PackSerializeListUserReply(vUser []model.User) (pack []ListUserReply) {
	for i := range vUser {
		pack = append(pack, _packSerializeListUserReply(vUser[i]))
	}
	return
}
func PSerializeLoginRequest(user *model.User, _password string) *LoginRequest {

    return &LoginRequest{
        Id: user.ID,
        Name: user.Name,
        Password: _password,
    }
}
func SerializeLoginRequest(user *model.User, _password string) LoginRequest {

    return LoginRequest{
        Id: user.ID,
        Name: user.Name,
        Password: _password,
    }
}
func _packSerializeLoginRequest(user *model.User, _password string) LoginRequest {

    return LoginRequest{
        Id: user.ID,
        Name: user.Name,
        Password: _password,
    }
}
func PackSerializeLoginRequest(user []*model.User, _password []string) (pack []LoginRequest) {
	for i := range user {
		pack = append(pack, _packSerializeLoginRequest(user[i], _password[i]))
	}
	return
}
func PSerializeLoginReply(_code int, user *model.User, _identity []string, _token string, _refreshToken string) *LoginReply {

    return &LoginReply{
        Code: _code,
        Id: user.ID,
        Name: user.Name,
        Identity: _identity,
        Token: _token,
        RefreshToken: _refreshToken,
    }
}
func SerializeLoginReply(_code int, user *model.User, _identity []string, _token string, _refreshToken string) LoginReply {

    return LoginReply{
        Code: _code,
        Id: user.ID,
        Name: user.Name,
        Identity: _identity,
        Token: _token,
        RefreshToken: _refreshToken,
    }
}
func _packSerializeLoginReply(_code int, user *model.User, _identity []string, _token string, _refreshToken string) LoginReply {

    return LoginReply{
        Code: _code,
        Id: user.ID,
        Name: user.Name,
        Identity: _identity,
        Token: _token,
        RefreshToken: _refreshToken,
    }
}
func PackSerializeLoginReply(_code []int, user []*model.User, _identity [][]string, _token []string, _refreshToken []string) (pack []LoginReply) {
	for i := range _code {
		pack = append(pack, _packSerializeLoginReply(_code[i], user[i], _identity[i], _token[i], _refreshToken[i]))
	}
	return
}
func PSerializeRegisterRequest(user *model.User, _password string) *RegisterRequest {

    return &RegisterRequest{
        Name: user.Name,
        Password: _password,
    }
}
func SerializeRegisterRequest(user *model.User, _password string) RegisterRequest {

    return RegisterRequest{
        Name: user.Name,
        Password: _password,
    }
}
func _packSerializeRegisterRequest(user *model.User, _password string) RegisterRequest {

    return RegisterRequest{
        Name: user.Name,
        Password: _password,
    }
}
func PackSerializeRegisterRequest(user []*model.User, _password []string) (pack []RegisterRequest) {
	for i := range user {
		pack = append(pack, _packSerializeRegisterRequest(user[i], _password[i]))
	}
	return
}
func PSerializeRegisterReply(_code int, user *model.User) *RegisterReply {

    return &RegisterReply{
        Code: _code,
        Id: user.ID,
    }
}
func SerializeRegisterReply(_code int, user *model.User) RegisterReply {

    return RegisterReply{
        Code: _code,
        Id: user.ID,
    }
}
func _packSerializeRegisterReply(_code int, user *model.User) RegisterReply {

    return RegisterReply{
        Code: _code,
        Id: user.ID,
    }
}
func PackSerializeRegisterReply(_code []int, user []*model.User) (pack []RegisterReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRegisterReply(_code[i], user[i]))
	}
	return
}
func PSerializeChangePasswordRequest(_oldPassword string, _newPassword string) *ChangePasswordRequest {

    return &ChangePasswordRequest{
        OldPassword: _oldPassword,
        NewPassword: _newPassword,
    }
}
func SerializeChangePasswordRequest(_oldPassword string, _newPassword string) ChangePasswordRequest {

    return ChangePasswordRequest{
        OldPassword: _oldPassword,
        NewPassword: _newPassword,
    }
}
func _packSerializeChangePasswordRequest(_oldPassword string, _newPassword string) ChangePasswordRequest {

    return ChangePasswordRequest{
        OldPassword: _oldPassword,
        NewPassword: _newPassword,
    }
}
func PackSerializeChangePasswordRequest(_oldPassword []string, _newPassword []string) (pack []ChangePasswordRequest) {
	for i := range _oldPassword {
		pack = append(pack, _packSerializeChangePasswordRequest(_oldPassword[i], _newPassword[i]))
	}
	return
}
func PSerializeInspectUserReply(_code int, _user *model.User) *InspectUserReply {

    return &InspectUserReply{
        Code: _code,
        User: _user,
    }
}
func SerializeInspectUserReply(_code int, _user *model.User) InspectUserReply {

    return InspectUserReply{
        Code: _code,
        User: _user,
    }
}
func _packSerializeInspectUserReply(_code int, _user *model.User) InspectUserReply {

    return InspectUserReply{
        Code: _code,
        User: _user,
    }
}
func PackSerializeInspectUserReply(_code []int, _user []*model.User) (pack []InspectUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectUserReply(_code[i], _user[i]))
	}
	return
}
func PSerializeGetUserReply(_code int, user *model.User) *GetUserReply {

    return &GetUserReply{
        Code: _code,
        Id: user.ID,
        LastLogin: user.LastLogin,
    }
}
func SerializeGetUserReply(_code int, user *model.User) GetUserReply {

    return GetUserReply{
        Code: _code,
        Id: user.ID,
        LastLogin: user.LastLogin,
    }
}
func _packSerializeGetUserReply(_code int, user *model.User) GetUserReply {

    return GetUserReply{
        Code: _code,
        Id: user.ID,
        LastLogin: user.LastLogin,
    }
}
func PackSerializeGetUserReply(_code []int, user []*model.User) (pack []GetUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetUserReply(_code[i], user[i]))
	}
	return
}
func PSerializePutUserRequest() *PutUserRequest {

    return &PutUserRequest{

    }
}
func SerializePutUserRequest() PutUserRequest {

    return PutUserRequest{

    }
}
func _packSerializePutUserRequest() PutUserRequest {

    return PutUserRequest{

    }
}
func PackSerializePutUserRequest() (pack []PutUserRequest) {
	return
}
