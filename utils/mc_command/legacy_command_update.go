package mc_command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OmineDev/neomega-core/neomega/blocks"
	"github.com/OmineDev/neomega-core/utils/mc_command/token"
)

func UpdateBlockDescribe(blockName, blockValueString string) (string, error) {
	blockValueString = strings.TrimSpace(blockValueString)
	blockValue, err := strconv.Atoi(blockValueString)
	if err != nil {
		return fmt.Sprintf("[ERROR BLOCK: %v %v]", blockName, blockValueString), fmt.Errorf("%v %v is not a legal legacy block description", blockName, blockValueString)
	}
	if blockValue == -1 {
		return fmt.Sprintf("%v []", blockName), nil
	}
	rtid, found := blocks.LegacyBlockToRuntimeID(blockName, uint16(blockValue))
	if !found {
		return fmt.Sprintf("[ERROR BLOCK: %v %v]", blockName, blockValueString), fmt.Errorf("cannot find new block represent for block %v %v", blockName, blockValueString)
	}
	newName, newState, found := blocks.RuntimeIDToBlockNameAndStateStr(rtid)
	if !found {
		return fmt.Sprintf("[ERROR BLOCK: %v %v]", blockName, blockValueString), fmt.Errorf("cannot find new block represent for block %v %v", blockName, blockValueString)
	}
	if newState == "" {
		newState = "[]"
	}
	newState = strings.ReplaceAll(strings.ReplaceAll(newState, " ", ""), ":", "=")
	return newName + " " + newState, nil
}

func UpdateLegacyExecuteCommand(command string) string {
	origCommand := command
	c := ParseLegacyMCExecuteCommand(command)
	if c == nil {
		return command
	}
	newCommand := fmt.Sprintf("execute as %v at @s positioned %v", c.Selector, c.Pos)
	if c.DetectPosIfAny != "" {
		if strings.ContainsAny(c.DetectBlockValueIfAny, "[]") {
			return origCommand
		}
		updateBlock, err := UpdateBlockDescribe(c.DetectBlockNameIfAny, c.DetectBlockValueIfAny)
		if err != nil {
			fmt.Println(err)
			return command
		}
		newCommand += fmt.Sprintf(" if block %v %v", c.DetectPosIfAny, updateBlock)
	}
	newCommand += fmt.Sprintf(" run %v", UpdateLegacyCommand(c.SubCommand))
	return newCommand
}

func UpdateLegacySetBlockCommand(command string) string {
	if strings.ContainsAny(command, "[]") {
		return command
	}
	c := ParseLegacySetBlockCommand(command)
	if c == nil {
		return command
	}
	newCommand := fmt.Sprintf("setblock %v ", c.Pos)
	if c.BlockValueIfAny == "" {
		c.BlockValueIfAny = "0"
	}
	if strings.ContainsAny(c.BlockName, "[]") {
		return command
	}
	blockStr, err := UpdateBlockDescribe(c.BlockName, c.BlockValueIfAny)
	if err != nil {
		fmt.Println(err)
		return command
	}
	if !strings.HasPrefix(c.OtherOptions, " ") {
		c.OtherOptions = " " + c.OtherOptions
	}
	return newCommand + blockStr + c.OtherOptions
}

func UpdateLegacyFillCommand(command string) string {
	if strings.ContainsAny(command, "[]") {
		return command
	}
	c := ParseLegacyFillCommand(command)
	if c == nil {
		return command
	}
	newCommand := fmt.Sprintf("fill %v %v ", c.StartPos, c.EndPos)
	if c.BlockValueIfAny == "" {
		c.BlockValueIfAny = "0"
	}
	if strings.ContainsAny(c.BlockName, "[]") {
		return command
	}
	blockStr, err := UpdateBlockDescribe(c.BlockName, c.BlockValueIfAny)
	if err != nil {
		fmt.Println(err)
		return command
	}
	if c.ReplaceBlockNameIfAny == "" {
		if !strings.HasPrefix(c.OtherOptions, " ") {
			c.OtherOptions = " " + c.OtherOptions
		}
		return newCommand + blockStr + c.OtherOptions
	} else {
		if c.ReplaceBlockValueIfAny == "" {
			c.ReplaceBlockValueIfAny = "-1"
		}
		if strings.ContainsAny(c.ReplaceBlockNameIfAny, "[]") {
			return command
		}
		replaceBlockStr, err := UpdateBlockDescribe(c.ReplaceBlockNameIfAny, c.ReplaceBlockValueIfAny)
		if err != nil {
			fmt.Println(err)
			return command
		}
		return newCommand + blockStr + " replace " + replaceBlockStr
	}
}

func UpdateLegacyCloneCommand(command string) string {
	if strings.ContainsAny(command, "[]") {
		return command
	}
	c := ParseLegacyCloneCommand(command)
	// fmt.Println(c)
	if c == nil {
		return command
	}
	newCommand := fmt.Sprintf("clone %v %v %v", c.StartPos, c.EndPos, c.TargetPos)
	if !c.IsFiltered {
		if !strings.HasPrefix(c.OtherOptions, " ") {
			c.OtherOptions = " " + c.OtherOptions
		}
		return newCommand + c.OtherOptions
	}

	if c.BlockValueIfFiltered == "" {
		c.BlockValueIfFiltered = "-1"
	}
	if strings.ContainsAny(c.BlockNameIfFiltered, "[]") {
		return command
	}
	blockStr, err := UpdateBlockDescribe(c.BlockNameIfFiltered, c.BlockValueIfFiltered)
	if err != nil {
		fmt.Println(err)
		return command
	}
	return newCommand + " filtered " + c.ModeIfFiltered + " " + blockStr
}

