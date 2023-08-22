package main

import "fmt"

type Contact struct {
	name     string
	email    string
	skype    string
	telegram string
}

func (self *Contact) present() string {
	return fmt.Sprintf("Hello, my name is %s. My contacts (email:%s, telegram:%s, skype:%s)",
		self.name, self.email, self.telegram, self.skype)
}

func main() {
	c := Contact{
		name:     "Alexander Vakhov",
		email:    "alex.vakhov@gmail.com",
		telegram: "@vakhov",
		skype:    "alex.vakhov",
	}
	fmt.Println(c.present())
}
