# gogettable
Process of making golang project gettable.

1. module 名稱格式設置為 github.com/[github 帳戶名稱]/[專案名稱]
2. 設置 tag (例如 v0.1.0)，並新增 Release，選擇要釋出的 tag 編號(例如 v0.1.0)
3. 此時其他人便可透過 `import "github.com/[github 帳戶名稱]/[專案名稱]"` 來取得你的套件了
4. 要 import 下一層資料夾中的 package 時，使用方式也和一般使用相同，往後面接資料夾名稱即可
* 例如：`import "github.com/j32u4ukh/gogettable/pkg"` 或 `import "github.com/j32u4ukh/gogettable/pkg/utils"`
	

## 更新套件

可執行 `$ go get -u github.com/[github 帳戶名稱]/[專案名稱]@[tag 編號]`

例如: `$ go get -u github.com/j32u4ukh/gogettable@v0.2.0`

## Tag

Tag 名稱建議格式為 v1.0.0 或 v2.3.4；若為非正式版可用 v0.2.0-alpha 或 v5.9-beta.3。

原本我使用 v-1.0 的格式，其他專案在引用時，go.mod 檔案中的版本編號就會與預期的不同，不是 v-1.0 而是 commit 編號