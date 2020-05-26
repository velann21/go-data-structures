package solid


// wrong approach here the SavetoDisk and SentTONw is not a responsible of Journal,
// Journal should do only one responsibility that is AddEntry, RemoveEntry
// SavaToDisk and SaveToNW should be seperate entity
//SRP is that there should be only one reason to change the class
type Journal struct {
	entries []string
}

func (jrl *Journal) AddEntry(){

}

func (jrl *Journal) RemoveEntry(){

}




//correct approach
// Here in below approach you don't require to change the Jouranl if you requierd any chnages on Persistent
//or Network, this is example of SRP
type JournalCorrect struct {
	entries []string
}

func (jrl *JournalCorrect) AddEntry(){

}

func (jrl *JournalCorrect) RemoveEntry(){

}

type Persistent struct {

}

func (per *Persistent) SaveToDisk(){

}

type NetworkCall struct {

}

func (nw *NetworkCall) SendToNW(){

}

