package api

import (
	"errors"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Question structure for a question
type Question struct {
	Key          *datastore.Key `json:"id" datastore:"-"`
	CTime        time.Time      `json:"created" datastore:",noindex"`
	Question     string         `json:"question" datastore:",noindex"`
	User         UserCard       `json:"user"`
	AnswersCount int            `json:"answers_count"`
}

// OK validate the question
func (q Question) OK() error {
	if len(q.Question) < 10 {
		return errors.New("question is too short")
	}
	return nil
}

// Create to save question to datastore
func (q *Question) Create(ctx context.Context) error {
	log.Debugf(ctx, "Saving question: %s", q.Question)
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return err
	}
	q.User = user.Card()
	q.CTime = time.Now()
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}
	return nil
}

// Update to update a question
func (q *Question) Update(ctx context.Context) error {
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}
	var err error
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}
	return nil
}

// GetQuestion to get a question
func GetQuestion(ctx context.Context, key *datastore.Key) (*Question, error) {
	var q Question
	err := datastore.Get(ctx, key, &q)
	if err != nil {
		return nil, err
	}
	q.Key = key
	return &q, nil
}

// TopQuestions to get the top questions
func TopQuestions(ctx context.Context) ([]*Question, error) {
	var questions []*Question
	questionKeys, err := datastore.NewQuery("Question").Order("-AnswerCount").Order("-CTime").Limit(25).GetAll(ctx, &questions)
	if err != nil {
		return nil, err
	}
	for i := range questions {
		questions[i].Key = questionKeys[i]
	}
	return questions, nil
}

// QuestionCard a question card
type QuestionCard struct {
	Key      *datastore.Key `json:"id" datastore:",noindex"`
	Question string         `json:"question" datastore:",noindex"`
	User     UserCard       `json:"user" datastore:",noindex"`
}

// Card return a QuestionCard from a card
func (q Question) Card() QuestionCard {
	return QuestionCard{
		Key:      q.Key,
		Question: q.Question,
		User:     q.User,
	}
}
