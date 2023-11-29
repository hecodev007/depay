package com

import (
	"time"
)

const (
	epoch     int64 = 1609459200000 // 2021-01-01 00:00:00 的时间戳，可以根据实际需求调整
	machineID uint8 = 111           // 机器 ID，可以根据实际需求调整
)

type Snowflake struct {
	sequence      uint16
	lastTimestamp int64
}

func (s *Snowflake) nextID() int64 {
	timestamp := time.Now().UnixNano()/1000000 - epoch
	if timestamp < s.lastTimestamp {
		panic("时钟回拨")
	}
	if timestamp == s.lastTimestamp {
		s.sequence++
		if s.sequence >= 1<<12 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano()/1000000 - epoch
			}
			s.sequence = 0
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = timestamp
	return (timestamp << 22) | (int64(machineID) << 12) | int64(s.sequence)
}

func GenerateOrderNumber() int64 {
	sf := Snowflake{}
	return sf.nextID()
}
