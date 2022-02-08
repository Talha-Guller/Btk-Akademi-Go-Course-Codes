package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// o iki kesme işaretini altgr+virgülün olduğu tuşa ve space bastıktan sonra oluyor

type Todo struct {
	UserId    int    `json:"userId"` //apı deki karşılığı ne ise json a onu yazıyoruz
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//get operasyonu
func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	//burda deferı kullanmalıyız çünkü diğer kodlarda hata çıksa bile get operasyonlarını kapatmalıyız
	//açık bi get sonrasında bize zorluklar çıkara bilir.

	//gelen veriyi anlamlı hale getirmek için şunları yaparız.
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	//bu işlem string bir biçimde bize sunar ama bizim todo şekline ihtiyacımız olduğu için şunları da yaparız

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
	// bize direk nesne gerektiği için bunu kullanmak daha mantıklı

}

//post operasyonu
func Demo2() {
	todo := Todo{1, 2, "Alışveriş Yapılacak", false}
	//elimizdeki struct'ı altta json formatına çeviriyoruz Marshal İle
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}

	//application hangi formatta data gönderceğini yazıyor ve alttaki gibi ikili göndere biliyorsun
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)

}
