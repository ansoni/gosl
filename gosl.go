package gosl

import (
	"github.com/ansoni/termination"	
	"os"
	"time"
)

func deadTrain(term *termination.Termination, entity *termination.Entity) {
	term.Close()
	os.Exit(0)
}

func train1(term *termination.Termination) {

	locY:=(term.Height-trainTopHeight)/2
	locX:=term.Width+1
	
	train := term.NewEntity(termination.Position{locX, locY, 0})
	train.Shape=trainTopShape
	train.MovementCallback=termination.LeftMovement

	trainBottom := term.NewEntity(termination.Position{locX, locY+trainTopHeight, 0})
	trainBottom.Shape=trainBottomShape
	trainBottom.MovementCallback=termination.LeftMovement	

	trainSmoke := term.NewEntity(termination.Position{locX, locY-trainSmokeHeight, 0})
	trainSmoke.Shape=trainSmokeShape
	trainSmoke.FramesPerSecond=4
	trainSmoke.MovesPerSecond=25
	trainSmoke.MovementCallback=termination.LeftMovement	

	trainCarriage := term.NewEntity(termination.Position{locX+trainTopWidth, locY, 0})
	trainCarriage.Shape=trainCarriageShape
	trainCarriage.DeathOnOffScreen = true
	trainCarriage.DeathCallback=deadTrain
	trainCarriage.MovementCallback=termination.LeftMovement	
}

func failSafe(term *termination.Termination) {
	time.Sleep(60 * time.Second)
	term.Close()
	os.Exit(0)
}


func Sl() {
	term := termination.New()
	go failSafe(term)
	term.FramesPerSecond = 25
	//term.Debug = "./debug.out"
	train1(term)
	term.Animate()
}

func main() {
	Sl()
}
