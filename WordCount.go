//功能要求如下：
//1，输入一个文件后，自动截取前1000个字符，统计其中“a”“c”“d”“e”出现的次数；
//2，输出格式要求：a c d e
//12 15 21 9

package main
import (
    "fmt"
    "os"
)
func main() {
    if len(os.Args) < 2 { 
        fmt.Println("Please Input File Name!")
        return
    }   
    file, err := os.Open(os.Args[1])
    if err != nil {
        return
    }   
    buff := make([]byte, 1000)
    _, err = file.Read(buff)
    if err != nil {
        return
    }   
    a, c, d, e, _ := Sum(buff)
    fmt.Printf("a   c   d   e   \n%d    %d  %d  %d", a, c, d, e)
}
func Sum(buff []byte) (a, c, d, e, def int) {
    for i := 0; i < len(buff); i++ {
        switch buff[i] {
        case 'a':
            a++
        case 'c':
            c++
        case 'd':
            d++
        case 'e':
            e++
        default:
            def++
        }
    }   
    return
}