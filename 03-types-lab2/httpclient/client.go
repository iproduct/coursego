package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Repo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func main() {
	//resp, err := http.Get("http://localhost:8080/headers")
	//resp, err := http.Get("http://google.com")
	//req, err := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	//req, err := http.NewRequest("GET", "https://api.github.com/users/iproduct/repos", nil)
	//req.Header.Add("Accept", `Accept: application/json`)
	//req.Header.Add("Custom-Header", `Custom Value`)
	//resp, err := http.DefaultClient.Do(req)

	resp, err := http.Get("https://api.github.com/users/iproduct/repos")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	// Results will be unmarshalled here
	var repos []Repo

	//// 1) Reading response.Body using bufio.NewScanner - not very efficient
	//scanner := bufio.NewScanner(resp.Body)
	//scanner.Split(bufio.ScanBytes)
	//jsonBuilder := strings.Builder{}
	//for scanner.Scan() {
	//	jsonBuilder.Write(scanner.Bytes())
	//}
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
	//err = json.Unmarshal([]byte(jsonBuilder.String()), &repos)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 2) Reading response.Body using bytes.Buffer
	buff := bytes.Buffer{}
	_, err = buff.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes := buff.Bytes()
	err = json.Unmarshal(bodyBytes, &repos)
	if err != nil {
		log.Fatal(err)
	}

	//// 3) Reading response.Body using bufio.NewReader
	//reader := bufio.NewReader(resp.Body)
	//bodyBytes, err := reader.ReadBytes(0)
	//if err != io.EOF {
	//	log.Fatal(err)
	//}
	//err = json.Unmarshal(bodyBytes, &repos)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 4) Reading response.Body using ioutil.ReadAll
	//bodyBytes, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = json.Unmarshal(bodyBytes, &repos)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// 5) Reading response.Body directly using json.Decoder - Preferred!
	//decoder := json.NewDecoder(resp.Body)
	//err = decoder.Decode(&repos)
	//if err != nil {  log.Fatal(err) }

	//
	//err := getData("https://api.github.com/users/iproduct/repos", &repos)
	//if err != nil {  log.Fatal(err) }

	// Printing repos - information subset
	for _, repo := range repos {
		fmt.Printf("%v | %v\n", repo.Id, repo.Name)
	}

	//scanner := bufio.NewScanner(resp.Body)
	//for i := 0; scanner.Scan() && i < 10; i++ {
	//	fmt.Println(i+1, ": ", scanner.Text())
	//}
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

}

func getData(uri string, dataPtr interface{}) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&dataPtr)
	return err
}
