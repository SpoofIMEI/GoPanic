package keybind

import "golang.design/x/hotkey"

/*
If you want to change the panic hotkey, do it here :)

Options:
  hotkey.ModCtrl
  hotkey.ModShift
  hotkey.ModAlt
  hotkey.ModWin
  hotkey.Mod
*/
var HotKey = hotkey.New([]hotkey.Modifier{
	hotkey.ModCtrl,
	hotkey.ModShift,
},
	hotkey.KeyP /* <-- insert any key like A/B/7/L */)
