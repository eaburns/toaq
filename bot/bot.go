package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/eaburns/pretty"
	"github.com/eaburns/toaq/ast"
	"github.com/eaburns/toaq/logic"
	"github.com/velour/chat"
	"github.com/velour/chat/discord"
)

const (
	parsePrefix = ".miu"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("DISCORD_TOKEN")
	server := os.Getenv("DISCORD_SERVER")
	chans := os.Getenv("DISCORD_CHANS")

	cl, err := discord.Dial(ctx, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, c := range strings.Split(chans, ",") {
		fmt.Println("joining", server, c)
		ch, err := cl.Join(ctx, server, c)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go run(ctx, ch)
	}
	fmt.Println("ready")
	select {}
}

func run(ctx context.Context, ch chat.Channel) {
	for {
		ev, err := ch.Receive(ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		evstr := pretty.String(ev)
		evstr = strings.Replace(evstr, "\n", " ", -1)
		evstr = strings.Replace(evstr, "\t", "", -1)
		fmt.Println(evstr)

		switch ev := ev.(type) {
		case chat.Message:
			fixBridgeMessage(&ev)
			switch {
			case strings.HasPrefix(ev.Text, parsePrefix):
				input := strings.TrimPrefix(ev.Text, parsePrefix)
				input = strings.TrimSpace(input)
				say(ctx, ch, parse(input))
			}
		}
	}
}

func parse(input string) string {
	input = strings.TrimSpace(input)
	p := ast.NewParser(input)
	text, err := p.Text()
	if err != nil {
		return "syntax error " + err.Error()
	}
	parse := ast.BracesString(text)
	if parse != "" {
		parse += "\n"
	}
	var math string
	stmt := logic.Interpret(text)
	if stmt == nil {
		math = "fragment"
	} else {
		math = logic.PrettyString(stmt)
	}
	return parse + math
}

func fixBridgeMessage(msg *chat.Message) {
	if msg.From == nil ||
		msg.From.Nick != "ToaqBridgeBot" ||
		strings.HasPrefix(msg.From.Nick, "Toaq Telegram") {
		return
	}
	s := strings.SplitN(msg.Text, ": ", 2)
	if len(s) != 2 {
		return
	}
	msg.Text = s[1]
}

func say(ctx context.Context, ch chat.Channel, text string) {
	if _, err := chat.Say(ctx, ch, text); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
