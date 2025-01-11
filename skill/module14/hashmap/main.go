package main

import "fmt"

type user struct {
	name, email, login string
	age                int
}

type userIdentity string

func main() {
	// Объявляем переменную с помощью var и сразу заполняем несколькими значениями
	var hashmap = map[userIdentity]user{
		"vanya_identity": {
			name: "Vanya", email: "vanya@gmail.com", login: "vanator", age: 19,
		},
		"lera_identity": {
			name: "Valery", email: "lera@gmail.com", login: "leraholera", age: 32,
		},
	}

	// Инициализируем переменную с использованием функции make
	// В качестве первого аргумента передаем тип, а авторым указываем ожидаемый размер.
	// Благодаря второму аргументу, go сгенерирует хеш-функцию, оптимизированную на 2000 элемнтов
	makedHashMap := make(map[userIdentity]user, 2000)

	fmt.Println(hashmap)      // Output is map[lera_identity:{Valery lera@gmail.com leraholera 32} vanya_identity:{Vanya vanya@gmail.com vanator 19}]
	fmt.Println(makedHashMap) // Output is map[]

	// Прочитаем значение по ключу:

	// Такой ключ существует
	fmt.Println(hashmap["lera_identity"]) // Output is {Valery lera@gmail.com leraholera 32}

	// Такого ключа не существует, поэтому будет выведено "пустое" значение
	fmt.Println(hashmap["georgy_identity"]) // Output is {   0}

	// Когда мы читаем значение из мапы, мы получаем кортеж из двух элементов: первый -- это значение,
	// а второй -- bool, которое принимает значение true если переменая существует и false  в противном случае
	vanya, ok := hashmap["vanya_identity"]
	fmt.Println(vanya, ok) // Output is {Vanya vanya@gmail.com vanator 19} true

	dasha, ok := hashmap["dasha_identity"]
	fmt.Println(dasha, ok) // Output is {   0} false

	// Мы можем удалить значение из мапы:
	delete(hashmap, "lera_identity")
	fmt.Println(hashmap) // Output is map[vanya_identity:{Vanya vanya@gmail.com vanator 19}]

	// Мы можем добавить новое значение в мапу
	hashmap["kolya_identity"] = user{name: "Nikolay", email: "kolian@gmail.com", login: "kolokol", age: 28}
	fmt.Println(hashmap) // Output is map[kolya_identity:{Nikolay kolian@gmail.com kolokol 28} vanya_identity:{Vanya vanya@gmail.com vanator 19}]

	// Мы можем получить количество сохраненных в мапе элементов:
	hashmapLen := len(hashmap)
	fmt.Println(hashmapLen) // Output is 2

	// Мы можем перебрать все элементы мапы в цикле
	// ВНИМАНИЕ. Go __не гарантирует__ порядок. Т.е. очередность получения элементов будет случайной
	for identity, u := range hashmap {
		fmt.Printf("[%s] %s's email is %s\n", identity, u.name, u.email)
	}
	/**
	Output is:
		[vanya_identity] Vanya's email is vanya@gmail.com
		[kolya_identity] Nikolay's email is kolian@gmail.com
	OR
		[kolya_identity] Nikolay's email is kolian@gmail.com
		[vanya_identity] Vanya's email is vanya@gmail.com
	*/
}
