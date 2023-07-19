# 设想一个一般的业务场景：

店铺用商品发布活动，用户购买商品。

# 业务需求，店铺端高级筛选：

* 按照是否购买过自己的商品筛选用户（购买过）；
* 按照最近一段时间内是否购买过自己的商品筛选用户（一段时间内购买过）；
* 按照是否购买过某个商品筛选用户（某商品购买过）；

# 方案：

## 支持[购买过]、[一段时间内购买过]筛选

每个店铺对应一个 redis hash, redis key 为 xxx:{{shop_id}}，其中 {{shop_id}} 为店铺 ID。

hash 中存储的 field-value 对定义为：

* field：用户 ID
* value: 用户最近一次购买的时间

使用 `HMGET` 进行查询。

## 支持[某商品购买过]筛选

每个商品对应一个 redis hash, redis key 为 xxx:product:{{product_id}}，其中 {{product_id}} 为商品 ID。

hash 中存储的 field-value 对定义为：

* field：用户 ID
* value: 空字符串

使用 `HMGET` 进行查询。
