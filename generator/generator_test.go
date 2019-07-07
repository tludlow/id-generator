package generator

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringLength(t *testing.T) {
	//We will generate strings of the required lengths, removing the prefix and measuring the length to make sure its correct.
	//This will run 1000 times, with a new string each time.
	for i := 0; i < 1000; i++ {
		//Random number for the string length generating.
		length := rand.Intn(30-1) + 1

		//Generate a new string, the prefix is random as its not important
		testString, _ := GenerateID(length, "tst")

		//Split the string generated so that we only have all the characters following the "_"
		splitString := strings.Split(testString, "_")[1]

		//assert.Equal(t, len(splitString), length, "Expected string length and actual length differ")
		assert.Len(t, splitString, length, "Expected string length and actual length differ")
	}
}

func TestPrefixError(t *testing.T) {
	//A id with a prefix which is NOT 3 characters in length will fail. Let's test this.

	//Should error, prefix 5 chars.
	id1, err1 := GenerateID(10, "hello")
	assert.NotNil(t, err1, "Error should exist but does not.")
	assert.Empty(t, id1, "ID should be empty as prefix incorrect.")

	//Should error, prefix 4 chars
	id2, err2 := GenerateID(15, "dogs")
	assert.NotNil(t, err2, "Error should exist but does not.")
	assert.Empty(t, id2, "ID should be empty as prefix incorrect.")

	//Should NOT error, prefix 3 chars
	id3, err3 := GenerateID(21, "usr")
	assert.Nil(t, err3, "Error should NOT exist but does.")
	assert.NotEmpty(t, id3, "ID should NOT be empty as prefix incorrect.")

	//Should error, prefix 2 chars
	id4, err4 := GenerateID(16, "hi")
	assert.NotNil(t, err4, "Error should exist but does not.")
	assert.Empty(t, id4, "ID should be empty as prefix incorrect.")

}
