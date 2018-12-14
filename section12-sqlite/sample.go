package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // コード上使用しないがドライバとして使用するのでimportしておく
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// DB接続
	DbConnection, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatalln(err)
	}

	// 最後にクローズ
	defer DbConnection.Close()

	// テーブル作成
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	log.Println(cmd)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// // INSERT
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Nancy", 20) // パラメータ代入
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// UPDATE
	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	// SELECT
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {

		// 構造体のメンバにSELECT結果を代入
		var p Person
		err := rows.Scan(&p.Name, &p.Age)

		// エラー取得（1行ずつ）
		if err != nil {
			log.Println(err)
		}

		// 配列に追加
		pp = append(pp, p)
	}

	// エラー取得（一括）
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	// SELECT結果出力
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
