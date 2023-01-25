# go-gettable
Process of making golang project gettable.

1. module 名稱格式設置為 github.com/[github 帳戶名稱]/[專案名稱]
2. 設置 tag (例如 v-0.1)，並新增 Release，選擇要釋出的 tag 編號(例如 v-0.1)
3. 此時其他人便可透過 import "github.com/[github 帳戶名稱]/[專案名稱]" 來取得你的套件了

> 要更新套件時，可執行 $ go get -u github.com/[github 帳戶名稱]/[專案名稱]@[tag 編號]
>> 例如: $ go get -u github.com/j32u4ukh/gogettable@v-0.2
