package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

//check files for error in permissions
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var uploaded bool
	uploaded = false
	fmt.Println(" ENTER FILE NAME ")
	var file_name string
	fmt.Scanln(&file_name)
	pause := make(chan int)

	terminate := make(chan int)
	//read from this file

	wg2 := new(sync.WaitGroup)
	wg2.Add(1)

	dat, err := ioutil.ReadFile(file_name)
	check(err)
	data := string(dat)

	//write and append to file
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	//gorutine to uload data (for now copy data in another file)
	go func() {
		var flag bool
		flag = false
		fmt.Println("UPLOADING  ... ... \n")

		for c := range data {
			//fmt.Println(i)
			//countinue upload/coping of data
			if flag == true {
				break
			}

			_, err2 := f.WriteString(string(c))

			if err2 != nil {
				log.Fatal(err2)
			}
			select {
			case msg := <-terminate:
				_ = msg
				flag = true
			case msg2 := <-pause:
				_ = msg2
				fmt.Println("ENTER 1 TO RESUME AND 2 TO TERMINATE \n")
				var temp int
				fmt.Scanln(&temp)
				if temp == 1 {
					flag = false
				}
				if temp == 2 {
					flag = true
				}

			default:
				flag = false
			}

		}
		if flag == true {
			fmt.Println("FILE TERMINATED \n")
			uploaded = true
		}
		if flag == false {
			fmt.Println("FILE UPLOADED .\n")
			uploaded = true
		}
		defer wg2.Done()
	}()
	fmt.Println("ENTER 1 TO PAUSE AND 2 TO TERMINATE.\n")
	for uploaded == false {

		var user_input int
		fmt.Scanln(&user_input)

		if user_input == 1 {
			pause <- 1
		}
		if user_input == 2 {
			terminate <- 1
		}

	}
	wg2.Wait()
}
