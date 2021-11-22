package npc

func getTemplate() map[string]string {
	return map[string]string{
		"class": "string",
		"object": "string",
		"position": "array", // [posWorld,vectorDir,vectorUp]
		"animations": "array",
		"action": "string",
		"name": "string",
		
		"loadout": "class",
		"mainWeapon": "string",
		"handWeapon": "string",
		"launcherWeapon": "string",
		"face": "string",
		"headgear": "string",
		"uniform": "string",
		"vest": "string",
		"backpack": "string",
	}
}

