package models

import "time"

type ModCollection struct {
	Game                 string        `json:"game"`
	Description          string        `json:"description"`
	Restorebeforeinstall bool          `json:"restorebeforeinstall"`
	Mods                 []Mod         `json:"mods"`
	Asimods              []Asimod      `json:"asimods"`
	Texturemodfiles      []interface{} `json:"texturemodfiles"`
	Queuename            string        `json:"queuename"`
	ImporterDescription  string        `json:"ImporterDescription"`
}

type Mod struct {
	Moddescpath       string             `json:"moddescpath"`
	Configurationtime time.Time          `json:"configurationtime"`
	Moddeschash       string             `json:"moddeschash"`
	Moddescsize       int                `json:"moddescsize"`
	Downloadlink      string             `json:"downloadlink"`
	Userchosenoptions []UserChosenOption `json:"userchosenoptions"`
	Allchosenoptions  []string           `json:"allchosenoptions"`
	Haschosenoptions  bool               `json:"haschosenoptions"`
	IsStandalone      bool               `json:"IsStandalone"`
	Modname           string             `json:"modname"`
}

type UserChosenOption struct {
	IsPlus        bool   `json:"IsPlus"`
	Key           string `json:"Key"`
	OriginalValue string `json:"OriginalValue"`
}

type Asimod struct {
	Updategroup int `json:"updategroup"`
	Version     int `json:"version"`
}
