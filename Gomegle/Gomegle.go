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

	chatTextBox, err := g.Driver.FindElement(selenium.ByCSSSelector, ".newtopicinput ")
	err = chatTextBox.SendKeys(`Gomegle
	`)
	if err != nil {
		panic(err)
	}
	elem.Click()
}

func (g *Gomegle) SendMessage(message string) bool {

	partnerConnected := g.checkIfPartnerConnected()

	if !partnerConnected{
		g.findNewPartner();
		return false
	}

	if message == "DISCONNECT" {
		g.findNewPartner()
		g.findNewPartner()
		g.findNewPartner()
		return false
	}

	// Enter some new message in text box.
	chatTextBox, err := g.Driver.FindElement(selenium.ByCSSSelector, ".chatmsg ")
	err = chatTextBox.SendKeys(message + `	
	`)
	if err != nil {
		panic(err)
	}
	return true
} 

func (g *Gomegle) checkIfPartnerConnected() bool {
	strangerMessages, err := g.Driver.FindElements(selenium.ByCSSSelector, "textarea[class='chatmsg disabled']")
	if err != nil {
		panic(err)
	}
	if (len(strangerMessages) == 0) {
		return true;
	}
	return false;
}

func (g *Gomegle) findNewPartner() {
	
	elem, err := g.Driver.FindElement(selenium.ByCSSSelector, ".disconnectbtn")

	if err != nil {
		panic(err)
	}

	elem.Click();

}

func (g *Gomegle) CheckForNewMessage (lastLatest int) (string, int) {

	partnerConnected := g.checkIfPartnerConnected()

	if !partnerConnected{
		g.findNewPartner();
		return "DISCONNECTED", 0
	}

	var returnMessages string = ""
	strangerMessages, err := g.Driver.FindElements(selenium.ByCSSSelector, ".strangermsg ")
	if err != nil {
		panic(err)
	}
	var thisLatest int
	for i := lastLatest; i < len(strangerMessages); i++ {
		strangerSpan, err := strangerMessages[i].FindElement(selenium.ByTagName, "span")
		if err != nil {
			return returnMessages, 0
		}
		strangerSpanText, err := strangerSpan.GetAttribute("innerHTML")
		if err != nil {
			return returnMessages, 0
		}
		returnMessages += strangerSpanText
		thisLatest = i
	} 
	return returnMessages, thisLatest+1
} 
