package skillissue

import (
	"encoding/json"

	"github.com/Tnze/go-mc/chat"
	"github.com/tidwall/gjson"
)

func (b *Bot) handleChatDeath(c chat.Message) error {
	target := getDeathTarget(c)
	b.Logger.Println("target:", target)
	if b.UncoolPeople.IsIgnored(target) || target == b.Client.Name {
		return nil
	}
	skillIssue := getSkillIssue(c)
	if skillIssue == "" {
		return nil
	}
	return b.sendSkillIssue(skillIssue)
}

func getDeathTarget(c chat.Message) string {
	var m chat.Message
	if err := json.Unmarshal(c.With[0], &m); err != nil {
		return ""
	}
	target := m.Text
	return target
}

func isKillerPlayer(c chat.Message) bool {
	if len(c.With) < 2 {
		return false
	}
	return gjson.ParseBytes(c.With[1]).Get("hoverEvent.contents.type").String() == "minecraft:player"
}

/*
  {
    "translate": "death.fell.accident.generic",
    "with": [
      {
        "insertion": "Honbra",
        "clickEvent": {
          "action": "suggest_command",
          "value": "/tell Honbra "
        },
        "hoverEvent": {
          "action": "show_entity",
          "contents": {
            "type": "minecraft:player",
            "id": "107c32ae-8838-4a99-b3d3-f8b6836992ba",
            "name": {
              "text": "Honbra"
            }
          }
        },
        "text": "Honbra"
      }
    ]
  }

  {
    "translate": "death.attack.player.item",
    "with": [
      {
        "insertion": "HonbraDev",
        "clickEvent": {
          "action": "suggest_command",
          "value": "/tell HonbraDev "
        },
        "hoverEvent": {
          "action": "show_entity",
          "contents": {
            "type": "minecraft:player",
            "id": "979b5829-2d98-4fb8-9d02-ac23113043d2",
            "name": {
              "text": "HonbraDev"
            }
          }
        },
        "text": "HonbraDev"
      },
      {
        "insertion": "Honbra",
        "clickEvent": {
          "action": "suggest_command",
          "value": "/tell Honbra "
        },
        "hoverEvent": {
          "action": "show_entity",
          "contents": {
            "type": "minecraft:player",
            "id": "107c32ae-8838-4a99-b3d3-f8b6836992ba",
            "name": {
              "text": "Honbra"
            }
          }
        },
        "text": "Honbra"
      },
      {
        "color": "aqua",
        "hoverEvent": {
          "action": "show_item",
          "contents": {
            "id": "minecraft:netherite_axe",
            "tag": "{CustomModelData:133702,Damage:3,Enchantments:[{id:\"minecraft:efficiency\",lvl:5s},{id:\"minecraft:mending\",lvl:1s},{id:\"minecraft:
  harpness\",lvl:5s},{id:\"minecraft:unbreaking\",lvl:3s}],RepairCost:15,display:{Name:\u0027{\"text\":\"Honbra\\\u0027s Axe\"}\u0027}}"
          }
        },
        "translate": "chat.square_brackets",
        "with": [
          {
            "italic": true,
            "extra": [
              {
                "text": "Honbra\u0027s Axe"
              }
            ],
            "text": ""
          }
        ]
      }
    ]
  }
  {
    "translate": "death.attack.player",
    "with": [
      {
        "insertion": "e4t_",
        "clickEvent": {
          "action": "suggest_command",
          "value": "/tell e4t_ "
        },
        "hoverEvent": {
          "action": "show_entity",
          "contents": {
            "type": "minecraft:player",
            "id": "c1f987d9-ad1d-410a-b45b-e037e3e0b672",
            "name": {
              "text": "e4t_"
            }
          }
        },
        "text": "e4t_"
      },
      {
        "insertion": "Honbra",
        "clickEvent": {
          "action": "suggest_command",
          "value": "/tell Honbra "
        },
        "hoverEvent": {
          "action": "show_entity",
          "contents": {
            "type": "minecraft:player",
            "id": "107c32ae-8838-4a99-b3d3-f8b6836992ba",
            "name": {
              "text": "Honbra"
            }
          }
        },
        "text": "Honbra"
      }
    ]
  }

*/
