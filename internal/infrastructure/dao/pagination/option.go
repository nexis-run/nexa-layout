// Copyright (C) moneta. 2025-present.
//
// Created at 2025-09-18, by liasica

package pagination

// Options 配置项
type Options struct {
	ignoreOrderErr bool // 排序字段映射错误时忽略该排序字段
	failOnCountErr bool // Count出错时返回错误
}

// Option 配置选项
type Option func(*Options)

var _ = IgnoreOrderErr

// IgnoreOrderErr 当排序字段映射错误时忽略该排序字段, 而是跳过该排序字段
// 当传入的排序字段不合法时, 不返回错误
// 适用于前端传入的排序字段不受控制的场景
// 注意: 使用该选项后, 可能会导致返回的数据顺序不符合预期, 请谨慎使用
func IgnoreOrderErr() Option {
	return func(o *Options) {
		o.ignoreOrderErr = true
	}
}

var _ = FailOnCountErr

// FailOnCountErr Count出错时返回错误
func FailOnCountErr() Option {
	return func(o *Options) {
		o.failOnCountErr = true
	}
}

// GetOptions 获取配置项
func GetOptions(opts ...Option) (options *Options) {
	options = &Options{}

	for _, o := range opts {
		o(options)
	}

	return options
}
