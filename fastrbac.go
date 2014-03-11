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

func HasPermission(repository Repository, owner Owner, object Object, access string) bool {
	// Owner always has full permission to the object
	if object.GetOwner() == owner.GetID() {
		return true
	}

	trust, _ := repository.GetTrust(owner, object)

	if trust != nil {
		for _, permission := range trust.Permissions {
			if permission == access {
				return true
			}
		}
	}

	return false
}

func GrantPermission(repository Repository, owner Owner, object Object, access string) error {
	if HasPermission(repository, owner, object, access) {
		return nil
	}

	_, err := repository.AddPermission(owner, object, access)

	return err
}

func GetTrustsByObjectType(repository *Repository, owner Owner, objectType string) {
	return
}
