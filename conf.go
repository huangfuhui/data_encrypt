package main

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	IsEncrypt      bool   // 加密或解密
	EncryptKey     string // 加密密钥
	DataSourcePath string // 数据源路径
	DataOutPutPath string // 加密后数据输出路径
	DataBackupPath string // 数据备份路径
}

// 读取配置文件，并将配置数据装载进Conf结构体
func (c *Conf) LoadConf(configFileName string, v interface{}) error {
	data, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

// 获取key，默认长度16字节，对应AES-128加密
func (c *Conf) GetEncryptKey() []byte {
	if len(c.EncryptKey) >= 16 {
		c.EncryptKey = c.EncryptKey[0:16]
	} else {
		// key长度不够，默认填充"0"
		for i := 16 - len(c.EncryptKey); i > 0; i-- {
			c.EncryptKey += "0"
		}
	}

	return []byte(c.EncryptKey)
}