func UpdateLegacyTestForBlockCommand(command string) string {
	if strings.ContainsAny(command, "[]") {
		return command
	}
	c := ParseLegacyTestForBlockCommand(command)
	// fmt.Println(c)
	if c == nil {
		return command
	}
	newCommand := fmt.Sprintf("testforblock %v ", c.Pos)
	if c.BlockValueIfAny == "" {
		c.BlockValueIfAny = "-1"
	}
	if strings.ContainsAny(c.BlockName, "[]") {
		return command
	}
	blockStr, err := UpdateBlockDescribe(c.BlockName, c.BlockValueIfAny)
	if err != nil {
		fmt.Println(err)
		return command
	}
	return newCommand + " " + blockStr
}

func UpdateLegacySummonCommand(command string) string {
	c := ParseLegacySummonCommand(command)
	// fmt.Println(c)
	if c == nil {
		return command
	}
	newCommand := fmt.Sprintf("summon %v %v facing %v %v %v", c.EntityType, c.Pos, c.FacingPos, c.Event, c.EntityName)
	return newCommand
}

func UpdateLegacyTitleCommand(command string) string {
	convetString := func(in string) string {
		out := strings.ReplaceAll(in, "\\", "\\\\")
		out = strings.ReplaceAll(in, "\n", "\\n")
		return out
	}
	if !strings.Contains(command, "\n") {
		return command
	}
	origCommand := command
	command = strings.TrimSpace(origCommand)
	reader := CleanStringAndNewSimpleTextReader(command)
	token.ReadSpecific(reader, "/", true)
	token.ReadWhiteSpace(reader)
	token.ReadSpecific(reader, "title", true)
	newCommand := "titleraw "
	token.ReadWhiteSpace(reader)
	ok, t := token.ReadMCSelector(reader)
	if !ok {
		return origCommand
	}
	newCommand += t
	token.ReadWhiteSpace(reader)
	ok, t = token.ReadNonWhiteSpace(reader)
	if !ok {
		return origCommand
	}
	newCommand += " " + t
	token.ReadWhiteSpace(reader)
	_, t = token.ReadUntilEnd(reader)
	reader = CleanStringAndNewSimpleTextReader(t)
	newCommand += ` {"rawtext": [`
	first := true
	res := ""
	for reader.Current() != "" {
		_, t := token.ReadAnyExcept(reader, "@")
		if t != "" {
			// fmt.Println("non-selector ", t, " END")
			res += convetString(t)
		} else {
			// possible a selector
			ok, t := token.ReadMCSelector(reader)
			if ok {
				// is selector
				if res != "" {
					if first {
						first = false
					} else {
						newCommand += ","
					}
					newCommand += `{"text":"` + res + `"}`
					res = ""
				}

				if first {
					first = false
				} else {
					newCommand += ","
				}
				newCommand += `{"selector":"` + convetString(t) + `"}`
			} else {
				// something else
				token.ReadSpecific(reader, "@", false)
				res += "@"
			}
		}
	}
	if res != "" {
		if first {
			first = false
		} else {
			newCommand += ","
		}
		newCommand += `{"text":"` + res + `"}`
		res = ""
	}
	return newCommand + "]}"
}

func UpdateLegacyCommand(command string) string {
	lCommand := strings.ToLower(command)
	lCommand = strings.TrimPrefix(lCommand, "/")
	if strings.HasPrefix(lCommand, "execute") {
		return UpdateLegacyExecuteCommand(command)
	} else if strings.HasPrefix(lCommand, "setblock") {
		return UpdateLegacySetBlockCommand(command)
	} else if strings.HasPrefix(lCommand, "fill") {
		return UpdateLegacyFillCommand(command)
	} else if strings.HasPrefix(lCommand, "clone") {
		return UpdateLegacyCloneCommand(command)
	} else if strings.HasPrefix(lCommand, "testforblock") {
		return UpdateLegacyTestForBlockCommand(command)
	} else if strings.HasPrefix(lCommand, "summon") {
		return UpdateLegacySummonCommand(command)
	} else if strings.HasPrefix(lCommand, "title") && !strings.HasPrefix(lCommand, "titleraw") {
		return UpdateLegacyTitleCommand(command)
	}
	return command
}

func IsUpdatableLegacyCommand(command string) string {
	lCommand := strings.ToLower(command)
	lCommand = strings.TrimPrefix(lCommand, "/")
	if strings.HasPrefix(lCommand, "execute") {
		return "execute"
	} else if strings.HasPrefix(lCommand, "setblock") {
		return "setblock"
	} else if strings.HasPrefix(lCommand, "fill") {
		return "fill"
	} else if strings.HasPrefix(lCommand, "clone") {
		return "clone"
	} else if strings.HasPrefix(lCommand, "testforblock") {
		return "testforblock"
	} else if strings.HasPrefix(lCommand, "summon") {
		return "summon"
	} else if strings.HasPrefix(lCommand, "title") && !strings.HasPrefix(lCommand, "titleraw") {
		return "title"
	}
	return ""
}
