package authorization

import (
	"auth-service/pkg/token/handler/jwt"
	"sync"
)

const (
	PERMISSIONS = "permissions"
	ROLES       = "roles"
	GROUPS      = "groups"
)

type Authorization struct {
	RequiredPermissions []string
	RequiredRoles       []string
	RequiredGroups      []string
}

type AccessResult struct {
	AccessType string
	HaveAccess bool
}

func sliceToMap(s []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, k := range s {
		m[k] = struct{}{}
	}
	return m
}

func (a *Authorization) HaveAccess(userAccess *jwt.UserAccess) bool {

	permissionsMap := sliceToMap(userAccess.Permissions)
	rolesMap := sliceToMap(userAccess.Roles)
	groupsMap := sliceToMap(userAccess.Groups)

	if userAccess.Superuser {
		return true
	}

	resultChan := make(chan AccessResult, 3)
	defer close(resultChan)

	var wg sync.WaitGroup
	wg.Add(3)
	go a.checkAccess(resultChan, &wg, permissionsMap, a.RequiredPermissions, PERMISSIONS)
	go a.checkAccess(resultChan, &wg, rolesMap, a.RequiredRoles, ROLES)
	go a.checkAccess(resultChan, &wg, groupsMap, a.RequiredGroups, GROUPS)
	wg.Wait()

	for i := 0; i < 3; i += 1 {
		result := <-resultChan
		if !result.HaveAccess {
			return false
		}
	}

	return true
}

func (a *Authorization) checkAccess(
	resultChan chan AccessResult,
	wg *sync.WaitGroup,
	accessMap map[string]struct{},
	requirements []string,
	accessType string,
) {
	defer wg.Done()
	for _, r := range requirements {
		_, exist := accessMap[r]
		if !exist {
			resultChan <- AccessResult{AccessType: accessType, HaveAccess: false}
			return
		}
	}
	resultChan <- AccessResult{AccessType: accessType, HaveAccess: true}
}
