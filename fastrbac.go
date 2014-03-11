package fastrbac

// A trust is a list of permissions for a holder on an object that it isn't owner for.
type Trust struct {
	HolderId    int64
	HolderType  string
	ObjectId    int64
	ObjectType  string
	Permissions []string
}

// A role defines what permissions that a holder has on objects that are owned by another holder
type Role struct {
	Title        string
	HolderIds    []int64
	RoleTargetId int64
}

func HasPermission(repository Repository, owner Owner, entity Entity, access string) bool {
	// Owner always has full permission to the object
	if entity.GetOwner() == owner.GetID() {
		return true
	}

	trust, _ := repository.GetTrust(owner, entity)

	if trust != nil {
		for _, permission := range trust.Permissions {
			if permission == access {
				return true
			}
		}
	}

	return false
}

func GrantPermission(repository Repository, owner Owner, entity Entity, access string) error {
	if HasPermission(repository, owner, entity, access) {
		return nil
	}

	_, err := repository.AddPermission(owner, entity, access)

	return err
}

func GetTrustsByObjectType(repository *Repository, owner Owner, objectType string) {
	return
}
