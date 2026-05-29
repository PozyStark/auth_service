package app

import (
	"auth-service/internal/services"
	"auth-service/pkg/token/handler/jwt"
	"context"
	"errors"
	"sync"
)

type Error struct {
	Operation string
	Err       error
}

func GetUserInfo(
	ctx context.Context,
	userPermissionService services.UserPermissionService,
	userRoleService services.UserRoleService,
	userGroupService services.UserGroupService,
	userId string,
) (jwt.UserAccess, []error) {

	var userAccess jwt.UserAccess
	errorChanel := make(chan Error, 3)
	defer close(errorChanel)

	var wg sync.WaitGroup
	wg.Add(3)
	func(ctx context.Context, wg *sync.WaitGroup, errorChanel chan Error) {
		defer wg.Done()
		result, err := userPermissionService.GetAllUserPermissionsByUserIdAsStringSlice(ctx, userId)
		userAccess.Permissions = result
		errorChanel <- Error{Operation: "GetAllUserPermissionsByUserIdAsStringSlice", Err: err}
	}(ctx, &wg, errorChanel)

	func(ctx context.Context, wg *sync.WaitGroup, errorChanel chan Error) {
		defer wg.Done()
		result, err := userRoleService.GetUserRolesByUserIdAsStringSlice(ctx, userId)
		userAccess.Roles = result
		errorChanel <- Error{Operation: "GetUserRolesByUserIdAsStringSlice", Err: err}
	}(ctx, &wg, errorChanel)

	func(ctx context.Context, wg *sync.WaitGroup, errorChanel chan Error) {
		defer wg.Done()
		result, err := userGroupService.GetUserGroupsByUserIdAsStringSlice(ctx, userId)
		userAccess.Groups = result
		errorChanel <- Error{Operation: "GetUserGroupsByUserIdAsStringSlice", Err: err}
	}(ctx, &wg, errorChanel)
	wg.Wait()

	var resultError []error
	for i := 0; i < 3; i++ {
		err := <-errorChanel
		if err.Err != nil {
			detailedError := errors.New(err.Operation + ": " + err.Err.Error())
			resultError = append(resultError, detailedError)
		}
	}

	return userAccess, resultError

}
