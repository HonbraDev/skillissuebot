package skillissue

import (
	"math/rand"
	"strings"

	"github.com/Tnze/go-mc/chat"
)

var skillIssuesGeneric = []string{
	"skill issue",
	"emotional damage!",
	"died of death",
	"TODO: git gud",
	"git clone skill --depth 1 --recursive",
	// "maybe something like buy a lootbox to get better", // ~ Philipp_DE1337
}

var skillIssuesFall = append(skillIssuesGeneric,
	"fix your hax",
	"didn't fix their hax",
)

func getSkillIssue(c chat.Message) string {
	deathType := strings.TrimPrefix(c.Translate, "death.")
	killerPlayer := isKillerPlayer(c)
	if (strings.HasPrefix(deathType, "fell") && !strings.Contains(deathType, "finish")) || deathType == "death.attack.fall" {
		return randomSkillIssue(skillIssuesFall)
	}
	if strings.HasPrefix(deathType, "attack") && !strings.HasSuffix(deathType, ".player") && !killerPlayer {
		return randomSkillIssue(skillIssuesGeneric)
	}
	return ""
}

func (b *Bot) sendSkillIssue(skillIssue string) error {
	return b.SendChatMsg("↑ " + skillIssue)
}

func randomSkillIssue(l []string) string {
	return l[rand.Intn(len(l))]
}
