我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：应该 wrap 这个 error，在业务调用过程中，不同层可能会对 sql.ErrNoRows 进行封装，以此来增加上下文信息，在这种操作以后，在最上面的业务层再去进行错误判断的时候，就不能直接用 err == sql.ErrNoRows 了，如果还想判断错误，就需要使用 wrap 的机制。

[wrap.go](https://github.com/shinwu/geektime-go/blob/main/week2/wrap/wrap.go) 是使用内制 errors 机制

[wrap2.go](https://github.com/shinwu/geektime-go/blob/main/week2/wrap/wrap2.go) 使用 github.com/pkg/errors 
