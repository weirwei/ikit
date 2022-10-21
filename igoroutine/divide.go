package igoroutine

// Divide 分批操作，使用场景为：需要获取下游全量数据，而下游必传分页参数
// multi 并发方法，用于并发请求
// page 当前分页，默认为1
// pageSize 页面大小，默认为100
type Divide struct {
	multi    *Multi
	page     int
	pageSize int
	total    int
	errs     []error
}

// NewDivide new a divide
func NewDivide(multi *Multi, opts ...DivideOption) *Divide {
	divide := &Divide{
		multi:    multi,
		page:     1,
		pageSize: 100,
	}
	for _, opt := range opts {
		opt(divide)
	}
	return divide
}

type DivideOption func(divide *Divide)

// OptTotal set total
func OptTotal(total int) DivideOption {
	if total < 0 {
		total = 0
	}
	return func(divide *Divide) {
		divide.total = total
	}
}

// OptPageSize set page size
func OptPageSize(pageSize int) DivideOption {
	if pageSize <= 0 {
		pageSize = 100
	}
	return func(divide *Divide) {
		divide.pageSize = pageSize
	}
}

// OptPage set page
func OptPage(page int) DivideOption {
	if page <= 0 {
		page = 1
	}
	return func(divide *Divide) {
		divide.page = page
	}
}

// GetTotal get total
func (d *Divide) GetTotal() int {
	return d.total
}

// Run 进行分组运行，入参为执行的函数，返回参数为错误信息
func (d *Divide) Run(f func(page, pageSize int) (total int, err error)) []error {
	page, pageSize := d.page, d.pageSize
	if d.total == 0 {
		total, err := f(1, pageSize)
		if err != nil {
			d.errs = append(d.errs, err)
			return d.errs
		}
		d.total = total
		page++
	}
	d.exec(func(page, pageSize int) error {
		_, err := f(page, pageSize)
		return err
	}, page, pageSize)
	return d.errs
}

func (d *Divide) exec(f func(page, pageSize int) error, page, pageSize int) {
	for (page-1)*pageSize < d.total {
		tmpPage := page
		d.multi.Run(func() error {
			return f(tmpPage, pageSize)
		})
		page++
	}
	d.errs = append(d.errs, d.multi.Wait()...)
}
