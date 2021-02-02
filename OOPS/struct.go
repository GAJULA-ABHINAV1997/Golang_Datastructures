package main 
  
import ( 
    "fmt"
) 
  
// declaring a struct 
type Book struct{ 
      
    // defining struct variables 
    name string 
    author string 
    pages int
} 
  
// function to print book details 
func (book Book) print_details(){ 
  
    fmt.Printf("Book %s was written by %s.", book.name, book.author) 
    fmt.Printf("\nIt contains %d pages.\n", book.pages) 
} 
  
// main function 
func main() { 
      
    // declaring a struct instance 
    book1 := Book{"Harry Potter", "J,K Rowling", 131} 
      
    // printing details of book1 
    book1.print_details() 
      
    // modifying book1 details 
    book1.name = "Electrical Drives"
    book1.pages = 162 
      
    // printing modified book1 
    book1.print_details() 
      
} 
