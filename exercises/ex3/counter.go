package ex3

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func TotalFile(fileName string) (total int, err error) {
	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer func() {
		err2 := f.Close()
		err = errors.Join(err, err2)
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		total += val
	}
	if err := s.Err(); err != nil {
		return 0, err
	}
	return total, nil
}
