package memory

import "fmt"

type Key struct {
	FeatureId int
	TagId     int
}

func (k Key) String() string {
	return fmt.Sprintf("key banner tag=%v feature=%v", k.TagId, k.FeatureId)
}
