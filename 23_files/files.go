package main

import "os"

func main() {

	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, errInfo := file.Stat()

	// if errInfo != nil {
	// 	panic(errInfo)
	// }
	// we can extract veriouse info from it
	// fmt.Println("file Name", fileInfo.Name())

	// defer file.Close()

	// buff := make([]byte, 12)

	// _, errBuff := file.Read(buff)

	// if errBuff != nil {
	// 	panic(errBuff)
	// }
	// fmt.Println("Buff:", string(buff))

	// //  get data by directly

	// data, errData := os.ReadFile("example.txt")

	// if errData != nil {
	// 	panic(errData)
	// }
	// fmt.Println(string(data))

	// read folder

	// open, openErr := os.Open(".")

	// if openErr != nil {
	// 	panic(openErr)
	// }

	// dir, err := open.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range dir {
	// 	fmt.Println("file", file.Name(), file.IsDir())

	// }

	// write in a file
	file, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Hi Eshan... ")
	file.WriteString("Hi GO... ")

	bytes := []byte("Hi, Golang")
	file.Write(bytes)

}
