package tests

import entityID "medicine/pkg/entity-id"

func GenerateEntityID() entityID.EntityID {
	id, err := entityID.NewGenerator().Generate()
	if err != nil {
		panic(err)
	}

	return id
}
