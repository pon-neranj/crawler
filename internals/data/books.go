package data
import(
	"fmt"
	"math/rand"
	"time"
)

type Book struct {
	ID        string     `json:"id"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Price   int    `json:"-"`
}
var Booklist []Book

func(b *Book)Insert()(Book){

	newUUID := time.Now().Unix()
	
    
	uuid := fmt.Sprintf("%d", newUUID) 
	
	book := Book{
		ID:uuid,
		Title:fmt.Sprintf("%s-neranj",uuid[:3]),
		Genres:[]string{"thriller","Comedy"},
		Price: rand.Intn(100),
	}

	Booklist = append(Booklist,book)

	return book
}

func(b *Book)FetchALL()[]Book{

	return Booklist

}