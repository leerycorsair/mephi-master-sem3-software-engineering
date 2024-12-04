package model

import "errors"

var (
	ErrResourceNotFound   = errors.New("Ресурс не найден")
	ErrResourceIsBusy     = errors.New("Ресурс занят")
	ErrResourceInvalid    = errors.New("Ресурс некорректный")
	ErrModelIncorrectFile = errors.New("Некорректный файл модели")
	ErrResourceCreateFail = errors.New("Ресурс не получилось создать")
)
