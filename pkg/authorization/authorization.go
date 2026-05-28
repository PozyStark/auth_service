package authorization

import (
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

type UserAccess struct {
	Permissions []string
	Roles       []string
	Groups      []string
}

type AccessResult struct {
	AccessType  string
	HaveAccess bool
}

func sliceToMap(s []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, k := range s {
		m[k] = struct{}{}
	}
	return m
}

func (a *Authorization) HaveAccess(userData *UserAccess) bool  {

	permissionsMap := sliceToMap(userData.Permissions)
	rolesMap := sliceToMap(userData.Roles)
	groupsMap := sliceToMap(userData.Groups)

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