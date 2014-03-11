package fastrbac

// GITHUB USER OR ORGANIZATION
type Owner interface {
	GetID() int64
	GetTypeName() string
}

// GITHUB REPOSITORY
type Entity interface {
	GetID() int64
	GetTypeName() string
	GetOwner() int64
}

// Datasource for the trusts
type Repository interface {
	AddPermission(owner Owner, entity Entity, access string) (*Trust, error)
	GetTrust(owner Owner, entity Entity) (*Trust, error)
	GetTrustsByObjectType(owner Owner, objectType string)
}
