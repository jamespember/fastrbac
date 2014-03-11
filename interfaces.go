package fastrbac

// GITHUB USER OR ORGANIZATION
type Owner interface {
	GetID() int64
	GetTypeName() string
}

// GITHUB REPOSITORY
type Object interface {
	GetID() int64
	GetTypeName() string
	GetOwner() int64
}

// Datasource for the trusts
type Repository interface {
	AddPermission(owner Owner, object Object, access string) (*Trust, error)
	GetTrust(owner Owner, object Object) (*Trust, error)
	GetTrustsByObjectType(owner Owner, objectType string)
}
