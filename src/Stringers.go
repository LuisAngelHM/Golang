package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM , %d GB de disco y es una %s ", myPC.ram, myPC.disk, myPC.brand)
}

func main() {

	MyPc := pc{ram: 16, brand: "lenovo", disk: 2000}
	fmt.Println(MyPc)
}
