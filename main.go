package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	public := os.Getenv("MARVEL_PUBLIC_KEY")
	private := os.Getenv("MARVEL_PRIVATE_KEY")
	clinet := marvelClient{
		private:    private,
		public:     public,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}
character, err := client.getCharacter ()
if err != nil {
log.Fatal(err)
}
spew.Dump(characters)
}

type marvelClient struct {
	public     string
	private    string
	httpClient *http.Client
}

func (c *marvelClient) md5Hash (ts int64)

func (c *marvelClient) signURL(url string) string {
//ts := time.Now ().Unix()
return ""
}

func (c *marvelClient) getCharacters() ([]Character, error){
	res, err:= c.httpClient.Get("https://gatewway.marvel.com/v1/public/characters")
	if err != nil {
		return nil , err
	}
	defer res.Body.Close()

	var characterResponse characterResponse
	if err := json.NewDecoder(res.Body).Decode(&characterResponse); err != nil {
		return nil, err
	}
return  characterResponse.Data.Results, nil
}
