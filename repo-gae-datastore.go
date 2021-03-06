package fastrbac

import (
	"appengine"
	"appengine/datastore"
	"fmt"
)

// TODO(robin): Implement memcache layer

const NAMESPACE_TRUST = "fastrbac.trust"

type datastoreRepository struct {
	Context appengine.Context
}

func NewDatastoreRepo(c appengine.Context) Repository {
	return &datastoreRepository{
		Context: c,
	}
}

// Generates the key for a trust
func (this *datastoreRepository) createTrustKey(owner Owner, object Object) *datastore.Key {
	keyString := fmt.Sprintf("%v_%v_%v_%v", owner.GetTypeName(), owner.GetID(), object.GetTypeName(), object.GetID())
	key := datastore.NewKey(this.Context, NAMESPACE_TRUST, keyString, 0, nil)
	return key
}

// Fetches a trust, returns nil as trust if none exist
func (this *datastoreRepository) GetTrust(owner Owner, object Object) (*Trust, error) {
	trust := new(Trust)
	if err := datastore.Get(this.Context, this.createTrustKey(owner, object), trust); err == datastore.ErrNoSuchEntity {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return trust, nil
}

// Add permission, append to existing trust or create a new one if one does not exist.
func (this *datastoreRepository) AddPermission(owner Owner, object Object, access string) (*Trust, error) {
	trust, _ := this.GetTrust(owner, object)

	if trust == nil {
		// Create new trust
		trust = &Trust{
			HolderId:    owner.GetID(),
			HolderType:  owner.GetTypeName(),
			ObjectId:    object.GetID(),
			ObjectType:  object.GetTypeName(),
			Permissions: []string{access},
		}
	} else {
		// Append permission to existing trust
		trust.Permissions = append(trust.Permissions, access)
	}

	if _, err := datastore.Put(this.Context, this.createTrustKey(owner, object), trust); err != nil { // Allocate key
		return nil, err
	}

	return trust, nil
}

// Get all trusts for a specific object type belonging to a specific owner
func (this *datastoreRepository) GetTrustsByObjectType(owner Owner, objectType string) {

	return
}
