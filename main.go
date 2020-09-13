package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	keymap := BuildEventToHidMap()

	r := bufio.NewReader(os.Stdin)
	for {
		event := InputEvent{}
		if err := binary.Read(r, binary.LittleEndian, &event); err != nil {
			os.Stderr.WriteString(err.Error())
			break
		}

		if _, ok := keymap[event.Code]; !ok {
			os.Stderr.WriteString(fmt.Sprintf("\tMissing event code from keymap: %d\n", event.Code))
			continue
		}

		if event.Type != 1 {
			continue // ignore
		}

		keymap[event.Code].Active = event.Value == 1 || event.Value == 2

		hidkey := keymap[event.Code]
		os.Stderr.WriteString(fmt.Sprintf("\tevent_code %d\tevent_value %d\thidkey_name %s\thidkey_code %d\n", event.Code, event.Value, hidkey.Name, hidkey.Code))

		keys := make([]byte, 6)
		mods := make([]byte, 8)
		keyIndex := 0
		modIndex := 0
		for _, hidkey := range keymap {
			if hidkey.IsNormalKey() {
				if hidkey.Active && keyIndex < len(keys) {
					keys[keyIndex] = hidkey.Code
					keyIndex++
				}
			} else {
				if hidkey.Active && modIndex < len(mods) {
					mods[modIndex] = hidkey.Mod
					modIndex++
				}
			}
		}

		var modByte byte
		for _, mod := range mods {
			modByte |= mod
		}

		report := []byte{modByte, 0, keys[0], keys[1], keys[2], keys[3], keys[4], keys[5]}
		if _, err := os.Stdout.Write(report); err != nil {
			panic(err)
		}
	}
}
