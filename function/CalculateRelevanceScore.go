package main

import (
	"MyStackoverflow/dao"
	"database/sql"
	"fmt"
)

func scanAndSum(rows *sql.Rows, m map[int]float64, maxScore float64, weight float64) map[int]float64 {

	for rows.Next() {
		var qid int
		var score float64
		err := rows.Scan(&qid, &score)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// normalization: divided by the max
		score = score / maxScore * weight
		if _, ok := m[qid]; !ok {
			m[qid] = score
		} else {
			m[qid] += score
		}
	}
	return m
}

func CalculateRelevanceScoreForQuestion(keyword string) map[int]float64 {
	/*
		Input keyword(if multiple separate with a single space), return a mapping(qid : relevance score)
	*/
	// could be adjusted
	weightMap := map[string]float64{
		"question_title": 2.0,
		"question_body":  1.0,
		"answer_body":    3.0,
		"topic":          3.0,
	}
	res := make(map[int]float64)
	db := dao.MyDB
	// relevance score for title of the question
	var maxScore float64
	err := db.Raw("select max(match(title) against('mysql')) from Questions").Scan(&maxScore).Error
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Raw("select qid, match(title) against(?) as score from Questions", keyword).Rows()
	if err != nil {
		fmt.Println(err)
	}
	res = scanAndSum(rows, res, maxScore, weightMap["question_title"])
	defer func() {
		_ = rows.Close()
	}()
	// relevance score for body of the question
	err = db.Raw("select max(match(body) against('mysql')) from Questions").Scan(&maxScore).Error
	if err != nil {
		fmt.Println(err)
	}
	rows, err = db.Raw("select qid, match(body) against(?) as score from Questions", keyword).Rows()
	if err != nil {
		fmt.Println(err)
	}
	res = scanAndSum(rows, res, maxScore, weightMap["question_body"])
	// relevance score for body of the answer
	var maxAnswerScore float64
	rows, err = db.Raw("select sum(match(body) against('mysql')) from Answers group by qid").Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var score float64
		err = rows.Scan(&score)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if score > maxAnswerScore {
			maxAnswerScore = score
		}
	}
	rows, err = db.Raw("select qid, sum(match(body) against('mysql')) from Answers group by qid").Rows()
	if err != nil {
		fmt.Println(err)
	}
	res = scanAndSum(rows, res, maxAnswerScore, weightMap["answer_body"])
	// relevance score for topic of the question
	rows, err = db.Raw("select tid, qid, count(*) from Topics join QuestionTopics using (tid) where topic_name = ? group by tid, qid", keyword).Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var tid int
		var qid int
		var score float64
		err = rows.Scan(&tid, &qid, &score)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		score = score * weightMap["topic"]
		if _, ok := res[qid]; !ok {
			res[qid] = score
		} else {
			res[qid] += score
		}
	}
	return res
}

func main() {

	m := CalculateRelevanceScoreForQuestion("mysql databases")
	fmt.Println(m)
}
