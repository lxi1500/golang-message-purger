package main
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os/signal"
	"syscall"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/lxi1500/gotitle"
	"strconv"


)



func banner() {
	fmt.Print("\033[H\033[2J")
	color.Red(` 
				██████╗ ██╗   ██╗██████╗  ██████╗ ███████╗██████╗ 
				██╔══██╗██║   ██║██╔══██╗██╔════╝ ██╔════╝██╔══██╗
				██████╔╝██║   ██║██████╔╝██║  ███╗█████╗  ██████╔╝
				██╔═══╝ ██║   ██║██╔══██╗██║   ██║██╔══╝  ██╔══██╗
				██║     ╚██████╔╝██║  ██║╚██████╔╝███████╗██║  ██║
				╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝
						[Made by Ixi#1400]														  
	`)

}

func main() {
	title.SetTitle("Purger | Login")
	banner() 
	fmt.Print("[!] Insert your token: ")
	fmt.Scan(&token)
	banner()
	selfbot, err := discordgo.New(token)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	
	selfbot.AddHandler(messageCreate)
	selfbot.AddHandler(func(dg *discordgo.Session, event *discordgo.Ready) {
		username, _ := dg.User("@me")
		title.SetTitle(fmt.Sprintf("Purger | Logged in as %s", username))
		color.Green(fmt.Sprintf("[!] Purger has been connected to %s!\n[?] Type .purge [amount] to purge.", username))
	})
	err = selfbot.Open()
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	selfbot.Close()
}



func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != s.State.User.ID {
		return
	}
	
	if strings.HasPrefix(m.Content, ".purge") {
		amount, err := strconv.Atoi(strings.Trim(m.Content, ".purge "))
		if err != nil {
			fmt.Printf("[!] Error: %s\n", err)
		}		
		deled := 0
		dID := "" // needed for fetching ? kinda odd but whatever
		for deled < amount {
			messages, err := s.ChannelMessages(m.ChannelID, 100, dID, "", "")
			if err != nil {
				fmt.Printf("[!] Error: %s\n", err)
			}
		
			for _, message := range messages {
				if deled >= amount {
					break
				}
		
				if message.Author.ID != s.State.User.ID {
					continue;
				}
		
				err = s.ChannelMessageDelete(m.ChannelID, message.ID)
				color.Yellow(fmt.Sprintf("[!] Deleted message: %v\n[?] Content: %s\n", message.ID, message.Content))
				if err != nil {
					fmt.Printf("[!] Error: %v", err)
					continue
				}
				deled += 1
				dID = message.ID
			}
		}
			color.Green(fmt.Sprintf("[!] Deleted %v messages!\n", deled))
	}
}



var (
	token string
)
