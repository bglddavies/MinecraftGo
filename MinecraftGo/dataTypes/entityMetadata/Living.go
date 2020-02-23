package entityMetadata

import entity ".."

type LivingEntityMetadata struct {
	entity.EntityMetadata
	HandState          *byte
	Health             *float32
	PotionAffectColour *int
	IsPotionAmbient    *bool
	ArrowCount         *int
	AbsorptionHealth   *int
	BedLocation        *int64 //Optional Position
}

func (em *LivingEntityMetadata) Write() []byte {
	output := em.EntityMetadata.Write()
	//TODO
	return output
}
