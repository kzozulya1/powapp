package wisdomquotes

import (
	"crypto/rand"
	"math/big"
)

// quotes set of quotes of Wisdom
var quotes = []string{
	"We must believe that we are gifted for something, and that this thing, at whatever cost, must be attained.",
	"The older I get, the greater power I seem to have to help the world; I am like a snowball - the further I am rolled the more I gain.",
	"Knowledge is love and light and vision.",
	"Wisdom is knowing what to do next. Skill is knowing how to do it. Virtue is doing it.",
	"When one door of happiness closes, another one opens; but we look so long at the closed door that we do not see the one which has opened for us.",
	"Poor eyes limit your sight; poor vision limits your deeds.",
	"I do not pray for success. I ask for faithfulness.",
	"I used to ask God to help me. Then I asked if I might help him.",
	"The wise person doesn't give the right answers, but poses the right questions.",
	"What happens is not as important as how you react to what happens.",
	"The journey of a thousand miles begins with one step.",
	"The only true wisdom is in knowing you know nothing.",
	"Just as treasures are uncovered from the earth, so virtue appears from good deeds, and wisdom appears from a pure and peaceful mind. To walk safely through the maze of human life, one needs the light of wisdom and the guidance of virtue.",
	"True wisdom comes to each of us when we realize how little we understand about life, ourselves, and the world around us.",
	"Tell me and I'll forget; show me and I may remember; involve me and I'll understand.",
	"Just remember the world is not a playground but a schoolroom. Life is not a holiday but an education. One eternal lesson for us all: to teach us how better we should love.",
	"Prayer is the raising of one's mind and heart to God or the requesting of good things from God.",
	"Prayer is not an old woman's idle amusement. Properly understood and applied, it is the most potent instrument of action.",
	"Perhaps instead of bombarding God with requests for what is not, we might try, instead, asking God to open our eyes to see what is.",
	"Listen to the wind, it talks. Listen to the silence, it speaks. Listen to your heart, it knows.",
	"Grace is not part of consciousness; it is the amount of light in our souls, not knowledge nor reason.",
	"There are two ways of spreading light: to be the candle or the mirror that reflects it.",
}

// Quote selects random quote
func Quote() string {
	r, err := random(int64(len(quotes)))
	if err != nil {
		return ""
	}

	return quotes[r]
}

// random yields random number in range [0..ceilValue]
func random(ceilValue int64) (int64, error) {
	bBig, err := rand.Int(rand.Reader, big.NewInt(ceilValue))
	if err != nil {
		return 0, err
	}

	return bBig.Int64(), nil
}
