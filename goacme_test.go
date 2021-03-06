

/*4:*/


//line goacme.w:99

package goacme

import(
"os/exec"
"9fans.net/go/plan9/client"
"testing"


/*12:*/


//line goacme.w:196

"fmt"
"time"
"9fans.net/go/plan9"




/*:12*/



/*17:*/


//line goacme.w:249

"bytes"
"errors"



/*:17*/


//line goacme.w:106

)

func prepare(t*testing.T){
_,err:=client.MountService("acme")
if err==nil{
t.Log("acme started already")
}else{
cmd:=exec.Command("acme")
err= cmd.Start()
if err!=nil{
t.Fatal(err)
}


/*13:*/


//line goacme.w:203

time.Sleep(time.Second)



/*:13*/


//line goacme.w:119

}
}



/*14:*/


//line goacme.w:207

func TestNewOpen(t*testing.T){
prepare(t)
w,err:=New()
if err!=nil{
t.Fatal(err)
}
defer w.Close()
defer w.Del(true)
if f,err:=fsys.Open(fmt.Sprintf("%d",w.id),plan9.OREAD);err!=nil{
t.Fatal(err)
}else{
f.Close()
}
}



/*:14*/



/*18:*/


//line goacme.w:254

func TestReadWrite(t*testing.T){
w,err:=New()
if err!=nil{
t.Fatal(err)
}
defer w.Close()
defer w.Del(true)
b1:=[]byte("test")
_,err= w.Write(b1)
if err!=nil{
t.Fatal(err)
}
w1,err:=Open(w.id)
if err!=nil{
t.Fatal(err)
}
defer w1.Close()
defer w1.Del(true)
b2:=make([]byte,10)
n,err:=w1.Read(b2)
if err!=nil{
t.Fatal(err)
}
if bytes.Compare(b1,b2[:n])!=0{
t.Fatal(errors.New("buffers don't match"))
}
}



/*:18*/



/*26:*/


//line goacme.w:355

func TestDel(t*testing.T){
w,err:=New()
if err!=nil{
t.Fatal(err)
}
w.Del(true)
w.Close()
if _,err:=Open(w.id);err==nil{
t.Fatal(errors.New(fmt.Sprintf("window %d is still opened",w.id)))
}
}




/*:26*/



/*33:*/


//line goacme.w:422

func TestDeleteAll(t*testing.T){
var l[10]int
for i:=0;i<len(l);i++{
w,err:=New()
if err!=nil{
t.Fatal(err)
}
l[i]= w.id
}
DeleteAll()
for _,v:=range l{
_,err:=Open(v)
if err==nil{
t.Fatal(errors.New(fmt.Sprintf("window %d is still opened",v)))
}
}
}



/*:33*/



/*61:*/


//line goacme.w:784

func TestEvent(t*testing.T){
w,err:=New()
if err!=nil{
t.Fatal(err)
}
defer w.Close()
defer w.Del(true)
msg:="Press left button of mouse on "
test:="Test"
if _,err:=w.Write([]byte(msg+test));err!=nil{
t.Fatal(err)
}
ch,err:=w.EventChannel(0,Look|Execute)
if err!=nil{
t.Fatal(err)
}
e,ok:=<-ch
if!ok{
t.Fatal(errors.New("Channel is closed"))
}
if e.Origin!=Mouse||e.Type!=Look||e.Begin!=len(msg)||e.End!=len(msg)+len(test)||e.Text!=test{
t.Fatal(errors.New(fmt.Sprintf("Something wrong with event: %#v",e)))
}
if _,err:=w.Write([]byte("\nChording test: select argument, press middle button of mouse on Execute and press left button of mouse without releasing middle button"));err!=nil{
t.Fatal(err)
}
e,ok= <-ch
if!ok{
t.Fatal(errors.New("Channel is closed"))
}
if e.Origin!=Mouse||e.Type!=(Execute)||e.Text!="Execute"||e.Arg!="argument"{
t.Fatal(errors.New(fmt.Sprintf("Something wrong with event: %#v",e)))
}
if err:=w.UnreadEvent(e);err!=nil{
t.Fatal(err)
}
if _,err:=w.Write([]byte("\nPress middle button of mouse on Del in the window's tag"));err!=nil{
t.Fatal(err)
}
e,ok= <-ch
if!ok{
t.Fatal(errors.New("Channel is closed"))
}
if e.Origin!=Mouse||e.Type!=(Execute|Tag)||e.Text!="Del"{
t.Fatal(errors.New(fmt.Sprintf("Something wrong with event: %#v",e)))
}
if err:=w.UnreadEvent(e);err!=nil{
t.Fatal(err)
}
}



/*:61*/



/*65:*/


//line goacme.w:875

func TestWriteReadAddr(t*testing.T){
w,err:=New()
if err!=nil{
t.Fatal(err)
}
defer w.Close()
defer w.Del(true)
if b,e,err:=w.ReadAddr();err!=nil{
t.Fatal(err)
}else if b!=0||e!=0{
t.Fatal(errors.New(fmt.Sprintf("Something wrong with address: %v, %v",b,e)))
}
if _,err:=w.Write([]byte("test"));err!=nil{
t.Fatal(err)
}
if err:=w.WriteAddr("0,$");err!=nil{
t.Fatal(err)
}
if b,e,err:=w.ReadAddr();err!=nil{
t.Fatal(err)
}else if b!=0||e!=4{
t.Fatal(errors.New(fmt.Sprintf("Something wrong with address: %v, %v",b,e)))
}
}



/*:65*/



/*68:*/


//line goacme.w:953

func TestWriteReadCtl(t*testing.T){
w,err:=New()
if err!=nil{
t.Fatal(err)
}
defer w.Close()
defer w.Del(true)
if _,err:=w.Write([]byte("test"));err!=nil{
t.Fatal(err)
}
if _,_,_,_,d,_,_,_,err:=w.ReadCtl();err!=nil{
t.Fatal(err)
}else if!d{
t.Fatal(errors.New(fmt.Sprintf("The window has to be dirty\n")))
}
if err:=w.WriteCtl("clean");err!=nil{
t.Fatal(err)
}
if _,_,_,_,d,_,_,_,err:=w.ReadCtl();err!=nil{
t.Fatal(err)
}else if d{
t.Fatal(errors.New(fmt.Sprintf("The window has to be clean\n")))
}
}



/*:68*/


//line goacme.w:123




/*:4*/


