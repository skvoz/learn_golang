package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	datesI18N "github.com/davidgs/datesi18n"
)

func main() {
	//read part of text
	//	- we must now order this part
	//translate this part (realize later)
	//write this part to new file
	copyData("anna_short.txt", "new_anna_short.txt")

	d := time.Now()
	d.Month()
	fmt.Println(d.Format("January 2, 2006"))
	year, month, day := d.Date()
	fr := datesI18N.NewDatesI18N("fr") // french
	fmt.Printf("French: %s, %s %d, %d\n", fr.DayName(int(d.Weekday())), fr.MonthName(int(month)), day, year)
	ru := datesI18N.NewDatesI18N("ru") // russian
	fmt.Printf("Russian: %s, %s %d, %d\n", ru.DayName(int(d.Weekday())), ru.MonthName(int(month)), day, year)
}

func copyData(srcFile string, dstFile string, datesI18N datesI18N) {
	data, err1 := ioutil.ReadFile(srcFile)
	if err1 != nil {
		log.Fatal(err1)
	}

	fileDst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer fileDst.Close()

	// Creating a buffered writer from the file variable using bufio.NewWriter()
	bufferedWriter := bufio.NewWriter(fileDst)

	// declaring a byte slice
	// bs := []byte{97, 98, 99}

	// writing the byte slice to the buffer in memory
	bytesWritten, err := bufferedWriter.Write(data)

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)

	// checking how much data is stored in buffer, just  waiting to be written to disk
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// -> 24 (3 bytes in the byte slice + 21 runes in the string, each rune is 1 byte)
	// The bytes have been written to buffer, not yet to file.

	// Writing from buffer to file.
	bufferedWriter.Flush()
}
