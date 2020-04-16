package main

import "fmt"

const globalLimit = 100
const maxCacheSize int = 10 * globalLimit

const (
	cacheKeyBook = "book_"
	cacheKeyCd = "cd_"
)


var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string){
	if len(cache) + 1 >= maxCacheSize {
		return
	}

	cache[key] = val
}

func getBook(isbn string) string {
	return cacheGet(cacheKeyBook + isbn)
}

func setBook(isbn string, name string) {
	cacheSet(cacheKeyBook+isbn, name)
}

func getCd(sku string) string {
	return cacheGet(cacheKeyCd+sku)
}

func setCd(sku string, name string) {
	cacheSet(cacheKeyCd+sku, name)
}

func main() {

	cache = make(map[string]string)

	setBook("1234-5678", "Get Ready To Go")
	setCd("1234-5678", "Get Ready To Go Audio Book")

	fmt.Println("Book: ", getBook("1234-5678"))
	fmt.Println("Cd: ", getCd("1234-5678"))


}