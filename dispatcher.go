package dispatcher

import (
	"errors"
	"reflect"
	"sync"
)

var (
	disLock = sync.Mutex{}
	disMap  = map[Signal][]IDis{}
)

type IDis interface {
	Dis(signal Signal, args ...interface{})
}

type Signal int64

func Dispatcher(signal Signal, args ...interface{}) error {
	defer disLock.Unlock()
	disLock.Lock()
	disList, ok := disMap[signal]
	if !ok {
		return errors.New("no found signal.")
	}
	for _, dis := range disList {
		dis := dis
		go func() {
			defer func() {
				if r := recover(); r != nil {
					return
				}
			}()
			dis.Dis(signal, args...)
		}()
	}
	return nil
}

func AddDisWithSignal(dis IDis, signals ...Signal) {
	defer disLock.Unlock()
	disLock.Lock()
	for _, signal := range signals {
		disList, ok := disMap[signal]
		if !ok {
			disList = []IDis{}
		}
		disList = append(disList, dis)
		disMap[signal] = disList
	}
}

func DelDis(dis IDis) {
	defer disLock.Unlock()
	disLock.Lock()
	for signal, disList := range disMap {
		for index, disEl := range disList {
			if reflect.DeepEqual(dis, disEl) {
				disListNew := disList[:index]
				if (index + 1) != len(disList) {
					disListNew = append(disListNew, disList[index+1:]...)
				}
				disMap[signal] = disListNew
				return
			}
		}
	}
}

func DelDisWithSignal(signals ...Signal) {
	defer disLock.Unlock()
	disLock.Lock()
	for _, signal := range signals {
		if disMap[signal] != nil {
			delete(disMap, signal)
		}
	}
}
