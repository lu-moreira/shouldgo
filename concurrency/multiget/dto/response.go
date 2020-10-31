package dto

import "github.com/lu-moreira/shouldgo/concurrency/multiget/model"

// GetUserResponse defines the expected response of GET /user/:id
type GetUserResponse model.User

// GetUserAssignedPermissionsResponse defines the expected response of GET user/:id/assigned-permissions
type GetUserAssignedPermissionsResponse []model.UserAssignedPermission

// GetUserAtrributesResponse defines the expected response of GET /user/:id/attributes
type GetUserAtrributesResponse []model.Attribute
