// package main

// import (
// 	"fmt"
// 	// "sort"
// )

// type Student struct {
// 	Name   string
// 	Grade  int
// 	Course string
// }

// func (s *Student) calculateAverage() {
// 	grades := []int{50, 70, s.Grade}

// 	total := 0
// 	for _, grade := range grades {
// 		total += grade
// 	}
// 	average := total / len(grades)

// 	fmt.Printf("Student: %s, Course: %s, Average Grade: %d\n", s.Name, s.Course, average)

// 	if average >= 60 {
// 		fmt.Println("You are passed.")
// 	} else {
// 		fmt.Println("You are failed.")
// 	}
// }

// func main() {

// 	student := Student{Name: "Umar", Grade: 55, Course: "IT"}
// 	student.calculateAverage()
// }





// func main(){
// 	filePath := ""
// 	fmt.Scan(&filePath)
// 	if checkFileExists(filePath + ".txt"){

// 	var rec1,rec2,rec3 string = filePath,filePath,filePath
//     rec1 += strconv.Itoa(rand.Intn(100))+ ".txt"
// 	rec2 += strconv.Itoa(rand.Intn(100))+ ".txt"
// 	rec3 += strconv.Itoa(rand.Intn(100))+ ".txt"
// 	fmt.Println("File already exist, you can create files with name like this:")
// 	if !checkFileExists(rec1){
// 		fmt.Printf("%v",rec1)
// 	}
// 	if !checkFileExists(rec2){
// 		fmt.Printf("%v",rec2)
// 	}
// 	if !checkFileExists(rec3){
// 		fmt.Printf("%v",rec3)
// 	}
// 	fmt.Println()
// }else {
// 	_,err := os.Create(filePath + ".txt")
// 	if err != nil{
// 		panic(err)
// 	}
// }
// }
// func checkFileExists(filePath string) bool {
// 	_,err := os.Stat(filePath)
// 	return !errors.Is(err,os.ErrNotExist)
// }




// func main(){
// 	var n int
// 	fmt.Println("enter the count of people:")
// 	fmt.Scanln(&n)
// 	file,err := os.Create("people.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	offSet := 0
// 	for i := 1; i <= n; i++{
// 		var name,mail,job string
// 		fmt.Println("enter name:")
// 		fmt.Scanln(&name)
// 		fmt.Println("enter mail:")
// 		fmt.Scanln(&mail)
// 		fmt.Println("enter job:")
// 		fmt.Scanln(&job)

// 		txt := fmt.Sprintf("%v name: %v mail: %v job: %v\n",i, name, mail,job)
// 		written, err := file.WriteAt([]byte(txt), int64(offSet))
// 		if err != nil{
// 			panic(err)
// 		}
// 		offSet += written
// 	}
// }




// type Person struct {
// 	Name string
// 	Age  int
// }

// type People []Person

// func (p People) Len() int           { return len(p) }
// func (p People) Less(i, j int) bool { return p[i].Age > p[j].Age }
// func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// func main() {
// 	people := People{
// 		{"John", 30},
// 		{"Alice", 25},
// 		{"Smith", 35},
// 	}
// 	sort.Sort(people)

// 	fmt.Println(people)
// }