package api

import (
	"errors"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Answer an answer
type Answer struct {
	Key    *datastore.Key `json:"id" datastore:"-"`
	Answer string         `json:"answer" datastore:",noindex"`
	CTime  time.Time      `json:"created"`
	User   UserCard       `json:"user" datastore:",noindex"`
	Score  int            `json:"score"`
}

// OK is answer OK
func (a Answer) OK() error {
	if len(a.Answer) < 10 {
		return errors.New("answer is too short")
	}
	return nil
}

// Create create the answer
func (a *Answer) Create(ctx context.Context, questionKey *datastore.Key) error {
	a.Key = datastore.NewIncompleteKey(ctx, "Answer", questionKey)
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return err
	}
	a.User = user.Card()
	a.CTime = time.Now()
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		q, err := GetQuestion(ctx, questionKey)
		if err != nil {
			return err
		}
		q.AnswersCount++
		err = q.Update(ctx)
		if err != nil {
			return err
		}
		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		return err
	}
	return nil
}

// GetAnswer get the answer
func GetAnswer(ctx context.Context, answerKey *datastore.Key) (*Answer, error) {
	var answer Answer
	err := datastore.Get(ctx, answerKey, &answer)
	if err != nil {
		return nil, err
	}
	answer.Key = answerKey
	return &answer, nil
}

// Put put the answer
func (a *Answer) Put(ctx context.Context) error {
	var err error
	a.Key, err = datastore.Put(ctx, a.Key, a)
	if err != nil {
		return err
	}
	return nil
}

// GetAnswers to get the answers
func GetAnswers(ctx context.Context, questionKey *datastore.Key) ([]*Answer, error) {
	var answers []*Answer
	answerKeys, err := datastore.NewQuery("Answer").Ancestor(questionKey).Order("-Score").Order("-CTime").GetAll(ctx, &answers)
	for i, answer := range answers {
		answer.Key = answerKeys[i]
	}
	if err != nil {
		return nil, err
	}
	return answers, nil
}

// AnswerCard represents an answer
type AnswerCard struct {
	Key    *datastore.Key `json:"id"`
	Answer string         `json:"answer"`
	User   UserCard       `json:"user"`
}

// Card get card from an answer
func (a Answer) Card() AnswerCard {
	return AnswerCard{
		Key:    a.Key,
		Answer: a.Answer,
		User:   a.User,
	}
}
