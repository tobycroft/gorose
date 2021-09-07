package gorose_pro

// NewBuilder 获取builder
func NewBuilder(driver string) IBuilder {
	return NewBuilderDriver().Getter(driver)
}
