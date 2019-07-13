## 数据加密解密

##### 1.配置说明

    {
      "isEncrypt": true,
      "encryptKey": "goatgames@123456",
      "dataSourcePath": "/Users/huangfuhui/go/test/input",
      "dataOutputPath": "/Users/huangfuhui/go/test/output",
      "dataBackupPath": "/Users/huangfuhui/go/test/backup"
    }
- isEncrypt - 加密或是解密，true代表加密，false代表解密
- encryptKey - 密钥，默认16个字节长的字符，过长会截断，过短会补"0"
- dataSourcePath - 需要加密解密的文件夹的绝对路径
- dataOutputPath - 文件加密解密后输出的绝对路径
- dataBackupPath - 文件临时备份的绝对路径
