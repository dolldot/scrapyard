let list_of_words = ["1abc2"; "pqr3stu8vwx"; "a1b2c3d4e5f"; "treb7uchet"]
let f word =
  let x letter = Printf.printf "%c\n" letter in
  String.iter x word in
  List.iter f list_of_words
