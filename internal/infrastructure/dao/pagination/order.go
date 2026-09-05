// Copyright (C) moneta. 2025-present.
//
// Created at 2025-09-18, by liasica

package pagination

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"go.uber.org/zap"

	"nexis.run/nexa-layout/internal/infrastructure/model"
)

type By string

const (
	ByAsc  By = "asc"
	ByDesc By = "desc"
)

type OrderOption interface {
	~func(*sql.Selector)
}

type OrderFields []*OrderField

// GetString 将 OrderFields 转换为字符串表示
// 格式为 `字段:排序规则|字段:排序规则`，例如 `id:desc|created_at:asc`
func (o OrderFields) GetString() string {
	var builder strings.Builder

	for i, field := range o {
		if i > 0 {
			builder.WriteString("|")
		}

		builder.WriteString(field.GetString())
	}

	return builder.String()
}

// OrderField 表示单个排序字段和规则
type OrderField struct {
	Field string // 字段名
	By    By     // 排序规则（asc/desc）
}

func NewOrderField(field string, by By) *OrderField {
	return &OrderField{
		Field: field,
		By:    by,
	}
}

func (o *OrderField) GetString() string {
	return o.Field + ":" + string(o.By)
}

// ParseOrderFields 解析排序字符串，返回 OrderField 列表
func ParseOrderFields(orderStr string) []OrderField {
	var result []OrderField

	if orderStr == "" {
		return result
	}

	items := splitAndTrim(orderStr, "|")

	for _, item := range items {
		parts := splitAndTrim(item, ":")
		if len(parts) == 2 {
			field := parts[0]

			by := By(parts[1])
			if field != "" && (by == ByAsc || by == ByDesc) {
				result = append(result, OrderField{Field: field, By: by})
			}
		}
	}

	return result
}

// splitAndTrim 按分隔符分割并保留空字段
func splitAndTrim(value, separator string) []string {
	parts := strings.Split(value, separator)
	for index := range parts {
		parts[index] = strings.TrimSpace(parts[index])
	}

	return parts
}

// GetOrder 获取排序函数
// fields 将请求字段映射为数据库列
// ignoreErr 为 true 时跳过未知字段
func (p *Paginator) GetOrder(fields map[string]string, ignoreErr bool) (orders []func(*sql.Selector), err error) {
	if p == nil || p.Order == nil || *p.Order == "" {
		return
	}

	ofs := ParseOrderFields(*p.Order)
	for _, of := range ofs {
		column, ok := fields[of.Field]

		if !ok || column == "" {
			zap.L().Error("排序字段非法", zap.String("field", of.Field), zap.Reflect("fields", fields))

			if !ignoreErr {
				orders = nil
				err = fmt.Errorf("%w：%s", model.ErrInvalidOrderField, of.Field)

				return
			}

			continue
		}

		switch of.By {
		case ByAsc:
			orders = append(orders, func(s *sql.Selector) {
				s.OrderBy(sql.Asc(s.C(column)))
			})
		case ByDesc:
			orders = append(orders, func(s *sql.Selector) {
				s.OrderBy(sql.Desc(s.C(column)))
			})
		}
	}

	return
}

// GetOrder 泛型获取排序函数
func GetOrder[T OrderOption](p *Paginator, fields map[string]string, ignoreErr bool) (orders []T, err error) {
	var list []func(*sql.Selector)
	list, err = p.GetOrder(fields, ignoreErr)

	orders = make([]T, len(list))
	for i, f := range list {
		orders[i] = T(f)
	}

	return
}
