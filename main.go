package main

/*
#include <stdio.h>
int add(int a,int b){
	return a+b;
}
*/
import "C"
import (
	"log"
)

func main() {
	log.Println(C.add(1, 2))
}
