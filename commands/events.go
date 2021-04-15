// Package commands implements discord commands that
// the user can interact with the discord bot

package commands

import (
	"discord-bot/data_layer"
	"errors"

	"github.com/bwmarrin/discordgo"
)

func Add(players []*discordgo.User, item string, quantity int) error {
	for _, player := range players {
		err := data_layer.AddToInventory(player.Username, item, quantity)
		if err != nil {
			return errors.New("failed to add inventory")
		}
	}
	return nil
}

func Get(player *discordgo.User) (map[string]int, error) {
	inventory, err := data_layer.GetAllInventory(player.Username)
	if err != nil {
		return map[string]int{}, errors.New("failed to retrieve inventory")
	}
	return inventory, nil
}

// A player uses an item by a certain quantity from inventory
func Use(player *discordgo.User, item string, quantity int) error {
	err := data_layer.UseItems(player.Username, item, quantity)
	if err != nil {
		return errors.New("can't use item. ")
	}
	return nil
}

// A player completely deletes an item from inventory
func Delete(player *discordgo.User, item string) {}

// A player updates an item from inventory to a target quantity
func Update(player *discordgo.User, item string, targetQuantity int) {}
