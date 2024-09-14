package models

import "fmt"

type FarmModifier interface {
	ID() string
	Modify(*Farm)
	Description() string
}

var MAX_BARN_LEVEL = 10
var STORAGE_PER_LEVEL = 50

type BarnLevelMod struct {
	level int
}

func (b *BarnLevelMod) ID() string {
	return fmt.Sprintf("BARN_LEVEL_%v", b.level)
}

func (b *BarnLevelMod) Description() string {
	return fmt.Sprintf("Barn Level: %v. Increases max. storage by %v", b.level, b.level*STORAGE_PER_LEVEL)
}

func (b *BarnLevelMod) Modify(f *Farm) {
	f.MaxStorage += b.level * STORAGE_PER_LEVEL
}

func GetBarnLevelModifiers() []*BarnLevelMod {
	mods := []*BarnLevelMod{}
	for level := 0; level < MAX_BARN_LEVEL; level += 1 {
		mod := &BarnLevelMod{
			level: level,
		}
		mods = append(mods, mod)
	}

	return mods
}
