package main 
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

type CharCount struct {
	EnCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main () {
	
	filename := "/home/ma/goCode/file/1.txt"
    file, err := os.Open(filename)

	if err != nil {
		fmt.Println("open file err %v", err)
		return 
	}

	defer file.Close()

	var cnt CharCount

	reader := bufio.NewReader(file)

	for {
		str, err2 := reader.ReadString('\n')
		fmt.Println(str)
		
		if err2 == io.EOF && str == ""{
			fmt.Printf("wenjianmowei %v", err2)
			break
		}
		for _, v := range str {
			if v >= '0' && v <= '9' {
				cnt.NumCount++
			} else if v >= 'a' && v <= 'z' {
				cnt.EnCount++
			} else if v >= 'A' && v <= 'Z' {
				cnt.EnCount++
			} else if v == ' ' || v == '\t'{
				cnt.SpaceCount++
			} else {
				cnt.OtherCount++
			}
		}

		
	}
	fmt.Printf("英文个数=%v 数字个数=%v 空格个数=%v\n", cnt.EnCount, cnt.NumCount, cnt.SpaceCount)

}