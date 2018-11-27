//Nate Salima

package main

import (
  "io/ioutil"
  "os"
  "strings"
  "fmt"
)

func main(){
  //Grabs File
  if len(os.Args) < 2 {
         fmt.Println("Missing parameter, provide file name!")
         return
     }
     var entry string = os.Args[1]
     b, err := ioutil.ReadFile(os.Args[1])
     if err != nil {
         fmt.Println("Can't read file:", os.Args[0])
         panic(err)
     }
     // Creates output file
     f, err := os.Create(entry[0:len(entry)-3] + "out")
      if err != nil {
          panic(err)
      }
      defer f.Close()
  str := ""
  bs := []byte(str)
  lexErr := false
  tokens := 0
  test := string(b)
  words := strings.Fields(test)
  fmt.Println("Processing Input File " + (os.Args[1]))
  //checks for tokens and outputs the tokens lexemes
  for i := 0; i < len(words); i++ {
    word := words[i]
    first := word[0:1]
    last := word[len(word)-1: len(word)]
    after := word[1:len(word)]
    if first == "$" {
      str = ("ID[STRING]: "+ after + "\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    } else if first == "#" {
      str = ("ID[INT]: "+ after + "\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if word == "BEGIN"{
      str = ("BEGIN\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if word == "END"{
      str = ("END\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "%"{
      str = ("ID[REAL]: "+ after  + "\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "W"{
      str = ("WRITE\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == ":"{
      str = ("COLON\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if word == "."{
      str = ("POINT\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == `"`{
      check := words[i-2]
      check2 := words[i-1]
      if check[0:1] == "$" || check2 == "WRITE"{
        str = ("STRING: ")
        bs = []byte(str)
        f.Write(bs)
        tokens++
      loop:for i=i ; i <= len(words)-1; i++{
            word = words[i]
            first = word[0:1]
            after = word[1:len(word)]
            last = word[len(word)-1:len(word)]
            if last == `"` {
              if first == `"` && len(word) > 2{
              str = (word[1:len(word)-1] + "\n")
              bs = []byte(str)
              f.Write(bs)
              }else{
              str =(word[0:len(word)-1] + "\n")
              bs = []byte(str)
              f.Write(bs)
              }
              break loop
            }else if first == `"`{
              str = (after + " ")
              bs = []byte(str)
              f.Write(bs)
            }else if first != `"` && last != `"`{
              str =(word + " ")
              bs = []byte(str)
              f.Write(bs)
            }
          }
        }else if check[0:1] != "$" || check2 != "WRITE"{
          str = ("Lexical Error, unrecognized symbol " + word + "\n")
          bs = []byte(str)
          f.Write(bs)
          lexErr = true
        }
    }else if first == "<" && after == "="{
      str = ("ASSIGN\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "*"{
      str = ("TIMES\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "/"{
      str =("DIVISION\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "+"{
      str = ("PLUS\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "-"{
      str = ("MINUS\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "^"{
      str = ("POWER\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "("{
      str = ("LPAREN\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == ")"{
      str = ("RPAREN\n")
      bs = []byte(str)
      f.Write(bs)
      tokens++
    }else if first == "1"||first == "2"||first == "3"||first == "4"||first == "5"||first == "6"||first == "7"||first == "8"||first == "9"||first == "0"{
      if strings.ContainsAny(word, "."){
        str = ("REAL_CONST: "+ word + "\n")
        bs = []byte(str)
        f.Write(bs)
        tokens++
      }else{
        str = ("INT_CONST: "+ word  + "\n")
        bs = []byte(str)
        f.Write(bs)
        tokens++
      }
    }else{
      str = ("Lexical Error, unrecognized symbol " + word + "\n")
      bs = []byte(str)
      f.Write(bs)
      lexErr = true
    }
  }
  //Prints final amount of Tokens
  fmt.Print(tokens)
  fmt.Println(" Tokens produced")
  fmt.Println("Results in Output File: " + entry[0:len(entry)-3] + "out")
  if lexErr == true {
  fmt.Println("There was a Lexical error processing the file")
  }
}
