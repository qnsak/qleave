
https://github.com/golang-standards/project-layout/blob/master/README_zh-TW.md
/cmd
/internal
    - domain
      - usecase - aggreates
      - delivery - adapter 
      - repository
/domain - model entity value object
/pkg
/api


Models Layer       - domain entiy
Repository Layer
Usecase Layer
Delivery Layer

Entity 	Domain
Use Case 	Application Service
Adapter 	Infrastructure
Framework & Driver 	External Service (IO)

go test -v  ./internal/user/usecase/test/

命名規範
  package:
    最好保持package 的名字和目錄保持一致。
    盡量採取有意義的包名，簡短，有意義，盡量和標準庫不要衝突。
    包名應該為小寫單詞，盡量不要使用下劃線或者混合大小寫。
    包名以及包所在的目錄名，不要使用複數，例如，是net/utl，而不是net/urls。
    不要用common、util、shared 或者lib 這類寬泛的、無意義的包名。
    包名要簡單明了，例如net、time、log。
  file
    盡量採取有意義的文件名，簡短，有意義，應該為小寫單詞，使用下劃線分隔各個單詞。
  struct
    採用駝峰命名法，首字母根據訪問控制大寫或者小寫。
    type UserInfo struct {
      Name      string `gorm:"type:varchar(20);not null"`
      Telephone string `gorm:"type:varchar(11);not null;unique"`
      Password  string `gorm:"size:255;not null"`
    }
  interface
    名字仍然以單個詞為優先。命名規則基本和上面的結構體類型類似。對於擁有唯一方法或通過多個擁有唯一方法的接口組合而成的接口，Go 語言的慣例是一般用"方法名+er"的方式為interface 命名，例如Reader , Writer 。
  變量命名
    和結構體類似，變量名稱一般遵循駝峰法，首字母根據訪問控制原則大寫或者小寫，但遇到特有名詞時，需要遵循以下規則：
    如果變量為私有，且特有名詞為首個單詞，則使用小寫，如appService；
    若變量類型為bool 類型，則名稱應以Has, Is, Can 或Allow 開頭。
  Error
    var ErrFormat = errors.New("unknown format")
https://google.github.io/styleguide/go/decisions


identification
leave
approved leave
He took a sick leave for three days last week.

$ mockgen -source=./domain/user.go -destination=./internal/identification/repository/mocks/user_repository.go -package=mocks


部署
https://medium.com/beyn-technology/production-ready-go-applications-1c876f013e78
https://codesahara.com/blog/how-to-deploy-golang-to-production-step-by-step/?fbclid=IwAR1kPvtPCFsOlxlZ63NQQi7Z4UyLFcdysLcnjRi5rMfH-_yeHIojb3HwMmU

