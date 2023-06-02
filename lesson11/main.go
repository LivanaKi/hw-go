 package main

 import(
	"fmt"
	"time"
 )

 func checkRead(ch int) {

 }
 func checkReadHang(ch int){

 }
 func checkReadPanic(ch int){

 }

 func TestReadEmpty(t *testing.T){
	ch := make(chan int, 5)
	checkReadHang(t, ch)
 }

 func TestReadClosedEmpty(t *testing.T){
	ch := make( chan int, 5)
	close(ch)
	checkReadHang(t, ch)
 }

 func TestReadClosed(t *testing.T){
	ch := make(chan int, 5)
	ch <- 5
	close(ch)
	checkReadOk(t, ch)
 }

 func TestWriterFull(t *testing.T){
	ch := make(chan int, 5)
	ch <- 5
	checkWriteHang(t, ch)
 }

 func TestWriterClosed(t *testing.T){
	ch := make(chan int, 5)
	close(ch)
	checkWritePanic(t, ch)
 }