package locales_test

import (
	"fmt"
	"testing"

	"github.com/hebcal/locales"
	"github.com/stretchr/testify/assert"
)

func TestHebrewStripNikkud(t *testing.T) {
	assert := assert.New(t)
	src := "חֲנוּכָּה יוֹם ד׳ (בְּשַׁבָּת)"
	dest := "חנוכה יום ד׳ (בשבת)"
	assert.Equal(dest, locales.HebrewStripNikkud(src))
	assert.Equal("א־ב", locales.HebrewStripNikkud("א־ב"))
	assert.Equal("אב", locales.HebrewStripNikkud("אֿב"))
}

func ExampleHebrewStripNikkud() {
	src := "וְהָאָ֗רֶץ הָיְתָ֥ה תֹ֙הוּ֙ וָבֹ֔הוּ"
	dest := locales.HebrewStripNikkud(src)
	fmt.Println(dest)
	// Output:
	// והארץ היתה תהו ובהו
}
