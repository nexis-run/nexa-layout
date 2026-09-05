// Copyright (C) moneta. 2025-present.
//
// Created at 2025-09-17, by liasica

package pagination

import (
	"context"
	"math"

	"entgo.io/ent/dialect/sql"
)

type Counter interface {
	Count(ctx context.Context) (int, error)
}

type Querier[Q Counter, D any] interface {
	Clone() Q
	Offset(offset int) Q
	Limit(limit int) Q
	All(ctx context.Context) ([]*D, error)
}

type OrderedQuerier[Q Counter, O OrderOption, D any] interface {
	Clone() Q
	Order(o ...O) Q
	Offset(offset int) Q
	Limit(limit int) Q
	All(ctx context.Context) ([]*D, error)
}

// Paginator 分页实体
type Paginator struct {
	PageSize int     `json:"pageSize,omitempty" query:"pageSize" validate:"gte=0,lte=100"` // 每页数量，最大为 100
	PageNum  int     `json:"pageNum,omitempty" query:"pageNum" validate:"gte=0"`           // 页码，从 1 开始
	Order    *string `json:"order,omitempty" query:"order"`                                // 排序字段用竖线分隔，例如 `id:desc|createdAt:asc`
}

// Parse 解析并设置默认值
func (p *Paginator) Parse() *Paginator {
	if p == nil {
		p = &Paginator{}
	}

	if p.PageNum <= 0 {
		p.PageNum = 1
	}

	if p.PageSize <= 0 {
		p.PageSize = 20
	}

	if p.PageSize > 100 {
		p.PageSize = 100
	}

	return p
}

// GetNextPageNum 获取下一页页码
func (p *Paginator) GetNextPageNum() int {
	p = p.Parse()
	if p.PageNum == math.MaxInt {
		return math.MaxInt
	}

	return p.PageNum + 1
}

// GetOffset 计算偏移量
func (p *Paginator) GetOffset() int {
	p = p.Parse()
	if p.PageNum-1 > math.MaxInt/p.PageSize {
		return math.MaxInt
	}

	return (p.PageNum - 1) * p.PageSize
}

// GetLimit 获取限制数量
func (p *Paginator) GetLimit() int {
	return p.Parse().PageSize
}

// Result 分页结果实体
type Result[T any] struct {
	Total    int  `json:"total"`    // 总数量
	PageSize int  `json:"pageSize"` // 每页数量
	PageNum  int  `json:"pageNum"`  // 页码，从 1 开始
	Items    []*T `json:"items"`    // 列表数据
}

var _ = PageList[Counter, any]

// PageList 通用分页列表查询
func PageList[Q Counter, D any](ctx context.Context, query Querier[Q, D], p *Paginator) (res *Result[D], err error) {
	// 预处理分页器
	p = p.Parse()

	res = &Result[D]{
		PageNum:  p.PageNum,
		PageSize: p.PageSize,
		Items:    make([]*D, 0),
	}

	res.Total, err = query.Clone().Count(ctx)
	if err != nil {
		return
	}

	offset := p.GetOffset()

	if res.Total == 0 || offset >= res.Total {
		return
	}

	query.Offset(offset)
	query.Limit(p.PageSize)
	res.Items, err = query.All(ctx)

	return
}

var _ = OrderedPageList[Counter, func(*sql.Selector), any]

// OrderedPageList 通用排序分页列表查询
func OrderedPageList[Q Counter, O OrderOption, D any](ctx context.Context, query OrderedQuerier[Q, O, D], p *Paginator, m map[string]string, opts ...Option) (res *Result[D], err error) {
	// 预处理分页器
	p = p.Parse()

	options := GetOptions(opts...)

	res = &Result[D]{
		PageNum:  p.PageNum,
		PageSize: p.PageSize,
		Items:    make([]*D, 0),
	}

	res.Total, err = query.Clone().Count(ctx)
	if err != nil {
		return
	}

	offset := p.GetOffset()

	if res.Total == 0 || offset >= res.Total {
		return
	}

	var orders []O

	orders, err = GetOrder[O](p, m, options.ignoreOrderErr)
	if err != nil {
		return
	}

	query.Order(orders...)
	query.Offset(offset)
	query.Limit(p.PageSize)
	res.Items, err = query.All(ctx)

	return
}
