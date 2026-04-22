# DAO

## 命名规则

> 获取单个实体`[Entity]`时，实体`[Entity]`可忽略，例如：`GetUserByID`可简化为`GetByID`。

- DAO接口命名：`[Entity]DAO`，例如：`UserDAO`
- 确定字段获取命名：`[Entity]DAO.Get[Entity]By[Field]`，例如：`UserDAO.GetUserByID`
- 更新字段命名：`[Entity]DAO.Update[Entity]By[Field]`，例如：`UserDAO.UpdateUserByID`
- 删除命名：`[Entity]DAO.Delete[Entity]By[Field]`，例如：`UserDAO.DeleteUserByID`
- 创建命名：`[Entity]DAO.Create[Entity]`，例如：`UserDAO.CreateUser`
- 模糊搜索命名：`[Entity]DAO.Search[Entity]By[Field]`，例如：`UserDAO.SearchUsersByName`
- 批量操作命名：`[Entity]DAO.Batch[Operation][Entity]By[Field]`，例如：`UserDAO.BatchDeleteUsersByIDs`
- 关联查询命名：`[Entity]DAO.Get[RelatedEntity]By[Field]`，例如：`UserDAO.GetOrdersByUserID`
- 包含关联查询命名：`[Entity]DAO.Get[Entity]With[RelatedEntity]By[Field]`，例如：`UserDAO.GetUserWithOrdersByID`
- 统计命名：`[Entity]DAO.CountBy[Field]`，例如：`UserDAO.CountByStatus`
- 分页查询命名：`[Entity]DAO.Get[Entities]By[Field]WithPagination`，例如：`UserDAO.GetUsersByStatusWithPagination`
