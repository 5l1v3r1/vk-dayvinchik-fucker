package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func spam(photo string, message string, name string, delay time.Duration, curage int, maxrandage int, randomAge bool) {

	var file, _ = ioutil.ReadFile("tokens.txt")
	var age int
	if randomAge == true {
		rand.Seed(time.Now().UnixNano())
		age = rand.Intn(maxrandage)
	} else {
		age = curage
	}

	tokens := strings.Split(string(file), "\n")
	for d, token := range tokens {
		green := color.New(color.FgGreen)
		boldGreen := green.Add(color.Bold)
		tlen := strconv.Itoa(len(tokens))
		cur := strconv.Itoa(d + 1)

		var url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=Начать&random_id=0&v=5.95&access_token=" + token
		_, err := http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 1 сообщение отправлено")
		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=👍🏻&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 2 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=2&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 3 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=Москва&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 4 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=" + message + "&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 5 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&attachment=" + photo + "&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 6 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=2&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 7 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 8 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=" + strconv.Itoa(age) + "&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 9 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 10 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=2&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 11 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=Москва&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 12 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=" + name + "&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 13 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=" + message + "&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 14 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 15 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 16 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 17 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=2&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 18 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 19 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 20 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 21 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 22 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 23 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 24 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 25 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 26 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 27 сообщение отправлено")

		url = "https://api.vk.com/method/messages.send?user_id=-91050183&message=1&random_id=0&v=5.95&access_token=" + token
		_, err = http.Get(url)
		time.Sleep(delay * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		boldGreen.Println("[ " + cur + "/" + tlen + " ]" + " 28 сообщение отправлено")

		boldGreen.Println(" Перехожу на новый аккаунт...")

	}
}
