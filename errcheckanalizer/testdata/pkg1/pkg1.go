package pkg1

import (
	"errors"
)

func mulfunc(i int) (int, error) {
	res := i * 2
	if i < 0 && res > 0 {
		return 0, errors.New("error")
	}
	return res, nil
}

func errCheckFunc() (int, error) {
	// формулируем ожидания: анализатор должен находить ошибку,
	// описанную в комментарии want
	res, _ := mulfunc(5) // want "assignment with unchecked error"
	// fmt.Println(res)
	return res, nil
}
