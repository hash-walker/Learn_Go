/*
Message Formatter
As Textio evolves, the team has decided to introduce a new feature for custom message formats. Depending on the user’s preferences, messages can be sent in different formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you’ll implement a system using interfaces.

Assignment
Implement the Formatter interface with a method Format that returns a formatted string.
Define structs that satisfy the Formatter interface: PlainText, Bold, Code.
The structs must all have a message field of type string
PlainText should return the message as is.
Bold should wrap the message in two asterisks (**) to simulate bold text (e.g., message).
Code should wrap the message in a single backtick (`) to simulate code block (e.g., message)
*/

package main

import "fmt"

type Formatter interface {
	Format() string
}

type PlainText struct{
	message string
}

type Bold struct{
	message string
}

type Code struct{
	message string
}

func (p PlainText) Format() string{
	return p.message
}

func (b Bold) Format() string{
	return fmt.Sprintf("**%v**", b.message)
}

func (c Code) Format() string{
	return fmt.Sprintf("`%v`", c.message)
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
