module.exports = {
	'extends': [
		'stylelint-config-standard',
		'stylelint-config-standard-scss',
		'stylelint-config-recess-order',
		'stylelint-config-prettier',
	],
	'rules': {
		// font-family中每一项是否需要引号
		'font-family-name-quotes': null,
		// font-family中不能缺少字体族
		'font-family-no-missing-generic-family-keyword': null,
		// 字符串为单引号
		'string-quotes': 'single',
		// 样式嵌套最大层数
		'max-nesting-depth': [
			3,
			{
				'ignore': ['blockless-at-rules', 'pseudo-classes'],
			},
		],
		// 每行允许的最大字符数
		'max-line-length': 100,
		// 在声明块中禁止重复属性
		'declaration-block-no-duplicate-properties': true,
		// 禁止重复选择器
		'no-duplicate-selectors': true,
		'no-descending-specificity': null,
		// 不验证@未知的名字，为了兼容scss的函数
		'at-rule-no-unknown': null,
		// 颜色指定大写
		'color-hex-case': 'upper',
		// 禁止供应商前缀?
		'property-no-vendor-prefix': null,
		// 禁止小于 1 的小数有一个前导零
		'number-leading-zero': 'never',
		// 禁止空第一行
		'no-empty-first-line': true,
	},
}