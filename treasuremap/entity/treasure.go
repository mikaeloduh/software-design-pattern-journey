package entity

import (
	"math/rand"
)

// Treasure
type Treasure struct {
	Position *Position
	Content  IContent
}

func (t *Treasure) SetPosition(p *Position) {
	t.Position = p
}

func (t *Treasure) GetPosition() *Position {
	return t.Position
}

func NewTreasure() *Treasure {
	// Define the contents of the treasure box and their percentages
	treasureContents := map[IContent]float64{
		SuperStar{}:          0.1,
		Poison{}:             0.25,
		AcceleratingPotion{}: 0.2,
		HealingPotion{}:      0.15,
		DevilFruit{}:         0.1,
		KingsRock{}:          0.1,
		DokodemoDoor{}:       0.1,
	}

	// Generate a random number between 0 and 1
	randomNumber := rand.Float64()

	// Initialize variables to keep track of the selected content and the cumulative probability
	cumulativeProbability := 0.0

	// Iterate through the treasure contents and check if the random number falls within the cumulative probability
	for content, probability := range treasureContents {
		cumulativeProbability += probability
		if randomNumber <= cumulativeProbability {
			return &Treasure{Content: content}
		}
	}
	return nil
}
