package main

import (
	"fmt"
)

const bucketSize = 100 // Размер массива для хранения пар ключ-значение

// HashMap структура, представляющая собой хеш-карту
type HashMap struct {
	buckets [][]pair // Массив списков пар ключ-значение
}

// pair структура для хранения одной пары ключ-значение
type pair struct {
	key   string
	value string
}

// NewHashMap конструктор для создания нового экземпляра HashMap
func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([][]pair, bucketSize),
	}
}

// hashstr функция для вычисления хеша строки
func hashstr(val string) uint64 {
	var sum int64 = 0
	for _, char := range val {
		sum += int64(char) // Суммируем ASCII значения символов
	}
	return uint64(sum % int64(bucketSize)) // Возвращаем остаток от деления на размер массива
}

// Set метод для добавления пары ключ-значение в карту
func (h *HashMap) Set(key, value string) {
	index := hashstr(key)
	pair := pair{key, value}

	// Находим список пар для данного индекса
	list := h.buckets[index]

	// Проверяем, есть ли пара с таким же ключом в списке
	found := false
	for i, p := range list {
		if p.key == key {
			list[i].value = value // Обновляем значение
			found = true
			break
		}
	}

	// Если пара не найдена, добавляем её в конец списка
	if !found {
		h.buckets[index] = append(list, pair)
	}
}

// Get метод для получения значения по ключу
func (h *HashMap) Get(key string) (string, bool) {
	index := hashstr(key)

	// Находим список пар для данного индекса
	list := h.buckets[index]

	// Проходим по списку и ищем пару с нужным ключом
	for _, p := range list {
		if p.key == key {
			return p.value, true
		}
	}

	return "", false // Ключ не найден
}

// Delete метод для удаления элемента по ключу
func (h *HashMap) Delete(key string) {
	index := hashstr(key)

	// Находим список пар для данного индекса
	list := h.buckets[index]

	// Проходим по списку и ищем пару с нужным ключом
	for i, p := range list {
		if p.key == key {
			// Удаляем элемент из списка
			h.buckets[index] = append(list[:i], list[i+1:]...)
			return
		}
	}
}

func main() {
	hashMap := NewHashMap()

	// Добавляем элементы
	hashMap.Set("key1", "value1")
	hashMap.Set("key2", "value2")
	hashMap.Set("sfee", "a;sjobwe")
	hashMap.Set("sarg", "thubj")
	fmt.Println(hashMap)
	fmt.Println(hashstr("key1"))

	// Получаем значение по ключу
	if value, exists := hashMap.Get("key1"); exists {
		fmt.Println(value) // выведет "value1"
	}

	// Удаляем элемент
	hashMap.Delete("key2")

	// Проверим наличие удаленного ключа
	if _, exists := hashMap.Get("key2"); !exists {
		fmt.Println("Key 'key2' was deleted successfully.")
	}
}
