package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var mode int
	var variant int
	var p, q, e, n, d int

	fmt.Println("Выберите режим: \nШифрование = 1, Дешифрование = 2")
	_, err := fmt.Scan(&mode)
	if err != nil {
		log.Fatal(err)
	}
	if mode != 1 && mode != 2 {
		log.Fatal("Выбран неверный режим\n")
	}

	fmt.Println("Ввод ключей = 1, ввод коэффициентов p,q = 2")
	_, err = fmt.Scan(&variant)
	if err != nil {
		log.Fatal(err)
	}

	if variant != 1 && variant != 2 {
		log.Fatal("Выбран неверный режим\n")
	}

	if variant == 1 && mode == 1 {
		fmt.Println("Введите коэффициенты открытого ключа (e, n)")
		_, err = fmt.Scan(&e, &n)
		if err != nil {
			log.Fatal(err)
		}

	} else if variant == 1 && mode == 2 {
		fmt.Println("Введите коэффициенты закрытого ключа (d, n)")
		_, err = fmt.Scan(&d, &n)
		if err != nil {
			log.Fatal(err)
		}

	} else if variant == 2 {
		fmt.Println("Введите коэффициенты p, q")
		_, err = fmt.Scan(&p, &q)

		if err != nil {
			log.Fatal(err)
		}

		if p == q {
			log.Fatal("Коэффициенты p, q равны")
		}
		isPrime(q)
		isPrime(p)

		n = p * q

		phi := phiFunc(n)

		e = eFind(phi)
		fmt.Printf("Открытый ключ: (e, n) = (%d, %d)\n", e, n)

		d = dFind(e, phi)
		fmt.Printf("Закрытый ключ: (d, n) = (%d, %d)\n", d, n)
	}

	if mode == 1 {
		fmt.Println("Введите сообщение")
		reader := bufio.NewReader(os.Stdin)

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		text = strings.TrimSpace(text)
		encrypt(text, e, n)

	} else if mode == 2 {
		var count int
		fmt.Println("Введите количество элементов")
		fmt.Scan(&count)

		var numbers []int
		fmt.Println("Введите зашифрованный текст")
		for i := 0; i < count; i++ {
			var n int
			fmt.Scan(&n)
			numbers = append(numbers, n)
		}
		decrypt(numbers, d, n)
	}

}
