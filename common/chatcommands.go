package common

import "strings"

const CommandNameSeparator = ","

type ChatCommandNames []string

func (c ChatCommandNames) String() string {
	return strings.Join(c, CommandNameSeparator)
}

// Names for commands
var (
	// User Commands
	CNMe     ChatCommandNames = []string{"me"}
	CNHelp   ChatCommandNames = []string{"help"}
	CNCount  ChatCommandNames = []string{"count"}
	CNColor  ChatCommandNames = []string{"color", "colour"}
	CNWhoAmI ChatCommandNames = []string{"w", "whoami"}
	CNAuth   ChatCommandNames = []string{"auth"}
	CNUsers  ChatCommandNames = []string{"users"}
	CNNick   ChatCommandNames = []string{"nick", "name"}
	// Mod Commands
	CNSv      ChatCommandNames = []string{"sv"}
	CNPlaying ChatCommandNames = []string{"playing"}
	CNUnmod   ChatCommandNames = []string{"unmod"}
	CNKick    ChatCommandNames = []string{"kick"}
	CNBan     ChatCommandNames = []string{"ban"}
	CNUnban   ChatCommandNames = []string{"unban"}
	CNPurge   ChatCommandNames = []string{"purge"}
	CNPin     ChatCommandNames = []string{"pin", "password"}
	// Admin Commands
	CNMod          ChatCommandNames = []string{"mod"}
	CNReloadPlayer ChatCommandNames = []string{"reloadplayer"}
	CNReloadEmotes ChatCommandNames = []string{"reloademotes"}
	CNModpass      ChatCommandNames = []string{"modpass"}
	CNNewPin       ChatCommandNames = []string{"newpin", "newpassword"}
	CNRoomAccess   ChatCommandNames = []string{"changeaccess", "hodor"}
)

var ChatCommands = []ChatCommandNames{
	// User
	CNMe, CNHelp, CNCount, CNColor, CNWhoAmI, CNAuth, CNUsers, CNNick,
	// Mod
	CNSv, CNPlaying, CNUnmod, CNKick, CNBan, CNUnban, CNPurge, CNPin,
	// Admin
	CNMod, CNReloadPlayer, CNReloadEmotes, CNModpass, CNRoomAccess,
}

func GetFullChatCommand(c string) string {
	for _, names := range ChatCommands {
		for _, n := range names {
			if c == n {
				return names.String()
			}
		}
	}
	return ""
}
