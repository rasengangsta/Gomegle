package Gomegle

import "github.com/tebeka/selenium"

type Gomegle struct {
	Driver selenium.WebDriver
	messages []string
}

func SetupOmegle() Gomegle {
	// Connect to the WebDriver instance running locally.

	g := Gomegle{}
	var err error
	caps := selenium.Capabilities{"browserName": "chrome"}
	g.Driver, err = selenium.NewRemote(caps, "http://localhost:9515")
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	
	// Navigate to the omegle.
	if err := g.Driver.Get("http://www.omegle.com/"); err != nil {
		panic(err)
	}
	return g
}

func (g *Gomegle) StartChatting () {
	// Get a reference to the text chat button.
	elem, err := g.Driver.FindElement(selenium.ByCSSSelector, "#textbtn")
	if err != nil {
		panic(err)
	}
	elem.Click()
}

func (g *Gomegle) SendMessage(message string) {
	// Enter some new message in text box.
	chatTextBox, err := g.Driver.FindElement(selenium.ByCSSSelector, ".chatmsg ")
	err = chatTextBox.SendKeys(message + `	
	`)
	if err != nil {
		panic(err)
	}
} 

func (g *Gomegle) CheckForNewMessage () []string {
	var returnMessages []string 
	strangerMessages, err := g.Driver.FindElements(selenium.ByCSSSelector, ".strangermsg ")
	if err != nil {
		panic(err)
	}
		
	for i := 0; i < len(strangerMessages); i++ {
		strangerSpan, err := strangerMessages[i].FindElement(selenium.ByTagName, "span")
		if err != nil {
			return returnMessages
		}
		strangerSpanText, err := strangerSpan.GetAttribute("innerHTML")
		if err != nil {
			return returnMessages
		}
		returnMessages = append(returnMessages, strangerSpanText)
	} 
	return returnMessages
} 
