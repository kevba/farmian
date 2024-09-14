package models

var farmModifiers map[string]FarmModifier = farmToMap()

func farmToMap() map[string]FarmModifier {
	modMap := map[string]FarmModifier{}
	mods := GetBarnLevelModifiers()
	for _, m := range mods {
		modMap[m.ID()] = m
	}
	return modMap
}
