package config

import (
	"github.com/bytedance/sonic"
)

func SetupSonic() {
}

func MarshalToString(v interface{}) (string, error) {
	return sonic.MarshalString(v)
}

func UnmarshalFromString(data string, v interface{}) error {
	return sonic.UnmarshalString(data, v)
}

const EnablePrefork = true