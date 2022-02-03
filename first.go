package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"math/rand"
	"time"

	
)
var moving int=0
var cx int =0
var cy int =0
var lx int =0
var ly int =0
var temp int=0
var drag int=0
var sx int =0
var sy int=0
var new_words int=0
var old_words int=0

func Home(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8"/>
	<meta name='viewport' content='width=device-width, initial-scale=1.0, maximum-scale=1.0, 
			 user-scalable=0' >
	<script type="text/javascript" charset="utf-8">
	function init() {
		var toucharea = document.getElementById("zone");
	    toucharea.addEventListener("touchend", clicke, false);
		toucharea.addEventListener("touchstart", tostart, false);		
		toucharea.addEventListener("touchmove", handle, false);
		
		var scrollarea= document.getElementById("scroll");
		scrollarea.addEventListener("touchmove", scrolle, false);
		scrollarea.addEventListener("touchstart", tostart, false);		


			}
	
	
		function sendData(a,b,g) {
		  var xhttp = new XMLHttpRequest();
		  
		  xhttp.open("POST", "receive", true);
		 xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
		 xhttp.send("a="+a+"&b="+b);
		}
		function on_input(){
			var string = document.getElementById("inpfield").value;
			var xhttp = new XMLHttpRequest();
			
			xhttp.open("POST", "typed", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.send("string="+string);
		  }

		function scrolle(event){
			var a    = parseInt(event.touches[0].pageX);
			var b     = parseInt(event.touches[0].pageY);
			var xhttp = new XMLHttpRequest();
			
			xhttp.open("POST", "scrolling", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.send("a="+a+"&b="+b);
		  }
		function S_shot(){
			var xhttp = new XMLHttpRequest();
			xhttp.open("POST", "ss", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.send();
		}
		function tostart(event){
			var a    = parseInt(event.touches[0].pageX);
			var b     = parseInt(event.touches[0].pageY);
			var xhttp = new XMLHttpRequest();
			
			xhttp.open("POST", "fix", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.send("a="+a+"&b="+b);
		  }
	   
	
		 function handle(event) {
		  var alpha    = parseInt(event.touches[0].pageX);
		  var beta     = parseInt(event.touches[0].pageY);
		  sendData(alpha,beta);
		  
		}

		function ENTER(){
			var xhttp = new XMLHttpRequest();
			xhttp.open("POST", "ENT", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.send();
		}

		function clicke(event) {
			var xhttp = new XMLHttpRequest();
			
			var a = parseInt(event.changedTouches[0].pageY);
			xhttp.open("POST", "click", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.send("a="+a);
		  }

		  function drag(){
			var btn = document.getElementById("drag");
			 var xhttp = new XMLHttpRequest();
			 xhttp.open("POST", "dragM", true);
			 xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			 xhttp.send();
			 
			 if (btn.innerHTML == "left mouse key down"){
				 btn.innerHTML = "left mouse key up";
			 }
			 else{
				 btn.innerHTML = "left mouse key down";
			 }
		 }
	
			</script>
			  
			<meta name="" content="">
			<title></title>
		  <style>
		   
		  .center{
			position: fixed;
			background-color: rgb(230,230,230);
			width:230px;
			
			top:85px;
			left:0%;
			color:rgb(200,200,200);
			
		   }
		   
		   html{
			touch-action:pan-down
			}
		  
			.corner{
				position: fixed;
				background-color: rgb(230,230,230);
				width:90px;
				
				top:175px;
				left:250px;
				color:rgb(200,200,200);

			}
			.Enter{
				position: fixed;
				top: 15px;
				padding: 5px 35px;
				left:250px;
				font-size: 20px;
				}
		 
			.lkeydown{
				position: fixed;
				top: 75px;
				padding: 5px 20px;
				left:250px;
				font-size: 16px;
					}
			.Scr_shot{
					position: fixed;
					top: 470px;
					padding: 5px 15px;
					left:250px;
					font-size: 16px;
				}
		 
		  </style>
		  </head>
		  <body onload="init()">

		  <div>
          <button class="lkeydown" id="drag" type="button" onclick="drag()">left mouse key down</button>
          
      </div>

	<p><button type="button" class="Enter" onclick="ENTER()">enter</button>
	<p><button type="button" class="Scr_shot" onclick="S_shot()">ScreenShot</button>
	<input style="height:30px;width:220px;font-size: 16pt;" id="inpfield"type="text" oninput=on_input()></input>
	</p>
	<div class="center" id="zone">
				<p align="center"><br><br><br><br><br><br><br>Left click area<br><br><br><br><br><br><br><hr></p>
				<p align="center"><br><br><br>Right click area<br><br><br><hr></p>
	</div>
				
	<div class="corner" id="scroll">
				<p align="center"><br><br><br><br><br>S<br>C<br>R<br>O<br>L<br>L</p> 
			   
	</div>
		   
		   
	</body>`

	w.Write([]byte(fmt.Sprintf(html)))

}

func set_cursor(w http.ResponseWriter, r *http.Request) {
	if r.Method== "POST" {
		if s, err := strconv.Atoi(r.FormValue("a")); err == nil {
			lx=s
			sx=s
			fmt.Println(s)
		}
		
		if v, err := strconv.Atoi(r.FormValue("b")); err == nil {
			ly=v
			sx=v
			fmt.Println(v)
		}
		fmt.Println("hi",lx,ly)		
	}
}
func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// var tx int = int(r.FormValue("a"))
		// var ty int = int(r.FormValue("b"))
		if s, err := strconv.Atoi(r.FormValue("a")); err == nil {
			cx=s
		}
		
		if v, err := strconv.Atoi(r.FormValue("b")); err == nil {
			cy=v
		}
		// fmt.Println(cy)
		// cy=ty
		// fmt.Println(cx,cy,lx,ly)
		fmt.Println("Receive ajax post data string ",cx-lx,cy-ly)
		robotgo.MoveRelative((cx-lx)*3,(cy-ly)*3)
		lx=cx
		ly=cy
		moving=1
		// w.Write([]byte("<h2>after<h2>"))
	}
}

func CLICK(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {

		if s, err := strconv.Atoi(r.FormValue("a")); err == nil {
			temp = s
			fmt.Println("click done")
		}
		if moving==0{
		if temp<450{
			robotgo.Click("left")
		}
		if temp>450{

			robotgo.Click("right")
		}
		
	}
	moving=0
	}

	}

func press_enter(w http.ResponseWriter, r *http.Request){
    if r.Method=="POST"{
		robotgo.KeyTap("enter")
	}
}
 
func leftkeydown(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
		if drag==0{
			robotgo.MouseToggle("down")
			drag=1

		}else{
			robotgo.MouseToggle("up")
			drag=0
		}

		
	}

}
func scrolled(w http.ResponseWriter, r *http.Request){
   if r.Method=="POST"{
	if s, err := strconv.Atoi(r.FormValue("a")); err == nil {
		cx = s
		fmt.Println("click done")
	if v, err := strconv.Atoi(r.FormValue("b")); err == nil {
		cy = v
		fmt.Println("click done")
	}
	if sy>cy{
		robotgo.Scroll(0,-1)
	}else{
		robotgo.Scroll(0,1)
	}
	sx=cx
	sy=cy
   }

}
}
func typing(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
		a:=r.FormValue("string")
		new_words=len(a)
		if new_words>old_words{
		robotgo.TypeStr(a[len(a)-1:])
	    }else if(old_words>new_words){
			robotgo.KeyTap("backspace")
		}
	old_words=new_words
	}

}
func shot_taken(w http.ResponseWriter, r *http.Request){
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)
		  
		j:=y1.Intn(200)
		fileName := fmt.Sprintf("%d_%dx%d.png", j, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}


}
func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/receive", receiveAjax)
	mux.HandleFunc("/fix", set_cursor)
	mux.HandleFunc("/click", CLICK)
	mux.HandleFunc("/ENT", press_enter)
    mux.HandleFunc("/dragM", leftkeydown)
	mux.HandleFunc("/scrolling", scrolled)
	mux.HandleFunc("/typed", typing)
	mux.HandleFunc("/ss", shot_taken)
	http.ListenAndServe(":8080", mux)
}