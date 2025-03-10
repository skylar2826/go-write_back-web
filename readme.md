## 功能介绍

使用write-back缓存模式：

init时，同步db数据至缓存；数据过期时删除缓存数据，调用onEvicted同步数据库删除数据；服务关闭时调用hook, 将新数据同步db

## 测试方案

/login post Id=? 查看数据

/register post Id=?Username=?Email=? 插入数据
