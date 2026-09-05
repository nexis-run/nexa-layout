// Copyright (C) moneta. 2025-present.
//
// Created at 2025-09-18, by liasica

package pagination

// Options 配置项
type Options struct {
	ignoreOrderErr bool // 排序字段映射错误时忽略该排序字段
}

// Option 配置选项
type Option func(*Options)

var _ = IgnoreOrderErr

// IgnoreOrderErr 跳过未映射的排序字段
func IgnoreOrderErr() Option {
	return func(o *Options) {
		o.ignoreOrderErr = true
	}
}

// GetOptions 获取配置项
func GetOptions(opts ...Option) (options *Options) {
	options = &Options{}

	for _, o := range opts {
		if o != nil {
			o(options)
		}
	}

	return options
}
