package stix2

type collectionReference struct {
	ref *Collection
}

func (c *collectionReference) SetReference(collection *Collection) {
	c.ref = collection
}

type referenceSetter interface {
	SetReference(collection *Collection)
}
