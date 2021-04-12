package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Necroforger/dgrouter"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"

	"discord-bot/commands"
)

// Command line flags
var (
	fToken    = flag.String("t", "", "bot token")
	fPrefix   = flag.String("p", "!", "bot prefix")
	Inventory = make(map[string]string)
)

func main() {
	flag.Parse()

	s, err := discordgo.New("Bot " + *fToken)
	if err != nil {
		log.Fatal(err)
	}

	router := exrouter.New()

	// Add some commands
	router.On("add", func(ctx *exrouter.Context) {

		// Print statement for debugging
		fmt.Println(ctx.Msg.Mentions)
		fmt.Println(ctx.Msg.Content)

		// Parse values from discord message
		contents := strings.Fields(ctx.Msg.Content)
		players := ctx.Msg.Mentions
		item := contents[2]
		quantityStr := contents[3]
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			ctx.Reply("quantity cannot be converted into integer.")
		}

		// Print statement for debugging
		fmt.Println(players, item, quantityStr)

		// Add items to players inventory
		err = commands.Add(players, item, quantity)
		if err != nil {
			ctx.Reply(("cannot add items into player inventory"))
		}

		ctx.Reply("Items added.")
	}).Desc("Add item into player inventory. E.g !add @[discord_name] [item] [quantity]")

	router.On("get", func(ctx *exrouter.Context) {
		players := ctx.Msg.Mentions
		if len(players) < 1 {
			ctx.Reply("Please mention valid player username.")
		}
		for _, player := range players {
			inventory, err := commands.Get(player)
			if err != nil {
				panic("failed to retrieve inventory")
			}
			ctx.Reply(player.Username + "'s inventory: \n")
			ctx.Reply(inventory)

		}
	}).Desc("View player inventory")

	router.On("avatar", func(ctx *exrouter.Context) {
		ctx.Reply(ctx.Msg.Author.AvatarURL("1024"))
	}).Desc("returns the user's avatar")

	// Match the regular expression user(name)?
	router.OnMatch("username", dgrouter.NewRegexMatcher("user(name)?"), func(ctx *exrouter.Context) {
		ctx.Reply("Your username is " + ctx.Msg.Author.Username)
	})

	router.Default = router.On("help", func(ctx *exrouter.Context) {
		var text = ""
		for _, v := range router.Routes {
			text += v.Name + " : \t" + v.Description + "\n"
		}
		ctx.Reply("```" + text + "```")
	}).Desc("prints this help menu")

	// Add message handler
	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		router.FindAndExecute(s, *fPrefix, s.State.User.ID, m.Message)
	})

	err = s.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bot is running...")
	// Prevent the bot from exiting
	<-make(chan struct{})
}
