package main

import "time"

type Task struct {
	Id      int64      `json:"id" db:"id"`
	Tag     string     `json:"tag" db:"tag"`
	URL     string     `json:"url" db:"url"`
	StartAt *time.Time `json:"start_at" db:"start_at"`
	EndAt   *time.Time `json:"end_at" db:"end_at"`
}

func (t Task) GetId() int64 {
	return t.Id
}

func (t Task) GetTag() string {
	return t.Tag
}

func (t Task) GetURL() string {
	return t.URL
}

func (t Task) GetStartAt() *time.Time {
	return t.StartAt
}

func (t Task) GetEndAt() *time.Time {
	return t.EndAt
}

func (t *Task) SetId(id int64) *Task {
	t.Id = id
	return t
}

func (t *Task) SetTag(tag string) *Task {
	t.Tag = tag
	return t
}

func (t *Task) SetURL(url string) *Task {
	t.URL = url
	return t
}

func (t *Task) SetStartAt(startAt *time.Time) *Task {
	t.StartAt = startAt
	return t
}

func (t *Task) SetEndAt(endAt *time.Time) *Task {
	t.EndAt = endAt
	return t
}

func (t *Task) Localtime() *Task {
	if t.GetStartAt() != nil {
		localtimeStartAt := t.GetStartAt().Local()
		t.SetStartAt(&localtimeStartAt)
	}
	if t.GetEndAt() != nil {
		localtimeEndAt := t.GetEndAt().Local()
		t.SetEndAt(&localtimeEndAt)
	}
	return t
}
