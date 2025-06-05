package conversion

import (
	"github.com/jinzhu/copier"

	"github.com/setcreed/fastgo/internal/apiserver/model"
	apiv1 "github.com/setcreed/fastgo/pkg/api/apiserver/v1"
)

// UserModelToUserV1 将模型层的 User（用户模型对象）转换为 Protobuf 层的 User（v1 用户对象）.
func UserModelToUserV1(userModel *model.User) *apiv1.User {
	var protoUser apiv1.User
	_ = copier.Copy(&protoUser, userModel)
	return &protoUser
}

// UserV1ToUserModel 将 Protobuf 层的 User（v1 用户对象）转换为模型层的 User（用户模型对象）.
func UserV1ToUserModel(protoUser *apiv1.User) *model.User {
	var userModel model.User
	_ = copier.Copy(&userModel, protoUser)
	return &userModel
}
