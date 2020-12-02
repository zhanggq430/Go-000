# 作业

我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

个人理解

1. 如果是sql.ErrNowRows应该向上抛，并且wrap这个error,这样上层才能判断错误的具体情况，如果是其他错误应该自己处理掉防止上层过多依赖dao层的信息

2. 我的例子里面是dao(warp这个error) -> service(直接向上抛)->api(通过is来判断根error)做出错误处理

