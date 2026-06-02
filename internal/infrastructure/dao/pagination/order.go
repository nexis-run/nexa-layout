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
// 格式为: `字段:排序规则|字段:排序规则`, 例如: `id:desc|created_at:asc`
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

// splitAndTrim 按分隔符分割并去除空格
func splitAndTrim(s, sep string) []string {
	raw := make([]string, 0)

	for _, part := range split(s, sep) {
		trimmed := trimSpace(part)

		if trimmed != "" {
			raw = append(raw, trimmed)
		}
	}

	return raw
}

// split 用于字符串分割（兼容 strings.Split）
func split(s, sep string) []string {
	// 推荐实际项目用 strings.Split
	var res []string

	start := 0

	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			res = append(res, s[start:i])
			start = i + len(sep)
			i = start - 1
		}
	}

	res = append(res, s[start:])

	return res
}

// trimSpace 去除首尾空格（兼容 strings.TrimSpace）
func trimSpace(s string) string {
	// 推荐实际项目用 strings.TrimSpace
	start, end := 0, len(s)

	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}

	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n') {
		end--
	}

	return s[start:end]
}

// GetOrder 获取排序函数
// fields: key-前端传参字段, value-数据库字段映射
// ignoreErr: 是否忽略非法字段错误, true-忽略, false-返回错误
func (p *Paginator) GetOrder(fields map[string]string, ignoreErr bool) (orders []func(*sql.Selector), err error) {
	if p.Order == nil || *p.Order == "" {
		return nil, nil
	}

	ofs := ParseOrderFields(*p.Order)
	for _, of := range ofs {
		column, ok := fields[of.Field]

		if !ok {
			zap.L().Error("排序字段非法", zap.String("field", of.Field), zap.Reflect("fields", fields))

			if !ignoreErr {
				return nil, fmt.Errorf("%w: %s", model.ErrInvalidOrderField, of.Field)
			}
		}

		switch of.By {
		case ByAsc:
			// TODO: 可使用 ent.Asc 替代
			orders = append(orders, func(s *sql.Selector) {
				s.OrderBy(sql.Asc(s.C(column)))
			})
		case ByDesc:
			// TODO: 可使用 ent.Desc 替代
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
