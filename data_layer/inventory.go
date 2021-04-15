// Package data_layer implements the data store
// to store inventory and player data
package data_layer

import "errors"

// Initialize inventory of a player
// looks like this in JSON
// inventory = {
// 	player: {
// 		item: quantity
// 	}
// }
var inventory = make(map[string]map[string]int)

// Returns all inventory a player currently has
func GetAllInventory(player string) (map[string]int, error) {
	// Checks if player exist
	if val, ok := inventory[player]; ok {
		return val, nil
	}
	return map[string]int{}, errors.New("this player does not exist")

}

// Add item into player's inventory
func AddToInventory(player string, item string, quantity int) error {
	// Checks if player, item and quantity have valid values
	if len(player) == 0 || len(item) == 0 || quantity < 0 {
		return errors.New("invalid values for player/item/quantity")
	}
	// Checks if player inventory is empty
	// If empty, then initialize a map.
	if inventory[player] == nil {
		inventory[player] = make(map[string]int)
	}

	// After initialized player inventory, add item into it.
	// If inventory is initialzed, add item into it.
	inventory[player][item] += quantity
	return nil
}

// Remove item from player's inventory
func UseItems(player string, item string, quantity int) error {
	// Check if player exists in inventory
	if val, ok := inventory[player]; ok {
		// Check if item exists in player inventory
		if _, ok := val[item]; ok && val[item] > 1 {
			// If player and item exists in inventory
			// then reduce by the quantity used
			inventory[player][item] -= quantity
			return nil
		} else {
			return errors.New("this player does not have this item")
		}
	}
	return errors.New("this player does not exist")
}
