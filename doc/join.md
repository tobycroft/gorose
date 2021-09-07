# 更多方法

## Join方法
```go
func (self *Interface) Api_join_select(group_id, user_id interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Join("coin on coin.id=cid")
	db.Where("amount", ">", 0)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
```
