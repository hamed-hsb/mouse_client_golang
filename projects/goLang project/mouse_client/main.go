package main
 
import (
    ."fmt"
    "github.com/go-vgo/robotgo"
)


func isInside(x, y, w, h, mx, my int) bool {
	rx := x + w -1
	ry := y + h -1

	return mx >= x && mx <= rx && my >= y && my <=ry

}


func main(){

	var mouseX int
	var mouseY int

	name := "gedit" //say
	fpid, err := robotgo.FindIds(name)
	if err == nil && len(fpid) > 0 {

		pid := fpid[0]
		robotgo.ActivePID(pid) //make gedit active window


		x, y, w, h := robotgo.GetBounds(pid)
		Printf("The active window's top left corner is at (%d, %d) \n",x,y)
		Printf("Its width is %d and its height is %d \n", w, h)
		mx, my := robotgo.GetMousePos()
		Printf("The screen location of the mouse cursor is (%d, %d)\n", mx, my)


		if !isInside(x,y,w,h,mx,my){
			Println("The mouse cursor is outside the active window")
		} else {

			wx := mx -x
			wy := my -y
			Printf("The window location of the mouse cursor is (%d, $d)\n", wx, wy)

		}
	}else{

		Println("Problem when finding PID(s) of ", name)

		for {
		mx, my := robotgo.GetMousePos()
		if (mx != mouseX || my != mouseY){
			
		mouseX = mx
		mouseY = my

        	Printf("The screen location of the mouse cursor is (%d, %d)\n", mx, my)

			}
		}
	}


}
